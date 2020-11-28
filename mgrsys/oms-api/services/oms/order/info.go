package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/order"
)

//InfoHandler 订单记录接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   order.IDbInfo
}

//NewInfoHandler 创建订单记录对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   order.NewDbInfo(container),
	}
}

//GetHandle 获取订单记录单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单记录单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("order_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("order_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//GetVdsHandle 获取上有发货数据数据
func (u *InfoHandler) QueryVdsHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取上有发货数据数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("delivery_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.GetVds(ctx.Request.GetString("delivery_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取订单记录数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单记录数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfo
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
		"data":  data,
		"count": count,
	}
}

//DeliveryHandle 获取订单发货记录
func (u *InfoHandler) DeliveryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单发货记录--------")
	ctx.Log.Info("1.参数校验")
	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Delivery(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//AuditHandle 获取订单审核记录
func (u *InfoHandler) AuditHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单审核记录--------")
	ctx.Log.Info("1.参数校验")
	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Audit(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//NotifyHandle 获取订单通知记录
func (u *InfoHandler) NotifyHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单通知记录--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Notify(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//NotifyHandle 获取订单退款通知记录
func (u *InfoHandler) RefundNotifyHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单通知记录--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.RefundNotify(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//RefundHandle 获取订单退款记录
func (u *InfoHandler) RefundHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单退款记录--------")
	ctx.Log.Info("1.参数校验")
	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Refund(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//ReturnHandle 获取订单退货记录
func (u *InfoHandler) ReturnHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单退货记录--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Return(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//LifetimeHandle 获取订单生命周期
func (u *InfoHandler) LifetimeHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取订单生命周期--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.Lifetime(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//DownPayHandle 获取下游支付记录
func (u *InfoHandler) DownPayHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取下游支付记录--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.DownPay(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}

//UpPayHandle 获取上游支付记录
func (u *InfoHandler) UpPayHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取上游支付记录--------")
	ctx.Log.Info("1.参数校验")

	var input order.QueryInfos
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.InfoLib.UpPay(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
	}
}
