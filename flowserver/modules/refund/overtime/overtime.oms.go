package overtime

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// OmsOverTime 退款超时
type OmsOverTime struct {
	*DBOverTime
	*OverTime
	*task.QTask
}

// NewOmsOverTime 构建NewOmsOverTime
func NewOmsOverTime(c component.IContainer) *OmsOverTime {
	oot := &OmsOverTime{}
	oot.DBOverTime = NewDBOverTime(c)
	oot.QTask = task.NewQTask(c)
	oot.OverTime = NewOverTime(c, oot, oot)
	return oot
}
