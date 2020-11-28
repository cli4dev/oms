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

//QueryRecord 查询账户余额变动信息
type QueryRecord struct {
	AccountId  string `json:"account_id" form:"account_id" m2s:"account_id"`    //AccountId 帐户名称
	ChangeType string `json:"change_type" form:"change_type" m2s:"change_type"` //ChangeType 变动类型
	TradeType  string `json:"trade_type" form:"trade_type" m2s:"trade_type"`    //TradeType 交易类型
	Types      string `json:"types" form:"types" m2s:"types"`
	StartTime  string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime    string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi         string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps         string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}
type RedrushInfo struct {
	TradeType int    `json:"trade_type" form:"trade_type" m2s:"trade_type" valid:"required"`
	TradeNo   string `json:"trade_no" form:"trade_no" m2s:"trade_no" valid:"required"`
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"` //Eid 外部用户账户编号
	Types     string `json:"types" form:"types" m2s:"types" valid:"required"`
}

type FlatAccountInfo struct {
	Amount    int    `json:"amount" form:"amount" m2s:"amount" valid:"required"`
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"` //Eid 外部用户账户编号
	Types     string `json:"types" form:"types" m2s:"types" valid:"required"`
}

//IDbRecord  账户余额变动信息接口
type IDbRecord interface {
	//Get 单条查询
	Get(recordId string) (db.QueryRow, error)

	//红冲
	Redrush(input *RedrushInfo) error

	RedrushDraw(input *RedrushInfo) error

	//交易平账
	FlatAccount(input *FlatAccountInfo) error

	//Query 列表查询
	Query(input *QueryRecord) (data db.QueryRows, count int, err error)
}

//DbRecord 账户余额变动信息对象
type DbRecord struct {
	c component.IContainer
}

//NewDbRecord 创建账户余额变动信息对象
func NewDbRecord(c component.IContainer) *DbRecord {
	return &DbRecord{
		c: c,
	}
}

//Get 查询单条数据账户余额变动信息
func (d *DbRecord) Get(recordId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetBeanpayAccountRecord, map[string]interface{}{
		"record_id": recordId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取账户余额变动信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取账户余额变动信息列表
func (d *DbRecord) Query(input *QueryRecord) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryBeanpayAccountRecordCount, map[string]interface{}{
		"account_id":  input.AccountId,
		"change_type": input.ChangeType,
		"trade_type":  input.TradeType,
		"start_time":  input.StartTime,
		"end_time":    input.EndTime,
		"types":       input.Types,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取账户余额变动信息列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryBeanpayAccountRecord, map[string]interface{}{
		"account_id":  input.AccountId,
		"change_type": input.ChangeType,
		"trade_type":  input.TradeType,
		"start_time":  input.StartTime,
		"end_time":    input.EndTime,
		"pi":          input.Pi,
		"ps":          input.Ps,
		"types":       input.Types,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取账户余额变动信息数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Redrush 加款红冲
func (d *DbRecord) Redrush(input *RedrushInfo) error {
	account := beanpay.GetAccount(d.c.GetPlatName(), input.Types)
	result, err := account.ReverseAddAmount(d.c, input.ChannelNO, utility.GetGUID(), input.TradeNo, input.TradeType)
	if err != nil || result.GetCode() != 200 {
		return err
	}
	return nil
}

//Redrush 提款红冲
func (d *DbRecord) RedrushDraw(input *RedrushInfo) error {
	account := beanpay.GetAccount(d.c.GetPlatName(), input.Types)
	result, err := account.ReverseDrawingAmount(d.c, input.ChannelNO, utility.GetGUID(), input.TradeNo, input.TradeType)
	if err != nil || result.GetCode() != 200 {
		return err
	}
	return nil
}

//FlatAccount 交易平账
func (d *DbRecord) FlatAccount(input *FlatAccountInfo) error {
	TradeType := 5
	account := beanpay.GetAccount(d.c.GetPlatName(), input.Types)
	result, err := account.TradeFlatAmount(d.c, input.ChannelNO, utility.GetGUID(), TradeType, input.Amount*100)
	if err != nil || result.GetCode() != 200 {
		return err
	}
	return nil
}
