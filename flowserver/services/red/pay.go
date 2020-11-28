package red

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/invoice/pay"
)

// RefundHandler 完成上游开票申请
type RefundHandler struct {
	c    component.IContainer
	fLib pay.IInvoicePay
}

//NewRefundHandler 构建PayHandler
func NewRefundHandler(c component.IContainer) *RefundHandler {
	return &RefundHandler{
		c:    c,
		fLib: pay.NewInvoiceRedPayOms(c),
	}
}

// Handle 开票支付
func (f *RefundHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------冲红退款---------")
	ctx.Log.Info("1.校验参数")

	if err := ctx.Request.Check("order_id", "invoice_id", "task_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.冲红退款")
	err := f.fLib.Pay(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("invoice_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "冲红退款", ctx.Request.GetString("order_id"), ctx.Request.GetString("invoice_id")).Start("冲红退款").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
