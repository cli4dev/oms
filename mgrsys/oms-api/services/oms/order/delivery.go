package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/order"
)

//DeliveryHandler 订单发货表接口
type DeliveryHandler struct {
	container component.IContainer
	DeliveryLib   order.IDbDelivery
}

//NewDeliveryHandler 创建订单发货表对象
func NewDeliveryHandler(container component.IContainer) (u *DeliveryHandler) {
	return &DeliveryHandler{
		container: container,
		DeliveryLib:   order.NewDbDelivery(container),
	}
}
//GetHandle 获取订单发货表单条数据
func (u *DeliveryHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单发货表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("delivery_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.DeliveryLib.Get(ctx.Request.GetString("delivery_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取订单发货表数据列表
func (u *DeliveryHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单发货表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryDelivery
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.DeliveryLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
