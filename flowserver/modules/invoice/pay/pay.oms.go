package pay

import (
	"github.com/micro-plat/hydra/component"
)

// InvoicePayOms oms开票支付处理
type InvoicePayOms struct {
	*InvoicePay
	*DBInvoicePay
	c component.IContainer
}

// NewInvoicePayOms 构建InvoicePayPayOms
func NewInvoicePayOms(c component.IContainer) *InvoicePayOms {
	bo := &InvoicePayOms{c: c}
	bo.DBInvoicePay = NewDBInvoicePay(c)
	bo.InvoicePay = NewInvoicePay(c, bo)
	return bo
}
