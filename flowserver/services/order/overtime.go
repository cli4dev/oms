package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// OvertimeHandler 超时结构体
type OvertimeHandler struct {
	container component.IContainer
	oLib      order.IOvertime
}

// NewOvertimeHandler 构建OvertimeHandler
func NewOvertimeHandler(container component.IContainer) *OvertimeHandler {
	return &OvertimeHandler{
		container: container,
		oLib:      order.NewOvertimeOms(container),
	}
}

// Handle 订单超时处理
func (o *OvertimeHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------订单超时处理--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("order_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.订单超时处理")
	orderID := ctx.Request.GetInt64("order_id")
	err = o.oLib.OrderOvertime(ctx.Request.GetInt64("task_id"), orderID)
	defer lcs.New(ctx, "订单超时处理", types.GetString(orderID)).Start("订单超时处理").End(err)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果")
	return "success"
}

// DeliveryHandle 发货超时处理
func (o *OvertimeHandler) DeliveryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------发货超时处理--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("delivery_id", "order_id", "task_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.发货超时处理")
	deliveryID := ctx.Request.GetInt64("delivery_id")
	orderID := ctx.Request.GetInt64("order_id")
	err = o.oLib.DeliveryOvertime(deliveryID, orderID,
		ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "发货超时处理", types.GetString(orderID), types.GetString(deliveryID)).Start("发货超时处理").End(err)
	if err != nil {
		return err
	}
	ctx.Log.Info("3.返回结果")
	return "success"
}
