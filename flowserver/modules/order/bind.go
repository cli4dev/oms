package order

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
)

// IBind 绑定接口
type IBind interface {
	BindMethod(orderID, taskID int64) error
}

// Bind 绑定结构体
type Bind struct {
	c    component.IContainer
	db   IDBBind
	task task.IQTask
}

// NewBind 构建Bind
func NewBind(c component.IContainer, db IDBBind, task task.IQTask) *Bind {
	return &Bind{
		c:    c,
		db:   db,
		task: task,
	}
}

// BindMethod 绑定
func (b *Bind) BindMethod(orderID, taskID int64) error {

	trans, err := b.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	// 1.开始任务
	if err = qtask.Processing(trans, taskID); err != nil {
		trans.Rollback()
		return err
	}

	// 2.锁定单
	orderInfo, err := b.db.LockOrderForBindByDB(trans, orderID)
	if err != nil {
		trans.Rollback()
		return err
	}
	if orderInfo == nil {
		if qerr := qtask.Finish(trans, taskID); qerr != nil {
			trans.Rollback()
			return err
		}
		trans.Commit()
		return fmt.Errorf("绑定锁订单失败，订单不存在或已完成绑定,order_id:%d", orderID)
	}

	// 3.查询上游商品
	product, err := b.db.QueryProductByDB(trans, orderInfo)
	if err != nil {
		trans.Rollback()
		return err
	}

	// 4.绑定
	tasks, info, completeBind, err := b.db.BindUpProductByDB(trans, product, orderInfo)
	if err != nil {
		trans.Rollback()
		return err
	}
	// 5.创建发货流程任务
	queueFuncs, err := b.task.CreateTask(trans, tasks, info)
	if err != nil {
		trans.Rollback()
		return err
	}
	// 6.业务处理成功，修改任务状态为完成(任务不再放入队列)
	if completeBind {
		if err := qtask.Finish(trans, taskID); err != nil {
			trans.Rollback()
			return err
		}
	}
	trans.Commit()
	return utils.CallBackFuncs(b.c, queueFuncs...)
}
