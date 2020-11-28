package enum

// 审核状态
var (
	AuditSuccess        = 0
	AuditFailed         = 90
	AuditPartialSuccess = 91
	AuditFalseSucc      = 92
)

// 审核类型
var (
	Delivery = 1
	Return   = 2
	Order    = 3
	Refund   = 4
)

// 渠道类型
var (
	DownChannel    = "down_channel"
	DownCommission = "down_commission"
	DownService    = "down_service"
	UpChannel      = "up_channel"
	UpCommission   = "up_commission"
	UpService      = "up_service"
)

var (
	UpChannelGroup   = "upchannel"
	DownChannelGroup = "downchannel"
)
