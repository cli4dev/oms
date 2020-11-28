package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/jfpps"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
)

// JFGetReturnHandler 订单退货接入层对象
type JFGetReturnHandler struct {
	container component.IContainer
}

// NewJFGetReturnHandler 构建对象
func NewJFGetReturnHandler(container component.IContainer) (u *JFGetReturnHandler) {
	return &JFGetReturnHandler{
		container: container,
	}
}

//FinishHandle 接入处理
func (u *JFGetReturnHandler) FinishHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------积分获取退货完成--------------------")

	ctx.Log.Info("1.参数校验")
	var input order.DeliveryResult
	err := ctx.Request.Bind(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.完成退货处理,input:%+v", input)
	defer lcs.New(ctx, "完成退货处理", input.DeliveryID).Start("开始完成退货处理").End(err)
	jf := jfpps.NewJFReturn(u.container)
	err = jf.Finish(&input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 完成订单退货流程结束")

	return "success"
}
