package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"
)

// JFRefund JFRefund
type JFRefund struct {
	*refund.Refund
	c component.IContainer
}

// NewJFRefund oms支付处理
func NewJFRefund(c component.IContainer) *JFRefund {
	ro := &JFRefund{c: c}
	ro.Refund = refund.NewRefund(c, NewDBJFRefund(c))
	return ro
}

//UpRefund 积分获取上游退款
func (o *JFRefund) UpRefund(returnID string, taskID int64) (err error) {
	// 开启任务
	if err = qtask.Processing(o.c, taskID); err != nil {
		return
	}

	// 关闭任务
	return qtask.Finish(o.c, taskID)
}
