package order

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// OvertimeOms oms超时处理
type OvertimeOms struct {
	*DBOvertime
	*Overtime
	*task.QTask
}

// NewOvertimeOms 构建OvertimeOms
func NewOvertimeOms(c component.IContainer) *OvertimeOms {
	oo := &OvertimeOms{}
	oo.DBOvertime = NewDBOvertime(c)
	oo.QTask = task.NewQTask(c)
	oo.Overtime = NewOvertime(c, oo, oo)
	return oo
}
