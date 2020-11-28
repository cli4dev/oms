package order

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// DeliveryOms oms开始发货处理
type DeliveryOms struct {
	*DBDelivery
	*Delivery
	*task.QTask
}

// NewDeliveryOms 构建DeliveryOms
func NewDeliveryOms(c component.IContainer) *DeliveryOms {
	do := &DeliveryOms{}
	do.DBDelivery = NewDBDelivery(c)
	do.QTask = task.NewQTask(c)
	do.Delivery = NewDelivery(c, do, do)
	return do
}
