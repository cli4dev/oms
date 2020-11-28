package refund

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// Oms RefundOms
type Oms struct {
	*Refund
	*DBRefund
	*task.QTask
}

// NewRefundOms oms退款处理
func NewRefundOms(c component.IContainer) *Oms {
	ro := &Oms{}
	ro.DBRefund = NewDBRefund(c)
	ro.QTask = task.NewQTask(c)
	ro.Refund = NewRefund(c, ro, ro)
	return ro
}
