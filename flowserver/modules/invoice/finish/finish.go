package finish

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"
)

//IInvoiceFinish 完成上游发票
type IInvoiceFinish interface {
	Finish(result *InvoiceResult) error
}

//InvoiceFinish 完成上游发票申请结构体
type InvoiceFinish struct {
	c  component.IContainer
	db IDBFinish
	t  task.IQTask
}

//NewInvoiceFinish 构建Finish
func NewInvoiceFinish(c component.IContainer, db IDBFinish, t task.IQTask) *InvoiceFinish {
	return &InvoiceFinish{
		c:  c,
		db: db,
		t:  t,
	}
}

// Finish 完成开票
func (f *InvoiceFinish) Finish(result *InvoiceResult) error {
	param, err := types.Struct2Map(result)
	if err != nil {
		return err
	}

	//1.检查订单和开票申请
	data, canFinishTask, tasks, err := f.db.CheckOrderAndInvoiceByDB(param)
	if canFinishTask {
		scallback, verr := jcsdk.SaveNotifyResult(f.c, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
		if verr != nil {
			return verr
		}
		if serr := scallback(f.c); serr != nil {
			return serr
		}
	}
	if err != nil {
		return err
	}

	trans, err := f.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	if err := f.db.FinishInvoiceByDB(trans, data); err != nil {
		trans.Rollback()
		return err
	}

	if result.Status == enum.InvoiceUnderwayResult {
		scallback, err := jcsdk.SaveNotifyResult(trans, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
		if err != nil {
			trans.Rollback()
			return err
		}
		trans.Commit()
		if serr := scallback(f.c); serr != nil {
			return serr
		}
		return context.NewErrorf(context.ERR_NO_CONTENT, "发货结果未知")
	}

	//5.创建开票通知任务
	queueFuncs, err := f.t.CreateTask(trans, tasks, data)
	if err != nil {
		trans.Rollback()
		return err
	}
	// 6.发货系统发货完成函数
	scallback, err := jcsdk.SaveNotifyResult(trans, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
	if err != nil {
		trans.Rollback()
		return err
	}
	queueFuncs = append(queueFuncs, scallback)
	trans.Commit()

	return utils.CallBackFuncs(f.c, queueFuncs...)
}
