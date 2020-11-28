package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// PayHandler 支付结构体
type PayHandler struct {
	container component.IContainer
	pLib      order.IPay
}

// NewPayHandler 构建PayHandler
func NewPayHandler(container component.IContainer) *PayHandler {
	return &PayHandler{
		container: container,
		pLib:      order.NewPayOms(container),
	}
}

// DownHandle 下游支付
func (u *PayHandler) DownHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------下游支付--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "order_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.下游支付")
	err = u.pLib.DownPay(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "下游支付", types.GetString(ctx.Request.GetInt64("order_id"))).Start("下游支付").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

// UpHandle 上游支付
func (u *PayHandler) UpHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------上游支付--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "delivery_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.上游支付")
	err = u.pLib.UpPay(ctx.Request.GetInt64("delivery_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "上游支付", types.GetString(ctx.Request.GetInt64("order_id")), types.GetString(ctx.Request.GetInt64("delivery_id"))).Start("上游支付").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
