package start

import (
	"github.com/micro-plat/hydra/component"
)

// InvoiceStartOms oms绑定处理
type InvoiceStartOms struct {
	*Invoice
	*DBInvoice
	c component.IContainer
}

// NewInvoiceStartOms 构建InvoiceStartOms
func NewInvoiceStartOms(c component.IContainer) *InvoiceStartOms {
	bo := &InvoiceStartOms{c: c}
	bo.DBInvoice = NewDBInvoice(c)
	bo.Invoice = NewInvoice(c, bo)
	return bo
}
