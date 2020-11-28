package refund

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// Handler 退款接入层对象
type Handler struct {
	container component.IContainer
	r         refund.IRefund
}

// NewRefundHandler 构建对象
func NewRefundHandler(container component.IContainer) (u *Handler) {
	return &Handler{
		container: container,
		r:         refund.NewOmsRefund(container),
	}
}

// UpHandle 上游退款接入处理
func (u *Handler) UpHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------上游退款--------------------")
	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "return_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行上游退款")
	defer lcs.New(ctx, "上游退款处理", ctx.Request.GetString("return_id")).Start("开始上游退款处理").End(err)
	err = u.r.UpRefund(ctx.Request.GetString("return_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}

//DownHandle 下游退款接入处理
func (u *Handler) DownHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------支付退款--------------------")
	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "refund_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行支付退款")
	defer lcs.New(ctx, "支付退款处理", ctx.Request.GetString("refund_id")).Start("开始支付退款处理").End(err)
	err = u.r.DownRefund(ctx.Request.GetString("refund_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}

// OrderFailHandle 订单失败退款
func (u *Handler) OrderFailHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------订单失败退款--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "order_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行订单失败退款")
	defer lcs.New(ctx, "订单失败退款处理", ctx.Request.GetString("order_id")).Start("开始订单失败退款处理").End(err)
	err = u.r.OrderRefund(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}
