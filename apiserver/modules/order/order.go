package order

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

// IOrder 订单接口
type IOrder interface {
	Request(info *RequestInfo) (types.IXMap, error)
	QueryOrder(info *QueryInfo) (types.IXMap, error)
}

// Order 请求结构体
type Order struct {
	c    component.IContainer
	db   IDBOrder
	task task.IQTask
}

// NewOrder 构建Order结构体
func NewOrder(c component.IContainer, db IDBOrder, task task.IQTask) *Order {
	return &Order{
		c:    c,
		db:   db,
		task: task,
	}
}

// Request 下单请求
func (o *Order) Request(info *RequestInfo) (types.IXMap, error) {
	// 1.检查订单
	orderInfo, err := o.db.CheckOrderByDB(info.ChannelNO, info.RequestNO)
	if err != nil {
		return nil, err
	}
	if orderInfo != nil {
		return o.db.BuildResult(orderInfo), nil
	}

	//2.检查下单收款账户
	err = o.db.CheckChannelAccount(info.ChannelNO, info.AccountNo)
	if err != nil {
		return nil, err
	}

	// 3.查询产品
	tasks, param, err := o.db.QueryPorductByDB(info)
	if err != nil {
		return nil, err
	}

	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	// 4.锁产品
	orderInfo, err = o.db.LockProductByDB(dbTrans, param)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}
	if orderInfo != nil {
		dbTrans.Rollback()
		return o.db.BuildResult(orderInfo), nil
	}

	// 5.创建订单和通知
	order, err := o.db.CreateOrderAndNotifyByDB(dbTrans, param)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	// 6.创建下游支付任务和超时任务
	queueFuncs, err := o.task.CreateTask(dbTrans, tasks, order)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	dbTrans.Commit()
	if queueFuncs != nil {
		if err := utils.CallBackFuncs(o.c, queueFuncs...); err != nil {
			return nil, err
		}
	}
	return o.db.BuildResult(order), nil
}

// QueryOrder 查询订单
func (o *Order) QueryOrder(info *QueryInfo) (types.IXMap, error) {

	data, err := o.db.QueryByDB(info)
	if err != nil {
		return nil, err
	}

	// 2.判断订单状态是否为正在通知,是,修改订单通知状态为查询成功
	if err := o.db.ProcessNotifyStatus(data); err != nil {
		return nil, err
	}

	return o.db.BuildQueryResult(data)
}
