package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryInfo 查询发货订单信息表
type QueryInfo struct {
	CarrierNo    string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`          //CarrierNo 运营商
	ChannelNo    string `json:"channel_no" form:"channel_no" m2s:"channel_no"`          //ChannelNo 上游渠道编号
	CoopId       string `json:"coop_id" form:"coop_id" m2s:"coop_id"`                   //CoopId 下游商户编号
	CoopOrderId  string `json:"coop_order_id" form:"coop_order_id" m2s:"coop_order_id"` //CoopOrderId 下游商户订单号
	CreateTime   string `json:"create_time" form:"create_time" m2s:"create_time"`       //CreateTime 创建时间
	ResultSource string `json:"result_source" form:"result_source" m2s:"result_source"` //ResultSource 发货结果来源
	ServiceClass string `json:"service_class" form:"service_class" m2s:"service_class"` //ServiceClass 服务类型
	Status       string `json:"status" form:"status" m2s:"status"`                      //Status 发货状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  发货订单信息表接口
type IDbInfo interface {
	//Get 单条查询
	Get(orderNo string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
}

//DbInfo 发货订单信息表对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建发货订单信息表对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Get 查询单条数据发货订单信息表
func (d *DbInfo) Get(orderNo string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsOrderInfo, map[string]interface{}{
		"order_no": orderNo,
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货订单信息表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取发货订单信息表列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsOrderInfoCount, map[string]interface{}{
		"carrier_no":    input.CarrierNo,
		"channel_no":    input.ChannelNo,
		"coop_id":       input.CoopId,
		"coop_order_id": input.CoopOrderId,
		"create_time":   input.CreateTime,
		"result_source": input.ResultSource,
		"service_class": input.ServiceClass,
		"status":        input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货订单信息表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsOrderInfo, map[string]interface{}{
		"carrier_no":    input.CarrierNo,
		"channel_no":    input.ChannelNo,
		"coop_id":       input.CoopId,
		"coop_order_id": input.CoopOrderId,
		"create_time":   input.CreateTime,
		"result_source": input.ResultSource,
		"service_class": input.ServiceClass,
		"status":        input.Status,
		"pi":            input.Pi,
		"ps":            input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货订单信息表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
