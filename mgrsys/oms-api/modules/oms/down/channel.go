package down

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/model"
)

//CreateChannel 创建下游渠道
type CreateChannel struct {
	ChannelNo   string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`       //ChannelNo 编号
	ChannelName string `json:"channel_name" form:"channel_name" m2s:"channel_name" valid:"required"` //ChannelName 名称

}

//UpdateChannel 添加下游渠道
type UpdateChannel struct {
	ChannelNo   string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`       //ChannelNo 编号
	ChannelName string `json:"channel_name" form:"channel_name" m2s:"channel_name" valid:"required"` //ChannelName 名称
	Status      string `json:"status" form:"status" m2s:"status" valid:"required"`                   //CreateTime 创建时间

}

//QueryChannel 查询下游渠道
type QueryChannel struct {
	ChannelNo   string `json:"channel_no" form:"channel_no" m2s:"channel_no"`       //ChannelNo 编号
	ChannelName string `json:"channel_name" form:"channel_name" m2s:"channel_name"` //ChannelName 名称

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbChannel  下游渠道接口
type IDbChannel interface {
	//Create 创建
	Create(input *CreateChannel) error
	//Get 单条查询
	Get(channelNo string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryChannel) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateChannel) (err error)
	//GetChannelDictionary 获取数据字典
	GetChannelDictionary() (db.QueryRows, error)
	//设置秘钥
	SetSecret(channelNo string) (err error)
	GetSecret(channelNo string) (secret string, err error)
}

//DbChannel 下游渠道对象
type DbChannel struct {
	c component.IContainer
}

//NewDbChannel 创建下游渠道对象
func NewDbChannel(c component.IContainer) *DbChannel {
	return &DbChannel{
		c: c,
	}
}

//GetChannelDictionary 获取数据字典
func (d *DbChannel) GetChannelDictionary() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsDownChannelDictionary, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取下游渠道数据字典发生错误")
	}
	return data, nil
}

//Create 添加下游渠道
func (d *DbChannel) Create(input *CreateChannel) error {

	// 添加下游渠道
	dbTrans, err := d.c.GetRegularDB().Begin()

	if err != nil {
		err = fmt.Errorf("开启事务失败:%v", err)
		return err
	}
	lastInsertID, affectedRow, q, a, err := dbTrans.Executes(sql.InsertOmsDownChannel, map[string]interface{}{
		"channel_no":   input.ChannelNo,
		"channel_name": input.ChannelName,
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加下游渠道数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}

	lastInsertID, affectedRow, q, a, err = dbTrans.Executes(sql.InsertBeanpayAccountInfo, map[string]interface{}{
		"account_name": fmt.Sprintf("%s渠道账户", input.ChannelName),
		"eid":          input.ChannelNo,
		"groups":       enum.DownChannel,
		"ident":        d.c.GetPlatName(),
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加渠道账户信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	lastInsertID, affectedRow, q, a, err = dbTrans.Executes(sql.InsertBeanpayAccountInfo, map[string]interface{}{
		"account_name": fmt.Sprintf("%s佣金账户", input.ChannelName),
		"eid":          input.ChannelNo,
		"groups":       enum.DownCommission,
		"ident":        d.c.GetPlatName(),
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加佣金账户信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	lastInsertID, affectedRow, q, a, err = dbTrans.Executes(sql.InsertBeanpayAccountInfo, map[string]interface{}{
		"account_name": fmt.Sprintf("%s服务费账户", input.ChannelName),
		"eid":          input.ChannelNo,
		"groups":       enum.DownService,
		"ident":        d.c.GetPlatName(),
	})
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("添加服务费账户信息数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	dbTrans.Commit()
	return nil
}

//Get 查询单条数据下游渠道
func (d *DbChannel) Get(channelNo string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsDownChannel, map[string]interface{}{
		"channel_no": channelNo,
	})
	if err != nil {
		return nil, fmt.Errorf("获取下游渠道数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取下游渠道列表
func (d *DbChannel) Query(input *QueryChannel) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsDownChannelCount, map[string]interface{}{
		"channel_no":   input.ChannelNo,
		"channel_name": input.ChannelName,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游渠道列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsDownChannel, map[string]interface{}{
		"channel_no":   input.ChannelNo,
		"channel_name": input.ChannelName,
		"pi":           input.Pi,
		"ps":           input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游渠道数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	fmt.Printf("data:%+v, q:%v, a:%v, err:%v ", data, q, a, err)
	return data, types.GetInt(c, 0), nil
}

//Update 更新下游渠道
func (d *DbChannel) Update(input *UpdateChannel) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsDownChannel, map[string]interface{}{
		"channel_no":   input.ChannelNo,
		"channel_name": input.ChannelName,
		"status":       input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新下游渠道数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}

func (d *DbChannel) SetSecret(channelNo string) (err error) {
	conf := model.GetConf(d.c)
	status, resultStr, params, err := d.c.Request(fmt.Sprintf("/%s/%s/%s/generate@%s.%s", conf.Ident, enum.DownChannelGroup, conf.Type, conf.ServerName, conf.PlatName),
		"GET",
		nil,
		map[string]interface{}{
			"euid": channelNo,
		}, true)
	if err != nil || status != 200 {
		err = fmt.Errorf("下游设置秘钥失败,state:%d,err:%+v,param:%+v,resultStr:%s", status, err, params, resultStr)
		return
	}
	return nil
}

func (d *DbChannel) GetSecret(channelNo string) (secret string, err error) {
	conf := model.GetConf(d.c)
	status, resultStr, params, err := d.c.Request(fmt.Sprintf("/%s/%s/%s/get@%s.%s", conf.Ident, enum.DownChannelGroup, conf.Type, conf.ServerName, conf.PlatName),
		"GET",
		nil,
		map[string]interface{}{
			"euid": channelNo,
		}, true)
	if err != nil || status != 200 {
		err = fmt.Errorf("下游获取秘钥秘钥失败,state:%d,err:%+v,param:%+v,resultStr:%s", status, err, params, resultStr)
		return
	}
	return resultStr, nil
}
