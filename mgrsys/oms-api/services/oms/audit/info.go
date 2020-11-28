package audit

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/sso/sdk/sso"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/oms/audit"
	au "gitlab.100bm.cn/micro-plat/oms/oms/modules/audit"
)

//InfoHandler 发货人工审核表接口
type InfoHandler struct {
	container component.IContainer
	InfoLib   audit.IDbInfo
}

//NewInfoHandler 创建发货人工审核表对象
func NewInfoHandler(container component.IContainer) (u *InfoHandler) {
	return &InfoHandler{
		container: container,
		InfoLib:   audit.NewDbInfo(container),
	}
}

//GetHandle 获取发货人工审核表单条数据
func (u *InfoHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货人工审核表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("audit_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.InfoLib.Get(ctx.Request.GetString("audit_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取发货人工审核表数据列表
func (u *InfoHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取发货人工审核表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input audit.QueryInfo
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

//PutHandle 更新发货人工审核表数据
func (u *InfoHandler) PutHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------人工审核发货结果--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("audit_status", "change_type"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	var input au.RequestParams
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}
	input.AuditBy = types.GetString(sso.GetMember(ctx).UserID)

	ctx.Log.Info("2.执行操作")
	err := u.InfoLib.Audit(&input, ctx.Request.GetInt("audit_status"), ctx.Request.GetInt("change_type"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return "success"

}
