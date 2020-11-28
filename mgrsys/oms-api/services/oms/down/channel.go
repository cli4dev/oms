package down

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/down"
)

//ChannelHandler 下游渠道接口
type ChannelHandler struct {
	container  component.IContainer
	ChannelLib down.IDbChannel
}

//NewChannelHandler 创建下游渠道对象
func NewChannelHandler(container component.IContainer) (u *ChannelHandler) {
	return &ChannelHandler{
		container:  container,
		ChannelLib: down.NewDbChannel(container),
	}
}

//GetDictionary 获取数据字典
func (u *ChannelHandler) GetDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游渠道数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.ChannelLib.GetChannelDictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

//PostHandle 添加下游渠道数据
func (u *ChannelHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加下游渠道数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.CreateChannel
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ChannelLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//GetHandle 获取下游渠道单条数据
func (u *ChannelHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游渠道单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("channel_no"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.ChannelLib.Get(ctx.Request.GetString("channel_no"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取下游渠道数据列表
func (u *ChannelHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游渠道数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input down.QueryChannel
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.ChannelLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//PutHandle 更新下游渠道数据
func (u *ChannelHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新下游渠道数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.UpdateChannel
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ChannelLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

func (u *ChannelHandler) SetSecretHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------设置秘钥--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("channel_no"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ChannelLib.SetSecret(ctx.Request.GetString("channel_no"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

func (u *ChannelHandler) GetSecretHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取秘钥--------")
	ctx.Log.Info("1.参数校验")
	if err := ctx.Request.Check("channel_no"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	secret, err := u.ChannelLib.GetSecret(ctx.Request.GetString("channel_no"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return secret
}
