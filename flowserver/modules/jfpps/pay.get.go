package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// JFPay JFPay
type JFPay struct {
	*order.Pay
	c component.IContainer
}

// NewJFPay oms支付处理
func NewJFPay(c component.IContainer) *JFPay {
	ro := &JFPay{c: c}
	ro.Pay = order.NewPay(c, NewDBJFPay(c), task.NewQTask(c))
	return ro
}

//UpPay 上游支付
func (o *JFPay) UpPay(deliveryID, taskID int64) error {
	// 开启任务
	if err := qtask.Processing(o.c, taskID); err != nil {
		return err
	}

	// 关闭任务
	return qtask.Finish(o.c, taskID)
}
