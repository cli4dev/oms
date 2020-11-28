package finish

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

//IDBFinish 完成开票数据层接口
type IDBFinish interface {
	CheckOrderAndInvoiceByDB(param map[string]interface{}) (types.XMap, bool, []string, error)
	FinishInvoiceByDB(trans db.IDBTrans, param types.XMap) error
}

//DBFinish 完成开票数据层
type DBFinish struct {
	c component.IContainer
}

//NewDBFinish 构建DBFinish
func NewDBFinish(c component.IContainer) *DBFinish {
	return &DBFinish{c: c}
}

// CheckOrderAndInvoiceByDB 1.检查订单和开票申请
func (f *DBFinish) CheckOrderAndInvoiceByDB(param map[string]interface{}) (types.XMap, bool, []string, error) {
	db := f.c.GetRegularDB()
	//1.检查开票申请
	rows, q, a, err := db.Query(sql.CheckInvoicing, map[string]interface{}{"invoice_id": param["coop_order_id"]})
	if err != nil {
		return nil, false, nil, fmt.Errorf("检查开票中信息发生异常,err:%v,sql:%v,args:%v", err, q, a)
	}

	if rows.IsEmpty() {
		return nil, true, nil, fmt.Errorf("开票信息不存在,invoice_id:%d", types.GetInt64(param["coop_order_id"]))
	}

	// 2.检查订单
	invoice := rows.Get(0)
	data, q, a, err := db.Query(sql.CheckFinishOrderInfo, invoice)
	if err != nil {
		return nil, false, nil, fmt.Errorf("检查订单信息发生异常,err:%v,sql:%v,args:%v", err, q, a)
	}

	if data.IsEmpty() {
		return nil, true, nil, fmt.Errorf("订单记录不存在，order_id:%d", invoice.GetInt64("order_id"))
	}
	info := data.Get(0)
	tasks := []string{}
	if invoice.GetString("notify_url") != "" {
		tasks = []string{task.TaskType.InvoiceNotifyTask}
	}

	fmt.Printf("param:%+v", param)
	if strings.EqualFold(types.GetString(param["status"]), enum.InvoiceSuccessResult) && invoice.GetInt("invoice_method") == enum.InvoiceMethod.PlatInvoice {
		tasks = append(tasks, task.TaskType.InvoicePayTask)
	}
	if param["result_params"] != "" {
		extend, err := types.NewXMapByJSON(types.GetString(param["result_params"]))
		if err != nil {
			return nil, false, nil, fmt.Errorf("完成开票,json转map失败%v", err)
		}
		info.Merge(extend)
	}
	info.Merge(data.Get(0))
	info.MergeMap(param)
	info.Merge(invoice)
	info.SetValue("invoice_id", param["coop_order_id"])
	return info, false, tasks, nil
}

//FinishInvoiceByDB 完成发票
func (f *DBFinish) FinishInvoiceByDB(trans db.IDBTrans, param types.XMap) error {
	if strings.EqualFold(param.GetString("status"), enum.InvoiceSuccessResult) ||
		strings.EqualFold(param.GetString("status"), enum.InvoiceFailedResult) {
		count, q, a, err := trans.Execute(sql.UpdateInvoiceFinish, param)
		if err != nil || types.GetInt(count) <= 0 {
			return fmt.Errorf("完成开票修改发生异常,err:%v,sql:%v,args:%v", err, q, a)
		}
	}
	return nil
}
