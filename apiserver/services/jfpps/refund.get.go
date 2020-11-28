package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/jfpps"
)

// JFRefundHandler 退款结构体
type JFRefundHandler struct {
	c component.IContainer
}

// NewJFRefundHandler 构建Handler
func NewJFRefundHandler(c component.IContainer) *JFRefundHandler {
	return &JFRefundHandler{
		c: c,
	}
}

//Handle 下单请求
func (o *JFRefundHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------积分发放退款请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &jfpps.JFRefundRequestBody{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求下单")
	or := jfpps.NewJFRefund(o.c)
	data, err := or.RefundRequest(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "积分发放退款", data.GetString("refund_id")).Create("积分发放退款")

	ctx.Log.Info("3.返回数据")
	return data
}
