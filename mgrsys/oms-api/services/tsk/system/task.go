package system

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/tsk/system"
)

//TaskHandler 任务表接口
type TaskHandler struct {
	container component.IContainer
	TaskLib   system.IDbTask
}

//NewTaskHandler 创建任务表对象
func NewTaskHandler(container component.IContainer) (u *TaskHandler) {
	return &TaskHandler{
		container: container,
		TaskLib:   system.NewDbTask(container),
	}
}
//GetHandle 获取任务表单条数据
func (u *TaskHandler) GetHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取任务表单条数据--------")
	ctx.Log.Info("1.参数校验")

	if err := ctx.Request.Check("task_id"); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, err := u.TaskLib.Get(ctx.Request.GetString("task_id"))
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return data
}

//QueryHandle  获取任务表数据列表
func (u *TaskHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取任务表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input system.QueryTask
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, err := u.TaskLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":data,
		"count":count,
	}
}
