package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// JFReturn JFReturn
type JFReturn struct {
	*refund.Return
}

// NewJFReturn NewJFReturn退货处理
func NewJFReturn(c component.IContainer) *JFReturn {
	ro := &JFReturn{}
	ro.Return = refund.NewReturn(c, NewDBJFReturn(c), task.NewQTask(c))
	return ro
}
