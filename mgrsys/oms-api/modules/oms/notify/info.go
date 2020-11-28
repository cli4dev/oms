package notify

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryInfo 查询订单通知表
type QueryInfo struct {
	CreateTime   string `json:"create_time" form:"create_time" m2s:"create_time"`       //CreateTime 创建时间
	NotifyStatus string `json:"notify_status" form:"notify_status" m2s:"notify_status"` //NotifyStatus 通知状态（0成功，10未开始,20等待通知，30正在通知）
	NotifyType   string `json:"notify_type" form:"notify_type" m2s:"notify_type"`       //NotifyType 通知类型（1.订单通知，2.退款通知）
	OrderId      string `json:"order_id" form:"order_id" m2s:"order_id"`                //OrderId 订单编号
	RefundId     string `json:"notify_id" form:"notify_id" m2s:"notify_id"`             //RefundId 退款编号
	StartTime    string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime      string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi           string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps           string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  订单通知表接口
type IDbInfo interface {
	//Get 单条查询
	Get(notifyId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
}

//DbInfo 订单通知表对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建订单通知表对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Get 查询单条数据订单通知表
func (d *DbInfo) Get(notifyId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsNotifyInfo, map[string]interface{}{
		"notify_id": notifyId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取订单通知表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取订单通知表列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsNotifyInfoCount, map[string]interface{}{
		"end_time":      input.EndTime,
		"start_time":    input.StartTime,
		"notify_status": input.NotifyStatus,
		"notify_type":   input.NotifyType,
		"order_id":      input.OrderId,
		"notify_id":     input.RefundId,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单通知表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsNotifyInfo, map[string]interface{}{
		"end_time":      input.EndTime,
		"start_time":    input.StartTime,
		"notify_status": input.NotifyStatus,
		"notify_type":   input.NotifyType,
		"order_id":      input.OrderId,
		"notify_id":     input.RefundId,
		"pi":            input.Pi,
		"ps":            input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取订单通知表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
