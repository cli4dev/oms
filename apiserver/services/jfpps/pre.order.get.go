package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/jfpps"
)

// PreOrderHandler 订单结构体
type PreOrderHandler struct {
	c  component.IContainer
	jf *jfpps.JFPreOrder
}

// NewPreOrderHandler 构建Handler
func NewPreOrderHandler(c component.IContainer) *PreOrderHandler {
	return &PreOrderHandler{
		c:  c,
		jf: jfpps.NewJFPreOrder(c),
	}
}

//RequestHandle 下单请求
func (o *PreOrderHandler) RequestHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------待激活积分下单请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &jfpps.JFPreRequestInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求下单")
	pre, err := o.jf.Request(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return pre
}

//CancelHandle 取消预下单
func (o *PreOrderHandler) CancelHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------待激活积分取消请求----------------")

	ctx.Log.Info("1.参数检验")
	if err := ctx.Request.Check("channel_no", "pre_request_no"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求下单")
	err := o.jf.Cancel(ctx.Request.GetString("channel_no"), ctx.Request.GetString("pre_request_no"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return "success"
}
