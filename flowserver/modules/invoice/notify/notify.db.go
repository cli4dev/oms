package notify

import (
	"fmt"
	"strings"
	"time"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/sas/sas/md5"
)

// IDBInvoiceNotify 开票通知数据层接口
type IDBInvoiceNotify interface {
	CheckOrderAndInvoiceByDB(invoiceID int64) (types.XMap, bool, error)
	StartNotifyInvoiceByDB(invoiceID int64) error
	RemoteNotify(c component.IContainer, data types.XMap, param types.XMap) (content string, err error)
	FailedByDB(param map[string]interface{}) error
	BuildNotifyParamByDB(notify types.XMap) (map[string]interface{}, error)
	SuccessByDB(param map[string]interface{}) error
}

//DBInvoiceNotify DBInvoiceNotify
type DBInvoiceNotify struct {
	c component.IContainer
}

//NewDBInvoiceNotify 构建DBInvoiceNotify
func NewDBInvoiceNotify(c component.IContainer) *DBInvoiceNotify {
	return &DBInvoiceNotify{c: c}
}

//CheckOrderAndInvoiceByDB 通知检查发票状态和订单
func (n *DBInvoiceNotify) CheckOrderAndInvoiceByDB(invoiceID int64) (types.XMap, bool, error) {
	db := n.c.GetRegularDB()

	//1.检查开票申请记录
	rows, q, a, err := db.Query(sql.NotifyCheckInvoiceApplyInfo, map[string]interface{}{
		"invoice_id": invoiceID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查开票申请记录发生异常,err:%v,sql:%v,args:%v", err, q, a)
	}
	if rows.IsEmpty() {
		return nil, true, fmt.Errorf("开票记录不存在,invoice_id:%d", invoiceID)
	}

	return rows.Get(0), false, nil
}

// StartNotifyInvoiceByDB 开始通知
func (n *DBInvoiceNotify) StartNotifyInvoiceByDB(invoiceID int64) error {
	db := n.c.GetRegularDB()
	count, q, a, err := db.Execute(sql.StartInvoiceNotify, map[string]interface{}{
		"invoice_id": invoiceID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("开始通知,修改开票通知状态失败,err:%v,sql:%v,args:%v", err, q, a)
	}
	return nil
}

// RemoteNotify 通知
func (n *DBInvoiceNotify) RemoteNotify(c component.IContainer, data types.XMap, param types.XMap) (content string, err error) {

	content, status, err := utils.HTTPRequest(param, "post", types.GetString(data.GetString("notify_url")))
	if err != nil || status != context.ERR_OK || !strings.Contains(content, "success") {
		return "", fmt.Errorf("调用通知接口错误,err:%v", err)
	}
	return
}

// FailedByDB 通知失败,修改通知次数,记录失败信息(超过最大次数,通知失败,并结束通知任务)
func (n *DBInvoiceNotify) FailedByDB(param map[string]interface{}) error {
	db := n.c.GetRegularDB()
	row, sqlStr, args, err := db.Execute(sql.FailedInvoiceNotify, param)
	if err != nil || row != 1 {
		return fmt.Errorf("修改发票结果通知状态发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// SuccessByDB 通知成功,修改发票结果通知状态成功,结束通知任务
func (n *DBInvoiceNotify) SuccessByDB(param map[string]interface{}) error {
	db := n.c.GetRegularDB()
	row, q, a, err := db.Execute(sql.SuccessNotify, param)
	if err != nil || row != 1 {
		return fmt.Errorf("修改发票通知状态为成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, q, a)
	}

	return nil
}

// BuildNotifyParamByDB 构建通知参数
func (n *DBInvoiceNotify) BuildNotifyParamByDB(notify types.XMap) (map[string]interface{}, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(notify.GetInt("invoice_status"), errorcode.RequestFlowType.Invoice, notify.GetString("fail_code"), notify.GetString("fail_msg"))

	param := map[string]interface{}{
		"channel_no":  notify.GetString("channel_no"),
		"request_no":  notify.GetString("request_no"),
		"invoice_id":  notify.GetInt64("invoice_id"),
		"invoice_no":  notify.GetString("invoice_no"),
		"can_red":     notify.GetInt("can_red"),
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
		"timestamp":   time.Now().Format("20060102150405"),
	}
	sign, err := md5.Sign(enum.Group, notify.GetString("channel_no"), param, "&", "=")
	if err != nil {
		return nil, err
	}
	param["sign"] = sign
	return param, nil
}
