package refund

import (
	"fmt"
	"strings"

	"gitlab.100bm.cn/micro-plat/fas/fas/sdk"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBRefund 上游退款DB接口
type IDBRefund interface {
	LockByDB(dbTrans db.IDBTrans, refundID string) error
	CheckUpRefundStatusByDB(returnID string) (types.XMap, error)
	UpdateUpRefundToSuccessByDB(returnID, refundID string, dbTrans db.IDBTrans) error
	CheckDownRefundStatusByDB(refundID string) (types.XMap, error)
	UpdateDownRefundToSuccessByDB(refundID string, dbTrans db.IDBTrans) error
	CheckOrderStatusByDB(orderID int64) (types.XMap, error)
	UpdateOrderRefundStatusByDB(dbTrans db.IDBTrans, orderID string) error
	UpRefundByDB(trans db.IDBTrans, data types.IXMap, platName string) error
	DownRefundByDB(trans db.IDBTrans, data types.IXMap, platName, tradeNo string) error
}

// DBRefund 上游退款对象
type DBRefund struct {
	c component.IContainer
}

//NewDBRefund 创建上游退款信息对象
func NewDBRefund(c component.IContainer) *DBRefund {
	return &DBRefund{
		c: c,
	}
}

// LockByDB 锁退款记录
func (n *DBRefund) LockByDB(dbTrans db.IDBTrans, refundID string) error {
	rows, q, a, err := dbTrans.Query(sql.SqlLockCurrentRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return fmt.Errorf("锁退款记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	return nil
}

// CheckUpRefundStatusByDB 检查上游退款状态
func (n *DBRefund) CheckUpRefundStatusByDB(returnID string) (types.XMap, error) {
	dbt := n.c.GetRegularDB()
	ups, q, a, err := dbt.Query(sql.SqlCheckUpRefundStatus, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil {
		return nil, fmt.Errorf("检查上游退款状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, ups.Len())
	}
	if ups.IsEmpty() {
		return nil, context.NewErrorf(errorcode.UP_REFUND_STATUS_ERROR.Code, "上游退款状态异常")
	}
	return ups.Get(0), nil
}

// UpdateUpRefundToSuccessByDB 更新上游退款为信息
func (n *DBRefund) UpdateUpRefundToSuccessByDB(returnID, refundID string, dbTrans db.IDBTrans) (err error) {
	// 将上游退款状态改为退款成功
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateUpRefundToSuccess, map[string]interface{}{
		"return_id": returnID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("上游退款状态改为退款成功错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	// 获取所有上游退款未成功记录
	failUpRefundList, q, a, err := dbTrans.Query(sql.SqlCheckUpRefundAllFinish, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return fmt.Errorf("获取所有上游退款未成功记录错误(err:%+v),sql:%s,input:%+v,failUpRefundList:%+v", err, q, a, failUpRefundList)
	}

	// 获取下游退款成功记录
	sucDownRefundList, q, a, err := dbTrans.Query(sql.SqlCheckDownRefundFinish, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return fmt.Errorf("获取下游退款成功记录错误(err:%+v),sql:%s,input:%+v,sucDownRefundList:%+v", err, q, a, sucDownRefundList)
	}

	// 判断退款环节是否结束
	if sucDownRefundList.IsEmpty() || !failUpRefundList.IsEmpty() {
		// 退款环节未结束
		return nil
	}

	// 退款环节结束，关闭退款
	count, q, a, err = dbTrans.Execute(sql.SqlCloseRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("关闭退款错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return nil
}

// CheckDownRefundStatusByDB 检查支付退款状态
func (n *DBRefund) CheckDownRefundStatusByDB(refundID string) (types.XMap, error) {
	dbt := n.c.GetRegularDB()
	// 检查支付退款状态
	rows, q, a, err := dbt.Query(sql.SqlCheckDownRefundStatus, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return nil, fmt.Errorf("获取支付退款信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	if rows.IsEmpty() {
		return nil, context.NewErrorf(errorcode.DOWN_REFUND_STATUS_ERROR.Code, "下游退款状态异常")
	}
	return rows.Get(0), nil
}

// UpdateDownRefundToSuccessByDB 更新成功下游退款信息
func (n *DBRefund) UpdateDownRefundToSuccessByDB(refundID string, dbTrans db.IDBTrans) error {
	// 获取所有上游退款未成功记录数
	len, q, a, err := dbTrans.Scalar(sql.SqlGetAllFailUpRefund, map[string]interface{}{
		"refund_id": refundID,
	})
	if err != nil {
		return fmt.Errorf("获取所有上游退款未成功记录数错误(err:%+v),sql:%s,input:%+v,len:%d", err, q, a, len)
	}

	// 更改退款记录下游退款状态
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateDownRefundToSuccess, map[string]interface{}{
		"refund_id": refundID,
		"code":      types.GetInt(len),
	})
	if err != nil {
		return fmt.Errorf("更改退款记录下游退款状态错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return nil
}

// CheckOrderStatusByDB 检查订单状态
func (n *DBRefund) CheckOrderStatusByDB(orderID int64) (types.XMap, error) {
	dbt := n.c.GetRegularDB()
	// 检查状态
	rows, q, a, err := dbt.Query(sql.SqlCheckOrderStatus, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, fmt.Errorf("检查订单状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, rows.Len())
	}
	if rows.IsEmpty() {
		return nil, context.NewErrorf(errorcode.ORDER_REFUND_STATUS_ERROR.Code, "订单退款状态异常")
	}
	return rows.Get(0), nil
}

// UpdateOrderRefundStatusByDB 更改订单失败退款状态
func (n *DBRefund) UpdateOrderRefundStatusByDB(dbTrans db.IDBTrans, orderID string) error {
	// 更改订单退款状态为成功
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateOrderStatus, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("更改订单退款状态为成功错误(err:%+v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}

	return nil
}

// DownRefundByDB 下游退款
func (n *DBRefund) DownRefundByDB(trans db.IDBTrans, data types.IXMap, platName, tradeNo string) error {
	config, err := n.c.GetVarConf("fd", "fd")
	if err != nil {
		return err
	}

	downChannel := data.GetString("down_channel_no")
	if data.GetString("down_account_no") != "" {
		downChannel = fmt.Sprintf("%s_%s", data.GetString("down_channel_no"), data.GetString("down_account_no"))
	}

	input := &sdk.DownRefund{
		Ident:                      n.c.GetPlatName(),
		OrderSource:                config.GetInt("order_source"),
		DownChannelNo:              downChannel,
		TradeOrderNo:               data.GetString("order_id"),
		TradeRefundNo:              tradeNo,
		RefundUnit:                 data.GetInt("refund_unit"),
		RefundFace:                 data.GetInt("refund_unit"),
		DownRefundAmount:           data.GetFloat64("sell_amount"),
		DownRefundCommissionAmount: data.GetFloat64("commission_amount"),
		DownRefundServiceAmount:    data.GetFloat64("service_amount"),
		DownRefundFeeAmount:        data.GetFloat64("fee_amount"),
		OrderDate:                  data.GetString("create_time"),
		Memo:                       "下游退款",
	}
	result, err := sdk.DownChannelRefund(n.c, input)
	if err != nil {
		return err
	}

	if !strings.Contains(result.Result, "success") {
		return fmt.Errorf("result:%v,msg:%v", result.Result, result.Msg)
	}
	return nil
}

// UpRefundByDB 上游退款
func (n *DBRefund) UpRefundByDB(trans db.IDBTrans, data types.IXMap, platName string) error {
	config, err := n.c.GetVarConf("fd", "fd")
	if err != nil {
		return err
	}

	downChannel := data.GetString("down_channel_no")
	if data.GetString("down_account_no") != "" {
		downChannel = fmt.Sprintf("%s_%s", data.GetString("down_channel_no"), data.GetString("down_account_no"))
	}

	input := &sdk.UpRefund{
		Ident:                    n.c.GetPlatName(),
		OrderSource:              config.GetInt("order_source"),
		DownChannelNo:            downChannel,
		UpChannelNo:              data.GetString("up_channel_no"),
		TradeOrderNo:             data.GetString("order_id"),
		TradeDeliveryNo:          data.GetString("delivery_id"),
		TradeRefundNo:            data.GetString("refund_id"),
		BillType:                 1,
		BusinessType:             data.GetInt("line_id"),
		CarrierNo:                data.GetString("carrier_no"),
		ProvinceNo:               data.GetString("province_no"),
		UpRefundUnit:             data.GetInt("up_refund_face"),
		UpRefundFace:             data.GetInt("up_refund_face"),
		UpRefundAmount:           data.GetFloat64("return_cost_amount", 0),
		UpRefundCommissionAmount: data.GetFloat64("return_commission_amount", 0),
		UpRefundServiceAmount:    data.GetFloat64("return_service_amount", 0),
		UpRefundFeeAmount:        data.GetFloat64("return_fee_amount", 0),
		OrderTime:                data.GetString("create_time"),
		Memo:                     "上游退款",
	}

	result, err := sdk.UpChannelRefund(n.c, input)
	if err != nil {
		return err
	}

	if !strings.Contains(result.Result, "success") {
		return fmt.Errorf("result:%v,msg:%v", result.Result, result.Msg)
	}

	return nil
}
