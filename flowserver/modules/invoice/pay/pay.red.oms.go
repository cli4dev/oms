package pay

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// InvoiceRedPayOms oms开票支付处理
type InvoiceRedPayOms struct {
	*InvoicePay
	*DBInvoicePay
	c component.IContainer
}

// NewInvoiceRedPayOms 构建InvoicePayPayOms
func NewInvoiceRedPayOms(c component.IContainer) *InvoiceRedPayOms {
	bo := &InvoiceRedPayOms{c: c}
	bo.DBInvoicePay = NewDBInvoicePay(c)
	bo.InvoicePay = NewInvoicePay(c, bo)
	return bo
}

// InvoicePayByDB 开票记账
func (d *InvoiceRedPayOms) InvoicePayByDB(trans db.IDBTrans, data types.XMap) error {
	return nil
}
