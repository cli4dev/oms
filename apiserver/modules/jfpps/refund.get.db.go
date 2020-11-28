package jfpps

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/refund"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//JFRefundExtParams 积分退款拓展参数
type JFRefundExtParams struct {
	JFAmount float64
}

// DBJFRefund 订单结构体
type DBJFRefund struct {
	*refund.DBRefund
	c  component.IContainer
	ps *JFRefundExtParams
}

// NewDBJFRefund 构建DBOrder结构体
func NewDBJFRefund(c component.IContainer, ps *JFRefundExtParams) *DBJFRefund {
	db := &DBJFRefund{
		c:  c,
		ps: ps,
	}
	db.DBRefund = refund.NewDBRefund(c)
	return db
}

//CheckRefundAndOrderGeneralByDb 检查退款与订单
func (r *DBJFRefund) CheckRefundAndOrderGeneralByDb(dbTrans db.IDBTrans, input *refund.RequestBody) (types.IXMap, types.IXMap, error) {
	refundInfo, order, err := r.DBRefund.CheckRefundAndOrderGeneralByDb(dbTrans, input)
	if order != nil && order.GetInt("num") < input.RefundNum+types.GetInt(r.ps.JFAmount*100) {
		return nil, nil, context.NewError(errorcode.REFUND_NUM_ERROR.Code, errorcode.REFUND_NUM_ERROR.Msg)
	}

	return refundInfo, order, err
}

// CreateRefundAndNotifyByDB 创建退款和通知
func (r *DBJFRefund) CreateRefundAndNotifyByDB(dbTrans db.IDBTrans, orderID string, extendInfo string, input *refund.RequestBody) (types.IXMap, error) {
	//查询订单拓展信息
	orderExtInfo, q, a, err := dbTrans.Query(sql.QueryOrderExtendInfo, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || orderExtInfo.IsEmpty() {
		return nil, fmt.Errorf("查询订单拓展信息错误(err:%v),sql:%s,input:%+v,list:%+v", err, q, a, orderExtInfo.Len())
	}
	orderExt, err := types.NewXMapByJSON(orderExtInfo.Get(0).GetString("extend_info"))
	if err != nil {
		return nil, err
	}

	//构建退款拓展信息
	ext := types.NewXMapByMap(map[string]interface{}{
		"jf_amount":   r.ps.JFAmount,
		"delivery_id": orderExtInfo.Get(0).GetInt64("delivery_id"),
	})
	ext.Merge(orderExt)

	if extendInfo != "" {
		extend, err := types.NewXMapByJSON(extendInfo)
		if err != nil {
			return nil, err
		}
		ext.Merge(extend)
	}

	extStr, err := json.Marshal(ext)
	if err != nil {
		return nil, err
	}

	// 创建退款记录
	refundID, count, sqlStr, args, err := utils.Insert(dbTrans, sql.SQLGetRefundID, sql.SQLCreateJFRefund, map[string]interface{}{
		"down_channel_no":   input.ChannelNo,
		"down_refund_no":    input.RefundNo,
		"refund_type":       enum.RefundType.General,
		"request_no":        input.RequestNo,
		"refund_num":        input.RefundNum,
		"notify_url":        input.NotifyURL,
		"order_id":          orderID,
		"refund_mer_amount": input.RefundMerAmount,
		"extend_info":       string(extStr),
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
