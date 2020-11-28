package order

// RequestInfo 下单结构体
type RequestInfo struct {
	ChannelNO       string  `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`
	AccountNo       string  `json:"account_no" form:"account_no" m2s:"account_no"`
	RequestNO       string  `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
	LineID          int     `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`
	CarrierNO       string  `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`
	ProvinceNo      string  `json:"province_no" form:"province_no" m2s:"province_no"`
	CityNO          string  `json:"city_no" form:"city_no" m2s:"city_no"`
	Num             int     `json:"num" form:"num" m2s:"num" valid:"required,range(0|1000)"`
	Face            float64 `json:"face" form:"face" m2s:"face" valid:"required"`
	Amount          float64 `json:"amount" form:"amount" m2s:"amount" valid:"required"`
	NotifyURL       string  `json:"notify_url" form:"notify_url" m2s:"notify_url"`
	RechargeAccount string  `json:"recharge_account" form:"recharge_account" m2s:"recharge_account"`
	CourierAmount   float64 `json:"courier_amount" form:"courier_amount" m2s:"courier_amount"`
	PointNum        int     `json:"point_num" form:"point_num" m2s:"point_num"`
	PayOrderNo      string  `json:"pay_order_no" form:"pay_order_no" m2s:"pay_order_no"`
}

// QueryInfo 订单查询结构体
type QueryInfo struct {
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`
	RequestNO string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
}
