package invoice

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/invoice"
)

// InvoiceHandler 订单结构体
type InvoiceHandler struct {
	c component.IContainer
	r invoice.IInvoice
}

// NewInvoiceHandler 构建Handler
func NewInvoiceHandler(c component.IContainer) *InvoiceHandler {
	return &InvoiceHandler{
		c: c,
		r: invoice.NewInvoiceOms(c),
	}
}

// RequestHandle 下单请求
func (o InvoiceHandler) RequestHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------开票请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &invoice.RequestInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.开票请求")
	data, err := o.r.Request(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "开票请求", data.GetString("invoice_id")).Create("开票请求")

	ctx.Log.Info("3.返回数据")
	return data
}

// QueryHandle 开票查询
func (o InvoiceHandler) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------开票查询----------------")

	ctx.Log.Info("1.参数检验")
	input := &invoice.QueryInfo{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.开票查询")
	data, err := o.r.Query(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return data
}
