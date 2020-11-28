package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
)

// DBJFPay 支付
type DBJFPay struct {
	*order.DBPay
	c component.IContainer
}

// NewDBJFPay 构建DBPay
func NewDBJFPay(c component.IContainer) *DBJFPay {
	db := &DBJFPay{c: c}
	db.DBPay = order.NewDBPay(c)
	return db
}
