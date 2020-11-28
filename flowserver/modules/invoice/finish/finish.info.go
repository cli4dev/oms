package finish

// InvoiceResult 开票结果参数
type InvoiceResult struct {
	TaskID            int64   `json:"task_id" form:"task_id" m2s:"task_id" valid:"required"`
	NotifyID          string  `json:"notify_id" form:"notify_id" m2s:"notify_id" valid:"required"`
	ChannelNO         string  `json:"coop_id" form:"coop_id" m2s:"coop_id" valid:"required"`
	OrderNo           string  `json:"order_no" form:"order_no" m2s:"order_no" valid:"required"`
	InvoiceID         string  `json:"coop_order_id" form:"coop_order_id" m2s:"coop_order_id" valid:"required"`
	Status            string  `json:"status" form:"status" m2s:"delivery_status" valid:"required"`
	UpOrderNo         string  `json:"up_order_no" form:"up_order_no" m2s:"up_order_no"`
	ResultCode        string  `json:"result_code" form:"result_code" m2s:"result_code"`
	ResultMsg         string  `json:"result_desc" form:"result_desc" m2s:"result_desc"`
	CourierCostAmount float64 `json:"courier_cost_amount" form:"courier_cost_amount" m2s:"courier_cost_amount"`
	ResultParams      string  `json:"result_params" m2s:"result_params"`
}
