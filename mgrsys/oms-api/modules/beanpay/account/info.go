package account

import (
	"fmt"

	"github.com/micro-plat/beanpay/beanpay"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/lib4go/utility"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateInfo 创建账户信息
type CreateInfo struct {
	AccountName string `json:"account_name" form:"account_name" m2s:"account_name" valid:"required"` //AccountName 帐户名称
	Eid         string `json:"eid" form:"eid" m2s:"eid" valid:"required"`                            //Eid 外部用户账户编号
	Groups      string `json:"groups" form:"groups" m2s:"groups" valid:"required"`                   //Groups 账户类型
	Ident       string `json:"ident" form:"ident" m2s:"ident" valid:"required"`                      //Ident 系统标识

}

//UpdateInfo 添加账户信息
type UpdateInfo struct {
	AccountId   string `json:"account_id" form:"account_id" m2s:"account_id" valid:"required"`       //AccountId 帐户编号
	AccountName string `json:"account_name" form:"account_name" m2s:"account_name" valid:"required"` //AccountName 帐户名称
	Eid         string `json:"eid" form:"eid" m2s:"eid" valid:"required"`                            //Eid 外部用户账户编号
	Ident       string `json:"ident" form:"ident" m2s:"ident" valid:"required"`                      //Ident 系统标识
	Status      string `json:"status" form:"status" m2s:"status" valid:"required"`                   //Status 账户状态
}

type AddInfo struct {
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"` //Eid 外部用户账户编号
	Amount    int    `json:"amount" form:"amount" m2s:"amount" valid:"required"`
	Types     string `json:"types" form:"types" m2s:"types" valid:"required"`
}

//QueryInfo 查询账户信息
type QueryInfo struct {
	AccountName string `json:"account_name" form:"account_name" m2s:"account_name"` //AccountName 帐户名称
	Eid         string `json:"eid" form:"eid" m2s:"eid"`                            //Eid 外部用户账户编号
	Groups      string `json:"groups" form:"groups" m2s:"groups"`                   //Groups 账户类型
	Ident       string `json:"ident" form:"ident" m2s:"ident"`                      //Ident 系统标识
	Status      string `json:"status" form:"status" m2s:"status"`                   //Status 账户状态
	Types       string `json:"types" form:"types" m2s:"types"`
	Pi          string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps          string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  账户信息接口
type IDbInfo interface {
	//Create 创建
	Create(input *CreateInfo) error
	//Get 单条查询
	Get(accountId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateInfo) (err error)
	//GetDownInfoDictionary 获取数据字典
	GetDownInfoDictionary() (db.QueryRows, error)
	GetUpInfoDictionary() (db.QueryRows, error)

	Add(input *AddInfo) error
	Draw(input *AddInfo) error
}

//DbInfo 账户信息对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建账户信息对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//GetDownInfoDictionary 获取数据字典
func (d *DbInfo) GetDownInfoDictionary() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetBeanpayDownAccountInfoDictionary, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取账户信息数据字典发生错误")
	}
	return data, nil
}

func (d *DbInfo) GetUpInfoDictionary() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetBeanpayUpAccountInfoDictionary, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取账户信息数据字典发生错误")
	}
	return data, nil
}

//Create 添加账户信息
func (d *DbInfo) Create(input *CreateInfo) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertBeanpayAccountInfo, map[string]interface{}{
		"account_name": input.AccountName,
		"eid":          input.Eid,
		"groups":       input.Groups,
		"ident":        input.Ident,
	})
	if err != nil {
		return fmt.Errorf("添加账户信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据账户信息
func (d *DbInfo) Get(accountId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetBeanpayAccountInfo, map[string]interface{}{
		"account_id": accountId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取账户信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取账户信息列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryBeanpayAccountInfoCount, map[string]interface{}{
		"account_name": input.AccountName,
		"eid":          input.Eid,
		"groups":       input.Groups,
		"ident":        input.Ident,
		"status":       input.Status,
		"types":        input.Types,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取账户信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryBeanpayAccountInfo, map[string]interface{}{
		"account_name": input.AccountName,
		"eid":          input.Eid,
		"groups":       input.Groups,
		"ident":        input.Ident,
		"status":       input.Status,
		"types":        input.Types,
		"pi":           input.Pi,
		"ps":           input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取账户信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新账户信息
func (d *DbInfo) Update(input *UpdateInfo) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateBeanpayAccountInfo, map[string]interface{}{
		"account_id":   input.AccountId,
		"account_name": input.AccountName,
		"eid":          input.Eid,
		"ident":        input.Ident,
		"status":       input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新账户信息数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}

//Add 加款
func (d *DbInfo) Add(input *AddInfo) error {

	account := beanpay.GetAccount(d.c.GetPlatName(), input.Types)
	result, err := account.AddAmount(d.c, input.ChannelNO, utility.GetGUID(), input.Amount*100)
	if err != nil || result.GetCode() != 200 {
		return err
	}
	return nil
}

//Draw 提款
func (d *DbInfo) Draw(input *AddInfo) error {
	account := beanpay.GetAccount(d.c.GetPlatName(), input.Types)
	result, err := account.DrawingAmount(d.c, input.ChannelNO, utility.GetGUID(), input.Amount*100)
	if err != nil || result.GetCode() != 200 {
		return err
	}
	return nil
}
