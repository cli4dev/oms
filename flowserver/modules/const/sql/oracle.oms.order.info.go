// +build oracle

package sql

const SqlUpdateOrderStatus = `
update oms_order_info t
set t.refund_status = 0
where t.order_id = @order_id
  and t.order_status = 90
  and t.refund_status = 20	
`

const SqlCheckOrderStatus = `
select 
	a.order_id,
	a.down_channel_no,
	a.down_account_no,
	round(a.total_face,5) refund_unit,
	round(a.sell_amount,5) sell_amount,
	round(a.commission_amount ,5) commission_amount,
	round(a.service_amount,5) service_amount,
	round(a.fee_amount,5) fee_amount,
	a.sell_amount refund_sell,
	a.sell_amount - a.commission_amount + a.service_amount real_refund,
	a.fee_amount refund_fee,
	to_char(a.create_time,'yyyy-mm-dd hh24:mi:ss') create_time
	from oms_order_info a
where a.order_id = @order_id
  and a.refund_status = 20
  and a.order_status in (90,91)
`

const SqlUpdateOrder = `
update oms_order_info t
set t.is_refund = 0,
	t.order_status = decode(t.success_face - @return_total_face,0,90,91),
	t.fail_code=decode(t.success_face - @return_total_face , 0,@fail_code,''),
	t.fail_msg=decode(t.success_face - @return_total_face ,0,@fail_msg,''),
	t.success_face = t.success_face - @return_total_face,
	t.success_sell_amount = t.success_sell_amount - @refund_sell_amount,
	t.success_commission = t.success_commission - @refund_commission_amount,
	t.success_service = t.success_service - @refund_service_amount,
	t.success_fee = t.success_fee - @refund_fee_amount,
	t.success_cost_amount = t.success_cost_amount - @return_cost_amount,
	t.success_up_commission = t.success_up_commission - @return_commission_amount,
	t.success_up_service = t.success_up_service - @return_service_amount,
	t.profit = t.profit - ( @refund_sell_amount - @refund_commission_amount + @refund_service_amount - @refund_fee_amount - @return_cost_amount + @return_commission_amount + @return_service_amount) 
where t.order_id = @order_id
  and t.order_status in (0,91)
`
const QueryOrderInfo = `select t.point_num from oms_order_info t where t.order_id = @order_id
`

const SqlLockOrder = `
select 
	t.order_id 
from oms_order_info t
where t.order_id = @order_id
  and t.order_status in (0,91) 
for update 
`
