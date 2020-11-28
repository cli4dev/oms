package order

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBOvertime 数据层接口
type IDBOvertime interface {
	QueryOrderInfoByDB(orderID int64) (types.XMap, error)
	OrderOvertimePayFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error)
	DeliveryAuditByDB(trans db.IDBTrans, data types.XMap) error
	OvertimeNotifyByDB(trans db.IDBTrans, data types.XMap) (bool, error)
	LockOrderAndDeliveryByDB(trans db.IDBTrans, deliveryID, orderID int64) (types.XMap, error)
	BuildServiceClass() string
	BindOvertimeFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error)
	DeliveryOvertimeFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, types.XMaps, error)
	WaitingDeliveryOvertimeByDB(trans db.IDBTrans, data types.XMap, result *model.OrderQueryResult) ([]string, error)
	DeliverySuccessByDB(trans db.IDBTrans, data types.XMap, result *model.OrderQueryResult) ([]string, error)
	CheckBindFailedByDB(trans db.IDBTrans, data types.XMap) (allFail bool, err error)
}

// DBOvertime 超时数据层
type DBOvertime struct {
	c component.IContainer
}

// NewDBOvertime 构建DBOvertime
func NewDBOvertime(c component.IContainer) *DBOvertime {
	return &DBOvertime{
		c: c,
	}
}

// QueryOrderInfoByDB 查询订单信息
func (o *DBOvertime) QueryOrderInfoByDB(orderID int64) (types.XMap, error) {
	db := o.c.GetRegularDB()
	param := map[string]interface{}{"order_id": orderID}

	// 1.查询非结束订单信息
	datas, sqlStr, args, err := db.Query(sql.QueryOrderForOvertime, param)
	if err != nil {
		return nil, fmt.Errorf("查询订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", datas.Len(), err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil
	}

	// 2.通知查询
	orderInfo := datas.Get(0)
	notifys, sqlStr, args, err := db.Query(sql.CheckNotifyExist, param)
	if err != nil {
		return nil, fmt.Errorf("查询订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", notifys.Len(), err, sqlStr, args)
	}
	if notifys.IsEmpty() {
		return orderInfo, nil
	}
	orderInfo.SetValue("notify_id", notifys.Get(0).GetInt64("notify_id"))
	return orderInfo, nil
}

// OrderOvertimePayFailedByDB 订单超时失败
func (o *DBOvertime) OrderOvertimePayFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error) {

	// 1.修改订单状态为失败,支付状态为失败
	row, sqlStr, args, err := trans.Execute(sql.OrderOvertimeFailed, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	tasks := []string{}
	if data.GetInt("point_num") > 0 {
		tasks = append(tasks, task.TaskType.PointUseRefundFDTask)
	}
	// 2.修改通知状态
	if data.GetInt64("notify_id") != 0 {
		row, sqlStr, args, err = trans.Execute(sql.UpdateNotifyStatus, data)
		if err != nil || row != 1 {
			return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
		}
		return append(tasks, task.TaskType.NotifyTask), nil
	}
	return tasks, nil
}

// DeliveryAuditByDB 发货未知审核
func (o *DBOvertime) DeliveryAuditByDB(trans db.IDBTrans, data types.XMap) error {
	_, row, sqlStr, args, err := utils.Insert(trans, sql.GetNewID, sql.CreateAuditForDelivery, data)
	if err != nil || row != 1 {
		return fmt.Errorf("发货未知创建审核记录发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// LockOrderAndDeliveryByDB 锁订单和发货
func (o *DBOvertime) LockOrderAndDeliveryByDB(trans db.IDBTrans, deliveryID, orderID int64) (types.XMap, error) {
	// 1.锁订单
	orderInfos, sqlStr, args, err := trans.Query(sql.LockOrderForDeliveryOvertime, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, fmt.Errorf("发货锁订单发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if orderInfos.IsEmpty() {
		return nil, nil
	}
	orderInfo := orderInfos.Get(0)
	// 2.锁发货
	deliverys, sqlStr, args, err := trans.Query(sql.LockDeliveryIn, map[string]interface{}{
		"delivery_id":  deliveryID,
		"bind_face":    orderInfo.GetFloat64("bind_face"),
		"total_face":   orderInfo.GetFloat64("total_face"),
		"success_face": orderInfo.GetFloat64("success_face"),
	})
	if err != nil {
		return nil, fmt.Errorf("锁发货发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if deliverys.IsEmpty() {
		return nil, nil
	}
	orderInfo.Merge(deliverys.Get(0))

	products, sqlStr, args, err := trans.Query(sql.QueryDownProduct, map[string]interface{}{
		"product_id": orderInfo.GetString("down_product_id"),
	})
	if err != nil {
		return nil, fmt.Errorf("查询下游产品信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if products.IsEmpty() {
		return nil, nil
	}
	orderInfo.Merge(products.Get(0))

	// 3.通知查询
	notifys, sqlStr, args, err := trans.Query(sql.CheckNotifyExist, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, fmt.Errorf("查询订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", notifys.Len(), err, sqlStr, args)
	}
	if notifys.IsEmpty() {
		return orderInfo, nil
	}
	orderInfo.SetValue("notify_id", notifys.Get(0).GetInt64("notify_id"))
	return orderInfo, nil
}

// OvertimeNotifyByDB 超时通知
func (o *DBOvertime) OvertimeNotifyByDB(trans db.IDBTrans, data types.XMap) (bool, error) {
	// 2.通知查询
	notifys, sqlStr, args, err := trans.Query(sql.CheckNotifyExist, data)
	if err != nil {
		return false, fmt.Errorf("查询订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", notifys.Len(), err, sqlStr, args)
	}
	if notifys.IsEmpty() {
		return false, nil
	}

	// 3.修改通知状态
	row, sqlStr, args, err := trans.Execute(sql.UpdateNotifyStatus, data)
	if err != nil || row != 1 {
		return false, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	return true, nil
}

// BuildServiceClass 构建ServiceClass
func (o *DBOvertime) BuildServiceClass() string {
	return "10"
}

// CheckBindFailedByDB 判断是否时绑定失败
func (o *DBOvertime) CheckBindFailedByDB(trans db.IDBTrans, data types.XMap) (allFail bool, err error) {
	// 查询非失败发货
	bindResult, sqlStr, args, err := trans.Scalar(sql.QueryNoFailedDelivery, data)
	if err != nil {
		return false, fmt.Errorf("查询非失败发货记录发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	return types.GetBool(bindResult), nil
}

// BindOvertimeFailedByDB 绑定失败处理
func (o *DBOvertime) BindOvertimeFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error) {
	// 1.修改订单状态为失败,支付状态为失败
	row, sqlStr, args, err := trans.Execute(sql.OrderRefund, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	tasks := []string{task.TaskType.OrderFailRefundTask}
	if data.GetInt("point_num") > 0 {
		tasks = append(tasks, task.TaskType.PointUseRefundFDTask)
	}

	// 2.修改通知状态
	if data.GetInt64("notify_id", 0) == 0 {
		return tasks, nil
	}

	row, sqlStr, args, err = trans.Execute(sql.UpdateNotifyStatus, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return append(tasks, task.TaskType.NotifyTask), nil
}

// DeliveryOvertimeFailedByDB 发货中处理
func (o *DBOvertime) DeliveryOvertimeFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, types.XMaps, error) {

	//查询发货中的发货信息
	deliverys, sqlStr, args, err := trans.Query(sql.QueryDeliveryInfo, map[string]interface{}{"order_id": data.GetInt64("order_id")})
	if err != nil {
		return nil, nil, fmt.Errorf("查询发货信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", deliverys.Len(), err, sqlStr, args)
	}
	if !deliverys.IsEmpty() {
		return []string{task.TaskType.DeliveryUnknownTask}, deliverys, nil
	}

	//没有发货中的记录,订单处于部分成功部分失败，处理订单信息
	row, sqlStr, args, err := trans.Execute(sql.UpdateOrderOvertime, map[string]interface{}{"order_id": data.GetInt64("order_id")})
	if err != nil || row != 1 {
		return nil, nil, fmt.Errorf("修改超时发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil, nil, nil
}

// WaitingDeliveryOvertimeByDB 等待发货处理
func (o *DBOvertime) WaitingDeliveryOvertimeByDB(trans db.IDBTrans, data types.XMap, result *model.OrderQueryResult) ([]string, error) {
	if result != nil {
		data.SetValue("msg", result.ResultDesc)
		data.SetValue("up_delivery_no", result.OrderNo)
		data.SetValue("courier_cost_amount", result.CourierCostAmount)
	}
	// 1.修改订单信息
	row, sqlStr, args, err := trans.Execute(sql.UpdateOrderForDeliveryFailed, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("发货失败修改订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 2.修改发货记录为失败
	rows, sqlStr, args, err := trans.Execute(sql.DeliveryFailed, data)
	if err != nil || rows != 1 {
		return nil, fmt.Errorf("发货失败修改发货记录发生异常,cnt:%d,err:%v,sql:%v,args:%v", rows, err, sqlStr, args)
	}

	// 3.全部失败,创建退款
	if !data.GetBool("all_failed") {
		return nil, nil
	}

	// 4.修改订单状态为失败,支付状态为失败
	row, sqlStr, args, err = trans.Execute(sql.OrderRefund, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	tasks := []string{task.TaskType.OrderFailRefundTask}
	if data.GetInt("point_num") > 0 {
		tasks = append(tasks, task.TaskType.PointUseRefundFDTask)
	}

	if data.GetInt64("notify_id") == 0 {
		return tasks, nil
	}
	// 5.修改通知状态
	row, sqlStr, args, err = trans.Execute(sql.UpdateNotifyStatus, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("订单超时失败发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return append(tasks, task.TaskType.NotifyTask), nil
}

// DeliverySuccessByDB 发货成功处理
func (o *DBOvertime) DeliverySuccessByDB(trans db.IDBTrans, data types.XMap, result *model.OrderQueryResult) ([]string, error) {
	data.SetValue("msg", result.ResultDesc)
	data.SetValue("up_delivery_no", result.OrderNo)
	data.SetValue("courier_cost_amount", result.CourierCostAmount)
	notifyID := data.GetInt64("notify_id", 0)

	// 1.修改订单信息
	extend, err := utils.GetExtendParamsByString(data.GetString("extend_info"), data.GetString("order_extend_info"))
	if err != nil {
		return nil, fmt.Errorf("合并订单拓展信息,err:%v", err)
	}
	data.SetValue("extend_info", extend)
	row, sqlStr, args, err := trans.Execute(sql.UpdateOrderForDeliverySuccess, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("发货成功修改订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 2.修改发货记录为成功
	row, sqlStr, args, err = trans.Execute(sql.DeliverySuccess, data)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("发货成功修改发货记录发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 3.判断是否完成发货
	if !data.GetBool("all_success") || notifyID == 0 {
		return []string{task.TaskType.UpPaymentTask}, nil
	}

	// 4.修改通知状态
	row, sqlStr, args, err = trans.Execute(sql.UpdateNotifyStatus, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || row != 1 {
		return nil, fmt.Errorf("修改通知发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return []string{task.TaskType.UpPaymentTask, task.TaskType.NotifyTask}, nil
}
