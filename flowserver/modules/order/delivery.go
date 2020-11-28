package order

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/qtask"
)

// IDelivery 发货接口
type IDelivery interface {
	Start(deliveryID, taskID int64) error
	Finish(result *DeliveryResult) (orderID string, err error)
}

// Delivery 发货结构体
type Delivery struct {
	c    component.IContainer
	db   IDBDelivery
	task task.IQTask
}

// NewDelivery 构建Delivery
func NewDelivery(c component.IContainer, db IDBDelivery, task task.IQTask) *Delivery {
	return &Delivery{
		c:    c,
		db:   db,
		task: task,
	}
}

// Start 开始发货
func (d *Delivery) Start(deliveryID, taskID int64) error {

	trans, err := d.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 1.调用发货流程任务开始
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 2.锁订单和发货
	data, canFinishTask, err := d.db.CheckDeliveryAndLockOrderByDB(trans, deliveryID)
	if canFinishTask {
		if err := qtask.Finish(trans, taskID); err != nil {
			trans.Rollback()
			return err
		}
		trans.Commit()
	}

	if err != nil {
		return err
	}

	// 3.修改发货记录状态为正在发货
	if err = d.db.UpdateDeliveryStatusByDB(trans, deliveryID); err != nil {
		trans.Rollback()
		return err
	}

	// 4.调用发货系统发货任务
	input, err := d.db.DeliveryBuildParam(trans, data)
	if err != nil {
		trans.Rollback()
		return err
	}
	result, deliveryQueueFunc, err := jcsdk.CreateOrder(trans, input)
	if err != nil {
		trans.Rollback()
		return err
	}
	if data.GetInt64("delivery_id") != types.GetInt64(result.CoopOrderID) {
		trans.Rollback()
		return fmt.Errorf("发货订单号不相同")
	}

	// 5.结束发货流程任务
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}
	trans.Commit()

	return deliveryQueueFunc(d.c)
}

// Finish 完成发货
func (d *Delivery) Finish(result *DeliveryResult) (orderID string, err error) {
	// 1.查询订单和发货
	data, orderID, canDeliveryFinish, err := d.db.CheckOrderAndDeliveryByDB(result)
	if canDeliveryFinish {
		scallback, serr := jcsdk.SaveNotifyResult(d.c, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
		if serr != nil {
			return orderID, serr
		}
		if qerr := scallback(d.c); qerr != nil {
			return orderID, qerr
		}
	}
	if err != nil {
		return orderID, err
	}
	trans, err := d.c.GetRegularDB().Begin()
	if err != nil {
		return orderID, err
	}

	var tasks []string
	switch result.Status {
	case enum.DeliveryResult.Failed:
		// 2.发货失败处理
		tasks, err = d.db.DeliveryFailedByDB(trans, data)
		if err != nil {
			trans.Rollback()
			return orderID, err
		}

	case enum.DeliveryResult.Success:
		// 3.发货成功处理
		notifyID, taskList, err := d.db.DeliverySuccessByDB(trans, data)
		if err != nil {
			trans.Rollback()
			return orderID, err
		}
		data["notify_id"] = notifyID
		tasks = d.task.AppendTasks(tasks, taskList...)
	}

	//4.创建后续流程任务
	queueFuncs, err := d.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return orderID, err
	}

	// 5.发货系统发货完成函数
	scallback, err := jcsdk.SaveNotifyResult(d.c, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
	if err != nil {
		trans.Rollback()
		return orderID, err
	}
	queueFuncs = append(queueFuncs, scallback)
	trans.Commit()

	return orderID, utils.CallBackFuncs(d.c, queueFuncs...)
}
