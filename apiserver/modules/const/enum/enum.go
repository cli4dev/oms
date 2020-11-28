package enum

// CanRefund 是否支持退款
var CanRefund = &struct {
	CanRefund    int //支持退款
	CanNotRefund int //不支持退款
}{0, 1}

// RefundType 退款类型
var RefundType = &struct {
	General   int
	Mandatory int
}{1, 2}

// OrderStatus 订单状态
var OrderStatus = &struct {
	Success      int
	Payment      int
	BindDelivery int
	Failed       int
	PartSuccess  int
}{0, 10, 20, 90, 91}

// NotifyStatus 通知状态
var NotifyStatus = &struct {
	Processing int
	Success    int
}{30, 0}

//JFPointType 积分类型
var JFPointType = &struct {
	BuyPoint        string
	ActivityPoint   string
	SubstitutePoint string
}{"buy_point", "activity_point", "substitute_point"}

//JFFDType 积分记账分类
var JFFDType = &struct {
	OrderUse    int
	OrderFail   int
	OrderRefund int
	Void        int
}{1, 2, 3, 4}

// InvoiceType 开票类型
var InvoiceType = &struct {
	NotInvoice      int //不开发票
	PlatInvoiceDown int //开发票
	UpInvoiceUser   int //通知上游给用户开票
}{1, 2, 3}

// InvoiceRole 开票角色
var InvoiceRole = &struct {
	PlatInvoice   int // 平台开票
	UpInvoiceUser int //通知上游给用户开票
}{1, 2}

// InvoiceObjType 开票主体类型
var InvoiceObjType = &struct {
	OrderType   int
	ConsumeType int
}{1, 2}

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
