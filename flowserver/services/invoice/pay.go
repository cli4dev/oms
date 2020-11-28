package invoice

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/invoice/pay"
)

// PayHandler 完成上游开票申请
type PayHandler struct {
	c    component.IContainer
	fLib pay.IInvoicePay
}

//NewPayHandler 构建PayHandler
func NewPayHandler(c component.IContainer) *PayHandler {
	return &PayHandler{
		c:    c,
		fLib: pay.NewInvoicePayOms(c),
	}
}

// Handle 开票支付
func (f *PayHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------开票支付---------")
	ctx.Log.Info("1.校验参数")

	if err := ctx.Request.Check("order_id", "invoice_id", "task_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.开票支付")
	err := f.fLib.Pay(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("invoice_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "开票支付", ctx.Request.GetString("order_id"), ctx.Request.GetString("invoice_id")).Start("开票支付").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
