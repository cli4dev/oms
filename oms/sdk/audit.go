package sdk

import "gitlab.100bm.cn/micro-plat/oms/oms/modules/audit"

//DeliveryUnknownAuditSucc 发货未知审核为成功
//i-->*Context/IContainer/IDB/IDBTrans
//input-->审核参数
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func DeliveryUnknownAuditSucc(i interface{}, input *audit.RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.DeliveryUnknownAuditSucc(i, input, platName, data)
}

//DeliveryUnknownAuditFail 发货未知审核为失败
//i-->*Context/IContainer/IDB/IDBTrans
//input-->审核参数
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func DeliveryUnknownAuditFail(i interface{}, input *audit.RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.DeliveryUnknownAuditFail(i, input, platName, data)
}

//ReturnUnknownAuditSucc 退货未知审核为成功
//i-->*Context/IContainer/IDB/IDBTrans
//input-->审核参数
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func ReturnUnknownAuditSucc(i interface{}, input *audit.RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.ReturnUnknownAuditSucc(i, input, platName, data)
}

//ReturnUnknownAuditFail 退货未知审核为失败
//i-->*Context/IContainer/IDB/IDBTrans
//input-->审核参数
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func ReturnUnknownAuditFail(i interface{}, input *audit.RequestParams, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.ReturnUnknownAuditFail(i, input, platName, data)
}

//OrderPartialSuccessAudit 审核订单部分成功
//i-->*Context/IContainer/IDB/IDBTrans
//orderID-->订单编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func OrderPartialSuccessAudit(i interface{}, orderID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.OrderPartialSuccessAudit(i, orderID, platName, data)
}

//RefundPartialSuccessAudit 审核退款部分成功
//i-->*Context/IContainer/IDB/IDBTrans
//refundID-->退款编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func RefundPartialSuccessAudit(i interface{}, orderID, refundID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.RefundPartialSuccessAudit(i, orderID, refundID, platName, data)
}

//DeliveryFalseSuccAudit 审核发货假成功
//i-->*Context/IContainer/IDB/IDBTrans
//deliveryID-->发货编号
//data-->在i参数为IDB或IDBTrans时,必须传此参数,入参类型-*Context/IContainer
func DeliveryFalseSuccAudit(i interface{}, deliveryID int64, platName string, data ...interface{}) (funcs []func(c interface{}) error, err error) {
	return audit.DeliveryFalseSuccAudit(i, deliveryID, platName, data)
}
