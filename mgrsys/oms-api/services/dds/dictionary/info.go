package dictionary

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/dds/dictionary"
)

//InfoHandler 字典表接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   dictionary.IDbInfo
}

//NewInfoHandler 创建字典表对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   dictionary.NewDbInfo(container),
	}
}
//PostHandle 添加字典表数据
func (u *InfoHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加字典表数据--------")
	ctx.Log.Info("1.参数校验")

	var input dictionary.CreateInfo
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
//GetHandle 获取字典表单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取字典表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取字典表数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取字典表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input dictionary.QueryInfo
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
//PutHandle 更新字典表数据
func (u *InfoHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新字典表数据--------")
	ctx.Log.Info("1.参数校验")

	var input dictionary.UpdateInfo
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
