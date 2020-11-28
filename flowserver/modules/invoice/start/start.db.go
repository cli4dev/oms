package start

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"
)

// IDBInvoice 数据层接口
type IDBInvoice interface {
	CheckOrderAndInvoiceByDB(orderID, invoiceID int64) (types.XMap, bool, error)
	StartInvoiceByDB(trans db.IDBTrans, invoiceID int64) error
	BuildOrderRequestParamByDB(trans db.IDBTrans, info types.IXMap, invoiceID int64) (*model.OrderCreateParam, error)
}

//DBInvoice 请求上游申请开票数据层
type DBInvoice struct {
	c component.IContainer
}

//NewDBInvoice 构建DBInvoice
func NewDBInvoice(c component.IContainer) *DBInvoice {
	return &DBInvoice{c: c}
}

//CheckOrderAndInvoiceByDB 检查订单和申请
func (d *DBInvoice) CheckOrderAndInvoiceByDB(orderID, invoiceID int64) (types.XMap, bool, error) {

	//1.检查订单
	db := d.c.GetRegularDB()

	//2.获取相关数据
	info, q, a, err := db.Query(sql.SQLGetDeliveryInfo, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查发货信息发生异常,err:%v,sql:%v,args:%v", err, q, a)
	}
	if info.IsEmpty() {
		return nil, true, fmt.Errorf("发货记录不存在，order_id:%d", orderID)
	}
	data := info.Get(0)
	//3.检查申请
	invoices, q, a, err := db.Query(sql.SQLCheckInvoiceInfo, map[string]interface{}{
		"invoice_id": invoiceID,
	})

	if err != nil {
		return nil, false, fmt.Errorf("检查发票申请信息发生异常,err:%v,sql:%v,args:%v", err, q, a)
	}

	if invoices.IsEmpty() {
		return nil, true, fmt.Errorf("申请记录不存在或状态不对,err:%v,sql:%v,args:%v", err, q, a)
	}
	data.Merge(invoices.Get(0))
	return data, false, nil
}

// StartInvoiceByDB 修改发票申请表,申请状态由等待审核变为正在审核
func (d *DBInvoice) StartInvoiceByDB(trans db.IDBTrans, invoiceID int64) error {
	count, q, a, err := trans.Execute(sql.SQLStartInvoice, map[string]interface{}{
		"invoice_id": invoiceID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("修改发票申请信息发生异常,err:%v,sql:%v,args:%v,cnt:%v", err, q, a, count)
	}
	return nil
}

// BuildOrderRequestParamByDB 构建发货系统参数
func (d *DBInvoice) BuildOrderRequestParamByDB(trans db.IDBTrans, info types.IXMap, invoiceID int64) (*model.OrderCreateParam, error) {

	extendInfo := map[string]interface{}{
		"order_id":      info.GetString("order_id"),
		"tele_phone":    info.GetString("tele_phone"),
		"invoice_title": info.GetString("invoice_title"),
		"tax_no":        info.GetString("tax_no"),
		"address":       info.GetString("address"),
		"bank_account":  info.GetString("bank_account"),
		"bank_name":     info.GetString("bank_name"),
		"amount":        info.GetString("amount"),
		"push_type":     info.GetString("push_type"),
		"push_phone_no": info.GetString("push_phone_no"),
		"deduct_amount": info.GetString("deduct_amount"),
		"push_email":    info.GetString("push_email"),
		"product_id":    info.GetString("up_product_id"),
	}
	bytes, err := json.Marshal(extendInfo)
	if err != nil {
		return nil, err
	}

	return &model.OrderCreateParam{
		CoopID:        info.GetString("down_channel_no"),
		CoopOrderID:   types.GetString(invoiceID),
		ChannelNo:     types.DecodeString(info.GetString("invoice_channel_no"), "", info.GetString("up_channel_no"), info.GetString("invoice_channel_no")),
		ServiceClass:  42,
		CarrierNo:     info.GetString("carrier_no"),
		ProductFace:   info.GetInt("sell_price"),
		ProductNum:    info.GetInt("num"),
		NotifyURL:     queue.InvoiceFinish.GetName(d.c.GetPlatName()),
		OrderTimeout:  259200,
		RequestParams: string(bytes),
	}, nil
}
