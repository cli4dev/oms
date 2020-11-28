package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryDelivery 查询订单发货表
type QueryDelivery struct {
	DeliveryId      string `json:"delivery_id" form:"delivery_id" m2s:"delivery_id"`                   //DeliveryId 发货编号
	CarrierNo       string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`                      //CarrierNo 运营商
	CityNo          string `json:"city_no" form:"city_no" m2s:"city_no"`                               //CityNo 城市
	CreateTime      string `json:"create_time" form:"create_time" m2s:"create_time"`                   //CreateTime 创建时间
	DeliveryStatus  string `json:"delivery_status" form:"delivery_status" m2s:"delivery_status"`       //DeliveryStatus 发货状态（0.发货成功，20等待发货，30正在发货，90发货失败）
	DownChannelNo   string `json:"down_channel_no" form:"down_channel_no" m2s:"down_channel_no"`       //DownChannelNo 下游渠道
	InvoiceType     string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type"`                //InvoiceType 开票方式（1.不开发票，2.上游开发票）
	LineId          string `json:"line_id" form:"line_id" m2s:"line_id"`                               //LineId 产品线
	ProvinceNo      string `json:"province_no" form:"province_no" m2s:"province_no"`                   //ProvinceNo 省份
	UpPaymentStatus string `json:"up_payment_status" form:"up_payment_status" m2s:"up_payment_status"` //UpPaymentStatus 上游支付状态（0支付成功，10未开始,20.等待支付，30.正在支付，99.无需支付）
	StartTime       string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime         string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi              string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps              string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbDelivery  订单发货表接口
type IDbDelivery interface {
	//Get 单条查询
	Get(deliveryId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryDelivery) (data db.QueryRows, count int, err error)
}

//DbDelivery 订单发货表对象
type DbDelivery struct {
	c component.IContainer
}

//NewDbDelivery 创建订单发货表对象
func NewDbDelivery(c component.IContainer) *DbDelivery {
	return &DbDelivery{
		c: c,
	}
}

//Get 查询单条数据订单发货表
func (d *DbDelivery) Get(deliveryId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsOrderDelivery, map[string]interface{}{
		"delivery_id": deliveryId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单发货表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取订单发货表列表
func (d *DbDelivery) Query(input *QueryDelivery) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsOrderDeliveryCount, map[string]interface{}{
		"delivery_id":       input.DeliveryId,
		"carrier_no":        input.CarrierNo,
		"city_no":           input.CityNo,
		"end_time":          input.EndTime,
		"start_time":        input.StartTime,
		"delivery_status":   input.DeliveryStatus,
		"down_channel_no":   input.DownChannelNo,
		"invoice_type":      input.InvoiceType,
		"line_id":           input.LineId,
		"province_no":       input.ProvinceNo,
		"up_payment_status": input.UpPaymentStatus,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单发货表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsOrderDelivery, map[string]interface{}{
		"delivery_id":       input.DeliveryId,
		"carrier_no":        input.CarrierNo,
		"city_no":           input.CityNo,
		"end_time":          input.EndTime,
		"start_time":        input.StartTime,
		"delivery_status":   input.DeliveryStatus,
		"down_channel_no":   input.DownChannelNo,
		"invoice_type":      input.InvoiceType,
		"line_id":           input.LineId,
		"province_no":       input.ProvinceNo,
		"up_payment_status": input.UpPaymentStatus,
		"pi":                input.Pi,
		"ps":                input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单发货表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
