package account

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/beanpay/account"
)

//RecordHandler 账户余额变动信息接口
type RecordHandler struct {
	container component.IContainer
	RecordLib account.IDbRecord
}

//NewRecordHandler 创建账户余额变动信息对象
func NewRecordHandler(container component.IContainer) (u *RecordHandler) {
	return &RecordHandler{
		container: container,
		RecordLib: account.NewDbRecord(container),
	}
}

//GetHandle 获取账户余额变动信息单条数据
func (u *RecordHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取账户余额变动信息单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("record_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.RecordLib.Get(ctx.Request.GetString("record_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取账户余额变动信息数据列表
func (u *RecordHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取账户余额变动信息数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input account.QueryRecord
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.RecordLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//RedrushHandle 加款红冲数据
func (u *RecordHandler) RedrushHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------加款红冲--------")
	ctx.Log.Info("1.参数校验")

	var input account.RedrushInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.RecordLib.Redrush(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//RedrushDrawHandle 提款红冲数据
func (u *RecordHandler) RedrushDrawHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------提款红冲--------")
	ctx.Log.Info("1.参数校验")

	var input account.RedrushInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.RecordLib.RedrushDraw(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}

//FlatAccountHandle 交易平账
func (u *RecordHandler) FlatAccountHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------交易平账--------")
	ctx.Log.Info("1.参数校验")

	var input account.FlatAccountInfo
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	err := u.RecordLib.FlatAccount(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"
}
