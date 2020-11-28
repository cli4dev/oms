package account

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/beanpay/account"
)

//InfoHandler 账户信息接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   account.IDbInfo
}

//NewInfoHandler 创建账户信息对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   account.NewDbInfo(container),
	}
}

//GetDownDictionaryHandle 获取下游账户数据字典
func (u *InfoHandler) GetDownDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游账户数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.InfoLib.GetDownInfoDictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

//GetUpDictionaryHandle 获取上游账户数据字典
func (u *InfoHandler) GetUpDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取上游账户数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.InfoLib.GetUpInfoDictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

//PostHandle 添加账户信息数据
func (u *InfoHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加账户信息数据--------")
	ctx.Log.Info("1.参数校验")

	var input account.CreateInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.InfoLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//GetHandle 获取账户信息单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取账户信息单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("account_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("account_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取账户信息数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取账户信息数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input account.QueryInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//PutHandle 更新账户信息数据
func (u *InfoHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新账户信息数据--------")
	ctx.Log.Info("1.参数校验")

	var input account.UpdateInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.InfoLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//AddHandle 账户信息数据
func (u *InfoHandler) AddHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------账户加款--------")
	ctx.Log.Info("1.参数校验")

	var input account.AddInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.InfoLib.Add(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//AddHandle 账户信息数据
func (u *InfoHandler) DrawHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------账户提款--------")
	ctx.Log.Info("1.参数校验")

	var input account.AddInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.InfoLib.Draw(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
