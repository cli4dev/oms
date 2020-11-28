package overtime

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"
)

// IOverTime 退款超时扫描接口
type IOverTime interface {
	TimeoutDeal(refundID string, taskID int64) (err error)
	ReturnUnknownDeal(refundID int64, returnID int64, orderID int64, taskID int64) (err error)
}

// OverTime 退款超时扫描对象
type OverTime struct {
	c    component.IContainer
	db   IDBOverTime
	task task.IQTask
}

//NewOverTime 创建退款超时扫描对象
func NewOverTime(c component.IContainer, db IDBOverTime, task task.IQTask) *OverTime {
	return &OverTime{
		c:    c,
		db:   db,
		task: task,
	}
}

/*******************************超时处理***********************************/

// TimeoutDeal 处理超时记录
func (n *OverTime) TimeoutDeal(refundID string, taskID int64) (err error) {
	trans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%v", err)
	}

	// 1.查询退款和通知信息
	data, err := n.db.QueryRefundInfoByDB(trans, refundID)
	if err != nil {
		trans.Rollback()
		return
	}
	if data.IsEmpty() {
		trans.Rollback()
		return qtask.Finish(n.c, taskID)
	}

	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}
	var returnFuncs []func(c interface{}) error
	var tasks []string
	//2.退货已完成，退款进入资金退款流程时，处理订单信息
	if data.GetInt("refund_status") == enum.RefundStatus.Refund {
		if err := n.db.UpdateRefundOvertime(trans, refundID); err != nil {
			trans.Rollback()
			return err
		}
	} else if data.GetInt("refund_status") == enum.RefundStatus.Return {

		//3.退货未完成时，处理未开始的退货为失败，查询出正在进行的退货，并判断是否所有退货都失败
		isAllFail, returnTasks, ReturningList, err := n.db.UpdateWaitReturnToFailByDB(trans, data.GetString("refund_id"))
		if err != nil {
			trans.Rollback()
			return err
		}

		//4.有进行中的退货，创建对应处理任务(是否创建任务，创建方法中判断)
		returnFuncs, err = n.task.CreateBatchTasks(trans, returnTasks, ReturningList)
		if err != nil {
			trans.Rollback()
			return err
		}

		//5.所有退货失败，处理退款记录并开启通知
		if isAllFail {
			tasks, err = n.db.DealForAllFailReturnByDB(trans, data)
			if err != nil {
				trans.Rollback()
				return err
			}
		}
	}

	//6.创建后续处理任务，合并回调函数
	methods, err := n.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}
	methods = append(methods, returnFuncs...)

	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	trans.Commit()
	return utils.CallBackFuncs(n.c, methods...)
}

// ReturnUnknownDeal 退货未知处理
func (n *OverTime) ReturnUnknownDeal(refundID int64, returnID int64, orderID int64, taskID int64) (err error) {
	trans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%v", err)
	}

	//1.锁结果未知的退货对应的订单、退款、退货记录，并获取通知信息
	data, err := n.db.LockByDB(trans, orderID, refundID, returnID)
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
		return
	}

	//2.请求vds查询退货结果
	result, err := jcsdk.QueryOrder(trans, data.GetString("down_channel_no"), data.GetString("return_id"), n.db.BuildServiceClass())
	if err != nil || result == nil {
		trans.Rollback()
		return fmt.Errorf("err:%v,result:%v", err, result)
	}

	var tasks []string
	switch result.Status {
	case enum.DeliveryResult.Success:
		//3.处理退货为成功，并获取退款总成功退货信息
		upRefundTasks, returnSuccessAmountInfo, err := n.db.DoForSuccessByDB(trans, data, result)
		if err != nil {
			trans.Rollback()
			return err
		}
		if returnSuccessAmountInfo.GetBool("finish") {
			//4.退款记录已完全退货，修改退款与订单信息
			tasks, err = n.db.DoForAllSuccessByDB(trans, data, returnSuccessAmountInfo, result)
			if err != nil {
				trans.Rollback()
				return err
			}
		}
		tasks = n.task.AppendTasks(tasks, upRefundTasks...)

	case enum.DeliveryResult.Failed:
		//5.退货失败,处理记录为失败，判断是否完全失败，完全失败处理退款并开启通知
		notifyTasks, err := n.db.DoForFailByDB(trans, data, result)
		if err != nil {
			trans.Rollback()
			return err
		}
		tasks = n.task.AppendTasks(tasks, notifyTasks...)

	case enum.DeliveryResult.Underway:
		//6.结果未知创建人工审核
		if err := n.db.CreateManualByDB(data); err != nil {
			trans.Rollback()
			return err
		}
	}

	//7.创建后续处理任务并处理回调
	methods, err := n.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}
	// 关闭task
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	trans.Commit()
	return utils.CallBackFuncs(n.c, methods...)
}
