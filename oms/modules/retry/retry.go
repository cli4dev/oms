package retry

import (
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/oms/modules/util"
)

//DeliveryBindRetry 发货绑定重试
func DeliveryBindRetry(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}

	//修改订单超时时间
	if err := GetBindInfo(sdkDB.DBTrans, orderID); err != nil {
		return nil, err
	}

	//创建绑定任务
	var queueFuncs []func(c interface{}) error
	_, bindFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.BindTask,
		map[string]interface{}{"order_id": orderID},
		enum.Task.TaskIntervalTime,
		queue.OrderBind.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, bindFunc), nil
}

//UpPayRetry 上有支付重试
func UpPayRetry(i interface{}, deliveryID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}
	//修改上游支付状态
	if err := UpdateDeliveryUpPayStatus(sdkDB.DBTrans, deliveryID); err != nil {
		return nil, err
	}

	//创建上游支付任务
	var queueFuncs []func(c interface{}) error
	_, upPayFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.UpPaymentTask,
		map[string]interface{}{"delivery_id": deliveryID},
		enum.Task.TaskIntervalTime,
		queue.OrderUpPay.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}
	return append(queueFuncs, upPayFunc), nil
}

//OrderNotifyRetry 订单通知重试
func OrderNotifyRetry(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}
	//修改订单通知状态
	if err := UpdateOrderNotifyStatus(sdkDB.DBTrans, orderID); err != nil {
		return nil, err
	}

	//创建订单通知任务
	var queueFuncs []func(c interface{}) error
	_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.NotifyTask,
		map[string]interface{}{"order_id": orderID},
		enum.Task.TaskIntervalTime,
		queue.OrderNotify.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, notifyFunc), nil
}

//RefundReturnRetry 退款退货重试
func RefundReturnRetry(i interface{}, refundID int64, returnID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}

	//修改退款退货状态
	if err := UpdateReturnStatus(sdkDB.DBTrans, refundID, returnID); err != nil {
		return nil, err
	}

	//创建退款退货任务
	var queueFuncs []func(c interface{}) error
	_, returnFuncs, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.ReturnTask,
		map[string]interface{}{"return_id": returnID},
		enum.Task.TaskIntervalTime,
		queue.StartUpReturn.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, returnFuncs), nil
}

//UpRefundRetry 上游支付退款重试
func UpRefundRetry(i interface{}, returnID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}

	//修改上游支付状态
	if err := UpdateUpRefundStatus(sdkDB.DBTrans, returnID); err != nil {
		return nil, err
	}

	//创建上游支付退款任务
	var queueFuncs []func(c interface{}) error
	_, upRefundFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.UpRefundPaymentTask,
		map[string]interface{}{"return_id": returnID},
		enum.Task.TaskIntervalTime,
		queue.UpRefund.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, upRefundFunc), nil
}

//DownRefundRetry 下游支付退款重试
func DownRefundRetry(i interface{}, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}

	//修改下游支付状态
	if err := UpdateDownRefundStatus(sdkDB.DBTrans, refundID); err != nil {
		return nil, err
	}

	//创建下游支付退款任务
	var queueFuncs []func(c interface{}) error
	_, downRefundFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.DownRefundPaymentTask,
		map[string]interface{}{"refund_id": refundID},
		enum.Task.TaskIntervalTime,
		queue.DownRefund.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, downRefundFunc), nil
}

//RefundNotifyRetry 退款通知重试
func RefundNotifyRetry(i interface{}, orderID int64, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	//获取数据操作对象
	sdkDB, err := util.NewSDKDB(i, data)
	if err != nil {
		return nil, err
	}

	if !sdkDB.IsTrans {
		err = sdkDB.TransBegin()
		if err != nil {
			return nil, err
		}

		defer func() {
			sdkDB.TransClose(err)
		}()
	}

	//修改退款通知状态
	if err := UpdateRefundNotifyStatus(sdkDB.DBTrans, orderID, refundID); err != nil {
		return nil, err
	}

	//创建退款通知任务
	var queueFuncs []func(c interface{}) error
	_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.RefundNotifyTask,
		map[string]interface{}{"refund_id": refundID},
		enum.Task.TaskIntervalTime,
		queue.RefundNotify.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}

	return append(queueFuncs, notifyFunc), nil
}
