package jfpps

//JFPreRequestInfo 未激活积分获取结构体
type JFPreRequestInfo struct {
	ChannelNO    string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`             //渠道编号
	PreRequestNO string `json:"pre_request_no" form:"pre_request_no" m2s:"pre_request_no" valid:"required"` //预下单编号
	LineID       int    `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                      //产品线编号
	PointType    int    `json:"point_type" form:"point_type" m2s:"point_type" valid:"required"`             //积分类型(1.免费赠送，2.购买商品赠送)
	Num          int    `json:"num" form:"num" m2s:"num" valid:"required,range(0|10000000)"`                //积分数量
	UserNo       string `json:"user_no" form:"user_no" m2s:"user_no" valid:"required"`                      //用户编号
	OverTime     string `json:"over_time" form:"over_time" m2s:"over_time"`                                 //积分过期时间
	ActivityNo   string `json:"activity_no" form:"activity_no" m2s:"activity_no" valid:"required"`
	Memo         string `json:"memo" form:"memo" m2s:"memo" valid:"required"`
}

// JFRequestInfo 获取积分结构体
type JFRequestInfo struct {
	ChannelNO    string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"` //渠道编号
	RequestNO    string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"` //下单请求编号
	PreRequestNO string `json:"pre_request_no" form:"pre_request_no" m2s:"pre_request_no"`      //预下单编号（无预下单不传）
	LineID       int    `json:"line_id" form:"line_id" m2s:"line_id"`                           //产品线（无预下单必传）
	PointType    int    `json:"point_type" form:"point_type" m2s:"point_type"`                  //积分类型(1.免费赠送，2.购买商品赠送)（无预下单必传）
	Num          int    `json:"num" form:"num" m2s:"num" valid:"range(0|10000000)"`             //积分数量（无预下单必传）
	UserNo       string `json:"user_no" form:"user_no" m2s:"user_no" `                          //用户编号（无预下单必传）
	NotifyURL    string `json:"notify_url" form:"notify_url" m2s:"notify_url"`                  //通知地址
	OverTime     string `json:"over_time" form:"over_time" m2s:"over_time"`                     //积分过期时间
	Memo         string `json:"memo" form:"memo" m2s:"memo" valid:"required"`
}

//JFFDUseRequest 积分使用记账请求参数
type JFFDUseRequest struct {
	ChannelNO        string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`             //渠道编号
	RequestNo        string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`             //记账请求编号
	ChannelPayNo     string `json:"channel_pay_no" form:"channel_pay_no" m2s:"channel_pay_no" valid:"required"` //外部支付订单号
	BuyPointNum      int    `json:"buy_point_num" form:"buy_point_num" m2s:"buy_point_num"`                     //购买商品赠送积分数
	ActivityPointNum int    `json:"activity_point_num" form:"activity_point_num" m2s:"activity_point_num"`      //参与活动赠送积分数
}

//JFFDVoidRequest 积分作废记账请求参数
type JFFDVoidRequest struct {
	ChannelNO        string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`         //渠道编号
	RequestNo        string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`         //记账请求编号
	GainNo           string `json:"gain_no" form:"gain_no" m2s:"gain_no" valid:"required"`                  //外部获取编号
	BuyPointNum      int    `json:"buy_point_num" form:"buy_point_num" m2s:"buy_point_num"`                 //购买商品赠送积分数
	ActivityPointNum int    `json:"activity_point_num" form:"activity_point_num" m2s:"activity_point_num" ` //参与活动赠送积分数
}
