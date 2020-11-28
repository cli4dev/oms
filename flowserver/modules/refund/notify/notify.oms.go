package notify

import (
	"github.com/micro-plat/hydra/component"
)

// OmsNotify oms通知
type OmsNotify struct {
	*DBNotify
	*Notify
}

// NewOmsNotify 构建
func NewOmsNotify(c component.IContainer) *OmsNotify {
	on := &OmsNotify{}
	on.DBNotify = NewDBNotify(c)
	on.Notify = NewNotify(c, on)
	return on
}
