package errorcode

import "gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"

//ErrorCode 错误码信息
type ErrorCode struct {
	Code int
	Msg  string
}

// NOTIFY_STATUS_ERROR 退款通知状态错误
var NOTIFY_STATUS_ERROR = &ErrorCode{993, "退款通知状态错误"}

// UP_REFUND_STATUS_ERROR 上游退款状态异常
var UP_REFUND_STATUS_ERROR = &ErrorCode{984, "上游退款状态异常"}

// DOWN_REFUND_STATUS_ERROR 下游退款状态异常
var DOWN_REFUND_STATUS_ERROR = &ErrorCode{985, "下游退款状态异常"}

// ORDER_REFUND_STATUS_ERROR 订单退款状态异常
var ORDER_REFUND_STATUS_ERROR = &ErrorCode{986, "订单退款状态异常"}

// CAN_NOT_REFUND 不支持退款
var CAN_NOT_REFUND = &ErrorCode{980, "不支持退款"}

var JF_CAN_NOT_PART_REFUND = &ErrorCode{981, "使用积分订单不支持部分退款"}

// REFUND_NUM_ERROR 退款数量错误
var REFUND_NUM_ERROR = &ErrorCode{982, "退款数量错误"}

// REFUNDING_ERROR 退款中
var REFUNDING_ERROR = &ErrorCode{983, "退款中"}

// ORDER_STATUS_ERROR 订单状态错误
var ORDER_STATUS_ERROR = &ErrorCode{990, "订单状态错误"}

// DOWN_CHANNEL_NO_EXIST_PRODUCT 下游渠道不存在该商品
var DOWN_CHANNEL_NO_EXIST_PRODUCT = &ErrorCode{920, "下游渠道不存在该商品"}

// DISCOUNT_DIFF_ERROR 折扣差异错误
var DISCOUNT_DIFF_ERROR = &ErrorCode{921, "折扣差异错误"}

var INVOICE_NO_EXIST_ORDER = &ErrorCode{930, "订单不存在或订单状态不正确"}

var INVOICE_NO_EXIST_DELIVERY = &ErrorCode{931, "发货不存在或发货状态不正确"}

var INVOICE_UP_NO_INVOICE = &ErrorCode{932, "上游不支持开票"}

var INVOICE_DOWN_NO_INVOICE = &ErrorCode{933, "下游不支持开票"}

var INVOICE_STATION_NO_INVOICE = &ErrorCode{934, "油站不支持开票"}

var INVOICE_EXIST_REFUND = &ErrorCode{935, "开票存在退款记录"}

var INVOICE_AMOUNT_ERROR = &ErrorCode{936, " 开票金额异常,amount:%v,sell_amount:%v"}

var INVOICE_ORDER_ID_DIFFER = &ErrorCode{937, "订单号不一致"}

var INVOICE_NO_EXIST = &ErrorCode{938, "开票记录不存在"}

var INVOICE_OVERTIME = &ErrorCode{939, "开票超时"}

var INVOICE_AMOUNT_NO_EQUAL = &ErrorCode{940, "金额不一致异常"}

var INVOICE_NO_CONSUME = &ErrorCode{941, "电子券销券记录不存在"}

var INVOICE_DEDUCT_AMOUNT_ERROR = &ErrorCode{942, "抵扣金额不一致,传入抵扣金额:%v,订单抵扣金额:%v"}

// SetFlowRecordStatus 设置流程记录状态、错误码、错误信息
func SetFlowRecordStatus(status int, requestFlowType string, failCode string, failMsg string) (string, string, string) {
	switch status {
	case enum.OrderResultStatus.Failed:
		if failCode != "" {
			return "FAILED", failCode, failMsg
		}
		return "FAILED", "900", requestFlowType + "失败"
	case enum.OrderResultStatus.Success:
		return "SUCCESS", "000", requestFlowType + "成功"
	case enum.OrderResultStatus.PartSuccess:
		return "PARTSUCCESS", "800", requestFlowType + "部分成功"
	default:
		return "UNDERWAY", "000", requestFlowType + "处理中"
	}
}

//RequestFlowType 请求流程分类
var RequestFlowType = &struct {
	Order   string
	Refund  string
	Invoice string
}{"订单", "退款", "开票"}
