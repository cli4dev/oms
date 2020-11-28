package refund

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// IDBRefund 数据层退款接口
type IDBRefund interface {
	CheckRefundAndOrderGeneralByDb(dbTrans db.IDBTrans, input *RequestBody) (types.IXMap, types.IXMap, error)
	CreateRefundAndNotifyByDB(dbTrans db.IDBTrans, orderID string, extendInfo string, input *RequestBody) (types.IXMap, error)
	PrcessRefundNotifyStatusByDB(db db.IDBExecuter, input types.XMap) error
	QueryRefundBodyByDB(dbt db.IDBExecuter, input *QueryRequestBody) (types.XMap, error)
	LockAndCheckRefundByDB(dbTrans db.IDBTrans, input *RequestBody) error
	CreateProductReturnByDB(dbTrans db.IDBTrans, input *RequestBody, refundInfo types.IXMap) ([]string, []string, types.XMaps, error)
	BuildRefundResult(input types.IXMap) (types.XMap, error)
}

// DBRefund 普通退款对象
type DBRefund struct {
	c component.IContainer
}

//NewDBRefund 创建退款对象
func NewDBRefund(c component.IContainer) *DBRefund {
	return &DBRefund{
		c: c,
	}
}

//CheckRefundAndOrderGeneralByDb 普通退款检查退款和订单
func (r *DBRefund) CheckRefundAndOrderGeneralByDb(dbTrans db.IDBTrans, input *RequestBody) (types.IXMap, types.IXMap, error) {
	// 检查退款是否存在
	refunds, q, a, err := dbTrans.Query(sql.SQLQueryRefund, map[string]interface{}{
		"down_channel_no": input.ChannelNo,
		"down_refund_no":  input.RefundNo,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("获取退款记录错误(err:%v),sql:%s,input:%+v,rows:%+v", err, q, a, refunds)
	}
	if !refunds.IsEmpty() {
		// 存在则返回
		return refunds.Get(0), nil, nil
	}

	// 检查订单基本信息
	orders, q, a, err := dbTrans.Query(sql.SQLQueryOrderInfo, map[string]interface{}{
		"down_channel_no": input.ChannelNo,
		"request_no":      input.RequestNo,
		"num":             input.RefundNum,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("获取订单信息错误(err:%v),sql:%s,input:%+v,rows:%+v", err, q, a, orders)
	}

	if orders.IsEmpty() {
		return nil, nil, context.NewErrorf(errorcode.ORDER_STATUS_ERROR.Code, "可退款订单不存在")
	}
	order := orders.Get(0)

	if order.GetInt("can_refund") == enum.CanRefund.CanNotRefund {
		return nil, nil, context.NewErrorf(errorcode.CAN_NOT_REFUND.Code, "不支持退款")
	}

	// 设置下个函数所需参数
	return nil, order, nil

}

// CreateRefundAndNotifyByDB 创建退款和通知
func (r *DBRefund) CreateRefundAndNotifyByDB(dbTrans db.IDBTrans, orderID string, extendInfo string, input *RequestBody) (types.IXMap, error) {
	// 创建退款记录
	refundID, count, sqlStr, args, err := utils.Insert(dbTrans, sql.SQLGetRefundID, sql.SQLCreateRefund, map[string]interface{}{
		"down_channel_no":   input.ChannelNo,
		"down_refund_no":    input.RefundNo,
		"refund_type":       enum.RefundType.General,
		"request_no":        input.RequestNo,
		"refund_num":        input.RefundNum,
		"notify_url":        input.NotifyURL,
		"order_id":          orderID,
		"refund_mer_amount": input.RefundMerAmount,
		"refund_point_num":  input.RefundPointNum,
		"extend_info":       extendInfo,
	})
	if err != nil || count != 1 {
		return nil, fmt.Errorf("创建退款记录错误(err:%+v),sqlStr:%s,args:%+v,count:%d", err, sqlStr, args, count)
	}

	// 查询退款记录
	refunds, q, a, err := dbTrans.Query(sql.QueryCreateRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || refunds.IsEmpty() {
		return nil, fmt.Errorf("获取退款记录错误(err:%v),sql:%s,input:%+v,list:%+v", err, q, a, refunds.Len())
	}

	// 判断是否需要通知
	if input.NotifyURL == "" {
		return refunds.Get(0), nil
	}

	// 创建退款通知记录
	_, row, sqlStr, args, err := utils.Insert(dbTrans, sql.GetNewID, sql.SQLCreateNotifyInfo, map[string]interface{}{
		"order_id":   orderID,
		"refund_id":  refundID,
		"notify_url": input.NotifyURL,
		"code":       0,
	})
	if err != nil || row <= 0 {
		return nil, fmt.Errorf("创建退款通知记录错误(err:%+v),sql:%s,input:%+v,row:%d", err, sqlStr, args, row)
	}
	return refunds.Get(0), nil
}

// PrcessRefundNotifyStatusByDB 更新告知状态
func (r *DBRefund) PrcessRefundNotifyStatusByDB(db db.IDBExecuter, input types.XMap) error {
	if input.GetInt("refund_notify_status") != enum.NotifyStatus.Processing {
		return nil
	}
	// 更新退款告知状态
	count, q, a, err := db.Execute(sql.SQLUpdateRefundNotifyToQuerySuccess, map[string]interface{}{
		"refund_id": input.GetString("refund_id"),
	})
	if err != nil {
		return fmt.Errorf("更新退款告知状态错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return nil
}

// QueryRefundBodyByDB 查询退款信息
func (r *DBRefund) QueryRefundBodyByDB(dbt db.IDBExecuter, input *QueryRequestBody) (types.XMap, error) {
	// 查询退款信息
	rows, q, a, err := dbt.Query(sql.SQLQueryRefundInfo, map[string]interface{}{
		"down_channel_no": input.ChannelNo,
		"down_refund_no":  input.RefundNo,
	})
	if err != nil {
		return nil, fmt.Errorf("查询退款信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	if rows.IsEmpty() {
		return nil, context.ERR_DataNotExist
	}
	return rows.Get(0), nil
}

// LockAndCheckRefundByDB 检查是否正在退款
func (r *DBRefund) LockAndCheckRefundByDB(dbTrans db.IDBTrans, input *RequestBody) error {
	// 锁定订单
	orders, q, a, err := dbTrans.Query(sql.SQLLockOrder, map[string]interface{}{
		"channel_no": input.ChannelNo,
		"request_no": input.RequestNo,
		"num":        input.RefundNum,
	})
	if err != nil || orders.IsEmpty() {
		return fmt.Errorf("锁定订单错误(err:%+v),sql:%s,arg:%+v,count:%d", err, q, a, orders.Len())
	}
	order := orders.Get(0)

	// 检查是否有正在进行的退款记录
	rows, q, a, err := dbTrans.Query(sql.SQLQueryProcessRefund, map[string]interface{}{
		"order_id": order.GetString("order_id"),
	})
	if err != nil {
		return fmt.Errorf("获取正在退款的记录错误(err:%v),sql:%s,arg:%+v,rows:%+v", err, q, a, rows)
	}
	if !rows.IsEmpty() {
		return context.NewErrorf(errorcode.REFUNDING_ERROR.Code, "该订单正在退款中，暂时无法退款！")
	}

	// 检查退款面值是否超过订单面值
	if !order.GetBool("can_refund") {
		return context.NewErrorf(errorcode.REFUND_NUM_ERROR.Code, "退款数量错误")
	}
	return nil
}

// CreateProductReturnByDB 创建退货
func (r *DBRefund) CreateProductReturnByDB(dbTrans db.IDBTrans, input *RequestBody, refundInfo types.IXMap) ([]string, []string, types.XMaps, error) {
	//1.获取全部成功发货的可退款数量
	dRows, q, a, err := dbTrans.Query(sql.SQLGetDeliveryDetailList, map[string]interface{}{
		"order_id": refundInfo.GetString("order_id"),
	})
	if err != nil || dRows.IsEmpty() {
		return nil, nil, nil, fmt.Errorf("获取发货编号和可退货数量错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, dRows.Len())
	}

	//2.根据退款数量依次创建成功发货对应的退货
	canRefundNum := input.RefundNum
	for _, v := range dRows {
		if v.GetInt("num") <= 0 {
			continue
		}

		//2.1发货记录可退款数量大于等于请求的退款数量，直接按请求退款数量退款
		if canRefundNum <= v.GetInt("num") {
			// 创建退货记录
			_, count, q, a, err := utils.Insert(dbTrans, sql.GetNewID, sql.SQLCreateUpReturn, map[string]interface{}{
				"ref_num":     canRefundNum,
				"delivery_id": v.GetString("delivery_id"),
				"refund_id":   refundInfo.GetString("refund_id"),
				"extend_info": refundInfo.GetString("extend_info"),
			})
			if err != nil || count <= 0 {
				return nil, nil, nil, fmt.Errorf("创建退货记录错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
			}

			break
		}
		// 创建退货记录
		_, count, q, a, err := utils.Insert(dbTrans, sql.GetNewID, sql.SQLCreateUpReturn, map[string]interface{}{
			"ref_num":     v.GetInt("num"),
			"delivery_id": v.GetString("delivery_id"),
			"refund_id":   refundInfo.GetString("refund_id"),
			"extend_info": refundInfo.GetString("extend_info"),
		})
		if err != nil || count <= 0 {
			return nil, nil, nil, fmt.Errorf("创建退货记录错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
		}
		canRefundNum -= v.GetInt("num")
	}

	// 返回退货结果
	rows, q, a, err := dbTrans.Query(sql.SQLGetCreatedReturnList, map[string]interface{}{
		"refund_id": refundInfo.GetString("refund_id"),
	})
	if err != nil || rows.IsEmpty() {
		return nil, nil, nil, fmt.Errorf("查询创建的退货记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}

	return []string{task.TaskType.RefundOverTimeTask}, []string{task.TaskType.ReturnTask}, rows, nil
}

// BuildRefundResult 构建退款返回结果
func (r *DBRefund) BuildRefundResult(input types.IXMap) (types.XMap, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(input.GetInt("status"), errorcode.RequestFlowType.Refund, input.GetString("fail_code"), input.GetString("fail_msg"))
	res := map[string]interface{}{
		"channel_no":  input.GetString("channel_no"),
		"request_no":  input.GetString("request_no"),
		"refund_no":   input.GetString("refund_no"),
		"refund_id":   input.GetInt64("refund_id"),
		"refund_num":  input.GetInt("refund_num"),
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	}

	if input.GetString("account_no") != "" {
		res["account_no"] = input.GetString("account_no")
	}
	if input.GetInt("point_num") > 0 && input.GetInt("status") == enum.RefundResultStatus.Success {
		pointList := types.NewXMaps()
		if input.GetInt("activity_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": input.GetInt("activity_send_num"),
				"point_type":  enum.PointType.Activity,
			})
		}
		if input.GetInt("buy_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": input.GetInt("buy_send_num"),
				"point_type":  enum.PointType.Buy,
			})
		}
		if pointList.IsEmpty() {
			return nil, fmt.Errorf("积分数据异常")
		}
		bt, err := json.Marshal(pointList)
		if err != nil {
			return nil, err
		}

		res["ext_params"] = string(bt)
	}

	return res, nil
}
