package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/vds/order"
)

//NotifyHandler 发货通知记录表接口
type NotifyHandler struct {
	container component.IContainer
	NotifyLib   order.IDbNotify
}

//NewNotifyHandler 创建发货通知记录表对象
func NewNotifyHandler(container component.IContainer) (u *NotifyHandler) {
	return &NotifyHandler{
		container: container,
		NotifyLib:   order.NewDbNotify(container),
	}
}
//GetHandle 获取发货通知记录表单条数据
func (u *NotifyHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货通知记录表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.NotifyLib.Get(ctx.Request.GetString("id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取发货通知记录表数据列表
func (u *NotifyHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货通知记录表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryNotify
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.NotifyLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
