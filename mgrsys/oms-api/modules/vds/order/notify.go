package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryNotify 查询发货通知记录表
type QueryNotify struct {
	CoopId      string `json:"coop_id" form:"coop_id" m2s:"coop_id"`                   //CoopId 下游商户
	CoopOrderId string `json:"coop_order_id" form:"coop_order_id" m2s:"coop_order_id"` //CoopOrderId 下游商户订单号
	OrderNo     string `json:"order_no" form:"order_no" m2s:"order_no"`                //OrderNo 订单号
	Status      string `json:"status" form:"status" m2s:"status"`                      //Status 通知状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbNotify  发货通知记录表接口
type IDbNotify interface {
	//Get 单条查询
	Get(id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryNotify) (data db.QueryRows, count int, err error)
}

//DbNotify 发货通知记录表对象
type DbNotify struct {
	c component.IContainer
}

//NewDbNotify 创建发货通知记录表对象
func NewDbNotify(c component.IContainer) *DbNotify {
	return &DbNotify{
		c: c,
	}
}

//Get 查询单条数据发货通知记录表
func (d *DbNotify) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsOrderNotify, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货通知记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取发货通知记录表列表
func (d *DbNotify) Query(input *QueryNotify) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsOrderNotifyCount, map[string]interface{}{
		"coop_id":       input.CoopId,
		"coop_order_id": input.CoopOrderId,
		"order_no":      input.OrderNo,
		"status":        input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货通知记录表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsOrderNotify, map[string]interface{}{
		"coop_id":       input.CoopId,
		"coop_order_id": input.CoopOrderId,
		"order_no":      input.OrderNo,
		"status":        input.Status,
		"pi":            input.Pi,
		"ps":            input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货通知记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
