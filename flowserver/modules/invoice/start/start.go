package start

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/qtask"
)

//IInvoice 请求上游申请开票接口
type IInvoice interface {
	Start(orderID, invoiceID, taskID int64) error
}

//Invoice 请求上游申请开票结构体
type Invoice struct {
	c  component.IContainer
	db IDBInvoice
}

//NewInvoice 构建Invoice
func NewInvoice(c component.IContainer, db IDBInvoice) *Invoice {
	return &Invoice{
		c:  c,
		db: db,
	}
}

//Start 请求上游申请开票
func (i *Invoice) Start(orderID, invoiceID, taskID int64) error {

	//1.检查订单和申请记录
	info, canFinishTask, err := i.db.CheckOrderAndInvoiceByDB(orderID, invoiceID)

	if canFinishTask {
		if err := qtask.Finish(i.c, taskID); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	trans, err := i.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 2.调用发货流程任务开始
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	//3.修改发票申请表,申请状态由等待审核变为正在审核
	if err = i.db.StartInvoiceByDB(trans, invoiceID); err != nil {
		trans.Rollback()
		return err
	}

	input, err := i.db.BuildOrderRequestParamByDB(trans, info, invoiceID)
	fmt.Printf("input:%+v", input)
	if err != nil {
		trans.Rollback()
		return err
	}

	//4.调用发货系统发票申请
	fmt.Printf("info:%+v", info)
	result, deliveryQueueFunc, err := jcsdk.CreateOrder(trans, input)
	if err != nil {
		trans.Rollback()
		return err
	}

	if invoiceID != types.GetInt64(result.CoopOrderID) {
		trans.Rollback()
		return fmt.Errorf("发票申请号不相同")
	}

	//5.结束发货流程任务
	if err := qtask.Finish(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	trans.Commit()
	return deliveryQueueFunc(i.c)
}
