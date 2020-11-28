package start

import (
	"encoding/json"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"
)

// InvoiceRedStartOms oms绑定处理
type InvoiceRedStartOms struct {
	c component.IContainer
	*Invoice
	*DBInvoice
}

// NewInvoiceRedStartOms 构建InvoiceRedStartOms
func NewInvoiceRedStartOms(c component.IContainer) *InvoiceRedStartOms {
	bo := &InvoiceRedStartOms{c: c}
	bo.DBInvoice = NewDBInvoice(c)
	bo.Invoice = NewInvoice(c, bo)
	return bo
}

// BuildOrderRequestParamByDB 构建发货系统参数
func (d *InvoiceRedStartOms) BuildOrderRequestParamByDB(trans db.IDBTrans, info types.IXMap, invoiceID int64) (*model.OrderCreateParam, error) {

	extendInfo := map[string]interface{}{
		"orig_invoice_id": info.GetString("orig_invoice_id"),
		"order_id":        info.GetString("order_id"),
	}
	bytes, err := json.Marshal(extendInfo)
	if err != nil {
		return nil, err
	}
	return &model.OrderCreateParam{
		CoopID:        info.GetString("down_channel_no"),
		CoopOrderID:   types.GetString(invoiceID),
		ChannelNo:     types.DecodeString(info.GetInt("invoice_method"), 1, info.GetString("down_channel_no"), info.GetString("up_channel_no")),
		ServiceClass:  43,
		CarrierNo:     info.GetString("carrier_no"),
		ProductFace:   info.GetInt("sell_price"),
		ProductNum:    info.GetInt("num"),
		NotifyURL:     queue.InvoiceRedFinish.GetName(d.c.GetPlatName()),
		OrderTimeout:  259200,
		RequestParams: string(bytes),
	}, nil
}
