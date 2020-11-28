package audit

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/utils"
	au "gitlab.100bm.cn/micro-plat/oms/oms/modules/audit"
	"gitlab.100bm.cn/micro-plat/oms/oms/sdk"
)

//UpdateInfo 添加发货人工审核表
type UpdateInfo struct {
	AuditId     string `json:"audit_id" form:"audit_id" m2s:"audit_id" valid:"required"`             //AuditId 人工审核编号
	AuditBy     string `json:"audit_by" form:"audit_by" m2s:"audit_by"`                              //AuditBy 审核人
	AuditMsg    string `json:"audit_msg" form:"audit_msg" m2s:"audit_msg"`                           //AuditMsg 审核信息
	AuditStatus string `json:"audit_status" form:"audit_status" m2s:"audit_status" valid:"required"` //AuditStatus 审核状态
	AuditTime   string `json:"audit_time" form:"audit_time" m2s:"audit_time"`                        //AuditTime 审核时间

}

//QueryInfo 查询发货人工审核表
type QueryInfo struct {
	ChangeType string `json:"change_type" form:"change_type" m2s:"change_type"` //ChangeType 变动类型（1.发货，2.退货，3.订单，4.退款）
	CreateTime string `json:"create_time" form:"create_time" m2s:"create_time"` //CreateTime 创建时间
	DeliveryId string `json:"delivery_id" form:"delivery_id" m2s:"delivery_id"` //DeliveryId 发货记录编号
	OrderId    string `json:"order_id" form:"order_id" m2s:"order_id"`          //OrderId 订单编号
	RefundId   string `json:"refund_id" form:"refund_id" m2s:"refund_id"`       //RefundId 退款编号
	StartTime  string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime    string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi         string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps         string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbInfo  发货人工审核表接口
type IDbInfo interface {
	//Get 单条查询
	Get(auditId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryInfo) (data db.QueryRows, count int, err error)
	//Update 更新
	Audit(input *au.RequestParams, status int, changeType int) (err error)
}

//DbInfo 发货人工审核表对象
type DbInfo struct {
	c component.IContainer
}

//NewDbInfo 创建发货人工审核表对象
func NewDbInfo(c component.IContainer) *DbInfo {
	return &DbInfo{
		c: c,
	}
}

//Get 查询单条数据发货人工审核表
func (d *DbInfo) Get(auditId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsAuditInfo, map[string]interface{}{
		"audit_id": auditId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取发货人工审核表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取发货人工审核表列表
func (d *DbInfo) Query(input *QueryInfo) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsAuditInfoCount, map[string]interface{}{
		"change_type": input.ChangeType,
		"end_time":    input.EndTime,
		"start_time":  input.StartTime,
		"delivery_id": input.DeliveryId,
		"order_id":    input.OrderId,
		"refund_id":   input.RefundId,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货人工审核表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsAuditInfo, map[string]interface{}{
		"change_type": input.ChangeType,
		"end_time":    input.EndTime,
		"start_time":  input.StartTime,
		"delivery_id": input.DeliveryId,
		"order_id":    input.OrderId,
		"refund_id":   input.RefundId,
		"pi":          input.Pi,
		"ps":          input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取发货人工审核表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Audit 人工审核发货
func (d *DbInfo) Audit(input *au.RequestParams, status int, changeType int) (err error) {
	var queueFunc []func(c interface{}) error
	switch changeType {
	case enum.Delivery:
		switch status {
		case enum.AuditSuccess:
			queueFunc, err = sdk.DeliveryUnknownAuditSucc(d.c, input)
		case enum.AuditFailed:
			queueFunc, err = sdk.DeliveryUnknownAuditFail(d.c, input)
		case enum.AuditFalseSucc:
			queueFunc, err = sdk.DeliveryFalseSuccAudit(d.c, input.DeliveryID)
		default:
			return fmt.Errorf("审核单号:%v,审核状态:%v,审核类型:%v,错误", input.AuditID, status, changeType)
		}
	case enum.Return:
		switch status {
		case enum.AuditSuccess:
			queueFunc, err = sdk.ReturnUnknownAuditSucc(d.c, input)
		case enum.AuditFailed:
			queueFunc, err = sdk.ReturnUnknownAuditFail(d.c, input)
		default:
			return fmt.Errorf("审核单号:%v,审核状态:%v,审核类型:%v,错误", input.AuditID, status, changeType)
		}
	case enum.Order:
		switch status {
		case enum.AuditPartialSuccess:
			queueFunc, err = sdk.OrderPartialSuccessAudit(d.c, input.OrderID)
		default:
			return fmt.Errorf("审核单号:%v,审核状态:%v,审核类型:%v,错误", input.AuditID, status, changeType)
		}
	case enum.Refund:
		switch status {
		case enum.AuditPartialSuccess:
			queueFunc, err = sdk.RefundPartialSuccessAudit(d.c, input.OrderID, input.RefundID)
		default:
			return fmt.Errorf("审核单号:%v,审核状态:%v,审核类型:%v,错误", input.AuditID, status, changeType)
		}
	default:
		return fmt.Errorf("审核单号:%v,审核状态:%v,审核类型:%v,错误", input.AuditID, status, changeType)
	}
	if err != nil {
		return
	}
	return utils.CallBackFunc(d.c, queueFunc)
}
