package finish

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// InvoiceFinishOms oms开票通知处理
type InvoiceFinishOms struct {
	*InvoiceFinish
	IDBFinish
	*task.QTask
}

// NewInvoiceFinishOms 构建NewInvoiceFinishOms
func NewInvoiceFinishOms(c component.IContainer) *InvoiceFinishOms {
	bo := &InvoiceFinishOms{}
	bo.IDBFinish = NewDBFinish(c)
	bo.QTask = task.NewQTask(c)
	bo.InvoiceFinish = NewInvoiceFinish(c, bo, bo)
	return bo
}
