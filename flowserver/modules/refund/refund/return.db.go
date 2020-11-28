package refund

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBReturn 订单退货DB接口
type IDBReturn interface {
	CheckReturnAndRefundByDB(returnID int64) (data types.XMap, finish bool, err error)
	UpdateReturnAndRefundStatusByDB(dbTrans db.IDBTrans, refundID, returnID string) (err error)
	CheckReturnAndRefundStatusByDB(input *order.DeliveryResult) (types.IXMap, error)
	LockByDB(dbTrans db.IDBTrans, orderID, refundID, returnID string) error
	DoForSuccessByDB(dbTrans db.IDBTrans, data types.IXMap, input *order.DeliveryResult) ([]string, types.XMap, error)
	DoForAllSuccessByDB(dbTrans db.IDBTrans, data types.IXMap, returnSuccessAmountInfo types.XMap, input *order.DeliveryResult) ([]string, error)
	DoForFailByDB(dbTrans db.IDBTrans, data types.IXMap, input *order.DeliveryResult) ([]string, error)
	BuildReturnRequestParam(dbTrans db.IDBTrans, returnInfo types.IXMap) (*model.OrderCreateParam, error)
}

// DBReturn 订单退货对象
type DBReturn struct {
	c component.IContainer
}

//NewDBReturn 创建订单退货信息对象
func NewDBReturn(c component.IContainer) *DBReturn {
	return &DBReturn{
		c: c,
	}
}

// CheckReturnAndRefundByDB 检查退货和退款信息,返回退货信息
func (n *DBReturn) CheckReturnAndRefundByDB(returnID int64) (returnInfo types.XMap, finish bool, err error) {
	dbt := n.c.GetRegularDB()
	// 检查退货状态
	returnInfos, q, a, err := dbt.Query(sql.SqlCheckReturnInfo, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查退货状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, returnInfos.Len())
	}
	if returnInfos.IsEmpty() {
		return nil, true, fmt.Errorf("退货记录不存在")
	}

	// 检查退款状态
	refundInfos, q, a, err := dbt.Query(sql.SqlCheckRefundInfo, map[string]interface{}{
		"refund_id": returnInfos.Get(0).GetInt64("refund_id"),
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查退款状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, refundInfos.Len())
	}
	if refundInfos.IsEmpty() {
		return nil, true, fmt.Errorf("退款记录不存")
	}
	returnInfo = returnInfos.Get(0)
	return returnInfo, false, nil
}

// UpdateReturnAndRefundStatusByDB 修改退货状态
func (n *DBReturn) UpdateReturnAndRefundStatusByDB(dbTrans db.IDBTrans, refundID, returnID string) (err error) {
	// 修改退款信息表的上游退货状态为正在
	count, q, a, err := dbTrans.Execute(sql.SqlChangeRefundReturnStart, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return fmt.Errorf("修改退款信息表的上游退货错误(err:%v),sql:%s,input:%+v", err, q, a)
	}

	// 修改退货信息表的上游退货状态为正在
	count, q, a, err = dbTrans.Execute(sql.SqlChangeReturnToProgress, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("修改退货信息表的上游退货错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return
}

// CheckReturnAndRefundStatusByDB 检查退货状态和退款状态
func (n *DBReturn) CheckReturnAndRefundStatusByDB(input *order.DeliveryResult) (types.IXMap, error) {
	dbt := n.c.GetRegularDB()
	// 检查退货状态
	returns, q, a, err := dbt.Query(sql.SqlCheckReturnStatus, map[string]interface{}{
		"return_id": input.DeliveryID,
	})
	if err != nil || returns.IsEmpty() {
		return nil, fmt.Errorf("获取退货信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, returns.Len())
	}

	// 检查退款状态
	refunds, q, a, err := dbt.Query(sql.SqlCheckRefundStatus, map[string]interface{}{
		"refund_id": returns.Get(0).GetString("refund_id"),
	})
	if err != nil || refunds.IsEmpty() {
		return nil, fmt.Errorf("获取退款信息错误(err:%v),sql:%s,input:%+v,refunds:%+v", err, q, a, refunds)
	}
	refund := refunds.Get(0)

	// 查询通知记录
	notifys, q, a, err := dbt.Query(sql.SqlGetNotifyInfo, map[string]interface{}{
		"refund_id": refunds.Get(0).GetString("refund_id"),
	})
	if err != nil {
		return nil, fmt.Errorf("获取退款通知信息错误(err:%v),sql:%s,input:%+v,notifys:%+v", err, q, a, notifys)
	}

	refund.SetValue("return_id", input.DeliveryID)
	refund.SetValue("notify_id", "0")
	if !notifys.IsEmpty() {
		refund.SetValue("notify_id", notifys.Get(0).GetString("notify_id"))
	}
	refund.Merge(returns.Get(0))
	return refund, nil
}

// LockByDB 锁定记录
func (n *DBReturn) LockByDB(dbTrans db.IDBTrans, orderID, refundID, returnID string) error {
	// 锁定订单
	orders, q, a, err := dbTrans.Query(sql.SqlLockOrder, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || orders.IsEmpty() {
		return fmt.Errorf("锁定订单错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, orders.Len())
	}
	// 锁定退款记录
	refunds, q, a, err := dbTrans.Query(sql.SqlLockRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || refunds.IsEmpty() {
		return fmt.Errorf("锁定退款记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, refunds.Len())
	}
	// 锁定退货记录
	returns, q, a, err := dbTrans.Query(sql.SqlLockReturn, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil || returns.IsEmpty() {
		return fmt.Errorf("锁定退货记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, returns.Len())
	}
	return nil

}

// DoForSuccessByDB 成功保存函数
func (n *DBReturn) DoForSuccessByDB(dbTrans db.IDBTrans, data types.IXMap, input *order.DeliveryResult) ([]string, types.XMap, error) {
	// 1.修改当前退货为成功，当前上游退款为等待
	returnExt := ""
	if data.GetString("return_extend_info") != "" {
		ext, err := utils.GetExtendParamsByString(data.GetString("return_extend_info"), input.ResultParams)
		if err != nil {
			return nil, nil, err
		}
		returnExt = ext
	}

	count, q, a, err := dbTrans.Execute(sql.SqlUpdateCurrentReturnToSuccess, map[string]interface{}{
		"return_id":             input.DeliveryID,
		"return_msg":            input.ResultMsg,
		"courier_refund_amount": data.GetFloat64("courier_refund_amount"),
		"extend_info":           returnExt,
	})
	if err != nil || count <= 0 {
		return nil, nil, fmt.Errorf("修改当前退货状态为成功错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 2.获取退货成功总金额
	rows, q, a, err := dbTrans.Query(sql.SqlGetReturnSuccessAmount, map[string]interface{}{
		"refund_id":   data.GetString("refund_id"),
		"refund_face": data.GetFloat64("refund_face"),
	})
	if err != nil || rows.IsEmpty() {
		return nil, nil, fmt.Errorf("获取退货成功总金额错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	result := rows.Get(0)
	return []string{task.TaskType.UpRefundPaymentTask}, result, nil
}

// DoForAllSuccessByDB 全部退货成功处理
func (n *DBReturn) DoForAllSuccessByDB(dbTrans db.IDBTrans, data types.IXMap, returnSuccessAmountInfo types.XMap, input *order.DeliveryResult) ([]string, error) {
	// 1.处理拓展信息
	refundExt, err := utils.GetExtendParamsByString(data.GetString("refund_extend_info"), input.ResultParams)
	if err != nil {
		return nil, err
	}
	// 2.修改退款状态改为全部成功
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateRefundAllSuccess, map[string]interface{}{
		"refund_id":   data.GetString("refund_id"),
		"notify_id":   data.GetInt64("notify_id"),
		"extend_info": refundExt,
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("全部成功，处理退款信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 3.修改订单退款信息
	orderCount, q, a, err := dbTrans.Execute(sql.SqlUpdateOrder, map[string]interface{}{
		"order_id":                 data.GetString("order_id"),
		"return_total_face":        returnSuccessAmountInfo.GetFloat64("face"),
		"refund_sell_amount":       data.GetFloat64("refund_sell_amount"),
		"refund_commission_amount": data.GetFloat64("refund_commission_amount"),
		"refund_service_amount":    data.GetFloat64("refund_service_amount"),
		"refund_fee_amount":        data.GetFloat64("refund_fee_amount"),
		"return_cost_amount":       returnSuccessAmountInfo.GetFloat64("return_cost_amount"),
		"return_commission_amount": returnSuccessAmountInfo.GetFloat64("return_commission_amount"),
		"return_service_amount":    returnSuccessAmountInfo.GetFloat64("return_service_amount"),
		"fail_code":                enum.DeliveryFailCode.DownRefund.PlatCode,
		"fail_msg":                 enum.DeliveryFailCode.DownRefund.PlatMsg,
	})

	if err != nil || orderCount <= 0 {
		return nil, fmt.Errorf("处理订单信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	tasks := []string{task.TaskType.DownRefundPaymentTask}
	if data.GetInt("refund_point_num") > 0 {
		tasks = append(tasks, task.TaskType.PointUseRefundFDTask)
		return tasks, nil
	}
	if data.GetInt64("notify_id") == 0 {
		return tasks, nil
	}

	//5.将订单通知记录修改为等待通知
	notifyCount, q, a, err := dbTrans.Execute(sql.SqlUpdateRefundNotifyToWait, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
	})
	if err != nil || notifyCount <= 0 {
		return nil, fmt.Errorf("开启通知错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notifyCount)
	}

	return append(tasks, task.TaskType.RefundNotifyTask), nil
}

// DoForFailByDB 更新当前退货为失败
func (n *DBReturn) DoForFailByDB(dbTrans db.IDBTrans, data types.IXMap, input *order.DeliveryResult) ([]string, error) {
	// 修改当前退货记录为失败,记录失败原因
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateCurrentReturnToFail, map[string]interface{}{
		"return_id":  input.DeliveryID,
		"return_msg": input.ResultMsg,
		"code":       data.GetInt("refund_type"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("修改当前退货记录为失败错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 获取所有未失败的退货记录
	rows, q, a, err := dbTrans.Query(sql.SqlGetAllFailReturnList, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
	})
	if err != nil {
		return nil, fmt.Errorf("获取所有失败的退货记录错误(err:%v),sql:%s,input:%+v,rows:%+v", err, q, a, rows)
	}

	// 判断是否全部失败
	if !rows.IsEmpty() {
		// 未全部失败
		return nil, nil
	}

	// 退货全部失败处理退款记录为失败
	count, q, a, err = dbTrans.Execute(sql.SqlUpdateRefundToFail, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
		"notify_id": data.GetInt64("notify_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("全部失败更新错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	if data.GetString("notify_id") == "0" {
		return nil, nil
	}

	//退款失败后，将通知改为等待通知
	count, q, a, err = dbTrans.Execute(sql.SqlUpdateNotifyToWait, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("更新通知为等待错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return []string{task.TaskType.RefundNotifyTask}, nil
}

// BuildReturnRequestParam 构建退货参数
func (n *DBReturn) BuildReturnRequestParam(dbTrans db.IDBTrans, returnInfo types.IXMap) (*model.OrderCreateParam, error) {
	extendInfo, err := utils.GetExtendParams(returnInfo.GetString("extend_info"), map[string]interface{}{
		"ext_channel_no": returnInfo.GetString("ext_channel_no"),
	})
	if err != nil {
		return nil, err
	}
	return &model.OrderCreateParam{
		CoopID:        returnInfo.GetString("down_channel_no"),
		CoopOrderID:   returnInfo.GetString("return_id"),
		ChannelNo:     returnInfo.GetString("up_channel_no"),
		ServiceClass:  11,
		CarrierNo:     returnInfo.GetString("carrier_no"),
		ProductFace:   types.GetInt(returnInfo.GetFloat64("return_face") * 100),
		ProductNum:    returnInfo.GetInt("return_num"),
		NotifyURL:     queue.GetFinishReturnQueue(returnInfo.GetString("pre_tag")).GetName(n.c.GetPlatName()),
		OrderTimeout:  returnInfo.GetInt("return_overtime"),
		RequestParams: extendInfo,
	}, nil
}
