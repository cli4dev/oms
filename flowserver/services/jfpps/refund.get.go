package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/jfpps"
)

// JFGetRefundHandler 退款接入层对象
type JFGetRefundHandler struct {
	container component.IContainer
}

// NewJFGetRefundHandler 构建对象
func NewJFGetRefundHandler(container component.IContainer) (u *JFGetRefundHandler) {
	return &JFGetRefundHandler{
		container: container,
	}
}

// UpHandle 上游退款接入处理
func (u *JFGetRefundHandler) UpHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------积分获取上游退款--------------------")
	ctx.Log.Info("1.参数校验")
	err := ctx.Request.Check("task_id", "return_id")
	if err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 执行上游退款")
	defer lcs.New(ctx, "上游退款处理", ctx.Request.GetString("return_id")).Start("开始上游退款处理").End(err)
	jf := jfpps.NewJFRefund(u.container)
	err = jf.UpRefund(ctx.Request.GetString("return_id"), ctx.Request.GetInt64("task_id"))
	if err != nil {
		return err
	}

	ctx.Log.Info("3. 流程结束")

	return "success"
}
