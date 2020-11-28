package order

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// Oms OrderOms结构体
type Oms struct {
	*Order
	*DBOrder
	*task.QTask
}

// NewOrderOms oms订单处理
func NewOrderOms(c component.IContainer) *Oms {
	oo := &Oms{}
	oo.DBOrder = NewDBOrder(c)
	oo.QTask = task.NewQTask(c)
	oo.Order = NewOrder(c, oo, oo)
	return oo
}
