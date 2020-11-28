package order

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"
)

// IOvertime 超时接口
type IOvertime interface {
	OrderOvertime(taskID, orderID int64) error
	DeliveryOvertime(deliveryID, orderID, taskID int64) error
}

// Overtime 下游支付完成结构体
type Overtime struct {
	c    component.IContainer
	db   IDBOvertime
	task task.IQTask
}

// NewOvertime 构建Overtime
func NewOvertime(c component.IContainer, db IDBOvertime, task task.IQTask) *Overtime {
	return &Overtime{
		c:    c,
		db:   db,
		task: task,
	}
}

// OrderOvertime 订单超时
func (o *Overtime) OrderOvertime(taskID, orderID int64) error {

	// 1.查询非失败超时订单信息
	data, err := o.db.QueryOrderInfoByDB(orderID)
	if err != nil {
		return nil
	}
	if data == nil {
		if err := qtask.Finish(o.c, taskID); err != nil {
			return err
		}
		return nil
	}
	trans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 2.开始超时任务
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 3.判断卡单环节
	var tasks []string
	var deliOverFuncs []func(c interface{}) error
	switch data.GetInt("order_status") {
	// 3.1支付失败
	case enum.OrderStatus.DownPay:
		// 3.1.1处理订单状态为失败,订单通知为等待通知
		tasks, err = o.db.OrderOvertimePayFailedByDB(trans, data)
		if err != nil {
			trans.Rollback()
			return err
		}

	// 3.2绑定发货
	case enum.OrderStatus.BindDelivery:
		// 3.2.1 判断是否绑定失败
		allBindFail, err := o.db.CheckBindFailedByDB(trans, data)
		if err != nil {
			trans.Rollback()
			return err
		}
		// 3.2.2 绑定失败处理,修改订单为失败,修改通知为等待通知,如果全部失败,创建退款和通知任务
		if allBindFail {
			tasks, err = o.db.BindOvertimeFailedByDB(trans, data)
			if err != nil {
				trans.Rollback()
				return err
			}
			break
		}

		// 3.2.3 处于发货中,创建发货超时消息队列,部分成功部分失败,修改超时时间
		deliveryTask, deliverys, err := o.db.DeliveryOvertimeFailedByDB(trans, data)
		if err != nil {
			trans.Rollback()
			return err
		}

		//3.2.4 创建发货未知处理任务
		deliOverFuncs, err = o.task.CreateBatchTasks(trans, deliveryTask, deliverys)
		if err != nil {
			trans.Rollback()
			return err
		}
	}

	// 4.处理任务
	queueFuncs, err := o.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}
	queueFuncs = append(queueFuncs, deliOverFuncs...)

	// 5.结束任务
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	trans.Commit()
	return utils.CallBackFuncs(o.c, queueFuncs...)
}

// DeliveryOvertime 发货超时
func (o *Overtime) DeliveryOvertime(deliveryID, orderID, taskID int64) error {

	trans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}
	// 1.发货中,锁订单和发货
	data, err := o.db.LockOrderAndDeliveryByDB(trans, deliveryID, orderID)
	if err != nil {
		trans.Rollback()
		return err
	}

	if data.IsEmpty() {
		if qerr := qtask.Finish(trans, taskID); qerr != nil {
			trans.Rollback()
			return qerr
		}
		trans.Commit()
		return fmt.Errorf("订单和发货已完成或不存在")
	}

	// 2.开始超时任务
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 3.等待发货,处理订单发货和退款
	var tasks []string
	if data.GetInt("delivery_status") == enum.DeliveryStatus.Wait {
		// 3.1处理订单和发货为失败,修改通知状态
		tasks, err = o.db.WaitingDeliveryOvertimeByDB(trans, data, nil)
		if err != nil {
			trans.Rollback()
			return err
		}

	} else {
		// 4.查询发货结果
		result, err := jcsdk.QueryOrder(trans, data.GetString("down_channel_no"), types.GetString(deliveryID), o.db.BuildServiceClass())
		if err != nil || result == nil {
			trans.Rollback()
			return fmt.Errorf("err:%v,result:%v", err, result)
		}

		// 6.发货结果处理
		switch result.Status {
		// 6.1发货成功
		case enum.DeliveryResult.Success:
			// 6.1.1 处理订单成功金额,发货状态,判断是否完全发货成功,修改通知状态
			tasks, err = o.db.DeliverySuccessByDB(trans, data, result)
			if err != nil {
				trans.Rollback()
				return err
			}
		// 6.2 发货失败处理
		case enum.DeliveryResult.Failed:
			// 6.2.1处理订单和发货为失败,判断是否完全失败,修改通知状态
			tasks, err = o.db.WaitingDeliveryOvertimeByDB(trans, data, result)
			if err != nil {
				trans.Rollback()
				return err
			}
		//6.3 创建审核记录
		default:
			if err := o.db.DeliveryAuditByDB(trans, data); err != nil {
				trans.Rollback()
				return err
			}
		}
	}
	// 7.处理任务
	queueFuncs, err := o.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}

	// 8.发货超时任务结束
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}
	trans.Commit()
	return utils.CallBackFuncs(o.c, queueFuncs...)
}
