package refund

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

// IRefund 退款接口
type IRefund interface {
	GeneralRequest(input *RequestBody) (result types.IXMap, err error)
	QueryRefundInfo(input *QueryRequestBody) (result types.XMap, err error)
}

// Refund 普通退款对象
type Refund struct {
	c    component.IContainer
	db   IDBRefund
	task task.IQTask
}

//NewRefund 创建发货下单对象
func NewRefund(c component.IContainer, db IDBRefund, task task.IQTask) *Refund {
	return &Refund{
		c:    c,
		db:   db,
		task: task,
	}
}

// GeneralRequest 普通退款
func (n *Refund) GeneralRequest(input *RequestBody) (result types.IXMap, err error) {
	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return
	}

	//1.检查退款和订单信息
	refInfo, orderInfo, err := n.db.CheckRefundAndOrderGeneralByDb(dbTrans, input)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}
	if refInfo != nil {
		dbTrans.Commit()
		return n.db.BuildRefundResult(refInfo)
	}

	//2.锁订单并检查退款是否正在进行
	if err = n.db.LockAndCheckRefundByDB(dbTrans, input); err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	//3.创建退款及通知信息
	refundInfo, err := n.db.CreateRefundAndNotifyByDB(dbTrans, orderInfo.GetString("order_id"), orderInfo.GetString("extend_info"), input)
	if err != nil {
		dbTrans.Rollback()
		return
	}

	//4.创建并返回退货信息
	overTimeTask, returnTask, returns, err := n.db.CreateProductReturnByDB(dbTrans, input, refundInfo)
	if err != nil {
		dbTrans.Rollback()
		return
	}

	//5.创建超时处理流程任务
	fmt.Printf("refundInfo:%+v", refundInfo)
	funcs, err := n.task.CreateTask(dbTrans, overTimeTask, refundInfo.ToMap())
	if err != nil {
		dbTrans.Rollback()
		return
	}

	//6.创建退货流程任务
	fs, err := n.task.CreateBatchTasks(dbTrans, returnTask, returns)
	if err != nil {
		dbTrans.Rollback()
		return
	}
	dbTrans.Commit()

	// 回调
	funcs = append(funcs, fs...)
	res, err := n.db.BuildRefundResult(refundInfo)
	if err != nil {
		return nil, err
	}
	return res, utils.CallBackFuncs(n.c, funcs...)
}

// QueryRefundInfo 查询退款信息
func (n *Refund) QueryRefundInfo(input *QueryRequestBody) (result types.XMap, err error) {
	// 查询退款信息
	data, err := n.db.QueryRefundBodyByDB(n.c.GetRegularDB(), input)
	if err != nil {
		return
	}

	// 告知处理
	if err = n.db.PrcessRefundNotifyStatusByDB(n.c.GetRegularDB(), data); err != nil {
		return nil, err
	}

	// 返回
	res, err := n.db.BuildRefundResult(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}
