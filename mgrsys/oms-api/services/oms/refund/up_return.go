package refund

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/refund"
)

//UpReturnHandler 上游退货信息表接口
type UpReturnHandler struct {
	container component.IContainer
	UpReturnLib   refund.IDbUpReturn
}

//NewUpReturnHandler 创建上游退货信息表对象
func NewUpReturnHandler(container component.IContainer) (u *UpReturnHandler) {
	return &UpReturnHandler{
		container: container,
		UpReturnLib:   refund.NewDbUpReturn(container),
	}
}
//GetHandle 获取上游退货信息表单条数据
func (u *UpReturnHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取上游退货信息表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("return_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.UpReturnLib.Get(ctx.Request.GetString("return_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取上游退货信息表数据列表
func (u *UpReturnHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取上游退货信息表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input refund.QueryUpReturn
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.UpReturnLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
