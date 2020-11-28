package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//Refund 订单积分退款处理结构体
type Refund struct {
	c    component.IContainer
	db   *DBRefund
	task task.IQTask
}

// NewRefund 构建Refund
func NewRefund(c component.IContainer) *Refund {
	return &Refund{
		c:    c,
		db:   NewDBRefund(c),
		task: task.NewQTask(c),
	}
}

//RefundFDCreate 退款记账创建
func (o *Refund) RefundFDCreate(orderID, refundID, taskID int64) error {
	if err := qtask.Processing(o.c, taskID); err != nil {
		return err
	}

	//1.检查订单、退款及记账记录
	record, refPointInfo, err := o.db.RefundFDCheck(orderID, refundID)
	if err != nil {
		return err
	}
	if record != nil {
		if qerr := qtask.Finish(o.c, taskID); qerr != nil {
			return qerr
		}
		return nil
	}

	//2.查询记账渠道
	jfChannel, err := o.db.GetJFUpChannel()
	if err != nil {
		return err
	}

	//3.创建记账记录
	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	notify, tasks, err := o.db.CreateRefundFDRecord(dbTrans, orderID, refundID, jfChannel.GetString("channel_no"), refPointInfo)
	if err != nil {
		dbTrans.Rollback()
		return err
	}

	if notify != nil {
		methods, err := o.task.CreateTask(dbTrans, tasks, notify)
		if err != nil {
			dbTrans.Rollback()
			return err
		}
		err = qtask.Finish(dbTrans, taskID)
		if err != nil {
			dbTrans.Rollback()
			return err
		}

		dbTrans.Commit()
		return utils.CallBackFuncs(o.c, methods...)
	}

	err = qtask.Finish(dbTrans, taskID)
	if err != nil {
		dbTrans.Rollback()
		return err
	}

	dbTrans.Commit()
	return nil
}
