package invoice

// RequestInfo 开票请求参数
type RequestInfo struct {
	ChannelNO    string  `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`
	InvoiceNO    string  `json:"invoice_no" form:"invoice_no" m2s:"invoice_no" valid:"required"`
	RequestNO    string  `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
	InvoiceTitle string  `json:"invoice_title" form:"invoice_title" m2s:"invoice_title" valid:"required"`
	NotifyURL    string  `json:"notify_url" form:"notify_url" m2s:"notify_url"`
	Remark       string  `json:"remark" form:"remark" m2s:"remark"`
	TaxNO        string  `json:"tax_no" form:"tax_no" m2s:"tax_no"`
	TelePhone    string  `json:"tele_phone" form:"tele_phone" m2s:"tele_phone"`
	Address      string  `json:"address" form:"address" m2s:"address"`
	BankName     string  `json:"bank_name" form:"bank_name" m2s:"bank_name"`
	BankAccount  string  `json:"bank_account" form:"bank_account" m2s:"bank_account"`
	PushType     string  `json:"push_type" form:"push_type" m2s:"push_type" valid:"required"`
	PushPhoneNO  string  `json:"push_phone_no" form:"push_phone_no" m2s:"push_phone_no"`
	PushEmail    string  `json:"push_email" form:"push_email" m2s:"push_email"`
	Amount       float64 `json:"amount" form:"amount" m2s:"amount" valid:"required"`
	DeductAmount float64 `json:"deduct_amount" form:"deduct_amount" m2s:"deduct_amount"`
}

// QueryInfo 开票查询对象
type QueryInfo struct {
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"down_channel_no" valid:"required"`
	InvoiceNO string `json:"invoice_no" form:"invoice_no" m2s:"invoice_no" valid:"required"`
	RequestNO string `json:"request_no" form:"request_no" m2s:"request_no" valid:"required"`
}
