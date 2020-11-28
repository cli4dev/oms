package life

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryTime 查询生命周期记录表
type QueryTime struct {
	OrderNo    string `json:"order_no" form:"order_no" m2s:"order_no"`          //OrderNo 业务单据号
	BatchNo    string `json:"batch_no" form:"batch_no" m2s:"batch_no"`          //BatchNo 业务批次号
	Content    string `json:"content" form:"content" m2s:"content"`             //Content 内容
	CreateTime string `json:"create_time" form:"create_time" m2s:"create_time"` //CreateTime 创建时间
	Ip         string `json:"ip" form:"ip" m2s:"ip"`                            //Ip 服务器ip

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbTime  生命周期记录表接口
type IDbTime interface {
	//Get 单条查询
	Get(id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryTime) (data db.QueryRows, count int, err error)
}

//DbTime 生命周期记录表对象
type DbTime struct {
	c component.IContainer
}

//NewDbTime 创建生命周期记录表对象
func NewDbTime(c component.IContainer) *DbTime {
	return &DbTime{
		c: c,
	}
}

//Get 查询单条数据生命周期记录表
func (d *DbTime) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetLcsLifeTime, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取生命周期记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取生命周期记录表列表
func (d *DbTime) Query(input *QueryTime) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryLcsLifeTimeCount, map[string]interface{}{
		"order_no":    input.OrderNo,
		"batch_no":    input.BatchNo,
		"content":     input.Content,
		"create_time": input.CreateTime,
		"ip":          input.Ip,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取生命周期记录表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryLcsLifeTime, map[string]interface{}{
		"order_no":    input.OrderNo,
		"batch_no":    input.BatchNo,
		"content":     input.Content,
		"create_time": input.CreateTime,
		"ip":          input.Ip,
		"pi":          input.Pi,
		"ps":          input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取生命周期记录表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
