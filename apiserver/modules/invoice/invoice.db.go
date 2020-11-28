package invoice

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// IDBInvoice 开票数据层接口
type IDBInvoice interface {
	CheckOrderAndDeliveryByDB(channelNO, requestNO, invoiceNO string, amount, deductAmount float64) (types.XMap, error)
	CheckInvoiceByDB(db db.IDBExecuter, channelNO, invoiceNO string) (types.XMap, error)
	LockOrderByDB(trans db.IDBTrans, orderID int64) error
	CreateInvoiceByDB(dbTrans db.IDBTrans, info *RequestInfo, order types.XMap) (types.XMap, []string, error)
	BuildResultByDB(invoice types.XMap, input *RequestInfo) types.XMap
	BuildQueryByDB(invoice types.XMap, input *QueryInfo) types.XMap
}

// DBInvoice 开票数据层
type DBInvoice struct {
	c component.IContainer
}

//NewDBInvoice 创建开票对象
func NewDBInvoice(c component.IContainer) *DBInvoice {
	return &DBInvoice{
		c: c,
	}
}

// CheckOrderAndDeliveryByDB 检查订单和发货
func (i *DBInvoice) CheckOrderAndDeliveryByDB(channelNO, requestNO, invoiceNO string, amount, deductAmount float64) (types.XMap, error) {
	db := i.c.GetRegularDB()

	//1.获取订单相关信息
	rows, q, a, err := db.Query(sql.SQLCheckOrderByInvoice, map[string]interface{}{
		"request_no":    requestNO,
		"channel_no":    channelNO,
		"amount":        amount,
		"deduct_amount": deductAmount,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单信息错误(err:%v),sql:%s,param:%+v,rows:%+v", err, q, a, rows)
	}

	if rows.IsEmpty() {
		return nil, context.NewErrorf(errorcode.INVOICE_NO_EXIST_ORDER.Code, errorcode.INVOICE_NO_EXIST_ORDER.Msg)
	}
	order := rows.Get(0)

	if order.GetFloat64("sell_amount") <= 0 || !order.GetBool("can_sell_amount") {
		return nil, context.NewErrorf(errorcode.INVOICE_AMOUNT_ERROR.Code, errorcode.INVOICE_AMOUNT_ERROR.Msg, amount, order.GetFloat64("sell_amount"))
	}

	if !order.GetBool("can_deduct_amount") {
		return nil, context.NewErrorf(errorcode.INVOICE_DEDUCT_AMOUNT_ERROR.Code, errorcode.INVOICE_DEDUCT_AMOUNT_ERROR.Msg, deductAmount, order.GetFloat64("deduct_amount"))
	}

	//2.获取发货信息
	deliverys, q, a, err := db.Query(sql.SQLCheckDeliveryByInvoice, map[string]interface{}{
		"order_id": order.GetString("order_id"),
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货信息错误(err:%v),sql:%s,param:%+v,rows:%+v", err, q, a, rows)
	}

	if deliverys.IsEmpty() {
		return nil, context.NewErrorf(errorcode.INVOICE_NO_EXIST_DELIVERY.Code, errorcode.INVOICE_NO_EXIST_DELIVERY.Msg)
	}

	if order.GetInt("invoice_type") == enum.InvoiceType.NotInvoice {
		return nil, context.NewErrorf(errorcode.INVOICE_DOWN_NO_INVOICE.Code, errorcode.INVOICE_DOWN_NO_INVOICE.Msg)
	}
	if deliverys.Get(0).GetInt("invoice_type") == enum.InvoiceType.NotInvoice {
		return nil, context.NewErrorf(errorcode.INVOICE_UP_NO_INVOICE.Code, errorcode.INVOICE_UP_NO_INVOICE.Msg)
	}
	if order.GetBool("can_invoice") {
		return nil, context.NewErrorf(errorcode.INVOICE_OVERTIME.Code, errorcode.INVOICE_OVERTIME.Msg)
	}
	delivery := deliverys.Get(0)
	delete(delivery, "invoice_type")
	order.Merge(delivery)
	fmt.Printf("order:%+v", order)
	return order, nil
}

// CheckInvoiceByDB 检查开票是否存在
func (i *DBInvoice) CheckInvoiceByDB(db db.IDBExecuter, channelNO, invoiceNO string) (types.XMap, error) {
	invoices, _, _, err := db.Query(sql.SQLCheckInvoice, map[string]interface{}{
		"invoice_no": invoiceNO,
		"channel_no": channelNO,
	})
	if err != nil {
		return nil, fmt.Errorf("检查开票是否存在错误,err:%v", err)
	}
	return invoices.Get(0), nil
}

// LockOrderByDB 锁定单
func (i *DBInvoice) LockOrderByDB(trans db.IDBTrans, orderID int64) error {
	//1.锁定订单
	orders, _, _, err := trans.Query(sql.SQLLockOrderByInvoice, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || orders.IsEmpty() {
		return fmt.Errorf("锁定订单错误,err:%v,orderID:%v", err, orderID)
	}

	//2.检查是否存在非失败退款申请记录
	refunds, _, _, err := trans.Query(sql.SQLCheckRefundByInvoice, map[string]interface{}{"order_id": orderID})
	if err != nil {
		return fmt.Errorf("开票检查退款是否存在发生异常，err:%v,order_id:%v", err, orderID)
	}

	if !refunds.IsEmpty() {
		return context.NewErrorf(errorcode.INVOICE_EXIST_REFUND.Code, "该订单存在非失败退款申请记录,不允许开票")
	}
	return nil
}

//CreateInvoiceByDB 创建开票申请
func (i *DBInvoice) CreateInvoiceByDB(dbTrans db.IDBTrans, info *RequestInfo, order types.XMap) (types.XMap, []string, error) {
	param, err := types.Struct2Map(info)
	if err != nil {
		return nil, nil, err
	}
	order.MergeMap(param)
	count, q, a, err := dbTrans.Execute(sql.SQLCreateInvocieInfo, order)
	if err != nil || count < 0 {
		return nil, nil, fmt.Errorf("创建开票记录发生异常(err:%v),sql:%s,param:%+v,rows:%+v", err, q, a, count)
	}
	invoice, err := i.CheckInvoiceByDB(dbTrans, info.ChannelNO, info.InvoiceNO)
	if err != nil {
		return nil, nil, err
	}
	order.Merge(invoice)
	return invoice, []string{task.TaskType.InvoiceTask}, nil
}

// BuildResultByDB 构建返回参数
func (i *DBInvoice) BuildResultByDB(invoice types.XMap, input *RequestInfo) types.XMap {
	fmt.Printf("invoice:%+v", invoice)
	status, code, msg := errorcode.SetFlowRecordStatus(invoice.GetInt("status"), errorcode.RequestFlowType.Invoice, invoice.GetString("fail_code"), invoice.GetString("fail_msg"))
	return map[string]interface{}{
		"channel_no":  input.ChannelNO,
		"request_no":  input.RequestNO,
		"invoice_id":  invoice.GetInt64("invoice_id"),
		"invoice_no":  input.InvoiceNO,
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	}
}

// BuildQueryByDB 构建返回参数
func (i *DBInvoice) BuildQueryByDB(invoice types.XMap, input *QueryInfo) types.XMap {
	fmt.Printf("invoice:%+v", invoice)

	status, code, msg := errorcode.SetFlowRecordStatus(invoice.GetInt("status"), errorcode.RequestFlowType.Invoice, invoice.GetString("fail_code"), invoice.GetString("fail_msg"))
	return map[string]interface{}{
		"channel_no":  input.ChannelNO,
		"request_no":  input.RequestNO,
		"invoice_id":  invoice.GetInt64("invoice_id"),
		"invoice_no":  input.InvoiceNO,
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	}
}
