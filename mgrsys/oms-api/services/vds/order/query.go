package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/vds/order"
)

//QueryHandler 发货结果查询记录表接口
type QueryHandler struct {
	container component.IContainer
	QueryLib   order.IDbQuery
}

//NewQueryHandler 创建发货结果查询记录表对象
func NewQueryHandler(container component.IContainer) (u *QueryHandler) {
	return &QueryHandler{
		container: container,
		QueryLib:   order.NewDbQuery(container),
	}
}
//GetHandle 获取发货结果查询记录表单条数据
func (u *QueryHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货结果查询记录表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("order_no"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.QueryLib.Get(ctx.Request.GetString("order_no"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取发货结果查询记录表数据列表
func (u *QueryHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货结果查询记录表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryQuery
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.QueryLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
