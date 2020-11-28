package channel

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateErrorCode 创建渠道错误码
type CreateErrorCode struct {
	ChannelNo     string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                //ChannelNo 渠道名称
	DealCode      string `json:"deal_code" form:"deal_code" m2s:"deal_code" valid:"required"`                   //DealCode 处理码
	ErrorCode     string `json:"error_code" form:"error_code" m2s:"error_code" valid:"required"`                //ErrorCode 错误码
	ErrorCodeDesc string `json:"error_code_desc" form:"error_code_desc" m2s:"error_code_desc" valid:"required"` //ErrorCodeDesc 错误码描述
	ServiceClass  string `json:"service_class" form:"service_class" m2s:"service_class" valid:"required"`       //ServiceClass 服务类型

}

//UpdateErrorCode 添加渠道错误码
type UpdateErrorCode struct {
	Id            string `json:"id" form:"id" m2s:"id" valid:"required"`                                        //Id 编号
	ChannelNo     string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                //ChannelNo 渠道名称
	DealCode      string `json:"deal_code" form:"deal_code" m2s:"deal_code" valid:"required"`                   //DealCode 处理码
	ErrorCode     string `json:"error_code" form:"error_code" m2s:"error_code" valid:"required"`                //ErrorCode 错误码
	ErrorCodeDesc string `json:"error_code_desc" form:"error_code_desc" m2s:"error_code_desc" valid:"required"` //ErrorCodeDesc 错误码描述
	ServiceClass  string `json:"service_class" form:"service_class" m2s:"service_class" valid:"required"`       //ServiceClass 服务类型

}

//QueryErrorCode 查询渠道错误码
type QueryErrorCode struct {
	ChannelNo    string `json:"channel_no" form:"channel_no" m2s:"channel_no"`          //ChannelNo 渠道名称
	ServiceClass string `json:"service_class" form:"service_class" m2s:"service_class"` //ServiceClass 服务类型

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbErrorCode  渠道错误码接口
type IDbErrorCode interface {
	//Create 创建
	Create(input *CreateErrorCode) error
	//Get 单条查询
	Get(id string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryErrorCode) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateErrorCode) (err error)
}

//DbErrorCode 渠道错误码对象
type DbErrorCode struct {
	c component.IContainer
}

//NewDbErrorCode 创建渠道错误码对象
func NewDbErrorCode(c component.IContainer) *DbErrorCode {
	return &DbErrorCode{
		c: c,
	}
}

//Create 添加渠道错误码
func (d *DbErrorCode) Create(input *CreateErrorCode) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertVdsChannelErrorCode, map[string]interface{}{
		"channel_no":      input.ChannelNo,
		"deal_code":       input.DealCode,
		"error_code":      input.ErrorCode,
		"error_code_desc": input.ErrorCodeDesc,
		"service_class":   input.ServiceClass,
	})
	if err != nil {
		return fmt.Errorf("添加渠道错误码数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据渠道错误码
func (d *DbErrorCode) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetVdsChannelErrorCode, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取渠道错误码数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取渠道错误码列表
func (d *DbErrorCode) Query(input *QueryErrorCode) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryVdsChannelErrorCodeCount, map[string]interface{}{
		"channel_no":    input.ChannelNo,
		"service_class": input.ServiceClass,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取渠道错误码列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryVdsChannelErrorCode, map[string]interface{}{
		"channel_no":    input.ChannelNo,
		"service_class": input.ServiceClass,
		"pi":            input.Pi,
		"ps":            input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取渠道错误码数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新渠道错误码
func (d *DbErrorCode) Update(input *UpdateErrorCode) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateVdsChannelErrorCode, map[string]interface{}{
		"id":              input.Id,
		"channel_no":      input.ChannelNo,
		"deal_code":       input.DealCode,
		"error_code":      input.ErrorCode,
		"error_code_desc": input.ErrorCodeDesc,
		"service_class":   input.ServiceClass,
	})
	if err != nil {
		return fmt.Errorf("更新渠道错误码数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
