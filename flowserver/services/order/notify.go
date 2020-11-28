package order

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// NotifyHandler 通知结构体
type NotifyHandler struct {
	container component.IContainer
	nLib      order.INotify
}

// NewNotifyHandler 构建NotifyHandler
func NewNotifyHandler(container component.IContainer) *NotifyHandler {
	return &NotifyHandler{
		container: container,
		nLib:      order.NewNotifyOms(container),
	}
}

// Handle 通知
func (n *NotifyHandler) Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("---------------下游通知--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("notify_id", "task_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.下游通知")
	notifyID := ctx.Request.GetInt64("notify_id")
	err = n.nLib.NotifyRequest(notifyID, ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "下游通知", types.GetString(ctx.Request.GetInt64("order_id")), types.GetString(notifyID)).Start("下游通知").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
