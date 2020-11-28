package enum

// DeliveryStatus 发货状态
var DeliveryStatus = &struct {
	Success    int
	Wait       int
	Processing int
	Failed     int
}{0, 20, 30, 90}

// RefundStatus 退款状态
var RefundStatus = &struct {
	Return         int
	Refund         int
	ReturnSuccess  int
	ReturnFail     int
	ReturnPartFail int
}{10, 20, 0, 90, 91}

// NotifyResult 通知结果
var NotifyResult = &struct {
	Success string
	Failed  string
}{"success", "failed"}

// OrderStatus 订单状态
var OrderStatus = &struct {
	DownPay      int
	BindDelivery int
}{10, 20}

//DeliveryResult 发货结果
var DeliveryResult = &struct {
	Success  string
	Failed   string
	Underway string
}{"SUCCESS", "FAILED", "UNKNOWN"}

// CanSplitOrder 拆单
var CanSplitOrder = &struct {
	SplitOrder   int
	NoSplitOrder int
}{0, 1}

// OrderResultStatus 结果状态
var OrderResultStatus = &struct {
	Success     int
	Processing  int
	Failed      int
	PartSuccess int
}{0, 30, 90, 91}

// 开票申请结果
const (
	InvoiceSuccessResult  = "SUCCESS"
	InvoiceFailedResult   = "FAILED"
	InvoiceNoNeed         = "NoNeed"
	InvoiceUnderwayResult = "UNDERWAY"
)

// InvoiceMethod 开票方式
var InvoiceMethod = &struct {
	PlatInvoice int
	UpInvoice   int
}{1, 2}

// InvoiceType 开票类型
var InvoiceType = &struct {
	Normal int
	Red    int
}{1, 2}

// Group 签名分组
const Group = "downchannel"

//JFFDType 积分记账分类
var JFFDType = &struct {
	OrderUse    int
	OrderFail   int
	OrderRefund int
	Void        int
}{1, 2, 3, 4}

var PointType = &struct {
	Activity   int
	Buy        int
	Substitute int
}{1, 2, 3}

// RefundResultStatus 退款返回状态
var RefundResultStatus = &struct {
	Success     int
	Refunding   int
	Failed      int
	PartSuccess int
}{0, 30, 90, 91}

//FailCode 订单失败错误编码
type FailCode struct {
	UpCode   string //上游返回的错误编码
	PlatCode string //系统对应的错误编码
	PlatMsg  string //系统对应的错误信息
}

//DeliveryFailCode 发货失败错误编码信息配置
var DeliveryFailCode = &struct {
	DownRefund *FailCode
}{
	&FailCode{
		UpCode:   "900",
		PlatCode: "900",
		PlatMsg:  "订单已退款",
	},
}
