package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/order"
)

// OrderHandler 订单结构体
type OrderHandler struct {
	c component.IContainer
	r order.IOrder
}

// NewOrderHandler 构建Handler
func NewOrderHandler(c component.IContainer) *OrderHandler {
	return &OrderHandler{
		c: c,
		r: order.NewOrderOms(c),
	}
}

// RequestHandle 下单请求
func (o OrderHandler) RequestHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------下单请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &order.RequestInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.下单")
	data, err := o.r.Request(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "下单", data.GetString("order_id")).Create("下单")

	ctx.Log.Info("3.返回数据")
	return data
}

// QueryHandle 订单查询
func (o OrderHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------订单查询----------------")

	ctx.Log.Info("1.参数检验")
	input := &order.QueryInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.订单查询")
	data, err := o.r.QueryOrder(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return data
}
