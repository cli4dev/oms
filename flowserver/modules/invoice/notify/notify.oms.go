package notify

import "github.com/micro-plat/hydra/component"

// InvoiceNoitfyOms oms开票通知处理
type InvoiceNoitfyOms struct {
	*InvoiceNotify
	*DBInvoiceNotify
}

// NewInvoiceNoitfyOms 构建NewInvoiceNoitfyOms
func NewInvoiceNoitfyOms(c component.IContainer) *InvoiceNoitfyOms {
	bo := &InvoiceNoitfyOms{}
	bo.DBInvoiceNotify = NewDBInvoiceNotify(c)
	bo.InvoiceNotify = NewInvoiceNotify(c, bo)
	return bo
}
