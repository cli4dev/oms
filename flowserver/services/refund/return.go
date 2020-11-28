package refund

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
)

// ReturnHandler 订单退货接入层对象
type ReturnHandler struct {
	container component.IContainer
	o         refund.IReturn
}

// NewReturnHandler 构建对象
func NewReturnHandler(container component.IContainer) (u *ReturnHandler) {
	return &ReturnHandler{
		container: container,
		o:         refund.NewOmsReturn(container),
	}
}

//ReturnHandle 接入处理
func (u *ReturnHandler) ReturnHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------订单退货--------------------")

	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "return_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Infof("2. 开始退货")
	defer lcs.New(ctx, "退货处理", ctx.Request.GetString("return_id")).Start("开始退货").End(err)
	err = u.o.Start(ctx.Request.GetInt64("task_id"), ctx.Request.GetInt64("return_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 订单退货流程结束")

	return "success"
}

//FinishHandle 接入处理
func (u *ReturnHandler) FinishHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退货完成--------------------")

	ctx.Log.Info("1.参数校验")
	var input order.DeliveryResult
	err := ctx.Request.Bind(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.完成退货处理,input:%+v", input)
	defer lcs.New(ctx, "完成退货处理", input.DeliveryID).Start("开始完成退货处理").End(err)
	err = u.o.Finish(&input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 完成订单退货流程结束")

	return "success"
}
