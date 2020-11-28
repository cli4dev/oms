package refund

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/overtime"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// OverTimeHandler 退款超时处理接入层对象
type OverTimeHandler struct {
	container component.IContainer
	t         overtime.IOverTime
}

// NewOverTimeHandler 构建对象
func NewOverTimeHandler(container component.IContainer) (u *OverTimeHandler) {
	return &OverTimeHandler{
		container: container,
		t:         overtime.NewOmsOverTime(container),
	}
}

//Handle 退款超时处理
func (u *OverTimeHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退款超时处理--------------------")

	ctx.Log.Info("1. 检查参数")
	err := ctx.Request.Check("task_id", "refund_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行超时处理")
	defer lcs.New(ctx, "退款超时处理", ctx.Request.GetString("refund_id")).Start("开始退款超时处理").End(err)
	err = u.t.TimeoutDeal(ctx.Request.GetString("refund_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}

//UnknownHandle 退货未知处理
func (u *OverTimeHandler) UnknownHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退货未知处理--------------------")

	ctx.Log.Info("1. 检查参数")
	err := ctx.Request.Check("task_id", "return_id", "refund_id", "refund_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 开始处理")
	defer lcs.New(ctx, "退货未知处理",
		ctx.Request.GetString("refund_id"),
		ctx.Request.GetString("return_id")).Start("开始退货未知处理").End(err)
	err = u.t.ReturnUnknownDeal(ctx.Request.GetInt64("refund_id"),
		ctx.Request.GetInt64("return_id"),
		ctx.Request.GetInt64("order_id"),
		ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")
	return "success"
}
