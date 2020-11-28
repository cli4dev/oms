package audit

//RequestParams 审核参数信息
type RequestParams struct {
	OrderID    int64  `json:"order_id" form:"order_id" m2s:"order_id"`
	RefundID   int64  `json:"refund_id" form:"refund_id" m2s:"refund_id"`
	DeliveryID int64  `json:"delivery_id" form:"delivery_id" m2s:"delivery_id"`
	AuditID    int    `json:"audit_id" form:"audit_id" m2s:"audit_id"`
	AuditBy    string `json:"audit_by" form:"audit_by" m2s:"audit_by"`
	AuditMsg   string `json:"audit_msg" form:"audit_msg" m2s:"audit_msg"`
}
