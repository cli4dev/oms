package refund

import "github.com/micro-plat/hydra/component"

// OmsRefund 退款
type OmsRefund struct {
	*DBRefund
	*Refund
}

// NewOmsRefund 构建OmsRefund
func NewOmsRefund(c component.IContainer) *OmsRefund {
	or := &OmsRefund{}
	or.DBRefund = NewDBRefund(c)
	or.Refund = NewRefund(c, or)
	return or
}
