package overtime

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBOverTime 退款超时扫描DB接口
type IDBOverTime interface {
	UpdateWaitReturnToFailByDB(trans db.IDBTrans, refundID string) (flag bool, returnTask []string, list types.XMaps, err error)
	DealForAllFailReturnByDB(trans db.IDBTrans, data types.XMap) ([]string, error)
	CreateManualByDB(data types.XMap) (err error)
	GetReturnInfoByDB(returnID string) (types.XMap, error)
	QueryRefundInfoByDB(trans db.IDBTrans, refundID string) (types.XMap, error)
	UpdateRefundOvertime(trans db.IDBTrans, refundID string) error
	BuildServiceClass() string
	LockByDB(trans db.IDBTrans, orderID, refundID, returnID int64) (types.XMap, error)
	DoForSuccessByDB(trans db.IDBTrans, data types.IXMap, input *model.OrderQueryResult) ([]string, types.XMap, error)
	DoForAllSuccessByDB(trans db.IDBTrans, data types.IXMap, returnSuccessAmountInfo types.XMap, input *model.OrderQueryResult) ([]string, error)
	DoForFailByDB(trans db.IDBTrans, data types.IXMap, input *model.OrderQueryResult) ([]string, error)
}

// DBOverTime 退款超时扫描对象
type DBOverTime struct {
	c component.IContainer
}

//NewDBOverTime 创建退款超时扫描信息对象
func NewDBOverTime(c component.IContainer) *DBOverTime {
	return &DBOverTime{
		c: c,
	}
}

// QueryRefundInfoByDB 查询退款信息
func (n *DBOverTime) QueryRefundInfoByDB(trans db.IDBTrans, refundID string) (types.XMap, error) {

	// 1.查询退款信息
	resfunds, q, a, err := trans.Query(sql.SqlGetTimeOutRefundInfo, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取超时记录错误(err:%v),sql:%s,input:%+v,refundID:%v", err, q, a, refundID)
	}
	if resfunds.IsEmpty() {
		return nil, nil
	}
	refundInfo := resfunds.Get(0)

	// 获取通知记录
	notifyInfos, q, a, err := trans.Query(sql.SqlGetNotifyInfo, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取通知记录错误(err:%v),sql:%s,input:%+v,refundID:%v", err, q, a, refundID)
	}
	if notifyInfos.IsEmpty() {
		return refundInfo, nil
	}

	refundInfo.SetValue("notify_id", notifyInfos.Get(0).GetInt64("notify_id"))
	return refundInfo, nil
}

// UpdateRefundOvertime 更新退款超时
func (n *DBOverTime) UpdateRefundOvertime(trans db.IDBTrans, refundID string) error {
	// 关闭超时扫描
	count, q, a, err := trans.Execute(sql.SqlUpdateReturnOvertime, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("更新超时时间错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	return nil
}

// UpdateToFailByDB 更新为失败
func (n *DBOverTime) UpdateToFailByDB(data types.XMap) (err error) {
	dbt := n.c.GetRegularDB()
	count, q, a, err := dbt.Execute(sql.SqlUpdateToFail, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
	})
	if err != nil || count <= 0 {
		err = fmt.Errorf("修改下游退款状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
		return err
	}
	return
}

// UpdateWaitReturnToFailByDB 更新等待退货为失败
func (n *DBOverTime) UpdateWaitReturnToFailByDB(trans db.IDBTrans, refundID string) (allFail bool, returnTask []string, list types.XMaps, err error) {
	//等待退货修改为失败
	_, q, a, err := trans.Execute(sql.UpdateWaitReturnToFail, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		err = fmt.Errorf("更新等待退货为失败记录错误(err:%v),sql:%s,input:%+v", err, q, a)
		return
	}

	//获取正在退货的退货记录
	list, q, a, err = trans.Query(sql.SqlGetProcessReturnList, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		err = fmt.Errorf("获取正在退货的记录错误(err:%v),sql:%s,input:%+v,refundID:%v", err, q, a, refundID)
		return
	}
	if !list.IsEmpty() {
		return false, []string{task.TaskType.ReturnUnknownTask}, list, nil
	}

	//获取非失败的退货记录
	rows, q, a, err := trans.Query(sql.SqlGetAllFailReturnList, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		err = fmt.Errorf("获取全部退货失败信息错误(err:%v),sql:%s,input:%+v,rows:%+v", err, q, a, rows)
		return
	}

	return rows.IsEmpty(), nil, list, nil
}

// DealForAllFailReturnByDB 处理完全退货
func (n *DBOverTime) DealForAllFailReturnByDB(trans db.IDBTrans, data types.XMap) ([]string, error) {
	//全部失败
	count, q, a, err := trans.Execute(sql.SqlUpdateRefundToFail, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
		"notify_id": data.GetInt64("notify_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("修改退货状态为失败错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 检查是否需要通知
	if data.GetInt64("notify_id", 0) == 0 {
		return nil, nil
	}

	// 开启通知
	count, q, a, err = trans.Execute(sql.SqlUpdateNotifyToWait, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("开启通知错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return []string{task.TaskType.RefundNotifyTask}, nil
}

// CreateManualByDB 创建人工审核记录
func (n *DBOverTime) CreateManualByDB(data types.XMap) (err error) {
	// 创建人工审核记录
	_, count, q, a, err := utils.Insert(n.c.GetRegularDB(), sql.GetNewID, sql.SqlCreateManual, data)
	if err != nil || count <= 0 {
		err = fmt.Errorf("创建人工审核记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
		return err
	}
	return
}

// GetReturnInfoByDB 获取退货信息
func (n *DBOverTime) GetReturnInfoByDB(returnID string) (types.XMap, error) {
	dbt := n.c.GetRegularDB()
	// 查询退货记录
	returns, q, a, err := dbt.Query(sql.SqlGetReturnInfo, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil || returns.IsEmpty() {
		return nil, fmt.Errorf("查询退货记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, returns.Len())
	}
	data := returns.Get(0)

	// 查询退款记录
	refunds, q, a, err := dbt.Query(sql.SqlQueryRefundInfo, map[string]interface{}{
		"refund_id": data.GetString("refund_id"),
	})
	if err != nil || refunds.IsEmpty() {
		return nil, fmt.Errorf("查询退款记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, refunds.Len())
	}

	data.SetValue("down_channel_no", refunds.Get(0).GetString("down_channel_no"))
	return data, nil
}

// BuildServiceClass 构建ServiceClass
func (n *DBOverTime) BuildServiceClass() string {
	return "40"
}

// LockByDB 锁定记录
func (n *DBOverTime) LockByDB(trans db.IDBTrans, orderID, refundID, returnID int64) (types.XMap, error) {
	// 锁定订单
	orders, q, a, err := trans.Query(sql.SqlLockOrder, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, fmt.Errorf("锁定订单错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, orders.Len())
	}
	if orders.IsEmpty() {
		return nil, nil
	}
	// 锁定退款记录
	refunds, q, a, err := trans.Query(sql.SqlLockRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return nil, fmt.Errorf("锁定退款记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, refunds.Len())
	}
	if refunds.IsEmpty() {
		return nil, nil
	}
	// 锁定退货记录
	returns, q, a, err := trans.Query(sql.SqlLockReturn, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil {
		return nil, fmt.Errorf("锁定退货记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, returns.Len())
	}
	if returns.IsEmpty() {
		return nil, nil
	}

	//查询退款通知信息
	notifys, q, a, err := trans.Query(sql.GetRefundNotifyInfo, map[string]interface{}{"refund_id": refundID})
	if err != nil {
		return nil, fmt.Errorf("查询退款通知信息异常(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notifys.Len())
	}

	orderInfo := orders.Get(0)
	orderInfo.Merge(refunds.Get(0))
	orderInfo.Merge(returns.Get(0))
	if notifys.IsEmpty() {
		return orderInfo, nil
	}
	orderInfo.SetValue("notify_id", notifys.Get(0).GetInt64("notify_id"))
	return orderInfo, nil
}

// DoForSuccessByDB 成功保存函数
func (n *DBOverTime) DoForSuccessByDB(trans db.IDBTrans, data types.IXMap, input *model.OrderQueryResult) ([]string, types.XMap, error) {

	// 1.修改当前退货为成功，当前上游退款为等待
	returnExt := ""
	count, q, a, err := trans.Execute(sql.SqlUpdateCurrentReturnToSuccess, map[string]interface{}{
		"return_id":             data.GetInt64("return_id"),
		"return_msg":            input.ResultDesc,
		"courier_refund_amount": input.CourierCostAmount,
		"extend_info":           returnExt,
	})
	if err != nil || count <= 0 {
		return nil, nil, fmt.Errorf("修改当前退货状态为成功错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 2.获取退货成功总金额
	rows, q, a, err := trans.Query(sql.SqlGetReturnSuccessAmount, map[string]interface{}{
		"refund_id":   data.GetString("refund_id"),
		"refund_face": data.GetFloat64("return_total_face"),
	})
	if err != nil || rows.IsEmpty() {
		return nil, nil, fmt.Errorf("获取退货成功总金额错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	result := rows.Get(0)

	return []string{task.TaskType.UpRefundPaymentTask}, result, nil
}

// DoForAllSuccessByDB 全部退货成功处理
func (n *DBOverTime) DoForAllSuccessByDB(trans db.IDBTrans, data types.IXMap, returnSuccessAmountInfo types.XMap, input *model.OrderQueryResult) ([]string, error) {
	// 1.处理拓展信息
	refundExt, err := utils.GetExtendParamsByString(data.GetString("refund_extend_info"), input.ResultParams)
	if err != nil {
		return nil, err
	}
	// 2.修改退款状态改为全部成功
	count, q, a, err := trans.Execute(sql.SqlUpdateRefundAllSuccess, map[string]interface{}{
		"refund_id":   data.GetInt64("refund_id"),
		"notify_id":   data.GetInt64("notify_id"),
		"extend_info": refundExt,
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("全部成功，处理退款信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 3.修改订单退款信息
	orderCount, q, a, err := trans.Execute(sql.SqlUpdateOrder, map[string]interface{}{
		"order_id":                 data.GetInt64("order_id"),
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

	// 开启通知
	notifyCount, q, a, err := trans.Execute(sql.SqlUpdateRefundNotifyToWait, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
	})
	if err != nil || notifyCount <= 0 {
		return nil, fmt.Errorf("开启通知错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notifyCount)
	}

	return append(tasks, task.TaskType.RefundNotifyTask), nil
}

// DoForFailByDB 更新当前退货为失败
func (n *DBOverTime) DoForFailByDB(trans db.IDBTrans, data types.IXMap, input *model.OrderQueryResult) ([]string, error) {
	// 修改当前退货记录为失败,记录失败原因
	count, q, a, err := trans.Execute(sql.SqlUpdateCurrentReturnToFail, map[string]interface{}{
		"return_id":  data.GetInt64("return_id"),
		"return_msg": input.ResultDesc,
		"code":       data.GetInt("refund_type"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("修改当前退货记录为失败错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 获取所有失败的退货记录
	rows, q, a, err := trans.Query(sql.SqlGetAllFailReturnList, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
	})
	if err != nil {
		return nil, fmt.Errorf("获取所有失败的退货记录错误(err:%v),sql:%s,input:%+v,rows:%+v", err, q, a, rows)
	}

	// 判断是否全部失败
	if !rows.IsEmpty() {
		// 未全部失败
		return nil, nil
	}

	// 全部失败处理
	count, q, a, err = trans.Execute(sql.SqlUpdateRefundToFail, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
		"notify_id": data.GetInt64("notify_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("全部失败更新错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	if data.GetInt64("notify_id") == 0 {
		return nil, nil
	}

	// 更新通知为等待
	count, q, a, err = trans.Execute(sql.SqlUpdateNotifyToWait, map[string]interface{}{
		"refund_id": data.GetInt64("refund_id"),
	})
	if err != nil || count <= 0 {
		return nil, fmt.Errorf("更新通知为等待错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return []string{task.TaskType.RefundNotifyTask}, nil
}
