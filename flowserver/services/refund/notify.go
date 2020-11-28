package refund

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/notify"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// NotifyHandler 退款通知接入层对象
type NotifyHandler struct {
	container component.IContainer
	n         notify.INotify
}

// NewNotifyHandler 构建对象
func NewNotifyHandler(container component.IContainer) (u *NotifyHandler) {
	return &NotifyHandler{
		container: container,
		n:         notify.NewOmsNotify(container),
	}
}

//Handle 接入处理
func (u *NotifyHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退款通知--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "notify_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行退款通知")
	defer lcs.New(ctx, "退款通知处理", ctx.Request.GetString("notify_id")).Start("开始退款通知处理").End(err)
	err = u.n.RefundNotify(ctx.Request.GetString("notify_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}
