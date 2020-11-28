package order

import (
	"encoding/json"
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/model"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBDelivery 数据层接口
type IDBDelivery interface {
	UpdateDeliveryStatusByDB(trans db.IDBTrans, deliveryID int64) error
	CheckDeliveryAndLockOrderByDB(trans db.IDBTrans, deliveryID int64) (types.XMap, bool, error)
	DeliveryBuildParam(dbTrans db.IDBTrans, data types.IXMap) (*model.OrderCreateParam, error)
	DeliveryFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error)
	CheckOrderAndDeliveryByDB(deliveryResult *DeliveryResult) (types.XMap, string, bool, error)
	BuildDeliveryServiceClass() int
	DeliverySuccessByDB(trans db.IDBTrans, data types.XMap) (notifyID int64, taskList []string, err error)
}

// DBDelivery 发货数据层
type DBDelivery struct {
	c component.IContainer
}

// NewDBDelivery 构建DBDelivery
func NewDBDelivery(c component.IContainer) *DBDelivery {
	return &DBDelivery{c: c}
}

// CheckOrderAndDeliveryByDB 检查订单和发货
func (d *DBDelivery) CheckOrderAndDeliveryByDB(deliveryResult *DeliveryResult) (types.XMap, string, bool, error) {

	db := d.c.GetRegularDB()
	// 1.检查发货
	deliverys, sqlStr, args, err := db.Query(sql.CheckDeliveryIn, map[string]interface{}{
		"delivery_id": deliveryResult.DeliveryID,
	})
	if err != nil {
		return nil, "", false, fmt.Errorf("检查发货信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if deliverys.IsEmpty() {
		return nil, "", true, fmt.Errorf("订单发货记录不存在或已完成发货，delivery_id:%v", deliveryResult.DeliveryID)
	}
	// 2.检查订单
	delivery := deliverys.Get(0)
	orderID := delivery.GetString("order_id")
	orders, sqlStr, args, err := db.Query(sql.CheckOrderForDelivery, map[string]interface{}{
		"order_id":      orderID,
		"delivery_face": delivery.GetFloat64("delivery_face"),
	})
	if err != nil {
		return nil, orderID, false, fmt.Errorf("发货检查订单信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if orders.IsEmpty() {
		return nil, orderID, true, fmt.Errorf("订单记录不存在或已完成发货，order_id:%v", orderID)
	}

	// 3.检查是否有通知
	notifys, sqlStr, args, err := db.Query(sql.CheckNotifyExist, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, orderID, false, fmt.Errorf("检查通知是否存在发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	delivery.Merge(orders.Get(0))
	delivery.SetValue("msg", deliveryResult.ResultMsg)
	delivery.SetValue("up_delivery_no", deliveryResult.OrderNo)
	delivery.SetValue("courier_cost_amount", deliveryResult.CourierCostAmount)
	delivery.SetValue("result_params", deliveryResult.ResultParams)
	if !notifys.IsEmpty() {
		delivery.SetValue("notify_id", notifys.Get(0).GetInt64("notify_id"))
	}

	return delivery, orderID, false, nil
}

// UpdateDeliveryStatusByDB 修改发货记录状态
func (d *DBDelivery) UpdateDeliveryStatusByDB(trans db.IDBTrans, deliveryID int64) error {
	row, sqlStr, args, err := trans.Execute(sql.UpdateDeliveryStatus, map[string]interface{}{
		"delivery_id": deliveryID,
	})
	if err != nil || row != 1 {
		return fmt.Errorf("修改发货记录状态发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// CheckDeliveryAndLockOrderByDB 锁订单和发货
func (d *DBDelivery) CheckDeliveryAndLockOrderByDB(trans db.IDBTrans, deliveryID int64) (types.XMap, bool, error) {
	// 1.检查发货
	db := d.c.GetRegularDB()
	deliverys, sqlStr, args, err := db.Query(sql.CheckDeliveryStatusForWaiting, map[string]interface{}{
		"delivery_id": deliveryID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查发货信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if deliverys.IsEmpty() {
		return nil, true, fmt.Errorf("订单发货记录不存在或已完成发货，delivery_id:%d", deliveryID)
	}
	delivery := deliverys.Get(0)

	// 2.锁订单
	orders, sqlStr, args, err := trans.Query(sql.LockOrderForDelivery, map[string]interface{}{
		"order_id": delivery.GetInt64("order_id"),
	})
	if err != nil {
		return nil, false, fmt.Errorf("发货锁订单发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if orders.IsEmpty() {
		return nil, true, fmt.Errorf("订单记录不存在或已完成发货，order_id:%d", delivery.GetInt64("order_id"))
	}

	delivery.Merge(orders.Get(0))
	return delivery, false, nil
}

// DeliveryBuildParam 构建发货系统参数
func (d *DBDelivery) DeliveryBuildParam(dbTrans db.IDBTrans, data types.IXMap) (*model.OrderCreateParam, error) {
	param := types.NewXMap()
	param.SetValue("ext_product_no", data.GetString("ext_product_no"))
	param.SetValue("ext_channel_no", data.GetString("ext_channel_no"))
	if data.GetString("extend_info") != "" {
		extend, err := types.NewXMapByJSON(data.GetString("extend_info"))
		if err != nil {
			return nil, err
		}
		param.Merge(extend)
	}
	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return &model.OrderCreateParam{
		CoopID:        data.GetString("down_channel_no"),
		CoopOrderID:   data.GetString("delivery_id"),
		ChannelNo:     data.GetString("up_channel_no"),
		ServiceClass:  d.BuildDeliveryServiceClass(),
		CarrierNo:     data.GetString("carrier_no"),
		ProductFace:   types.GetInt(data.GetFloat64("face") * 100),
		ProductNum:    data.GetInt("num"),
		NotifyURL:     queue.OrderFinishDelivery.GetName(d.c.GetPlatName()),
		OrderTimeout:  utils.GetDeliveryOvertime(data.GetInt("flow_overtime"), data.GetInt("delivery_overtime")),
		RequestParams: string(bytes),
	}, nil
}

// BuildDeliveryServiceClass 构建ServiceClass
func (d *DBDelivery) BuildDeliveryServiceClass() int {
	return 10
}

// DeliveryFailedByDB 发货失败
func (d *DBDelivery) DeliveryFailedByDB(trans db.IDBTrans, data types.XMap) ([]string, error) {

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
	if data.GetBool("complete_bind") && data.GetInt("flow_overtime") > 0 {
		return []string{task.TaskType.BindTask}, nil
	}
	return nil, nil
}

// DeliverySuccessByDB 发货成功
func (d *DBDelivery) DeliverySuccessByDB(trans db.IDBTrans, data types.XMap) (notifyID int64, taskList []string, err error) {

	// 1.修改订单信息
	extend, err := utils.GetExtendParamsByString(data.GetString("extend_info"), data.GetString("order_extend_info"))
	if err != nil {
		return 0, nil, fmt.Errorf("合并订单拓展信息,err:%v", err)
	}
	data.SetValue("extend_info", extend)
	row, sqlStr, args, err := trans.Execute(sql.UpdateOrderForDeliverySuccess, data)
	if err != nil || row != 1 {
		return 0, nil, fmt.Errorf("发货成功修改订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 2.修改发货记录为成功
	row, sqlStr, args, err = trans.Execute(sql.DeliverySuccess, data)
	if err != nil || row != 1 {
		return 0, nil, fmt.Errorf("发货成功修改发货记录发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	//3.判断发货是否完成
	tasks := []string{task.TaskType.UpPaymentTask}
	if !data.GetBool("all_success") {
		return 0, tasks, nil
	}

	//4.判断是否需要通知
	if data.GetInt64("notify_id") == 0 {
		return 0, tasks, nil
	}
	row, sqlStr, args, err = trans.Execute(sql.UpdateNotifyStatus, data)
	if err != nil || row != 1 {
		return 0, nil, fmt.Errorf("修改通知发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	return data.GetInt64("notify_id"), append(tasks, task.TaskType.NotifyTask), nil
}
