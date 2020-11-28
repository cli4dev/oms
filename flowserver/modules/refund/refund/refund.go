package refund

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask/qtask"
)

// IRefund 退款接口
type IRefund interface {
	UpRefund(returnID string, taskID int64) (err error)
	DownRefund(refundID string, taskID int64) (err error)
	OrderRefund(orderID int64, taskID int64) (err error)
}

// Refund 退款对象
type Refund struct {
	c  component.IContainer
	db IDBRefund
}

//NewRefund 创建上游退款对象
func NewRefund(c component.IContainer, db IDBRefund) *Refund {
	return &Refund{
		c:  c,
		db: db,
	}
}

// UpRefund Refund 执行上游退款
func (n *Refund) UpRefund(returnID string, taskID int64) (err error) {
	// 开启任务
	if err = qtask.Processing(n.c, taskID); err != nil {
		return
	}

	// 检查状态
	data, err := n.db.CheckUpRefundStatusByDB(returnID)
	if context.GetCode(err) == errorcode.UP_REFUND_STATUS_ERROR.Code {
		// 关闭任务
		if qerr := qtask.Finish(n.c, taskID); qerr != nil {
			return
		}
		return
	}
	if err != nil {
		return
	}

	// 开启事务
	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return
	}

	// 锁退款
	if err = n.db.LockByDB(dbTrans, data.GetString("refund_id")); err != nil {
		dbTrans.Rollback()
		return
	}

	// 退款
	if err = n.db.UpRefundByDB(dbTrans, data, n.c.GetPlatName()); err != nil {
		dbTrans.Rollback()
		return
	}

	// 更新信息
	if err = n.db.UpdateUpRefundToSuccessByDB(returnID, data.GetString("refund_id"), dbTrans); err != nil {
		dbTrans.Rollback()
		return
	}

	dbTrans.Commit()
	// 关闭任务
	return qtask.Finish(n.c, taskID)
}

// DownRefund 执行退款
func (n *Refund) DownRefund(refundID string, taskID int64) (err error) {
	// 开启任务
	if err = qtask.Processing(n.c, taskID); err != nil {
		return
	}

	// 检查状态
	data, err := n.db.CheckDownRefundStatusByDB(refundID)
	if context.GetCode(err) == errorcode.DOWN_REFUND_STATUS_ERROR.Code {
		// 关闭任务
		if qerr := qtask.Finish(n.c, taskID); qerr != nil {
			return qerr
		}
	}
	if err != nil {
		return
	}

	// 开启事务
	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%+v", err)
	}

	// 锁退款
	if err = n.db.LockByDB(dbTrans, refundID); err != nil {
		dbTrans.Rollback()
		return
	}

	// 调用退款函数
	if err = n.db.DownRefundByDB(dbTrans, data, n.c.GetPlatName(), refundID); err != nil {
		dbTrans.Rollback()
		return
	}

	// 更新信息
	if err = n.db.UpdateDownRefundToSuccessByDB(refundID, dbTrans); err != nil {
		dbTrans.Rollback()
		return
	}

	dbTrans.Commit()

	// 关闭任务
	return qtask.Finish(n.c, taskID)
}

// OrderRefund 订单失败退款
func (n *Refund) OrderRefund(orderID int64, taskID int64) (err error) {
	// 开启任务
	if err = qtask.Processing(n.c, taskID); err != nil {
		return
	}

	// 检查订单
	data, err := n.db.CheckOrderStatusByDB(orderID)
	if context.GetCode(err) == errorcode.ORDER_REFUND_STATUS_ERROR.Code {
		// 关闭任务
		if qerr := qtask.Finish(n.c, taskID); qerr != nil {
			return
		}
		return
	}
	if err != nil {
		return
	}

	// 开启事务
	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%+v", err)
	}

	// 更改状态
	if err = n.db.UpdateOrderRefundStatusByDB(dbTrans, data.GetString("order_id")); err != nil {
		dbTrans.Rollback()
		return
	}

	// 调用退款函数
	if err = n.db.DownRefundByDB(dbTrans, data, n.c.GetPlatName(), fmt.Sprintf("r%v", data.GetString("order_id"))); err != nil {
		dbTrans.Rollback()
		return
	}

	// 后续操作
	dbTrans.Commit()

	// 关闭任务
	return qtask.Finish(n.c, taskID)
}
