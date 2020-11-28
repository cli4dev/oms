package refund

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// OmsReturn oms退款
type OmsReturn struct {
	*DBReturn
	*Return
	*task.QTask
}

// NewOmsReturn 构建OmsReturn
func NewOmsReturn(c component.IContainer) *OmsReturn {
	or := &OmsReturn{}
	or.DBReturn = NewDBReturn(c)
	or.QTask = task.NewQTask(c)
	or.Return = NewReturn(c, or, or)
	return or
}
