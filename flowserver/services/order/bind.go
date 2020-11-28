package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"

	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"

	"github.com/micro-plat/hydra/component"
)

// BindHandler 绑定结构体
type BindHandler struct {
	c    component.IContainer
	bLib order.IBind
}

// NewBindHandler 构建BindHandler
func NewBindHandler(c component.IContainer) *BindHandler {
	return &BindHandler{
		c:    c,
		bLib: order.NewBindOms(c),
	}
}

// Handle 最低价绑定
func (b *BindHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------最低价绑定--------------------")
	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "order_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.最低价绑定")
	orderID := ctx.Request.GetInt64("order_id")
	err = b.bLib.BindMethod(orderID, ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "最低价绑定", types.GetString(orderID)).Start("最低价绑定").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
