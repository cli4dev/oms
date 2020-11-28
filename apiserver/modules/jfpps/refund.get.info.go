package jfpps

// JFRefundRequestBody 退款请求参数
type JFRefundRequestBody struct {
	ChannelNo string  `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`
	RefundNo  string  `json:"refund_no" form:"refund_no" m2s:"refund_no" valid:"required"` //外部退款号
	RequestNo string  `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
	RefundNum int     `json:"refund_num" form:"refund_num" m2s:"refund_num"`
	NotifyURL string  `json:"notify_url" form:"notify_url" m2s:"notify_url"`
	JFAmount  float64 `json:"jf_amount" form:"jf_amount" m2s:"jf_amount"` //积分不够完成退款时，补的金额
}
