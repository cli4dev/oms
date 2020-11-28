package queue

// OrderDownPay 订单请求
var OrderDownPay = &Queue{"order:pay", "/order/pay/down"}

// OrderBind 订单绑定
var OrderBind = &Queue{"order:bind", "/order/bind"}

// OrderStartDelivery 订单开始发货
var OrderStartDelivery = &Queue{"order:delivery", "/order/delivery/start"}

// OrderFinishDelivery 订单发货完成
var OrderFinishDelivery = &Queue{"order:delivery:finish", "/order/delivery/finish"}

// OrderUpPay 上游支付
var OrderUpPay = &Queue{"order:uppay", "/order/pay/up"}

// OrderNotify 订单通知
var OrderNotify = &Queue{"order:notify", "/order/notify"}

// OrderOvertime 订单超时
var OrderOvertime = &Queue{"order:overtime", "/order/overtime"}

// OrderDeliveryOvertime 发货未知
var OrderDeliveryOvertime = &Queue{"order:delivery:unknown", "/order/overtime/delivery"}

// StartUpReturn 开始退货
var StartUpReturn = &Queue{"refund:return", "/refund/product/return"}

// FinishUpReturn 退货完成
var FinishUpReturn = &Queue{"refund:return:finish", "/refund/product/finish"}

// DownRefund 订单下游退款
var DownRefund = &Queue{"refund:pay", "/refund/pay/down"}

// UpRefund 上游退款
var UpRefund = &Queue{"refund:uppay", "/refund/pay/up"}

// OrderFailedRefund 订单失败退款
var OrderFailedRefund = &Queue{"refund:order:fail", "/refund/pay/orderfail"}

// RefundNotify 退款通知
var RefundNotify = &Queue{"refund:notify", "/refund/notify"}

// RefundOvertime 退款超时
var RefundOvertime = &Queue{"refund:overtime", "/refund/overtime"}

// ReturnOvertime 退货超时
var ReturnOvertime = &Queue{"refund:return:unknown", "/refund/overtime/unknown"}

// InvoiceStart 开始开票
var InvoiceStart = &Queue{"invoice:start", "/crp/invoice/start"}

// InvoiceFinish 完成开票
var InvoiceFinish = &Queue{"invoice:finish", "/crp/invoice/finish"}

// InvoiceNotify 开票通知
var InvoiceNotify = &Queue{"invoice:notify", "/crp/invoice/notify"}

// InvoicePay 开票支付
var InvoicePay = &Queue{"invoice:pay", "/crp/invoice/pay"}

// InvoiceRedStart 开始开票
var InvoiceRedStart = &Queue{"red:start", "/crp/red/start"}

// InvoiceRedFinish 发票冲红完成
var InvoiceRedFinish = &Queue{"red:finish", "/crp/red/finish"}

// InvoiceRedFefund 冲红退款
var InvoiceRedFefund = &Queue{"red:refund", "/crp/red/refund"}

// InvoiceRedNotify 开票通知
var InvoiceRedNotify = &Queue{"red:notify", "/crp/red/notify"}

// JFGetUpPay 积分获取上游支付
var JFGetUpPay = &Queue{"point:get:order:uppay", "/point/get/pay/up"}

//PointUseRefundFD 积分使用退款记账
var PointUseRefundFD = &Queue{"point:use:refund:fd", "/point/use/refund/fd"}

// JFGetFinishUpReturn 积分发放退货完成
var JFGetFinishUpReturn = &Queue{"point:get:return:finish", "/point/get/return/finish"}

// JFGetUpRefund 积分发放上游退款
var JFGetUpRefund = &Queue{"point:get:refund:uppay", "/point/get/refund/up"}
