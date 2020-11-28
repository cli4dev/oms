package invoice

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

// IInvoice 申请发票接口
type IInvoice interface {
	Request(param *RequestInfo) (types.IXMap, error)
	Query(input *QueryInfo) (types.XMap, error)
}

//Invoice 请求结构体
type Invoice struct {
	c    component.IContainer
	db   IDBInvoice
	task task.IQTask
}

// NewInvoice 构建Invoice结构体
func NewInvoice(c component.IContainer, db IDBInvoice, task task.IQTask) *Invoice {
	return &Invoice{
		c:    c,
		db:   db,
		task: task,
	}
}

// Request 请求开票
func (i *Invoice) Request(param *RequestInfo) (types.IXMap, error) {

	// 1.检查开票
	invoice, err := i.db.CheckInvoiceByDB(i.c.GetRegularDB(), param.ChannelNO, param.InvoiceNO)
	if !invoice.IsEmpty() {
		return i.db.BuildResultByDB(invoice, param), nil
	}

	//2.检查订单和发货
	fmt.Printf("param:%+v", param)
	order, err := i.db.CheckOrderAndDeliveryByDB(param.ChannelNO, param.RequestNO, param.InvoiceNO, param.Amount, param.DeductAmount)
	if err != nil {
		return nil, err
	}

	//2.开启事务
	dbTrans, err := i.c.GetRegularDB().Begin()
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	//3.锁订单
	if err = i.db.LockOrderByDB(dbTrans, order.GetInt64("order_id")); err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	//4.检查开票记录
	invoice, err = i.db.CheckInvoiceByDB(dbTrans, param.ChannelNO, param.InvoiceNO)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}
	if !invoice.IsEmpty() {
		if order.GetInt64("order_id") != invoice.GetInt64("order_id") {
			return nil, context.NewErrorf(errorcode.INVOICE_ORDER_ID_DIFFER.Code, "订单号不一致,order_id:%v,order_id:%v",
				order.GetInt64("order_id"), invoice.GetString("order_id"))
		}
		dbTrans.Commit()
		return i.db.BuildResultByDB(invoice, param), nil
	}

	//5.创建开票申请记录
	invoice, tasks, err := i.db.CreateInvoiceByDB(dbTrans, param, order)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}
	// 6.创建下游支付任务和超时任务
	fmt.Printf("tasks:%+v", tasks)
	queueFuncs, err := i.task.CreateTask(dbTrans, tasks, order)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	dbTrans.Commit()
	if queueFuncs != nil {
		if err := utils.CallBackFuncs(i.c, queueFuncs...); err != nil {
			return nil, err
		}
	}

	//9.构建返回参数
	return i.db.BuildResultByDB(invoice, param), nil
}

// Query 查询开票
func (i *Invoice) Query(input *QueryInfo) (types.XMap, error) {
	invoice, err := i.db.CheckInvoiceByDB(i.c.GetRegularDB(), input.ChannelNO, input.InvoiceNO)
	if err != nil {
		return nil, err
	}
	return i.db.BuildQueryByDB(invoice, input), nil
}
