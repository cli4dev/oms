package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/jfpps"
)

// PayHandler 支付结构体
type PayHandler struct {
	container component.IContainer
}

// NewPayHandler 构建PayHandler
func NewPayHandler(container component.IContainer) *PayHandler {
	return &PayHandler{
		container: container,
	}
}

// Handle 上游支付
func (u *PayHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------积分上游支付--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "delivery_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.积分上游支付")
	jf := jfpps.NewJFPay(u.container)
	err = jf.UpPay(ctx.Request.GetInt64("delivery_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "积分上游支付", types.GetString(ctx.Request.GetInt64("order_id")), types.GetString(ctx.Request.GetInt64("delivery_id"))).Start("上游支付").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
