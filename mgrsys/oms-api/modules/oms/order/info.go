package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryInfo 查询订单记录
type QueryInfo struct {
	StartTime          string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime            string `json:"end_time" form:"end_time" m2s:"end_time"`
	OrderID            string `json:"order_id" form:"order_id" m2s:"order_id"`
	CanSplitOrder      string `json:"can_split_order" form:"can_split_order" m2s:"can_split_order"`                //CanSplitOrder 是否拆单
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`                               //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no"`                            //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no"`                                        //CityNo 城市
	CompleteUpPay      string `json:"complete_up_pay" form:"complete_up_pay" m2s:"complete_up_pay"`                //CompleteUpPay 完成上游支付（0已完成，1未完成）
	CreateTime         string `json:"create_time" form:"create_time" m2s:"create_time"`                            //CreateTime 创建时间
	DeliveryBindStatus string `json:"delivery_bind_status" form:"delivery_bind_status" m2s:"delivery_bind_status"` //DeliveryBindStatus 发货绑定状态（0发货成功，10.未开始，20.等待绑定，30.正在发货，90.全部失败）
	DeliveryPause      string `json:"delivery_pause" form:"delivery_pause" m2s:"delivery_pause"`                   //DeliveryPause 发货暂停（0.是，1否）
	DownChannelNo      string `json:"down_channel_no" form:"down_channel_no" m2s:"down_channel_no"`                //DownChannelNo 下游渠道名称
	DownProductId      string `json:"down_product_id" form:"down_product_id" m2s:"down_product_id"`                //DownProductId 下游商品编号
	DownShelfId        string `json:"down_shelf_id" form:"down_shelf_id" m2s:"down_shelf_id"`                      //DownShelfId 下游货架名称
	InvoiceType        string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type"`                         //InvoiceType 开票方式（1.不开发票，0.不限制，2.需要发票）
	IsRefund           string `json:"is_refund" form:"is_refund" m2s:"is_refund"`                                  //IsRefund 用户退款（0.是，1否）
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id"`                                        //LineId 产品线
	NotifyStatus       string `json:"notify_status" form:"notify_status" m2s:"notify_status"`                      //NotifyStatus 订单信息告知状态（0通知成功，100查询成功，10.未开始，20.等待告知，30.正在告知）
	OrderStatus        string `json:"order_status" form:"order_status" m2s:"order_status"`                         //OrderStatus 订单状态（10.支付，20.绑定发货，0.成功，90.失败，91.部分成功）
	PaymentStatus      string `json:"payment_status" form:"payment_status" m2s:"payment_status"`                   //PaymentStatus 支付状态（0支付成功，10.未开始，20.等待支付，30.正在支付，90.支付超时）
	RechageAccount     string `json:"rechage_account" form:"rechage_account" m2s:"rechage_account"`                //RechageAccount 充值账户
	RefundStatus       string `json:"refund_status" form:"refund_status" m2s:"refund_status"`                      //RefundStatus 订单失败退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99.无需退款）
	RequestNo          string `json:"request_no" form:"request_no" m2s:"request_no"`                               //RequestNo 下游渠道订单编号

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}
type QueryInfos struct {
	OrderID int64  `json:"order_id" form:"order_id" m2s:"order_id" valid:"required"` //OrderId 订单编号
	Pi      string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps      string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  订单记录接口
type IDbInfo interface {
	//Get 单条查询
	Get(orderId string) (db.QueryRow, error)

	GetVds(delivery_id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
	// 获取订单发货信息
	Delivery(input *QueryInfos) (datas types.XMaps, count int, err error)

	// 获取订单发货信息
	Audit(input *QueryInfos) (datas types.XMaps, count int, err error)
	// 获取订单发货信息
	Notify(input *QueryInfos) (datas types.XMaps, count int, err error)

	// 获取订单退款发货信息
	RefundNotify(input *QueryInfos) (datas types.XMaps, count int, err error)
	// 获取订单发货信息
	Refund(input *QueryInfos) (datas types.XMaps, count int, err error)
	// 获取订单发货信息
	Return(input *QueryInfos) (datas types.XMaps, count int, err error)
	// 获取订单发货信息
	Lifetime(input *QueryInfos) (datas types.XMaps, count int, err error)

	DownPay(input *QueryInfos) (datas types.XMaps, count int, err error)

	UpPay(input *QueryInfos) (datas types.XMaps, count int, err error)
}

//DbInfo 订单记录对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建订单记录对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Get 查询单条数据订单记录
func (d *DbInfo) Get(orderId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsOrderInfo, map[string]interface{}{
		"order_id": orderId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单记录数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

func (d *DbInfo) GetVds(delivery_id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	fmt.Println(delivery_id, "55555")
	data, q, a, err := db.Query(sql.GetVdsOrderInfoList, map[string]interface{}{
		"coop_order_id": delivery_id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单记录数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取订单记录列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsOrderInfoCount, map[string]interface{}{
		"can_split_order":      input.CanSplitOrder,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"complete_up_pay":      input.CompleteUpPay,
		"end_time":             input.EndTime,
		"start_time":           input.StartTime,
		"delivery_bind_status": input.DeliveryBindStatus,
		"delivery_pause":       input.DeliveryPause,
		"down_channel_no":      input.DownChannelNo,
		"down_product_id":      input.DownProductId,
		"down_shelf_id":        input.DownShelfId,
		"invoice_type":         input.InvoiceType,
		"is_refund":            input.IsRefund,
		"line_id":              input.LineId,
		"notify_status":        input.NotifyStatus,
		"order_status":         input.OrderStatus,
		"payment_status":       input.PaymentStatus,
		"rechage_account":      input.RechageAccount,
		"refund_status":        input.RefundStatus,
		"request_no":           input.RequestNo,
		"order_id":             input.OrderID,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单记录列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsOrderInfo, map[string]interface{}{
		"can_split_order":      input.CanSplitOrder,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"complete_up_pay":      input.CompleteUpPay,
		"end_time":             input.EndTime,
		"start_time":           input.StartTime,
		"delivery_bind_status": input.DeliveryBindStatus,
		"delivery_pause":       input.DeliveryPause,
		"down_channel_no":      input.DownChannelNo,
		"down_product_id":      input.DownProductId,
		"down_shelf_id":        input.DownShelfId,
		"invoice_type":         input.InvoiceType,
		"is_refund":            input.IsRefund,
		"line_id":              input.LineId,
		"notify_status":        input.NotifyStatus,
		"order_status":         input.OrderStatus,
		"payment_status":       input.PaymentStatus,
		"rechage_account":      input.RechageAccount,
		"refund_status":        input.RefundStatus,
		"request_no":           input.RequestNo,
		"order_id":             input.OrderID,
		"pi":                   input.Pi,
		"ps":                   input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单记录数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

// Delivery 获取订单发货信息
func (d *DbInfo) Delivery(input *QueryInfos) (datas types.XMaps, count int, err error) {
	db := d.c.GetRegularDB()
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	c, q, a, err := db.Scalar(sql.GetOrderDeliveryInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单发货列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderDeliveryInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单发货列表数据发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)
	return
}

// Notify 获取订单通知信息
func (d *DbInfo) Notify(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderNotifyInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单通知列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderNotifyInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单通知列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)
	return
}

// Notify 获取订单退款通知信息
func (d *DbInfo) RefundNotify(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderRefundNotifyInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退款通知列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderRefundNotifyInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退款通知列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)
	return
}

// Audit 获取订单审核信息
func (d *DbInfo) Audit(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderAuditInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单审核列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderAuditInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单审核列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)
	return
}

// Refund 获取订单退款信息
func (d *DbInfo) Refund(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderRefundInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退款列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderRefundInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退款列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)

	return
}

// Return 获取订单退货信息
func (d *DbInfo) Return(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderReturnInfoCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退货列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderReturnInfo, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单退货列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)

	return
}

// Lifetime 获取订单生命周期信息
func (d *DbInfo) Lifetime(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetOrderLifeTimeCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单生命周期列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetOrderLifeTime, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单生命周期列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)

	return
}

// DownPay 获取下游支付信息
func (d *DbInfo) DownPay(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetDownPayCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游支付信息条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetDownPayTime, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游支付信息列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)

	return
}

// UpPay 获取上游支付信息
func (d *DbInfo) UpPay(input *QueryInfos) (datas types.XMaps, count int, err error) {
	param := map[string]interface{}{
		"order_id": input.OrderID,
		"pi":       input.Pi,
		"ps":       input.Ps,
	}
	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.GetUpPayCount, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游支付信息条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}
	datas, q, a, err = db.Query(sql.GetUpPayTime, param)
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游支付信息列表发生错误(err:%v),sql:%s,参数：%v", err, q, a)
	}
	count = types.GetInt(c)

	return
}
