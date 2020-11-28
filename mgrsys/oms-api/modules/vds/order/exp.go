package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryExp 查询发货异常订单记录表
type QueryExp struct {
	CarrierNo string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"` //CarrierNo 运营商
	ChannelNo string `json:"channel_no" form:"channel_no" m2s:"channel_no"` //ChannelNo 上游渠道
	CoopId    string `json:"coop_id" form:"coop_id" m2s:"coop_id"`          //CoopId 下游商户

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbExp  发货异常订单记录表接口
type IDbExp interface {
	//Get 单条查询
	Get(id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryExp) (data db.QueryRows, count int, err error)
}

//DbExp 发货异常订单记录表对象
type DbExp struct {
	c component.IContainer
}

//NewDbExp 创建发货异常订单记录表对象
func NewDbExp(c component.IContainer) *DbExp {
	return &DbExp{
		c: c,
	}
}

//Get 查询单条数据发货异常订单记录表
func (d *DbExp) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsOrderExp, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货异常订单记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取发货异常订单记录表列表
func (d *DbExp) Query(input *QueryExp) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsOrderExpCount, map[string]interface{}{
		"carrier_no": input.CarrierNo,
		"channel_no": input.ChannelNo,
		"coop_id":    input.CoopId,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货异常订单记录表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsOrderExp, map[string]interface{}{
		"carrier_no": input.CarrierNo,
		"channel_no": input.ChannelNo,
		"coop_id":    input.CoopId,
		"pi":         input.Pi,
		"ps":         input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货异常订单记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
