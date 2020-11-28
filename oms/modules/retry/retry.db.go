package retry

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/oms/modules/const/sql"

	"github.com/micro-plat/lib4go/db"
)

//GetBindInfo 获取绑定信息
func GetBindInfo(dbTrans db.IDBTrans, orderID int64) error {

	datas, _, _, err := dbTrans.Query(sql.GetBindInfo, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || datas.IsEmpty() {
		return fmt.Errorf("获取绑定信息失败，count:%d,err:%v", datas.Len(), err)
	}

	return nil
}

//UpdateDeliveryUpPayStatus 上游支付状态更新
func UpdateDeliveryUpPayStatus(dbTrans db.IDBTrans, deliveryID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateDeliveryUpPayStatus, map[string]interface{}{
		"delivery_id": deliveryID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改上游支付状态为等待状态失败，count:%d,err:%v", count, err)
	}

	return nil
}

//UpdateOrderNotifyStatus 修改订单通知状态
func UpdateOrderNotifyStatus(dbTrans db.IDBTrans, orderID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateOrderNotifyStatus, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改订单通知状态为等待状态失败，count:%d,err:%v", count, err)
	}
	return nil
}

//UpdateReturnStatus 修改退款退货状态
func UpdateReturnStatus(dbTrans db.IDBTrans, refundID int64, returnID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateReturnOvertime, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改退款超时时间失败，count:%d,err:%v", count, err)
	}

	count, _, _, err = dbTrans.Execute(sql.UpdateReturnStatus, map[string]interface{}{
		"refund_id": refundID,
		"return_id": returnID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改退款退货状态为等待状态失败，count:%d,err:%v", count, err)
	}

	return nil
}

//UpdateUpRefundStatus 修改上游退款状态
func UpdateUpRefundStatus(dbTrans db.IDBTrans, returnID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateUpRefundStatus, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改上游退款状态为等待状态失败，count:%d,err:%v", count, err)
	}
	return nil
}

//UpdateDownRefundStatus 修改下游退款状态
func UpdateDownRefundStatus(dbTrans db.IDBTrans, refundID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateDownRefundStatus, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改下游退款状态为等待状态失败，count:%d,err:%v", count, err)
	}

	return nil
}

//UpdateRefundNotifyStatus 修改退款通知状态
func UpdateRefundNotifyStatus(dbTrans db.IDBTrans, orderID int64, refundID int64) error {
	count, _, _, err := dbTrans.Execute(sql.UpdateRefundNotifyStatus, map[string]interface{}{
		"order_id":  orderID,
		"refund_id": refundID,
	})
	if err != nil || count != 1 {
		return fmt.Errorf("修改退款通知状态为等待状态失败，count:%d,err:%v", count, err)
	}

	return nil
}
