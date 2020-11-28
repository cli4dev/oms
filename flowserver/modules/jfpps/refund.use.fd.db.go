package jfpps

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//IDBRefund 退款数据层
type IDBRefund interface {
}

//DBRefund 退款
type DBRefund struct {
	c component.IContainer
}

//NewDBRefund NewDBRefund
func NewDBRefund(c component.IContainer) *DBRefund {
	return &DBRefund{c: c}
}

//RefundFDCheck 积分退款记账检查
func (o *DBRefund) RefundFDCheck(orderID int64, refundID int64) (types.XMap, types.XMap, error) {
	db := o.c.GetRegularDB()
	//1.检查订单的退款积分记账是否存在
	rows, _, _, err := db.Query(sql.CheckRefundFDRecord, map[string]interface{}{
		"order_id":  orderID,
		"refund_id": refundID,
	})
	if err != nil || !rows.IsEmpty() {
		return rows.Get(0), nil, err
	}

	//2.检查订单
	isUserRefund := types.DecodeInt(refundID, 0, 1, 0)
	orders, _, _, err := db.Query(sql.QueryRefundFDOrder, map[string]interface{}{
		"order_id":  orderID,
		"is_refund": isUserRefund,
	})
	if err != nil || orders.IsEmpty() {
		return nil, nil, fmt.Errorf("检查订单异常,count:%d,err:%v", orders.Len(), err)
	}
	refundPointNum := orders.Get(0).GetInt("point_num")

	//3.检查退款
	if isUserRefund == 0 {
		refunds, _, _, err := db.Query(sql.QueryRefundFDRefundRecord, map[string]interface{}{
			"order_id":  orderID,
			"refund_id": refundID,
		})
		if err != nil || refunds.IsEmpty() {
			return nil, nil, fmt.Errorf("检查退款记录异常,count:%d,err:%v", orders.Len(), err)
		}
		refundPointNum = refunds.Get(0).GetInt("refund_point_num")
	}

	//4.检查订单记账记录
	records, _, _, err := db.Query(sql.QueryJFOrderFDRecord, map[string]interface{}{"order_id": orderID})
	if err != nil || records.IsEmpty() || records.Get(0).GetInt("cn") == 0 {
		return nil, nil, fmt.Errorf("检查积分订单记账异常,count:%d,cn:%d,err:%v", records.Len(), records.Get(0).GetInt("cn"), err)
	}

	pointInfo := records.Get(0)
	refundBuy, refundAct, err := o.refundPointCalc(pointInfo.GetInt("last_buy_num"), pointInfo.GetInt("last_activity_num"), refundPointNum)
	if err != nil {
		return nil, nil, err
	}

	return nil, types.NewXMapByMap(map[string]interface{}{
		"refund_point_num":    refundPointNum,
		"refund_buy_num":      refundBuy,
		"refund_activity_num": refundAct,
	}), nil
}

//GetJFUpChannel 获取积分上游渠道信息
func (o *DBRefund) GetJFUpChannel() (types.XMap, error) {
	db := o.c.GetRegularDB()
	//查询积分上游
	JFChannels, _, _, err := db.Query(sql.GetJFUpChannel, map[string]interface{}{})
	if err != nil || JFChannels.Len() != 1 {
		return nil, fmt.Errorf("查询积分上游异常,count:%d,err:%v", JFChannels.Len(), err)
	}

	return JFChannels.Get(0), nil
}

//CreateRefundFDRecord 创建退款记账记录
func (o *DBRefund) CreateRefundFDRecord(dbTrans db.IDBTrans, orderID, refundID int64, fdChannelNo string, refPointInfo types.XMap) (types.XMap, []string, error) {

	//1.锁订单
	orders, _, _, err := dbTrans.Query(sql.LockRefundFDOrder, map[string]interface{}{"order_id": orderID})
	if err != nil || orders.IsEmpty() {
		return nil, nil, fmt.Errorf("锁定单异常,count:%d,err:%v", orders.Len(), err)
	}

	//2.检查退款记账记录是否存在
	rows, _, _, err := dbTrans.Query(sql.CheckRefundFDRecord, map[string]interface{}{
		"order_id":  orderID,
		"refund_id": refundID,
	})
	if err != nil || !rows.IsEmpty() {
		return nil, nil, err
	}

	//3.创建退款记账记录
	_, count, _, _, err := utils.Insert(dbTrans, sql.GetNewRefundFDID, sql.CreateRefundFDRecord, map[string]interface{}{
		"fd_channel_no":       fdChannelNo,
		"request_type":        types.DecodeInt(refundID, 0, enum.JFFDType.OrderFail, enum.JFFDType.OrderRefund),
		"order_id":            orderID,
		"refund_id":           refundID,
		"refund_point_num":    refPointInfo.GetInt("refund_point_num"),
		"refund_buy_num":      refPointInfo.GetInt("refund_buy_num"),
		"refund_activity_num": refPointInfo.GetInt("refund_activity_num"),
	})
	if err != nil || count != 1 {
		return nil, nil, fmt.Errorf("创建退款积分记账记录异常,count:%d,err:%v", count, err)
	}

	if refundID > 0 {
		//4.查询退款通知信息
		notifys, _, _, err := dbTrans.Query(sql.RefundNotifyInfo, map[string]interface{}{
			"refund_id": refundID,
		})
		if err != nil || notifys.IsEmpty() {
			return nil, nil, fmt.Errorf("查询退款通知异常,count:%d,err:%v", notifys.Len(), err)
		}

		//5.通知修改为等待
		notifyCount, q, a, err := dbTrans.Execute(sql.SqlUpdateRefundNotifyToWait, map[string]interface{}{
			"refund_id": refundID,
		})
		if err != nil || notifyCount <= 0 {
			return nil, nil, fmt.Errorf("开启通知错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notifyCount)
		}

		return notifys.Get(0), []string{task.TaskType.RefundNotifyTask}, nil
	}
	return nil, nil, nil
}

func (o *DBRefund) refundPointCalc(lastBuyNum int, lastActivityNum int, refundPointNum int) (refundBuyNum int, refundActivityNum int, err error) {
	if lastBuyNum >= refundPointNum {
		return refundPointNum, 0, nil
	}
	refundBuyNum = lastBuyNum
	refundActivityNum = (refundPointNum - lastBuyNum)
	if lastActivityNum < refundActivityNum {
		return 0, 0, fmt.Errorf("退款积分操过发货积分")
	}

	return refundBuyNum, refundActivityNum, nil
}
