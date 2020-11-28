package notify

import (
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/sas/sas/md5"
)

// InvoiceRedNoitfyOms oms开票通知处理
type InvoiceRedNoitfyOms struct {
	*InvoiceNotify
	*DBInvoiceNotify
}

// NewInvoiceRedNoitfyOms 构建NewInvoiceRedNoitfyOms
func NewInvoiceRedNoitfyOms(c component.IContainer) *InvoiceRedNoitfyOms {
	bo := &InvoiceRedNoitfyOms{}
	bo.DBInvoiceNotify = NewDBInvoiceNotify(c)
	bo.InvoiceNotify = NewInvoiceNotify(c, bo)
	return bo
}

// BuildNotifyParamByDB 构建通知参数
func (n *InvoiceRedNoitfyOms) BuildNotifyParamByDB(notify types.XMap) (map[string]interface{}, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(notify.GetInt("invoice_status"), errorcode.RequestFlowType.Invoice, notify.GetString("fail_code"), notify.GetString("fail_msg"))

	param := map[string]interface{}{
		"channel_no": notify.GetString("channel_no"),
		"request_no": notify.GetString("request_no"),
		"invoice_id": notify.GetInt64("orig_invoice_id"),
		"invoice_no": notify.GetString("invoice_no"),
		"can_red":    0,
		"status":     status,
		"code":       code,
		"msg":        msg,
		"timestamp":  time.Now().Format("20060102150405"),
	}
	sign, err := md5.Sign(enum.Group, notify.GetString("channel_no"), param, "&", "=")
	if err != nil {
		return nil, err
	}
	param["sign"] = sign
	return param, nil
}
