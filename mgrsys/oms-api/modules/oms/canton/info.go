package canton

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateInfo 创建省市信息
type CreateInfo struct {
	CantonCode   string `json:"canton_code" form:"canton_code" m2s:"canton_code" valid:"required"`    //CantonCode 区域编号
	ChineseName  string `json:"chinese_name" form:"chinese_name" m2s:"chinese_name" valid:"required"` //ChineseName 中文名称
	Spell        string `json:"spell" form:"spell" m2s:"spell" valid:"required"`                      //Spell 英文或全拼
	Grade        string `json:"grade" form:"grade" m2s:"grade" valid:"required"`                      //Grade 行政级别
	Parent       string `json:"parent" form:"parent" m2s:"parent" valid:"required"`                   //Parent 父级
	SimpleSpell  string `json:"simple_spell" form:"simple_spell" m2s:"simple_spell" valid:"required"` //SimpleSpell 简拼
	AreaCode     string `json:"area_code" form:"area_code" m2s:"area_code" valid:"required"`          //AreaCode 区号
	StandardCode string `json:"standard_code" form:"standard_code" m2s:"standard_code" valid:"required"`
}

//UpdateInfo 添加省市信息
type UpdateInfo struct {
	CantonCode   string `json:"canton_code" form:"canton_code" m2s:"canton_code" valid:"required"`    //CantonCode 区域编号
	ChineseName  string `json:"chinese_name" form:"chinese_name" m2s:"chinese_name" valid:"required"` //ChineseName 中文名称
	Spell        string `json:"spell" form:"spell" m2s:"spell" valid:"required"`                      //Spell 英文或全拼
	Grade        string `json:"grade" form:"grade" m2s:"grade" valid:"required"`                      //Grade 行政级别
	Parent       string `json:"parent" form:"parent" m2s:"parent" valid:"required"`                   //Parent 父级
	SimpleSpell  string `json:"simple_spell" form:"simple_spell" m2s:"simple_spell" valid:"required"` //SimpleSpell 简拼
	AreaCode     string `json:"area_code" form:"area_code" m2s:"area_code" valid:"required"`          //AreaCode 区号
	StandardCode string `json:"standard_code" form:"standard_code" m2s:"standard_code" valid:"required"`
}

//QueryInfo 查询省市信息
type QueryInfo struct {
	CantonCode  string `json:"canton_code" form:"canton_code" m2s:"canton_code"`    //CantonCode 区域编号
	ChineseName string `json:"chinese_name" form:"chinese_name" m2s:"chinese_name"` //ChineseName 中文名称
	Parent      string `json:"parent" form:"parent" m2s:"parent"`                   //Parent 父级
	SimpleSpell string `json:"simple_spell" form:"simple_spell" m2s:"simple_spell"` //SimpleSpell 简拼

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  省市信息接口
type IDbInfo interface {
	//Create 创建
	Create(input *CreateInfo) error
	//Get 单条查询
	Get(cantonCode string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateInfo) (err error)
	//GetInfoDictionary 获取数据字典
	GetInfoDictionary(grade int) (db.QueryRows, error)
	GetList() (db.QueryRows, error)
	GetInfoDictionaryByProvince(grade int, cantonCode string) (db.QueryRows, error)
}

//DbInfo 省市信息对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建省市信息对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

func (d *DbInfo) GetList() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsCantonInfoDictionaryQuery, map[string]interface{}{})

	if err != nil {
		return nil, fmt.Errorf("获取省市信息数据字典发生错误")
	}
	return data, nil
}

//GetInfoDictionary 获取数据字典
func (d *DbInfo) GetInfoDictionaryByProvince(grade int, cantonCode string) (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsCantonInfoDictionaryByProvince, map[string]interface{}{
		"grade":  grade,
		"parent": cantonCode,
	})

	if err != nil {
		return nil, fmt.Errorf("获取省市信息数据字典发生错误")
	}
	return data, nil
}

//GetInfoDictionary 获取数据字典
func (d *DbInfo) GetInfoDictionary(grade int) (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsCantonInfoDictionary, map[string]interface{}{
		"grade": grade,
	})
	if err != nil {
		return nil, fmt.Errorf("获取省市信息数据字典发生错误")
	}
	return data, nil
}

//Create 添加省市信息
func (d *DbInfo) Create(input *CreateInfo) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertOmsCantonInfo, map[string]interface{}{
		"canton_code":   input.CantonCode,
		"chinese_name":  input.ChineseName,
		"spell":         input.Spell,
		"grade":         input.Grade,
		"parent":        input.Parent,
		"simple_spell":  input.SimpleSpell,
		"area_code":     input.AreaCode,
		"standard_code": input.StandardCode,
	})
	if err != nil {
		return fmt.Errorf("添加省市信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据省市信息
func (d *DbInfo) Get(cantonCode string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsCantonInfo, map[string]interface{}{
		"canton_code": cantonCode,
	})
	if err != nil {
		return nil, fmt.Errorf("获取省市信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取省市信息列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsCantonInfoCount, map[string]interface{}{
		"canton_code":  input.CantonCode,
		"chinese_name": input.ChineseName,
		"parent":       input.Parent,
		"simple_spell": input.SimpleSpell,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取省市信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsCantonInfo, map[string]interface{}{
		"canton_code":  input.CantonCode,
		"chinese_name": input.ChineseName,
		"parent":       input.Parent,
		"simple_spell": input.SimpleSpell,
		"pi":           input.Pi,
		"ps":           input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取省市信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新省市信息
func (d *DbInfo) Update(input *UpdateInfo) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsCantonInfo, map[string]interface{}{
		"canton_code":   input.CantonCode,
		"chinese_name":  input.ChineseName,
		"spell":         input.Spell,
		"grade":         input.Grade,
		"parent":        input.Parent,
		"simple_spell":  input.SimpleSpell,
		"area_code":     input.AreaCode,
		"standard_code": input.StandardCode,
	})
	if err != nil {
		return fmt.Errorf("更新省市信息数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
