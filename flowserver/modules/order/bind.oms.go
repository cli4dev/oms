package order

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// BindOms oms绑定处理
type BindOms struct {
	*Bind
	*DBBind
	*task.QTask
}

// NewBindOms 构建BindOms
func NewBindOms(c component.IContainer) *BindOms {
	bo := &BindOms{}
	bo.DBBind = NewDBBind(c)
	bo.QTask = task.NewQTask(c)
	bo.Bind = NewBind(c, bo, bo)
	return bo
}
