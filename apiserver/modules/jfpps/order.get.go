package jfpps

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// JFOrder OrderOms结构体
type JFOrder struct {
	*order.Order
	c  component.IContainer
	db *DBJFOrder
}

// NewJFOrder 积分订单处理
func NewJFOrder(c component.IContainer) *JFOrder {
	oo := &JFOrder{
		c:  c,
		db: NewDBJFOrder(c, nil),
	}

	return oo
}

//Request 重写下单请求
func (o *JFOrder) Request(req *JFRequestInfo) (types.IXMap, error) {
	orderReq, ps, err := o.BuildJFParams(req)
	if err != nil {
		return nil, err
	}
	o.Order = order.NewOrder(o.c, NewDBJFOrder(o.c, ps), task.NewQTask(o.c))

	return o.Order.Request(orderReq)
}

//BuildJFParams 构建积分请求参数
func (o *JFOrder) BuildJFParams(req *JFRequestInfo) (*order.RequestInfo, *JFExtParams, error) {

	res := &order.RequestInfo{
		ChannelNO: req.ChannelNO,
		RequestNO: req.RequestNO,
		LineID:    req.LineID,
		Num:       req.Num,
		Face:      0.01,
		Amount:    0.01 * types.GetFloat64(req.Num),
		NotifyURL: req.NotifyURL,
	}
	ps := &JFExtParams{
		PreOrderID: 0,
		OverTime:   req.OverTime,
		UserNo:     req.UserNo,
		PointType:  req.PointType,
		Memo:       req.Memo,
	}
	if req.PreRequestNO != "" {
		preInfo, err := o.db.QueryJFPreOrderInfo(req)
		if err != nil {
			return nil, nil, err
		}

		res = &order.RequestInfo{
			ChannelNO: req.ChannelNO,
			RequestNO: req.RequestNO,
			LineID:    preInfo.GetInt("line_id"),
			Num:       preInfo.GetInt("num"),
			Face:      0.01,
			Amount:    0.01 * preInfo.GetFloat64("num"),
			NotifyURL: req.NotifyURL,
		}

		ps = &JFExtParams{
			PreOrderID: preInfo.GetInt64("pre_order_id"),
			OverTime:   preInfo.GetString("over_time"),
			UserNo:     preInfo.GetString("user_no"),
			PointType:  preInfo.GetInt("point_type"),
			Memo:       req.Memo,
		}
	}

	if res.LineID == 0 || res.Num == 0 || ps.UserNo == "" || ps.PointType == 0 {
		return nil, nil, fmt.Errorf("缺少必传参数")
	}

	return res, ps, nil
}
