package order

import (
	"fmt"
	"strings"

	"gitlab.100bm.cn/micro-plat/fas/fas/sdk"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBPay 支付数据层
type IDBPay interface {
	CheckOrderForDownPayByDB(orderID int64) ([]string, types.XMap, error)
	OrderDownPaySuccessByDB(trans db.IDBTrans, orderID int64) error
	CheckOrderAndDeliveryByDB(deliveryID int64) (types.XMap, bool, error)
	UpPaySuccessByDB(trans db.IDBTrans, param types.XMap) error
	DownChannelPayByDB(trans db.IDBTrans, data types.XMap) error
	UpChannelPayByDB(trans db.IDBTrans, data types.XMap) error
}

// DBPay 支付
type DBPay struct {
	c component.IContainer
}

// NewDBPay 构建DBPay
func NewDBPay(c component.IContainer) *DBPay {
	return &DBPay{c: c}
}

// CheckOrderForDownPayByDB 下游支付检查订单
func (p *DBPay) CheckOrderForDownPayByDB(orderID int64) ([]string, types.XMap, error) {
	db := p.c.GetRegularDB()
	datas, sqlStr, args, err := db.Query(sql.CheckOrderForDownPay, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("下游支付检查订单发生异常,cnt:%d,err:%v,sql:%v,args:%v", datas.Len(), err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil, nil
	}
	return []string{task.TaskType.BindTask}, datas.Get(0), nil
}

// OrderDownPaySuccessByDB 订单下游支付成功
func (p *DBPay) OrderDownPaySuccessByDB(trans db.IDBTrans, orderID int64) error {
	row, sqlStr, args, err := trans.Execute(sql.OrderDownPaySuccess, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || row != 1 {
		return fmt.Errorf("订单下游支付成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// CheckOrderAndDeliveryByDB 检查订单和订单发货
func (p *DBPay) CheckOrderAndDeliveryByDB(deliveryID int64) (types.XMap, bool, error) {

	// 1.检查订单发货
	db := p.c.GetRegularDB()
	datas, sqlStr, args, err := db.Query(sql.CheckDeliveryForUpPay, map[string]interface{}{
		"delivery_id": deliveryID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("上游支付检查发货信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", datas.Len(), err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, true, fmt.Errorf("上游支付时,订单发货不存在或已完成支付,delivery_id:%d", deliveryID)
	}

	// 2.检查订单
	data := datas.Get(0)
	orders, sqlStr, args, err := db.Query(sql.CheckOrderForUpPay, map[string]interface{}{
		"order_id": data.GetInt64("order_id"),
	})
	if err != nil {
		return nil, false, fmt.Errorf("上游支付检查订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", orders.Len(), err, sqlStr, args)
	}
	if orders.IsEmpty() {
		return nil, true, fmt.Errorf("上游支付时,订单不存在或已完成支付,order_id:%d", data.GetInt64("order_id"))
	}
	data.Merge(orders.Get(0))
	return data, false, nil
}

// UpPaySuccessByDB 上游支付成功
func (p *DBPay) UpPaySuccessByDB(trans db.IDBTrans, param types.XMap) error {
	// 1.判断是否完全支付
	deliverys, sqlStr, args, err := trans.Query(sql.QuerySuccessDeliveryFace, param)
	if err != nil || deliverys.IsEmpty() {
		return fmt.Errorf("查询发货成功总面值发生异常,cnt:%d,err:%v,sql:%v,args:%v", deliverys.Len(), err, sqlStr, args)
	}
	// 2.完全上游支付成功
	if deliverys.Get(0).GetFloat64("delivery_total_face")+param.GetFloat64("delivery_face") == param.GetFloat64("total_face") {
		row, sqlStr, args, err := trans.Execute(sql.OrderCompleteUpPay, param)
		if err != nil || row != 1 {
			return fmt.Errorf("完成上游支付成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
		}
	}
	// 3.发货支付成功
	row, sqlStr, args, err := trans.Execute(sql.DeliveryUpPaySuccess, param)
	if err != nil || row != 1 {
		return fmt.Errorf("上游支付成功,修改发货支付状态为支付成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	return nil
}

// DownChannelPayByDB 下游支付
func (p *DBPay) DownChannelPayByDB(trans db.IDBTrans, data types.XMap) error {
	config, err := p.c.GetVarConf("fd", "fd")
	if err != nil {
		return err
	}

	downChannel := data.GetString("channel_no")
	if data.GetString("account_no") != "" {
		downChannel = fmt.Sprintf("%s_%s", data.GetString("channel_no"), data.GetString("account_no"))
	}

	input := &sdk.DownPay{
		Ident:                p.c.GetPlatName(),
		OrderSource:          config.GetInt("order_source"),
		DownChannelNo:        downChannel,
		TradeOrderNo:         data.GetString("order_id"),
		ExtOrderNo:           data.GetString("request_no"),
		RechargeAccountNo:    data.GetString("recharge_account"),
		BusinessType:         data.GetInt("line_id"),
		CarrierNo:            data.GetString("carrier_no"),
		ProvinceNo:           data.GetString("province_no"),
		CityNo:               data.GetString("city_no"),
		TotalFace:            data.GetInt("total_face"),
		RechargeUnit:         data.GetInt("total_face"),
		OrderTime:            data.GetString("create_time"),
		DownSellAmount:       data.GetFloat64("sell_amount"),
		DownCommissionAmount: data.GetFloat64("commission_amount"),
		DownServiceAmount:    data.GetFloat64("service_amount"),
		DownFeeAmount:        data.GetFloat64("fee_amount"),
		Memo:                 "下游支付",
	}

	result, err := sdk.DownChannelPay(p.c, input)
	if err != nil {
		return err
	}

	if !strings.Contains(result.Result, "success") {
		return fmt.Errorf("result:%v,msg:%v", result.Result, result.Msg)
	}
	return nil
}

// UpChannelPayByDB 上游支付
func (p *DBPay) UpChannelPayByDB(trans db.IDBTrans, data types.XMap) error {
	config, err := p.c.GetVarConf("fd", "fd")
	if err != nil {
		return err
	}
	downChannel := data.GetString("down_channel_no")
	if data.GetString("down_account_no") != "" {
		downChannel = fmt.Sprintf("%s_%s", data.GetString("down_channel_no"), data.GetString("down_account_no"))
	}

	input := &sdk.UpPay{
		Ident:                p.c.GetPlatName(),
		OrderSource:          config.GetInt("order_source"),
		DownChannelNo:        downChannel,
		UpChannelNo:          data.GetString("up_channel_no"),
		TradeOrderNo:         data.GetString("order_id"),
		TradeDeliveryNo:      data.GetString("delivery_id"),
		ExtOrderNo:           data.GetString("request_no"),
		RechargeAccountNo:    data.GetString("recharge_account"),
		BillType:             1,
		BusinessType:         data.GetInt("line_id"),
		CarrierNo:            data.GetString("carrier_no"),
		ProvinceNo:           data.GetString("province_no"),
		CityNo:               data.GetString("city_no"),
		DownOrderUnit:        data.GetInt("down_order_face"),
		DownOrderFace:        data.GetInt("down_order_face"),
		DownDrawUnit:         data.GetInt("down_order_face"),
		DownDrawFace:         data.GetInt("down_order_face"),
		DownSellAmount:       data.GetFloat64("down_sell_amount"),
		DownCommissionAmount: data.GetFloat64("down_commission_amount"),
		DownServiceAmount:    data.GetFloat64("down_service_amount"),
		DownFeeAmount:        data.GetFloat64("down_fee_amount"),
		UpDrawUnit:           data.GetInt("up_face"),
		UpDrawFace:           data.GetInt("up_face"),
		UpCostAmount:         data.GetFloat64("cost_amount"),
		UpCommissionAmount:   data.GetFloat64("up_commission_amount"),
		UpServiceAmount:      data.GetFloat64("up_service_amount"),
		OrderTime:            data.GetString("order_time"),
		Memo:                 "上游支付",
	}
	result, err := sdk.UpChannelPay(p.c, input)
	if err != nil {
		return err
	}
	if !strings.Contains(result.Result, "success") {
		return fmt.Errorf("result:%v,msg:%v", result.Result, result.Msg)
	}

	return nil
}
