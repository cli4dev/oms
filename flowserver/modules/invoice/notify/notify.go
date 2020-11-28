package notify

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
)

//IInvoiceNotify 电子发票申请结果通知
type IInvoiceNotify interface {
	Notify(invoiceID, orderID, taskID int64) error
}

//InvoiceNotify 电子发票申请结果通知结构体
type InvoiceNotify struct {
	c  component.IContainer
	db IDBInvoiceNotify
}

//NewInvoiceNotify 构建InvoiceNotify
func NewInvoiceNotify(c component.IContainer, db IDBInvoiceNotify) *InvoiceNotify {
	return &InvoiceNotify{
		c:  c,
		db: db,
	}
}

//Notify 发票结果通知
func (n *InvoiceNotify) Notify(invoiceID, orderID, taskID int64) error {

	//1.检查订单和发票记录
	data, canFinishTask, err := n.db.CheckOrderAndInvoiceByDB(invoiceID)
	if canFinishTask {
		if err := qtask.Finish(n.c, taskID); err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}

	// 2.开始通知任务
	if err = qtask.Processing(n.c, taskID); err != nil {
		return err
	}

	//3.开始通知
	if err = n.db.StartNotifyInvoiceByDB(invoiceID); err != nil {
		return err
	}

	// 4.构建参数
	param, err := n.db.BuildNotifyParamByDB(data)
	if err != nil {
		return err
	}

	//5.通知
	content, err := n.db.RemoteNotify(n.c, data, param)
	data["msg"] = content

	//6.通知失败,修改通知次数,记录失败信息(超过最大次数,通知失败,并结束通知任务)
	if err != nil {
		if derr := n.db.FailedByDB(data); derr != nil {
			return derr
		}
		return fmt.Errorf("发票结果通知异常,res:%s,params:%v,err:%v", content, data, err)
	}
	// 7.通知成功,修改发票结果通知状态成功,结束通知任务
	if err = n.db.SuccessByDB(data); err != nil {
		return err
	}

	// 8.任务结束
	return qtask.Finish(n.c, taskID)
}
