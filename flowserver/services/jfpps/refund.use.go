package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/jfpps"
)

// RefundHandler 积分使用记账结构体
type RefundHandler struct {
	container component.IContainer
	jf        *jfpps.Refund
}

// NewRefundHandler 构建UseHandler
func NewRefundHandler(container component.IContainer) *RefundHandler {
	return &RefundHandler{
		container: container,
		jf:        jfpps.NewRefund(container),
	}
}

// Handle 积分使用记账
func (u *RefundHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------积分退款记账创建--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "order_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.积分退款记账")
	err = u.jf.RefundFDCreate(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("refund_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
