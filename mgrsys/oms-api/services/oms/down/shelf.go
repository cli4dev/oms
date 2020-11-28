package down

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/down"
)

//ShelfHandler 下游货架接口
type ShelfHandler struct {
	container component.IContainer
	ShelfLib  down.IDbShelf
}

//NewShelfHandler 创建下游货架对象
func NewShelfHandler(container component.IContainer) (u *ShelfHandler) {
	return &ShelfHandler{
		container: container,
		ShelfLib:  down.NewDbShelf(container),
	}
}

//GetDictionary 获取数据字典
func (u *ShelfHandler) GetDictionaryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游货架数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.ShelfLib.GetShelfDictionary()
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

func (u *ShelfHandler) GetChannelHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游货架数据字典--------")
	ctx.Log.Info("1.执行操作")
	data, err := u.ShelfLib.GetShelf(ctx.Request.GetString("channel_no"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("2.返回结果")
	return data
}

//PostHandle 添加下游货架数据
func (u *ShelfHandler) PostHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------添加下游货架数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.CreateShelf
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ShelfLib.Create(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//GetHandle 获取下游货架单条数据
func (u *ShelfHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游货架单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("shelf_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.ShelfLib.Get(ctx.Request.GetString("shelf_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取下游货架数据列表
func (u *ShelfHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游货架数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input down.QueryShelf
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.ShelfLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//PutHandle 更新下游货架数据
func (u *ShelfHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------更新下游货架数据--------")
	ctx.Log.Info("1.参数校验")

	var input down.UpdateShelf
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.ShelfLib.Update(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
