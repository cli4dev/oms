package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryQuery 查询发货结果查询记录表
type QueryQuery struct {
	ChannelNo string `json:"channel_no" form:"channel_no" m2s:"channel_no"` //ChannelNo 上游渠道
	Status    string `json:"status" form:"status" m2s:"status"`             //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbQuery  发货结果查询记录表接口
type IDbQuery interface {
	//Get 单条查询
	Get(orderNo string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryQuery) (data db.QueryRows, count int, err error)
}

//DbQuery 发货结果查询记录表对象
type DbQuery struct {
	c component.IContainer
}

//NewDbQuery 创建发货结果查询记录表对象
func NewDbQuery(c component.IContainer) *DbQuery {
	return &DbQuery{
		c: c,
	}
}

//Get 查询单条数据发货结果查询记录表
func (d *DbQuery) Get(orderNo string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsOrderQuery, map[string]interface{}{
		"order_no": orderNo,
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货结果查询记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取发货结果查询记录表列表
func (d *DbQuery) Query(input *QueryQuery) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsOrderQueryCount, map[string]interface{}{
		"channel_no": input.ChannelNo,
		"status":     input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货结果查询记录表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsOrderQuery, map[string]interface{}{
		"channel_no": input.ChannelNo,
		"status":     input.Status,
		"pi":         input.Pi,
		"ps":         input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货结果查询记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
