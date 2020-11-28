package refund

// RequestBody 退款请求参数
type RequestBody struct {
	ChannelNo       string  `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`
	RefundNo        string  `json:"refund_no" form:"refund_no" m2s:"refund_no" valid:"required"` //外部退款号
	RequestNo       string  `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
	RefundNum       int     `json:"refund_num" form:"refund_num" m2s:"refund_num" valid:"required"`
	NotifyURL       string  `json:"notify_url" form:"notify_url" m2s:"notify_url"`
	RefundMerAmount float64 `json:"refund_mer_amount" form:"refund_mer_amount" m2s:"refund_mer_amount"`
	RefundPointNum  int     `json:"refund_point_num" form:"refund_point_num" m2s:"refund_point_num"`
}

// QueryRequestBody 退款查询对象
type QueryRequestBody struct {
	ChannelNo string `json:"channel_no" form:"channel_no" m2s:"down_channel_no" valid:"required"`
	RefundNo  string `json:"refund_no" form:"refund_no" m2s:"down_refund_no" valid:"required"`
	RequestNo string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
}
