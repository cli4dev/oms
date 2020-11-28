package invoice

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// Oms InvoiceOms结构体
type Oms struct {
	*Invoice
	*DBInvoice
	*task.QTask
}

// NewInvoiceOms oms订单处理
func NewInvoiceOms(c component.IContainer) *Oms {
	oo := &Oms{}
	oo.QTask = task.NewQTask(c)
	oo.DBInvoice = NewDBInvoice(c)
	oo.Invoice = NewInvoice(c, oo, oo)
	return oo
}
