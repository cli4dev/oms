package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/jfpps"
)

// JFOrderHandler 订单结构体
type JFOrderHandler struct {
	c component.IContainer
}

// NewJFOrderHandler 构建Handler
func NewJFOrderHandler(c component.IContainer) *JFOrderHandler {
	return &JFOrderHandler{
		c: c,
	}
}

//Handle 下单请求
func (o *JFOrderHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------积分发放下单请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &jfpps.JFRequestInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求下单")
	od := jfpps.NewJFOrder(o.c)
	data, err := od.Request(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "下单", data.GetString("order_id")).Create("下单")

	ctx.Log.Info("3.返回数据")
	return data
}
