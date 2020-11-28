package pay

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
)

// IDBInvoicePay 数据层接口
type IDBInvoicePay interface {
	CheckInvoiceForPaymentByDB(invoiceID, orderID int64) (types.XMap, error)
	InvoicePaymentSuccessByDB(trans db.IDBTrans, invoiceID int64) error
	InvoicePayByDB(trans db.IDBTrans, data types.XMap) error
}

//DBInvoicePay 开票记账数据层
type DBInvoicePay struct {
	c component.IContainer
}

//NewDBInvoicePay 构建DBInvoicePay
func NewDBInvoicePay(c component.IContainer) *DBInvoicePay {
	return &DBInvoicePay{c: c}
}

// CheckInvoiceForPaymentByDB 开票支付检查
func (d *DBInvoicePay) CheckInvoiceForPaymentByDB(invoiceID, orderID int64) (types.XMap, error) {
	db := d.c.GetRegularDB()
	datas, q, a, err := db.Query(sql.SQLCheckInvoiceForPayment, map[string]interface{}{
		"invoice_id": invoiceID,
	})
	if err != nil {
		return nil, fmt.Errorf("开票支付检查信息发生异常,err:%v,sql:%v,args:%v,cnt:%v", err, q, a, datas.Len())
	}
	if datas.IsEmpty() {
		return nil, nil
	}
	data := datas.Get(0)
	return data, nil
}

// InvoicePaymentSuccessByDB 开票支付成功
func (d *DBInvoicePay) InvoicePaymentSuccessByDB(trans db.IDBTrans, invoiceID int64) error {
	count, q, a, err := trans.Execute(sql.SQLInvoicePaymentSuccess, map[string]interface{}{
		"invoice_id": invoiceID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("开票支付成功发生异常,err:%v,sql:%v,args:%v,cnt:%v", err, q, a, count)
	}
	return nil
}

// InvoicePayByDB 开票记账
func (d *DBInvoicePay) InvoicePayByDB(trans db.IDBTrans, data types.XMap) error {
	return nil
}
