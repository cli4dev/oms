package red

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/invoice/start"
)

//StartHandler 请求上游申请开票结构体
type StartHandler struct {
	c    component.IContainer
	iLib start.IInvoice
}

//NewStartHandler 构建StartHandler
func NewStartHandler(c component.IContainer) *StartHandler {
	return &StartHandler{
		c:    c,
		iLib: start.NewInvoiceRedStartOms(c),
	}
}

//Handle 申请上游开票
func (i *StartHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------开始发票冲红---------")
	ctx.Log.Info("1.校验参数")
	err := ctx.Request.Check("order_id", "invoice_id", "task_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.开始发票冲红")
	err = i.iLib.Start(ctx.Request.GetInt64("order_id"), ctx.Request.GetInt64("invoice_id"), ctx.Request.GetInt64("task_id"))
	defer lcs.New(ctx, "开始发票冲红", types.GetString(ctx.Request.GetInt64("order_id")), types.GetString("apply_id")).Start("开始发票冲红").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
