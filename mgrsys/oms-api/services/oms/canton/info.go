package canton

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/canton"
)

//InfoHandler 省市信息接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   canton.IDbInfo
}

//NewInfoHandler 创建省市信息对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   canton.NewDbInfo(container),
	}
}
//GetDictionary 获取数据字典 
func (u *InfoHandler) GetDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取省市信息数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.InfoLib.GetInfoDictionary(ctx.Request.GetInt("grade"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}
//PostHandle 添加省市信息数据
func (u *InfoHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加省市信息数据--------")
	ctx.Log.Info("1.参数校验")

	var input canton.CreateInfo
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
//GetHandle 获取省市信息单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取省市信息单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("canton_code"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("canton_code"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取省市信息数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取省市信息数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input canton.QueryInfo
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
		"data":data,
		"count":count,
	}
}
//PutHandle 更新省市信息数据
func (u *InfoHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新省市信息数据--------")
	ctx.Log.Info("1.参数校验")

	var input canton.UpdateInfo
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

//GetDictionaryByProvinceHandle 根据省份获取数据字典
func (u *InfoHandler) GetDictionaryByProvinceHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------根据省份获取数据字典--------")

	ctx.Log.Info("1.执行操作")
	data, err := u.InfoLib.GetInfoDictionaryByProvince(ctx.Request.GetInt("grade"), ctx.Request.GetString("canton_code"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

//GetDictionaryByProvinceHandle 根据省份获取数据字典
func (u *InfoHandler) GetListHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------根据省份获取数据字典--------")

	ctx.Log.Info("1.执行操作")
	data, err := u.InfoLib.GetList()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

