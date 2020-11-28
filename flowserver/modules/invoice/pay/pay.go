package pay

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
)

//IInvoicePay 请求上游申请开票接口
type IInvoicePay interface {
	Pay(orderID, invoiceID, taskID int64) error
}

//InvoicePay 请求上游申请开票结构体
type InvoicePay struct {
	c  component.IContainer
	db IDBInvoicePay
}

//NewInvoicePay 构建InvoicePay
func NewInvoicePay(c component.IContainer, db IDBInvoicePay) *InvoicePay {
	return &InvoicePay{
		c:  c,
		db: db,
	}
}

// Pay 开票记账
func (i *InvoicePay) Pay(orderID, invoiceID, taskID int64) error {
	data, err := i.db.CheckInvoiceForPaymentByDB(invoiceID, orderID)
	if err != nil {
		return err
	}
	if data.IsEmpty() {
		return qtask.Finish(i.c, taskID)
	}

	if err := qtask.Processing(i.c, taskID); err != nil {
		return err
	}

	trans, err := i.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	if err := i.db.InvoicePaymentSuccessByDB(trans, invoiceID); err != nil {
		trans.Rollback()
		return err
	}

	if err := i.db.InvoicePayByDB(trans, data); err != nil {
		trans.Rollback()
		return err
	}
	trans.Commit()
	return qtask.Finish(i.c, taskID)
}
