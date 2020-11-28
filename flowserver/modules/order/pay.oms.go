package order

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// PayOms PayOms
type PayOms struct {
	*Pay
	*DBPay
	*task.QTask
}

// NewPayOms oms支付处理
func NewPayOms(c component.IContainer) *PayOms {
	ro := &PayOms{}
	ro.DBPay = NewDBPay(c)
	ro.QTask = task.NewQTask(c)
	ro.Pay = NewPay(c, ro, ro)
	return ro
}
