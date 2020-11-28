package order

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
)

// INotify 通知接口
type INotify interface {
	NotifyRequest(notifyID, taskID int64) error
}

// Notify 通知结构体
type Notify struct {
	c  component.IContainer
	db IDBNotify
}

// NewNotify 构建Notify
func NewNotify(c component.IContainer, db IDBNotify) *Notify {
	return &Notify{
		c:  c,
		db: db,
	}
}

// NotifyRequest 通知
func (n *Notify) NotifyRequest(notifyID, taskID int64) error {

	// 1.查询订单信息
	data, canFinishTask, err := n.db.CheckOrderAndNotifyByDB(notifyID)
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

	// 3.开始通知
	if err = n.db.StartByDB(data); err != nil {
		return err
	}

	//4.失败处理函数
	defer func() error {
		if err != nil {
			if failErr := n.db.FailedByDB(data); failErr != nil {
				return failErr
			}
		}
		return nil
	}()
	// 5.构建通知参数
	param, err := n.db.BuildNotifyParamByDB(data)
	if err != nil {
		return err
	}

	// 6.远程通知
	content, err := n.db.RemoteNotify(param, data.GetString("notify_url"))
	data["msg"] = content
	if err != nil {
		return fmt.Errorf("订单通知异常,res:%s,params:%v,err:%v", content, param, err)
	}

	// 7.通知成功,修改订单通知状态成功,通知记录状态成功,结束通知任务
	if err = n.db.SuccessByDB(data); err != nil {
		return err
	}

	// 8.任务结束
	return qtask.Finish(n.c, taskID)
}
