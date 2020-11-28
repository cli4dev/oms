package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask/qtask"
)

// IPay 支付接口
type IPay interface {
	DownPay(orderID, taskID int64) error
	UpPay(deliveryID, taskID int64) error
}

// Pay 下游支付完成结构体
type Pay struct {
	c    component.IContainer
	db   IDBPay
	task task.IQTask
}

// NewPay 构建Pay
func NewPay(c component.IContainer, db IDBPay, task task.IQTask) *Pay {
	return &Pay{
		c:    c,
		db:   db,
		task: task,
	}
}

// DownPay 下游支付流程
func (p *Pay) DownPay(orderID, taskID int64) error {

	// 1.下游支付检查订单
	tasks, data, err := p.db.CheckOrderForDownPayByDB(orderID)
	if err != nil {
		return err
	}
	if data == nil {
		if err := qtask.Finish(p.c, taskID); err != nil {
			return err
		}
		return context.NewErrorf(context.ERR_NO_CONTENT, "下游支付时,订单不存在或已完成支付,order_id:%v", orderID)
	}

	trans, err := p.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 2.开始任务
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 3.成功，修改订单支付状态为成功，支付状态为绑定发货状态，将绑定状态设置为正在绑定,创建绑定流程任务
	if err = p.db.OrderDownPaySuccessByDB(trans, orderID); err != nil {
		trans.Rollback()
		return err
	}

	// 4.支付
	if err := p.db.DownChannelPayByDB(trans, data); err != nil {
		trans.Rollback()
		return err
	}

	// 5.支付流程处理
	queueFuncs, err := p.task.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}

	// 6.业务处理成功，修改任务状态为完成(任务不再放入队列)
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}
	trans.Commit()

	return utils.CallBackFuncs(p.c, queueFuncs...)
}

// UpPay 上游支付
func (p *Pay) UpPay(deliveryID, taskID int64) error {

	// 1.检查订单和订单发货
	data, canFinishTask, err := p.db.CheckOrderAndDeliveryByDB(deliveryID)
	if canFinishTask {
		if err := qtask.Finish(p.c, taskID); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	trans, err := p.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 2.开始任务
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 3.支付成功,修改发货记录的支付状态
	if err = p.db.UpPaySuccessByDB(trans, data); err != nil {
		trans.Rollback()
		return err
	}

	// 4.上游支付处理 支付系统支付,销售金额，佣金，服务费
	if err = p.db.UpChannelPayByDB(trans, data); err != nil {
		trans.Rollback()
		return err
	}

	// 5.任务结束
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}
	trans.Commit()
	return nil
}
