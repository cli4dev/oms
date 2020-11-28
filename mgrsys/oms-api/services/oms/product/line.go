package product

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/product"
)

//LineHandler 产品线接口
type LineHandler struct {
	container component.IContainer
	LineLib   product.IDbLine
}

//NewLineHandler 创建产品线对象
func NewLineHandler(container component.IContainer) (u *LineHandler) {
	return &LineHandler{
		container: container,
		LineLib:   product.NewDbLine(container),
	}
}
//GetDictionary 获取数据字典 
func (u *LineHandler) GetDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取产品线数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.LineLib.GetLineDictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}
//PostHandle 添加产品线数据
func (u *LineHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加产品线数据--------")
	ctx.Log.Info("1.参数校验")

	var input product.CreateLine
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.LineLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
//GetHandle 获取产品线单条数据
func (u *LineHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取产品线单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("line_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.LineLib.Get(ctx.Request.GetString("line_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取产品线数据列表
func (u *LineHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取产品线数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input product.QueryLine
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.LineLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
//PutHandle 更新产品线数据
func (u *LineHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新产品线数据--------")
	ctx.Log.Info("1.参数校验")

	var input product.UpdateLine
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.LineLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
