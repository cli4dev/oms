package notify

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/notify"
)

//InfoHandler 订单通知表接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   notify.IDbInfo
}

//NewInfoHandler 创建订单通知表对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   notify.NewDbInfo(container),
	}
}
//GetHandle 获取订单通知表单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单通知表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("notify_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("notify_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取订单通知表数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单通知表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input notify.QueryInfo
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
