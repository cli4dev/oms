package channel

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateInfo 创建渠道基本信息
type CreateInfo struct {
	ChannelNo            string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                                     //ChannelNo 渠道名称
	ExtParams            string `json:"ext_params" form:"ext_params" m2s:"ext_params"`                                                      //ExtParams 预留字段
	FirstQueryTime       int    `json:"first_query_time" form:"first_query_time" m2s:"first_query_time"`                                    //FirstQueryTime 首次查询时间
	NotifyUrl            string `json:"notify_url" form:"notify_url" m2s:"notify_url" valid:"required"`                                     //NotifyUrl 通知回调地址
	QueryReplenishTime   int    `json:"query_replenish_time" form:"query_replenish_time" m2s:"query_replenish_time"`                        //QueryReplenishTime 查询后补间隔时间
	QueryUrl             string `json:"query_url" form:"query_url" m2s:"query_url"`                                                         //QueryUrl 查询地址
	RequestReplenishTime string `json:"request_replenish_time" form:"request_replenish_time" m2s:"request_replenish_time" valid:"required"` //RequestReplenishTime 发货后补间隔时间
	RequestUrl           string `json:"request_url" form:"request_url" m2s:"request_url" valid:"required"`                                  //RequestUrl 上游请求地址
	ServiceClass         string `json:"service_class" form:"service_class" m2s:"service_class" valid:"required"`                            //ServiceClass 服务类型

}

//UpdateInfo 添加渠道基本信息
type UpdateInfo struct {
	Id                   string `json:"id" form:"id" m2s:"id" valid:"required"`                                                             //Id id
	CanQuery             string `json:"can_query" form:"can_query" m2s:"can_query" valid:"required"`                                        //CanQuery 是否支持查询
	ChannelNo            string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                                     //ChannelNo 渠道名称
	ExtParams            string `json:"ext_params" form:"ext_params" m2s:"ext_params"`                                                      //ExtParams 预留字段
	FirstQueryTime       string `json:"first_query_time" form:"first_query_time" m2s:"first_query_time"`                                    //FirstQueryTime 首次查询时间
	NotifyUrl            string `json:"notify_url" form:"notify_url" m2s:"notify_url" valid:"required"`                                     //NotifyUrl 通知回调地址
	QueryReplenishTime   string `json:"query_replenish_time" form:"query_replenish_time" m2s:"query_replenish_time"`                        //QueryReplenishTime 查询后补间隔时间
	QueryUrl             string `json:"query_url" form:"query_url" m2s:"query_url"`                                                         //QueryUrl 查询地址
	RequestReplenishTime string `json:"request_replenish_time" form:"request_replenish_time" m2s:"request_replenish_time" valid:"required"` //RequestReplenishTime 发货后补间隔时间
	RequestUrl           string `json:"request_url" form:"request_url" m2s:"request_url" valid:"required"`                                  //RequestUrl 上游请求地址
	ServiceClass         string `json:"service_class" form:"service_class" m2s:"service_class" valid:"required"`                            //ServiceClass 服务类型
	Status               string `json:"status" form:"status" m2s:"status" valid:"required"`                                                 //Status 状态

}

//QueryInfo 查询渠道基本信息
type QueryInfo struct {
	CanQuery     string `json:"can_query" form:"can_query" m2s:"can_query"`             //CanQuery 是否支持查询
	ChannelNo    string `json:"channel_no" form:"channel_no" m2s:"channel_no"`          //ChannelNo 渠道名称
	ServiceClass string `json:"service_class" form:"service_class" m2s:"service_class"` //ServiceClass 服务类型
	Status       string `json:"status" form:"status" m2s:"status"`                      //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  渠道基本信息接口
type IDbInfo interface {
	//Create 创建
	Create(input *CreateInfo) error
	//Get 单条查询
	Get(id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateInfo) (err error)
}

//DbInfo 渠道基本信息对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建渠道基本信息对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Create 添加渠道基本信息
func (d *DbInfo) Create(input *CreateInfo) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertVdsChannelInfo, map[string]interface{}{
		"channel_no":             input.ChannelNo,
		"ext_params":             input.ExtParams,
		"first_query_time":       input.FirstQueryTime,
		"notify_url":             input.NotifyUrl,
		"query_replenish_time":   input.QueryReplenishTime,
		"query_url":              input.QueryUrl,
		"request_replenish_time": input.RequestReplenishTime,
		"request_url":            input.RequestUrl,
		"service_class":          input.ServiceClass,
	})
	if err != nil {
		return fmt.Errorf("添加渠道基本信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据渠道基本信息
func (d *DbInfo) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsChannelInfo, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取渠道基本信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取渠道基本信息列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsChannelInfoCount, map[string]interface{}{
		"can_query":     input.CanQuery,
		"channel_no":    input.ChannelNo,
		"service_class": input.ServiceClass,
		"status":        input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取渠道基本信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsChannelInfo, map[string]interface{}{
		"can_query":     input.CanQuery,
		"channel_no":    input.ChannelNo,
		"service_class": input.ServiceClass,
		"status":        input.Status,
		"pi":            input.Pi,
		"ps":            input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取渠道基本信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新渠道基本信息
func (d *DbInfo) Update(input *UpdateInfo) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateVdsChannelInfo, map[string]interface{}{
		"id":                     input.Id,
		"can_query":              input.CanQuery,
		"channel_no":             input.ChannelNo,
		"ext_params":             input.ExtParams,
		"first_query_time":       input.FirstQueryTime,
		"notify_url":             input.NotifyUrl,
		"query_replenish_time":   input.QueryReplenishTime,
		"query_url":              input.QueryUrl,
		"request_replenish_time": input.RequestReplenishTime,
		"request_url":            input.RequestUrl,
		"service_class":          input.ServiceClass,
		"status":                 input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新渠道基本信息数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
