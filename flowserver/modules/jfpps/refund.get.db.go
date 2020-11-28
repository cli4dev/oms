package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"
)

// DBJFRefund 退款
type DBJFRefund struct {
	*refund.DBRefund
	c component.IContainer
}

// NewDBJFRefund 构建DBPay
func NewDBJFRefund(c component.IContainer) *DBJFRefund {
	db := &DBJFRefund{c: c}
	db.DBRefund = refund.NewDBRefund(c)
	return db
}
