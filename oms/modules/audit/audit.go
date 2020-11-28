package audit

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/oms/modules/util"

	"github.com/micro-plat/qtask/qtask"
)

//DeliveryUnknownAuditSucc 审核发货未知为发货成功
func DeliveryUnknownAuditSucc(i interface{}, input *RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//检查发货信息
	info, err := DeliveryAuditCheck(sdkDB.DBTrans, input)
	if err != nil {
		return nil, err
	}

	//处理发货成功-控制并发
	notifyID, err := DeliveryAuditSuccess(sdkDB.DBTrans, input, info)
	if err != nil {
		return nil, err
	}

	//处理审核记录审核成功
	err = SetAuditRecordSuccess(sdkDB.DBTrans, input, 1)
	if err != nil {
		return nil, err
	}

	//创建上游支付任务
	var queueFuncs []func(c interface{}) error
	_, payFunc, err := qtask.Create(sdkDB.DBTrans,
		enum.Task.UpPaymentTask,
		map[string]interface{}{"delivery_id": input.DeliveryID},
		enum.Task.TaskIntervalTime,
		queue.OrderUpPay.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}
	queueFuncs = append(queueFuncs, payFunc)

	//判断发货是否全部成功，创建通知任务
	if notifyID != 0 {
		// 6.创建通知任务
		_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.NotifyTask,
			map[string]interface{}{"notify_id": notifyID},
			enum.Task.TaskIntervalTime,
			queue.OrderNotify.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, notifyFunc)
	}

	return queueFuncs, nil
}

//DeliveryUnknownAuditFail 审核发货未知为发货失败
func DeliveryUnknownAuditFail(i interface{}, input *RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//检查发货信息
	_, err = DeliveryAuditCheck(sdkDB.DBTrans, input)
	if err != nil {
		return nil, err
	}

	//处理发货失败-控制并发
	isAllDeliveryFail, notifyID, err := DeliveryAuditFail(sdkDB.DBTrans, input)
	if err != nil {
		return nil, err
	}

	//处理审核记录审核失败
	err = SetAuditRecordFail(sdkDB.DBTrans, input, 1)
	if err != nil {
		return nil, err
	}

	//判断订单全部发货失败
	var queueFuncs []func(c interface{}) error
	if isAllDeliveryFail {
		//创建订单失败退款任务
		_, refundFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.RefundTask,
			map[string]interface{}{"order_id": input.OrderID},
			enum.Task.TaskIntervalTime,
			queue.OrderFailedRefund.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, refundFunc)
	}

	if notifyID != 0 {
		//创建订单失败通知任务
		_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.NotifyTask,
			map[string]interface{}{"notify_id": notifyID},
			enum.Task.TaskIntervalTime,
			queue.OrderNotify.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, notifyFunc)
	}

	return queueFuncs, nil
}

//ReturnUnknownAuditSucc 审核退货未知为退货成功
func ReturnUnknownAuditSucc(i interface{}, input *RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//处理退货成功-控制并发
	isAllReturnSucc, info, err := ReturnAuditSuccess(sdkDB.DBTrans, input)
	if err != nil {
		return nil, err
	}

	//处理审核记录审核成功
	err = SetAuditRecordSuccess(sdkDB.DBTrans, input, 2)
	if err != nil {
		return nil, err
	}

	var queueFuncs []func(c interface{}) error
	if isAllReturnSucc {
		//完全退货退款与订单数据处理
		notifyID, err := ReturnAuditComleteSuccess(sdkDB.DBTrans, info)
		if err != nil {
			return nil, err
		}

		//上游支付退款流程任务创建
		_, upRefundFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.UpRefundPaymentTask,
			map[string]interface{}{"return_id": input.DeliveryID},
			enum.Task.TaskIntervalTime,
			queue.UpRefund.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, upRefundFunc)

		//下游支付退款流程任务创建
		_, downRefundFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.DownRefundPaymentTask,
			map[string]interface{}{"refund_id": input.RefundID},
			enum.Task.TaskIntervalTime,
			queue.DownRefund.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, downRefundFunc)

		//通知流程任务创建
		if notifyID != 0 {
			_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
				enum.Task.RefundNotifyTask,
				map[string]interface{}{"notify_id": notifyID},
				enum.Task.TaskIntervalTime,
				queue.RefundNotify.GetName(platName),
				qtask.WithDeadline(900))
			if err != nil {
				return nil, err
			}
			queueFuncs = append(queueFuncs, notifyFunc)
		}
	}

	return queueFuncs, nil
}

//ReturnUnknownAuditFail 审核退货未知为退货失败
func ReturnUnknownAuditFail(i interface{}, input *RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//处理退货为失败-控制并发
	isAllReturnFail, err := ReturnAuditFail(sdkDB.DBTrans, input)
	if err != nil {
		return nil, err
	}

	//处理审核记录审核失败
	if err = SetAuditRecordFail(sdkDB.DBTrans, input, 2); err != nil {
		return nil, err
	}

	var queueFuncs []func(c interface{}) error
	if isAllReturnFail {
		//退货全部失败处理退款信息
		notifyID, err := ReturnAuditCompleteFail(sdkDB.DBTrans, input)
		if err != nil {
			return nil, err
		}
		if notifyID == 0 {
			return queueFuncs, nil
		}

		//创建通知任务
		_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
			enum.Task.RefundNotifyTask,
			map[string]interface{}{"notify_id": notifyID},
			enum.Task.TaskIntervalTime,
			queue.RefundNotify.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, notifyFunc)
	}

	return queueFuncs, nil
}

//OrderPartialSuccessAudit 订单部分成功审核
func OrderPartialSuccessAudit(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//判断并处理订单部分成功
	notifyID, err := UpdateOrderPartialSuccess(sdkDB.DBTrans, orderID)
	if err != nil {
		return nil, err
	}

	//创建支付退款任务
	var queueFuncs []func(c interface{}) error
	_, downRefundFunc, err := qtask.Create(sdkDB.DBTrans,
		"部分成功支付退款",
		map[string]interface{}{"order_id": orderID},
		enum.Task.TaskIntervalTime,
		queue.OrderFailedRefund.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}
	queueFuncs = append(queueFuncs, downRefundFunc)

	//创建订单通知任务
	if notifyID != 0 {
		_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
			"部分成功通知",
			map[string]interface{}{"notify_id": notifyID},
			enum.Task.TaskIntervalTime,
			queue.OrderNotify.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, notifyFunc)
	}

	return queueFuncs, nil
}

//RefundPartialSuccessAudit 退款部分成功审核
func RefundPartialSuccessAudit(i interface{}, orderID int64, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//判断并处理退款部分成功
	notifyID, err := UpdateRefundPartialSuccess(sdkDB.DBTrans, orderID, refundID)
	if err != nil {
		return nil, err
	}

	//创建下游退款任务
	var queueFuncs []func(c interface{}) error
	_, refundFunc, err := qtask.Create(sdkDB.DBTrans,
		"退款部分成功支付退款",
		map[string]interface{}{"refund_id": refundID},
		enum.Task.TaskIntervalTime,
		queue.DownRefund.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}
	queueFuncs = append(queueFuncs, refundFunc)

	//创建通知任务
	if notifyID != 0 {
		_, notifyFunc, err := qtask.Create(sdkDB.DBTrans,
			"退款部分成功通知",
			map[string]interface{}{"notify_id": notifyID},
			enum.Task.TaskIntervalTime,
			queue.RefundNotify.GetName(platName),
			qtask.WithDeadline(900))
		if err != nil {
			return nil, err
		}
		queueFuncs = append(queueFuncs, notifyFunc)
	}

	return queueFuncs, nil
}

//DeliveryFalseSuccAudit 发货假成功审核
func DeliveryFalseSuccAudit(i interface{}, deliveryID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
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

	//判断发货状态是否正常
	deliveryInfo, err := GetFalseSuccDeliveryInfo(sdkDB.DBTrans, deliveryID)
	if err != nil {
		return nil, err
	}

	//创建假成功退款记录与退货记录
	returnID, err := CreateFalseSuccReturn(sdkDB.DBTrans, deliveryInfo)
	if err != nil {
		return nil, err
	}

	//创建上游退款任务
	var queueFuncs []func(c interface{}) error
	_, upRefundFunc, err := qtask.Create(sdkDB.DBTrans,
		"假成功上游退款",
		map[string]interface{}{"return_id": returnID},
		enum.Task.TaskIntervalTime,
		queue.UpRefund.GetName(platName),
		qtask.WithDeadline(900))
	if err != nil {
		return nil, err
	}
	queueFuncs = append(queueFuncs, upRefundFunc)

	return queueFuncs, nil
}
