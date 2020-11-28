package refund

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/qtask"
)

// IReturn 订单退货接口
type IReturn interface {
	Start(taskID, returnID int64) (err error)
	Finish(result *order.DeliveryResult) (err error)
}

// Return 订单退货对象
type Return struct {
	c    component.IContainer
	db   IDBReturn
	task task.IQTask
}

//NewReturn 创建订单退货对象
func NewReturn(c component.IContainer, db IDBReturn, task task.IQTask) *Return {
	return &Return{
		c:    c,
		db:   db,
		task: task,
	}
}

// Start 执行退货
func (n *Return) Start(taskID, returnID int64) (err error) {

	err = qtask.Processing(n.c, taskID)
	if err != nil {
		return
	}

	// 1.检查退货和退款信息,返回退货信息
	returnInfo, finish, err := n.db.CheckReturnAndRefundByDB(returnID)
	if finish {
		// 关闭任务
		if qerr := qtask.Finish(n.c, taskID); qerr != nil {
			return qerr
		}
		return
	}
	if err != nil {
		return
	}

	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%+v", err)
	}

	// 2.修改退款信息表和退货信息表的上游退货状态为正在
	if err = n.db.UpdateReturnAndRefundStatusByDB(dbTrans, returnInfo.GetString("refund_id"), returnInfo.GetString("return_id")); err != nil {
		dbTrans.Rollback()
		return
	}

	// 3.构建发货系统参数
	param, err := n.db.BuildReturnRequestParam(dbTrans, returnInfo)
	if err != nil {
		return
	}

	// 4.请求发货系统发货
	result, returnBackFunc, err := jcsdk.CreateOrder(dbTrans, param)
	if err != nil {
		dbTrans.Rollback()
		return err
	}
	if returnInfo.GetString("return_id") != types.GetString(result.CoopOrderID) {
		dbTrans.Rollback()
		return fmt.Errorf("退货订单号不相同")
	}

	if err = qtask.Finish(dbTrans, taskID); err != nil {
		dbTrans.Rollback()
		return
	}
	dbTrans.Commit()

	// 5.回调函数
	return returnBackFunc(n.c)
}

// Finish 保存退货结果
func (n *Return) Finish(result *order.DeliveryResult) (err error) {
	// 1.检查退货状态和退款状态
	refundInfo, err := n.db.CheckReturnAndRefundStatusByDB(result)
	if err != nil {
		return
	}

	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败,err:%+v", err)
	}

	// 2.锁订单,退款记录,退货记录
	if err = n.db.LockByDB(dbTrans, refundInfo.GetString("order_id"), refundInfo.GetString("refund_id"), result.DeliveryID); err != nil {
		dbTrans.Rollback()
		return
	}

	var tasks []string
	switch result.Status {
	case enum.DeliveryResult.Success:
		// 3.成功退货处理,并返回退货成功总金额
		upRefundTasks, returnSuccessAmountInfo, err := n.db.DoForSuccessByDB(dbTrans, refundInfo, result)
		if err != nil {
			dbTrans.Rollback()
			return err
		}
		tasks = n.task.AppendTasks(tasks, upRefundTasks...)

		if returnSuccessAmountInfo.GetBool("finish") {
			// 4.完全退货成功处理,并修改订单退款信息
			allReturnTasks, err := n.db.DoForAllSuccessByDB(dbTrans, refundInfo, returnSuccessAmountInfo, result)
			if err != nil {
				dbTrans.Rollback()
				return err
			}
			tasks = n.task.AppendTasks(tasks, allReturnTasks...)
		}

	case enum.DeliveryResult.Failed:
		// 5.退货失败处理
		notifyTasks, err := n.db.DoForFailByDB(dbTrans, refundInfo, result)
		if err != nil {
			dbTrans.Rollback()
			return err
		}
		tasks = n.task.AppendTasks(tasks, notifyTasks...)
	}

	//6.创建后续流程任务
	methods, err := n.task.CreateTask(dbTrans, tasks, refundInfo.ToMap())
	if err != nil {
		return err
	}

	// 7.发货系统发货完成函数
	scallback, err := jcsdk.SaveNotifyResult(n.c, enum.NotifyResult.Success, result.NotifyID, result.TaskID)
	if err != nil {
		dbTrans.Rollback()
		return err
	}
	methods = append(methods, scallback)
	dbTrans.Commit()
	// 回调
	return utils.CallBackFuncs(n.c, methods...)
}
