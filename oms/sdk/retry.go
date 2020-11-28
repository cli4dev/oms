package sdk

import "gitlab.100bm.cn/micro-plat/oms/oms/modules/retry"

//DeliveryBindRetry 发货绑定重试
//i-->*Context/IContainer/IDB/IDBTrans
//orderID-->订单编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func DeliveryBindRetry(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.DeliveryBindRetry(i, orderID, platName, data)
}

//UpPayRetry 上游支付重试
//i-->*Context/IContainer/IDB/IDBTrans
//deliveryID-->上游发货编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func UpPayRetry(i interface{}, deliveryID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.UpPayRetry(i, deliveryID, platName, data)
}

//OrderNotifyRetry 订单通知重试
//i-->*Context/IContainer/IDB/IDBTrans
//orderID-->订单编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func OrderNotifyRetry(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.OrderNotifyRetry(i, orderID, platName, data)
}

//RefundReturnRetry 上游退货重试
//i-->*Context/IContainer/IDB/IDBTrans
//returnID-->退货编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func RefundReturnRetry(i interface{}, refundID int64, returnID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.RefundReturnRetry(i, refundID, returnID, platName, data)
}

//UpRefundRetry 上游退款重试
//i-->*Context/IContainer/IDB/IDBTrans
//returnID-->退货编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func UpRefundRetry(i interface{}, returnID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.UpRefundRetry(i, returnID, platName, data)
}

//DownRefundRetry 下游退款重试
//i-->*Context/IContainer/IDB/IDBTrans
//refundID-->退款编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func DownRefundRetry(i interface{}, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.DownRefundRetry(i, refundID, platName, data)
}

//RefundNotifyRetry 退款通知重试
//i-->*Context/IContainer/IDB/IDBTrans
//orderID-->订单编号
//refundID-->退款编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func RefundNotifyRetry(i interface{}, orderID int64, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return retry.RefundNotifyRetry(i, orderID, refundID, platName, data)
}
