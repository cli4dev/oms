package notify

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/qtask/qtask"
)

// INotify 退款通知接口
type INotify interface {
	RefundNotify(notifyID string, taskID int64) (err error)
}

// Notify 退款通知对象
type Notify struct {
	c  component.IContainer
	db IDBNotify
}

//NewNotify 创建退款通知对象
func NewNotify(c component.IContainer, db IDBNotify) *Notify {
	return &Notify{
		c:  c,
		db: db,
	}
}

// RefundNotify 进行通知
func (n *Notify) RefundNotify(notifyID string, taskID int64) (err error) {

	// 1.获取信息
	err = n.db.CheckByDB(notifyID)
	if context.GetCode(err) == errorcode.NOTIFY_STATUS_ERROR.Code {
		// 关闭任务
		if qerr := qtask.Finish(n.c, taskID); err != nil {
			return qerr
		}
		return err
	}
	if err != nil {
		return
	}
	// 2.开启通知任务
	if err = qtask.Processing(n.c, taskID); err != nil {
		return
	}

	// 3.开始通知
	if err = n.db.BeginNotifyByDB(notifyID); err != nil {
		return
	}

	// 4.获取通知信息
	data, err := n.db.GetNotifyInfoByDB(notifyID)
	if err != nil {
		return
	}

	//4.失败处理函数
	defer func() error {
		if err != nil {
			if failErr := n.db.UpdateNotifyToFailByDB(data.ToMap()); failErr != nil {
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

	// 6.远程通知处理
	content, err := n.db.RemoteNotify(param, data.GetString("notify_url"))
	data.SetValue("notify_msg", content)
	if err != nil {
		return fmt.Errorf("退款通知异常,params:%v,err:%v", data, err)
	}

	// 7.通知成功
	if err = n.db.UpdateNotifyToSuccessByDB(data.ToMap()); err != nil {
		return err
	}

	// 8.关闭任务
	return qtask.Finish(n.c, taskID)
}
