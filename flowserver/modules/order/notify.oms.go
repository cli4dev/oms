package order

import (
	"github.com/micro-plat/hydra/component"
)

// NotifyOms oms通知处理
type NotifyOms struct {
	*DBNotify
	*Notify
}

// NewNotifyOms 构建NotifyOms
func NewNotifyOms(c component.IContainer) *NotifyOms {
	no := &NotifyOms{}
	no.DBNotify = NewDBNotify(c)
	no.Notify = NewNotify(c, no)
	return no
}
