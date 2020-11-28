package channel

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/vds/channel"
)

//ErrorCodeHandler 渠道错误码接口
type ErrorCodeHandler struct {
	container component.IContainer
	ErrorCodeLib   channel.IDbErrorCode
}

//NewErrorCodeHandler 创建渠道错误码对象
func NewErrorCodeHandler(container component.IContainer) (u *ErrorCodeHandler) {
	return &ErrorCodeHandler{
		container: container,
		ErrorCodeLib:   channel.NewDbErrorCode(container),
	}
}
//PostHandle 添加渠道错误码数据
func (u *ErrorCodeHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加渠道错误码数据--------")
	ctx.Log.Info("1.参数校验")

	var input channel.CreateErrorCode
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ErrorCodeLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
//GetHandle 获取渠道错误码单条数据
func (u *ErrorCodeHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取渠道错误码单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.ErrorCodeLib.Get(ctx.Request.GetString("id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取渠道错误码数据列表
func (u *ErrorCodeHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取渠道错误码数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input channel.QueryErrorCode
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.ErrorCodeLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
//PutHandle 更新渠道错误码数据
func (u *ErrorCodeHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新渠道错误码数据--------")
	ctx.Log.Info("1.参数校验")

	var input channel.UpdateErrorCode
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ErrorCodeLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
