package audit

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/oms/oms/modules/const/sql"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

//DeliveryAuditCheck 发货审核检查
func DeliveryAuditCheck(dbTrans db.IDBTrans, input *RequestParams) (types.IXMap, error) {
	//查询发货信息
	rows, _, args, err := dbTrans.Query(sql.GetDeliveryInfo, map[string]interface{}{
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
	})
	if err != nil || rows.IsEmpty() {
		return nil, fmt.Errorf("发货审核获取发货信息失败，count:%d,args:%v,err:%v", rows.Len(), args, err)
	}

	return rows.Get(0), nil
}

//DeliveryAuditSuccess 发货审核成功
func DeliveryAuditSuccess(dbTrans db.IDBTrans, input *RequestParams, info types.IXMap) (notifyID int64, err error) {
	//处理订单信息
	count, _, _, err := dbTrans.Execute(sql.AuditOrderSuccess, info.ToMap())
	if err != nil || count != 1 {
		return 0, fmt.Errorf("发货审核成功修改订单失败，count:%d,err:%v", count, err)
	}

	//处理发货为成功
	count, _, _, err = dbTrans.Execute(sql.AuditDeliverySuccess, map[string]interface{}{
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
		"audit_msg":   input.AuditMsg,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("发货审核成功失败，count:%d,err:%v", count, err)
	}

	//判断订单是否全部发货成功
	orders, _, _, err := dbTrans.Query(sql.GetOrderStatus, map[string]interface{}{
		"order_id": input.OrderID,
	})
	if err != nil || orders.IsEmpty() {
		return 0, fmt.Errorf("获取订单状态信息失败，count:%d,err:%v", orders.Len(), err)
	}
	if orders.Get(0).GetInt("order_status") != 0 || orders.Get(0).GetInt("delivery_bind_status") != 0 {
		return 0, nil
	}

	//获取通知编号
	nid, sqlStr, args, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    input.OrderID,
		"notify_type": 1,
	})
	if err != nil {
		return 0, fmt.Errorf("获取审核订单通知编号失败，sql:%s,args:%v,err:%v", sqlStr, args, err)
	}
	notifyID = types.GetInt64(nid)

	if notifyID == 0 {
		return 0, nil
	}

	//开启通知状态
	count, sqlStr, args, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("开启审核订单通知失败，count:%d,sql:%s,args:%v,err:%v", count, sqlStr, args, err)
	}

	return notifyID, nil
}

//SetAuditRecordSuccess 审核记录处理成功
func SetAuditRecordSuccess(dbTrans db.IDBTrans, input *RequestParams, AuditType int) error {
	count, sqlStr, args, err := dbTrans.Execute(sql.SetAuditRecordSuccess, map[string]interface{}{
		"audit_id":    input.AuditID,
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
		"audit_by":    input.AuditBy,
		"audit_msg":   input.AuditMsg,
		"change_type": AuditType,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("审核记录处理成功失败，count:%d,sqlStr:%s,args:%v,err:%v", count, sqlStr, args, err)
	}
	return nil
}

//DeliveryAuditFail 发货审核为失败
func DeliveryAuditFail(dbTrans db.IDBTrans, input *RequestParams) (isAllDeliveryFail bool, notifyID int64, err error) {
	//锁订单
	orders, _, _, err := dbTrans.Query(sql.LockAuditOrder, map[string]interface{}{
		"order_id": input.OrderID,
	})
	if err != nil || orders.IsEmpty() {
		return false, 0, fmt.Errorf("锁审核订单发生异常，count:%d,err:%v", orders.Len(), err)
	}

	//处理发货为失败
	count, _, _, err := dbTrans.Execute(sql.AuditDeliveryFail, map[string]interface{}{
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
		"audit_msg":   input.AuditMsg,
	})
	if err != nil || count != 1 {
		return false, 0, fmt.Errorf("发货审核为失败异常，count:%d,err:%v", count, err)
	}

	//判断发货是否全部失败
	cn, _, _, err := dbTrans.Scalar(sql.GetNotFailDeliveryCount, map[string]interface{}{
		"order_id": input.OrderID,
	})
	if err != nil {
		return false, 0, fmt.Errorf("发货审核为失败异常，err:%v", err)
	}

	if types.GetInt(cn) != 0 {
		return false, 0, nil
	}

	//处理订单发货失败
	count, _, _, err = dbTrans.Execute(sql.UpdateOrderDeliveryFail, map[string]interface{}{
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
		"audit_msg":   input.AuditMsg,
	})
	if err != nil {
		return false, 0, fmt.Errorf("处理订单失败异常，count:%d,err:%v", count, err)
	}

	//获取通知编号
	nid, _, _, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    input.OrderID,
		"notify_type": 1,
	})
	if err != nil {
		return false, 0, fmt.Errorf("获取审核订单通知编号失败，err:%v", err)
	}
	notifyID = types.GetInt64(nid)

	if notifyID == 0 {
		return true, 0, nil
	}

	//开启通知状态
	count, _, _, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return false, 0, fmt.Errorf("获取审核订单通知编号失败，count:%d,err:%v", count, err)
	}

	return true, notifyID, nil
}

//SetAuditRecordFail 处理审核记录为失败
func SetAuditRecordFail(dbTrans db.IDBTrans, input *RequestParams, AuditType int) error {
	count, _, _, err := dbTrans.Execute(sql.SetAuditRecordFail, map[string]interface{}{
		"audit_id":    input.AuditID,
		"order_id":    input.OrderID,
		"delivery_id": input.DeliveryID,
		"audit_by":    input.AuditBy,
		"audit_msg":   input.AuditMsg,
		"change_type": AuditType,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("发货审核记录处理为失败异常，count:%d,err:%v", count, err)
	}
	return nil
}

//ReturnAuditSuccess 退货结果未知审核为成功
func ReturnAuditSuccess(dbTrans db.IDBTrans, input *RequestParams) (isAllReturnSucc bool, info types.IXMap, err error) {
	//锁订单
	orders, _, _, err := dbTrans.Query(sql.LockReturnOrder, map[string]interface{}{
		"order_id": input.OrderID,
	})
	if err != nil || orders.IsEmpty() {
		return false, nil, fmt.Errorf("锁订单失败，count:%d,err:%v", orders.Len(), err)
	}

	//退货记录处理为成功
	count, _, _, err := dbTrans.Execute(sql.AuditReturnSuccess, map[string]interface{}{
		"order_id":  input.OrderID,
		"return_id": input.DeliveryID,
		"refund_id": input.RefundID,
		"audit_msg": input.AuditMsg,
	})
	if err != nil || count != 1 {
		return false, nil, fmt.Errorf("审核未知退货为成功失败，count:%d,err:%v", count, err)
	}

	//判断退款记录的退货是否全部完成
	refunds, _, _, err := dbTrans.Query(sql.GetAuditRefundInfo, map[string]interface{}{
		"refund_id": input.RefundID,
	})
	if err != nil || refunds.IsEmpty() {
		return false, nil, fmt.Errorf("获取退款信息失败，count:%d,err:%v", refunds.Len(), err)
	}

	returns, _, _, err := dbTrans.Query(sql.GetAuditRetrunInfo, map[string]interface{}{
		"refund_id": input.RefundID,
	})
	if err != nil || returns.IsEmpty() {
		return false, nil, fmt.Errorf("获取退款信息失败，count:%d,err:%v", returns.Len(), err)
	}

	result := refunds.Get(0)
	result.Merge(returns.Get(0))

	if refunds.Get(0).GetInt("refund_face") != returns.Get(0).GetInt("up_return_face") {
		return false, result, nil
	}

	return true, result, nil
}

//ReturnAuditComleteSuccess 退款审核完全成功
func ReturnAuditComleteSuccess(dbTrans db.IDBTrans, params types.IXMap) (notifyID int64, err error) {
	//全部完成处理退款信息
	count, _, _, err := dbTrans.Execute(sql.UpdateRefundSuccess, params.ToMap())
	if err != nil || count != 1 {
		return 0, fmt.Errorf("审核未知退货为成功失败，count:%d,err:%v", count, err)
	}

	//全部完成处理订单信息
	count, sqlStr, args, err := dbTrans.Execute(sql.UpdateOrderRefundSuccess, params.ToMap())
	if err != nil || count != 1 {
		return 0, fmt.Errorf("审核未知退货为成功失败，count:%d,sql:%s,args:%v,err:%v", count, sqlStr, args, err)
	}

	//获取通知编号
	nid, _, _, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    params.GetInt64("order_id"),
		"refund_id":   params.GetInt64("refund_id"),
		"notify_type": 2,
	})
	if err != nil {
		return 0, fmt.Errorf("获取审核订单通知编号失败，err:%v", err)
	}
	notifyID = types.GetInt64(nid)

	if notifyID == 0 {
		return 0, nil
	}

	//开启通知状态
	count, _, _, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("获取审核订单通知编号失败，count:%d,err:%v", count, err)
	}

	return notifyID, nil
}

//ReturnAuditFail 退货记录审核为失败
func ReturnAuditFail(dbTrans db.IDBTrans, input *RequestParams) (isAllReturnFail bool, err error) {
	//锁订单
	orders, _, _, err := dbTrans.Query(sql.LockReturnOrder, map[string]interface{}{
		"order_id": input.OrderID,
	})
	if err != nil || orders.IsEmpty() {
		return false, fmt.Errorf("锁订单失败，count:%d,err:%v", orders.Len(), err)
	}

	//退货记录处理为失败
	count, _, _, err := dbTrans.Execute(sql.AuditReturnFail, map[string]interface{}{
		"order_id":  input.OrderID,
		"return_id": input.DeliveryID,
		"refund_id": input.RefundID,
		"audit_msg": input.AuditMsg,
	})
	if err != nil || count != 1 {
		return false, fmt.Errorf("审核未知退货为失败异常，count:%d,err:%v", count, err)
	}

	//判断退款记录的退货是否全部Fail
	refunds, _, _, err := dbTrans.Query(sql.GetFailRefundInfo, map[string]interface{}{
		"refund_id": input.RefundID,
	})
	if err != nil || refunds.IsEmpty() {
		return false, fmt.Errorf("审核未知退货为失败异常，count:%d,err:%v", refunds.Len(), err)
	}

	retruns, _, _, err := dbTrans.Query(sql.GetFailReturnInfo, map[string]interface{}{
		"refund_id": input.RefundID,
	})
	if err != nil || retruns.IsEmpty() {
		return false, fmt.Errorf("审核未知退货为失败异常，count:%d,err:%v", retruns.Len(), err)
	}

	if refunds.Get(0).GetInt("refund_face") != retruns.Get(0).GetInt("up_return_face") {
		return false, nil
	}

	return true, nil
}

//ReturnAuditCompleteFail 退货审核完全失败处理
func ReturnAuditCompleteFail(dbTrans db.IDBTrans, input *RequestParams) (notifyID int64, err error) {
	//全部完成处理退款信息
	count, _, _, err := dbTrans.Execute(sql.UpdateRefundFail, map[string]interface{}{
		"order_id":  input.OrderID,
		"refund_id": input.RefundID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("审核未知退货为成功失败，count:%d,err:%v", count, err)
	}

	//获取通知编号
	nid, _, _, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    input.OrderID,
		"refund_id":   input.RefundID,
		"notify_type": 2,
	})
	if err != nil {
		return 0, fmt.Errorf("获取审核订单通知编号失败，err:%v", err)
	}
	notifyID = types.GetInt64(nid)

	if notifyID == 0 {
		return 0, nil
	}

	//开启通知状态
	count, _, _, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("获取审核订单通知编号失败，count:%d,err:%v", count, err)
	}

	return notifyID, nil
}

//UpdateOrderPartialSuccess 修改订单为部分成功
func UpdateOrderPartialSuccess(dbTrans db.IDBTrans, orderID int64) (notifyID int64, err error) {
	//判断发货记录是否全部终结
	cn, _, _, err := dbTrans.Scalar(sql.GetNotEndDelivery, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return 0, err
	}
	if types.GetInt(cn) > 0 {
		return 0, fmt.Errorf("订单发货记录没有全部终结")
	}

	//处理订单为部分成功
	count, _, _, err := dbTrans.Execute(sql.UpdateOrderPartialSuccess, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return 0, err
	}
	if types.GetInt(count) != 1 {
		return 0, fmt.Errorf("订单不是部分成功订单")
	}

	//获取通知编号
	nid, _, _, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    orderID,
		"notify_type": 1,
	})
	if err != nil {
		return 0, fmt.Errorf("获取审核订单通知编号失败，err:%v", err)
	}
	notifyID = types.GetInt64(nid)

	if notifyID == 0 {
		return 0, nil
	}

	//开启通知状态
	count, _, _, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("获取审核订单通知编号失败，count:%d,err:%v", count, err)
	}

	return notifyID, nil
}

//UpdateRefundPartialSuccess 修改退款记录为部分成功
func UpdateRefundPartialSuccess(dbTrans db.IDBTrans, orderID int64, refundID int64) (int64, error) {
	//锁订单
	orders, _, _, err := dbTrans.Query(sql.LockReturnOrder, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || orders.IsEmpty() {
		return 0, fmt.Errorf("锁订单失败，count:%d,err:%v", orders.Len(), err)
	}

	//判断退货记录是否全部终结
	rows, _, _, err := dbTrans.Query(sql.GetPartialRetrunInfo, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || rows.IsEmpty() {
		return 0, fmt.Errorf("判断退款部分成功失败，count:%d,err:%v", rows.Len(), err)
	}
	if rows.Get(0).GetInt("not_end_count") > 0 || rows.Get(0).GetInt("not_refund_count") > 0 {
		return 0, fmt.Errorf("退款退货记录没有全部终结")
	}

	//处理退款为部分成功
	count, sqlStr, args, err := dbTrans.Execute(sql.UpdateReturnPartialSuccess, map[string]interface{}{
		"refund_id":           refundID,
		"success_retrun_face": rows.Get(0).GetInt("refund_face"),
	})
	fmt.Println(sqlStr, args)
	if err != nil {
		return 0, err
	}
	if types.GetInt(count) != 1 {
		return 0, fmt.Errorf("退款不是部分成功退款")
	}

	//处理订单信息
	count, _, _, err = dbTrans.Execute(sql.UpdateReturnPartialOrder, rows.Get(0))
	if err != nil || count != 1 {
		return 0, fmt.Errorf("部分退款处理订单失败，count:%d,err:%v", count, err)
	}

	//获取通知编号
	nid, _, _, err := dbTrans.Scalar(sql.GetOrderNotifyID, map[string]interface{}{
		"order_id":    orderID,
		"refund_id":   refundID,
		"notify_type": 2,
	})
	if err != nil {
		return 0, fmt.Errorf("获取审核订单通知编号失败，err:%v", err)
	}
	notifyID := types.GetInt64(nid)

	if notifyID == 0 {
		return 0, nil
	}
	//开启通知状态
	count, _, _, err = dbTrans.Execute(sql.StartOrderNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count != 1 {
		return 0, fmt.Errorf("获取审核订单通知编号失败，count:%d,err:%v", count, err)
	}

	return notifyID, nil
}

//GetFalseSuccDeliveryInfo 获取假成功发货信息
func GetFalseSuccDeliveryInfo(dbTrans db.IDBTrans, deliveryID int64) (types.IXMap, error) {
	rows, _, _, err := dbTrans.Query(sql.GetFalseSuccDeliveryInfo, map[string]interface{}{
		"delivery_id": deliveryID,
	})
	if err != nil {
		return nil, err
	}
	if rows.IsEmpty() {
		return nil, fmt.Errorf("发货记录状态错误或不存在")
	}

	return rows.Get(0), nil
}

//CreateFalseSuccReturn 创建假成功退货信息
func CreateFalseSuccReturn(dbTrans db.IDBTrans, info types.IXMap) (int64, error) {
	//锁订单
	rows, _, _, err := dbTrans.Query(sql.LockFalseSuccOrder, map[string]interface{}{
		"order_id": info.GetInt64("order_id"),
	})
	if err != nil || rows.IsEmpty() {
		return 0, fmt.Errorf("锁订单发生异常,count:%d,err:%v", rows.Len(), err)
	}

	//检查退款是否存在
	cn, _, _, err := dbTrans.Scalar(sql.CheckFalseSuccReturn, map[string]interface{}{
		"delivery_id": info.GetInt64("delivery_id"),
	})
	if err != nil {
		return 0, err
	}
	if types.GetInt(cn) > 0 {
		return 0, fmt.Errorf("退货已存在")
	}

	//创建退货记录
	id, count, _, _, err := utils.Insert(dbTrans, sql.GetReturnID, sql.CreateFalseSuccReturn, map[string]interface{}{
		"delivery_id": info.GetInt64("delivery_id"),
	})

	if err != nil || count != 1 {
		return 0, fmt.Errorf("创建假成功退货记录发生异常，count:%d,err:%v", count, err)
	}

	return id, nil
}
