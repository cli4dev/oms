package life

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/lcs/life"
)

//TimeHandler 生命周期记录表接口
type TimeHandler struct {
	container component.IContainer
	TimeLib   life.IDbTime
}

//NewTimeHandler 创建生命周期记录表对象
func NewTimeHandler(container component.IContainer) (u *TimeHandler) {
	return &TimeHandler{
		container: container,
		TimeLib:   life.NewDbTime(container),
	}
}
//GetHandle 获取生命周期记录表单条数据
func (u *TimeHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取生命周期记录表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.TimeLib.Get(ctx.Request.GetString("id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取生命周期记录表数据列表
func (u *TimeHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取生命周期记录表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input life.QueryTime
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.TimeLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
