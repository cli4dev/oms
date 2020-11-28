package invoice

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/invoice/finish"
)

// FinishHandler 完成上游开票申请
type FinishHandler struct {
	c    component.IContainer
	fLib finish.IInvoiceFinish
}

//NewFinishHandler 构建FinishHandler
func NewFinishHandler(c component.IContainer) *FinishHandler {
	return &FinishHandler{
		c:    c,
		fLib: finish.NewInvoiceFinishOms(c),
	}
}

// Handle 完成上游申请发票
func (f *FinishHandler) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------完成开票---------")
	ctx.Log.Info("1.校验参数")
	var input finish.InvoiceResult
	err := ctx.Request.Bind(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.完成开票,input:%+v", input)
	err = f.fLib.Finish(&input)
	defer lcs.New(ctx, "完成开票", types.GetString(ctx.Request.GetInt64("order_no"))).Start("完成开票").End(err)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
