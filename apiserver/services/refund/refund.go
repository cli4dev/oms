package refund

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/lcs/lcs"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/refund"
)

// RefundHandle 普通退款接入层对象
type RefundHandle struct {
	container component.IContainer
	r         refund.IRefund
}

// NewRefundHandle 构建对象
func NewRefundHandle(container component.IContainer) (u *RefundHandle) {
	return &RefundHandle{
		container: container,
		r:         refund.NewRefundOms(container),
	}
}

// GeneralHandle 普通退款接入处理
func (u *RefundHandle) GeneralHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------普通退款--------------------")

	ctx.Log.Info("1. 参数检查")
	input := &refund.RequestBody{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 创建退款")
	data, err := u.r.GeneralRequest(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "普通退款", data.GetString("refund_id")).Create("普通退款")

	ctx.Log.Info("3. 流程结束")

	return data
}

//QueryHandle 接入处理
func (u *RefundHandle) QueryHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退款查询--------------------")

	ctx.Log.Info("1. 检验参数")
	input := &refund.QueryRequestBody{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2. 获取退款信息")
	data, err := u.r.QueryRefundInfo(input)
	if err != nil {
		return err
	}
	defer lcs.New(ctx, "退款查询", types.GetString("refund_id")).Create("退款查询")
	ctx.Log.Info("3. 返回数据")

	return data
}
