package dictionary

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateInfo 创建字典表
type CreateInfo struct {
	Name   string `json:"name" form:"name" m2s:"name" valid:"required"`          //Name 名称
	SortNo string `json:"sort_no" form:"sort_no" m2s:"sort_no" valid:"required"` //SortNo 排序值
	Type   string `json:"type" form:"type" m2s:"type" valid:"required"`          //Type 类型
	Value  string `json:"value" form:"value" m2s:"value" valid:"required"`       //Value 值

}

//UpdateInfo 添加字典表
type UpdateInfo struct {
	Id     string `json:"id" form:"id" m2s:"id" valid:"required"`                //Id 编号
	Name   string `json:"name" form:"name" m2s:"name" valid:"required"`          //Name 名称
	SortNo string `json:"sort_no" form:"sort_no" m2s:"sort_no" valid:"required"` //SortNo 排序值
	Status string `json:"status" form:"status" m2s:"status" valid:"required"`    //Status 状态
	Type   string `json:"type" form:"type" m2s:"type" valid:"required"`          //Type 类型
	Value  string `json:"value" form:"value" m2s:"value" valid:"required"`       //Value 值

}

//QueryInfo 查询字典表
type QueryInfo struct {
	Name   string `json:"name" form:"name" m2s:"name"`       //Name 名称
	Status string `json:"status" form:"status" m2s:"status"` //Status 状态
	Type   string `json:"type" form:"type" m2s:"type"`       //Type 类型

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  字典表接口
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

//DbInfo 字典表对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建字典表对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Create 添加字典表
func (d *DbInfo) Create(input *CreateInfo) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertDdsDictionaryInfo, map[string]interface{}{
		"name":    input.Name,
		"sort_no": input.SortNo,
		"type":    input.Type,
		"value":   input.Value,
	})
	if err != nil {
		return fmt.Errorf("添加字典表数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据字典表
func (d *DbInfo) Get(id string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetDdsDictionaryInfo, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("获取字典表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取字典表列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryDdsDictionaryInfoCount, map[string]interface{}{
		"name":   input.Name,
		"status": input.Status,
		"type":   input.Type,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取字典表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryDdsDictionaryInfo, map[string]interface{}{
		"name":   input.Name,
		"status": input.Status,
		"type":   input.Type,
		"pi":     input.Pi,
		"ps":     input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取字典表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新字典表
func (d *DbInfo) Update(input *UpdateInfo) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateDdsDictionaryInfo, map[string]interface{}{
		"id":      input.Id,
		"name":    input.Name,
		"sort_no": input.SortNo,
		"status":  input.Status,
		"type":    input.Type,
		"value":   input.Value,
	})
	if err != nil {
		return fmt.Errorf("更新字典表数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
