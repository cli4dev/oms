package red

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/invoice/notify"
)

//NotifyHandler 发票结果通知
type NotifyHandler struct {
	c    component.IContainer
	nLib notify.IInvoiceNotify
}

//NewNotifyHandler 构建NotifyHandler
func NewNotifyHandler(c component.IContainer) *NotifyHandler {
	return &NotifyHandler{
		c:    c,
		nLib: notify.NewInvoiceRedNoitfyOms(c),
	}
}

// Handle 发票结果通知
func (n *NotifyHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------发票冲红通知---------")
	ctx.Log.Info("1.校验参数")
	err := ctx.Request.Check("invoice_id", "order_id", "task_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.发票冲红通知")
	invoiceID := ctx.Request.GetInt64("invoice_id")
	err = n.nLib.Notify(invoiceID, ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "发票冲红通知", types.GetString(ctx.Request.GetString("order_id")), types.GetString(invoiceID)).Start("发票冲红通知").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
