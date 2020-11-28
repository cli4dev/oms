package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// DeliveryHandler 发货结构体
type DeliveryHandler struct {
	c    component.IContainer
	dLib order.IDelivery
}

// NewDeliveryHandler 构建DeliveryHandler
func NewDeliveryHandler(c component.IContainer) *DeliveryHandler {
	return &DeliveryHandler{
		c:    c,
		dLib: order.NewDeliveryOms(c),
	}
}

// StartHandle 发货流程
func (u *DeliveryHandler) StartHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------开始发货--------------------")
	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "delivery_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.开始发货")
	deliveryID := ctx.Request.GetInt64("delivery_id")
	err = u.dLib.Start(deliveryID, ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "开始发货", types.GetString(ctx.Request.GetInt64("order_id")), types.GetString(deliveryID)).Start("开始发货").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"

}

// FinishHandle 发货完成
func (u *DeliveryHandler) FinishHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------发货完成--------------------")
	ctx.Log.Info("1.参数校验")
	var input order.DeliveryResult
	err := ctx.Request.Bind(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Infof("2.发货完成,input:%+v", input)
	orderID, err := u.dLib.Finish(&input)
	defer lcs.New(ctx, "发货完成", orderID, input.DeliveryID).Start("发货完成").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"

}
