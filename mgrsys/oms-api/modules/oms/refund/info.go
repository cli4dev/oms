package refund

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryInfo 查询退款记录
type QueryInfo struct {
	RefundId           string `json:"refund_id" form:"refund_id" m2s:"refund_id"`                                  //RefundId 退款编号
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`                               //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no"`                            //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no"`                                        //CityNo 城市
	CompleteUpRefund   string `json:"complete_up_refund" form:"complete_up_refund" m2s:"complete_up_refund"`       //CompleteUpRefund 已完成上游退款（0.已完成，1.未完成）
	CreateTime         string `json:"create_time" form:"create_time" m2s:"create_time"`                            //CreateTime 创建时间
	DownChannelNo      string `json:"down_channel_no" form:"down_channel_no" m2s:"down_channel_no"`                //DownChannelNo 下游渠道名称
	DownRefundStatus   string `json:"down_refund_status" form:"down_refund_status" m2s:"down_refund_status"`       //DownRefundStatus 下游退款状态（0成功，10.未开始，20.等待，30正在，99无需）
	DownShelfId        string `json:"down_shelf_id" form:"down_shelf_id" m2s:"down_shelf_id"`                      //DownShelfId 下游货架
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id"`                                        //LineId 产品线
	OrderId            string `json:"order_id" form:"order_id" m2s:"order_id"`                                     //OrderId 订单编号
	RefundNotifyStatus string `json:"refund_notify_status" form:"refund_notify_status" m2s:"refund_notify_status"` //RefundNotifyStatus 退款通知状态（0成功，10.未开始，20.等待，30正在，99无需）
	RefundStatus       string `json:"refund_status" form:"refund_status" m2s:"refund_status"`                      //RefundStatus 状态（10.退货，20.退款）
	RefundType         string `json:"refund_type" form:"refund_type" m2s:"refund_type"`                            //RefundType 退款方式（1.普通退款，2.强制退款,3.假成功退款）
	UpReturnStatus     string `json:"up_return_status" form:"up_return_status" m2s:"up_return_status"`             //UpReturnStatus 上游退货状态（0.退货成功，20.等待退货，30.正在退货，90.退款失败，91.部分退款）
	StartTime          string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime            string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi                 string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps                 string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  退款记录接口
type IDbInfo interface {
	//Get 单条查询
	Get(refundId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
}

//DbInfo 退款记录对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建退款记录对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Get 查询单条数据退款记录
func (d *DbInfo) Get(refundId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsRefundInfo, map[string]interface{}{
		"refund_id": refundId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取退款记录数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取退款记录列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsRefundInfoCount, map[string]interface{}{
		"refund_id":            input.RefundId,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"complete_up_refund":   input.CompleteUpRefund,
		"end_time":             input.EndTime,
		"start_time":           input.StartTime,
		"down_channel_no":      input.DownChannelNo,
		"down_refund_status":   input.DownRefundStatus,
		"down_shelf_id":        input.DownShelfId,
		"line_id":              input.LineId,
		"order_id":             input.OrderId,
		"refund_notify_status": input.RefundNotifyStatus,
		"refund_status":        input.RefundStatus,
		"refund_type":          input.RefundType,
		"up_return_status":     input.UpReturnStatus,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取退款记录列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsRefundInfo, map[string]interface{}{
		"refund_id":            input.RefundId,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"complete_up_refund":   input.CompleteUpRefund,
		"end_time":             input.EndTime,
		"start_time":           input.StartTime,
		"down_channel_no":      input.DownChannelNo,
		"down_refund_status":   input.DownRefundStatus,
		"down_shelf_id":        input.DownShelfId,
		"line_id":              input.LineId,
		"order_id":             input.OrderId,
		"refund_notify_status": input.RefundNotifyStatus,
		"refund_status":        input.RefundStatus,
		"refund_type":          input.RefundType,
		"up_return_status":     input.UpReturnStatus,
		"pi":                   input.Pi,
		"ps":                   input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取退款记录数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
