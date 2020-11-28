package down

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/down"
)

//ProductHandler 下游商品接口
type ProductHandler struct {
	container component.IContainer
	ProductLib   down.IDbProduct
}

//NewProductHandler 创建下游商品对象
func NewProductHandler(container component.IContainer) (u *ProductHandler) {
	return &ProductHandler{
		container: container,
		ProductLib:   down.NewDbProduct(container),
	}
}
//PostHandle 添加下游商品数据
func (u *ProductHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加下游商品数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.CreateProduct
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ProductLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
//GetHandle 获取下游商品单条数据
func (u *ProductHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游商品单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("product_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.ProductLib.Get(ctx.Request.GetString("product_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取下游商品数据列表
func (u *ProductHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游商品数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input down.QueryProduct
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.ProductLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
//PutHandle 更新下游商品数据
func (u *ProductHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新下游商品数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.UpdateProduct
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ProductLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
