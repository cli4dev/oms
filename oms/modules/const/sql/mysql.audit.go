// +build !oracle

package sql

const AuditDeliverySuccess = `UPDATE oms_order_delivery t
SET t.delivery_status   = 0,
   t.up_payment_status = 30,
   t.end_time          = NOW(),
   t.return_msg        = @audit_msg
WHERE t.delivery_id = @delivery_id
AND t.order_id = @order_id
AND t.delivery_status = 30
AND t.up_payment_status = 10`

const GetDeliveryInfo = `SELECT t.order_id,
t.face,
t.num,
t.total_face success_face,
t.cost_amount,
t.up_commission_amount up_commission,
t.service_amount up_service,
p.sell_discount,
p.sell_price,
p.commission_discount,
p.service_discount,
p.payment_fee_discount fee_discount
FROM oms_order_delivery t
INNER JOIN oms_down_product p
ON t.down_product_id = p.product_id
WHERE t.delivery_id = @delivery_id
AND t.order_id = @order_id
AND t.delivery_status = 30
AND t.up_payment_status = 10
`

const AuditOrderSuccess = `
update oms_order_info t
set t.order_status          = case when (t.success_face + @success_face =t.total_face) then 0 else 20 end,
	t.delivery_bind_status  = case when (t.success_face + @success_face =t.total_face) then 0 else 30 end,
	t.notify_status         = case when (t.success_face + @success_face =t.total_face) then 30 else 10 end,
	t.success_face          = t.success_face + @success_face,
	t.success_sell_amount   = t.success_sell_amount + @num * @sell_price,
	t.success_commission    = t.success_commission + @success_face * @commission_discount,
	t.success_service       = t.success_service + @success_face * @service_discount,
	t.success_fee           = t.success_fee + @num * @sell_price * @fee_discount,
	t.success_cost_amount   = t.success_cost_amount + @cost_amount,
	t.success_up_commission = t.success_up_commission + @up_commission,
	t.success_up_service    = t.success_up_service + @up_service,
	t.profit                = t.profit + @num * @sell_price - @success_face * @commission_discount 
							  + @success_face * @service_discount - @num * @sell_price * @fee_discount 
							  - @cost_amount + @up_commission + @up_service
where t.order_id = @order_id
and t.order_status = 20
and t.delivery_bind_status = 30
and t.notify_status = 10
`

const SetAuditRecordSuccess = `update oms_audit_info t
set t.audit_status = 0,
	t.audit_by     = @audit_by,
	t.audit_time   = NOW(),
	t.audit_msg    = @audit_msg
where t.audit_id = @audit_id
and t.order_id = @order_id
and t.delivery_id = @delivery_id
and t.change_type = @change_type
and t.audit_status=20
`

const LockAuditOrder = `select t.order_id from oms_order_info t where t.order_id=@order_id for update`

const AuditDeliveryFail = `update oms_order_delivery t
set t.delivery_status   = 90,
	t.up_payment_status = 99,
	t.end_time          = NOW(),
	t.return_msg        = @msg
where t.delivery_id = @delivery_id
and t.order_id = @order_id
and t.delivery_status = 30
and t.up_payment_status = 10`

const SetAuditRecordFail = `update oms_audit_info t
set t.audit_status = 90,
	t.audit_by     = @audit_by,
	t.audit_time   = NOW(),
	t.audit_msg    = @audit_msg
where t.audit_id = @audit_id
and t.order_id = @order_id
and t.delivery_id = @delivery_id
and t.change_type = @change_type
and t.audit_status=20
`

const LockReturnOrder = `select t.order_id
from oms_order_info t
where t.order_id = @order_id
 for update`

const AuditReturnSuccess = `update oms_refund_up_return t
set t.return_status    = 0,
	t.up_refund_status = 20,
	t.end_time         = NOW(),
	t.return_msg       = @audit_msg
where t.return_id = @return_id
and t.refund_id = @refund_id
and t.order_id = @order_id
and t.return_status = 30
and t.up_refund_status = 10
`

const GetAuditRefundInfo = `select t.order_id,
t.refund_id,
t.refund_face,
t.refund_sell_amount,
t.refund_commission_amount,
t.refund_service_amount,
t.refund_fee_amount
from oms_refund_info t
where t.refund_id = @refund_id`

const GetAuditRetrunInfo = `select sum(r.return_total_face) up_return_face,
sum(r.return_cost_amount) refund_cost,
sum(r.return_commission_amount) refund_up_commission,
sum(r.return_service_amount) refund_up_service
from oms_refund_up_return r
where r.refund_id = @refund_id
and r.return_status = 0`

const UpdateRefundSuccess = `update oms_refund_info t
set t.refund_status        = 20,
	t.up_return_status     = 0,
	t.down_refund_status   = 20,
	t.refund_notify_status = 30
where t.refund_id = @refund_id
and t.order_id = @order_id
and t.refund_status = 10
and up_return_status = 30
and t.down_refund_status = 10
and t.refund_notify_status = 10
`

const UpdateOrderRefundSuccess = `update oms_order_info t
set t.order_status          = (case when t.success_face = @refund_face then 90 else 91 end),
	t.is_refund             = 0,
	t.success_face          = t.success_face - @refund_face,
	t.success_sell_amount   = t.success_sell_amount - @refund_sell_amount,
	t.success_commission    = t.success_commission - @refund_commission_amount,
	t.success_service       = t.success_service - @refund_service_amount,
	t.success_fee           = t.success_fee - @refund_fee_amount,
	t.success_cost_amount   = t.success_cost_amount - @refund_cost,
	t.success_up_commission = t.success_up_commission -
							  @refund_up_commission,
	t.success_up_service    = t.success_up_service - @refund_up_service,
	t.profit                = t.profit -
							  (@refund_sell_amount - @refund_commission_amount +
							  @refund_service_amount - @refund_fee_amount -
							  @refund_cost + @refund_up_commission +
							  @refund_up_service)
where t.order_id = @order_id
and t.order_status in (0, 91)
and t.success_face >= @refund_face`

const AuditReturnFail = `update oms_refund_up_return t
set t.return_status    = 90,
	t.up_refund_status = 99,
	t.end_time         = NOW(),
	t.return_msg       = @audit_msg
where t.return_id = @return_id
and t.refund_id = @refund_id
and t.order_id = @order_id
and t.return_status = 30
and t.up_refund_status = 10`

const GetFailRefundInfo = `select t.order_id, t.refund_face
from oms_refund_info t
where t.refund_id = @refund_id
`

const GetFailReturnInfo = `select sum(t.return_total_face) up_return_face
from oms_refund_up_return t
where t.refund_id = @refund_id
 and t.return_status = 90
`

const UpdateRefundFail = `update oms_refund_info t
set t.refund_status        = 90,
	t.up_return_status     = 90,
	t.down_refund_status   = 99,
	t.refund_notify_status = 30
where t.refund_id = @refund_id
and t.order_id = @order_id
and t.refund_status = 10
and t.up_return_status = 30
and t.down_refund_status = 10
and t.refund_notify_status = 10`

const GetOrderStatus = `select t.order_status, t.delivery_bind_status,((t.order_overtime-NOW())*24*60*60-300) overtime
from oms_order_info t
where t.order_id = @order_id
`
const GetOrderNotifyID = `SELECT 
n.notify_id 
FROM
oms_notify_info n 
WHERE n.order_id = @order_id 
AND n.notify_status = 10 
AND n.notify_type = @notify_type
&refund_id`

const StartOrderNotify = `UPDATE 
oms_notify_info n 
SET
n.notify_status = 20,
n.start_time = NOW() 
WHERE n.notify_id = @notify_id 
AND n.notify_status = 10`

const GetNotFailDeliveryCount = `select count(1)
from oms_order_delivery t
where t.order_id = @order_id
 and t.delivery_status != 90
`

const UpdateOrderDeliveryFail = `update oms_order_info t
set t.delivery_bind_status = 90,
	t.refund_status        = 20,
	t.notify_status        = 30,
	t.order_status         = 90
where t.order_id=@order_id 
and t.delivery_bind_status = 30
and t.refund_status = 10
and t.notify_status = 10
and t.order_status  = 20
`
const GetNotEndDelivery = `select count(1)
from oms_order_delivery t
where t.order_id = @order_id
 and t.delivery_status not in (0, 90)`

const UpdateOrderPartialSuccess = `update oms_order_info t
 set t.order_status         = 91,
	 t.delivery_bind_status = 91,
	 t.refund_status        = 20,
	 t.notify_status        = 20
 where t.order_id = @order_id
 and t.order_status = 20
 and t.delivery_bind_status = 30
 and t.refund_status = 10
 and t.notify_status = 10
 and t.total_face > t.success_face
 and t.success_face > 0
 `

const GetPartialRetrunInfo = `
SELECT MAX(t.order_id) order_id,
SUM(CASE WHEN t.return_status=0 THEN 0 WHEN t.return_status=90 THEN 0 ELSE 1 END) not_end_count,
SUM(CASE WHEN t.up_refund_status=0 THEN 0 WHEN t.up_refund_status=90 THEN 0 ELSE 1 END) not_refund_count,
SUM(CASE WHEN t.return_status =0 THEN t.return_total_face ELSE 0 END) refund_face,
SUM(CASE WHEN t.return_status =0 THEN t.return_num * p.sell_price ELSE 0 END) refund_sell_amount,
SUM(CASE WHEN t.return_status =0 THEN t.return_total_face*p.commission_discount ELSE 0 END) refund_commission,
SUM(CASE WHEN t.return_status =0 THEN t.return_total_face*p.service_discount ELSE 0 END) refund_service,
SUM(CASE WHEN t.return_status =0 THEN t.return_num * p.sell_price * p.payment_fee_discount ELSE 0 END) refund_fee,
SUM(CASE WHEN t.return_status =0 THEN t.return_cost_amount ELSE 0 END) refund_cost,
SUM(CASE WHEN t.return_status =0 THEN t.return_commission_amount ELSE 0 END) refund_up_commission,
SUM(CASE WHEN t.return_status =0 THEN t.return_service_amount ELSE 0 END) refund_up_service
FROM oms_refund_up_return t
INNER JOIN oms_down_product p
ON t.down_product_id = p.product_id
WHERE t.refund_id = @refund_id`

const UpdateReturnPartialSuccess = `update oms_refund_info t
set t.refund_status        = 20,
	t.up_return_status     = 91,
	t.down_refund_status   = 20,
	t.refund_notify_status = 30
where t.refund_id = @refund_id
and t.refund_status = 10
and t.up_return_status = 30
and t.down_refund_status = 10
and t.refund_notify_status = 10
and t.refund_face > @success_retrun_face
and @success_retrun_face > 0
 `
const UpdateReturnPartialOrder = `update oms_order_info t
 set t.order_status          = 91,
	 t.is_refund             = 0,
	 t.success_face          = t.success_face - @refund_face,
	 t.success_sell_amount   = t.success_sell_amount - @refund_sell_amount,
	 t.success_commission    = t.success_commission - @refund_commission,
	 t.success_service       = t.success_service - @refund_service,
	 t.success_fee           = t.success_fee - @refund_fee,
	 t.success_cost_amount   = t.success_cost_amount - @refund_cost,
	 t.success_up_commission = t.success_up_commission -
							   @refund_up_commission,
	 t.success_up_service    = t.success_up_service - @refund_up_service,
	 t.profit                = t.profit -
							   (@refund_sell_amount - @refund_commission +
							   @refund_service - @refund_fee -
							   @refund_cost + @refund_up_commission +
							   @refund_up_service)
 where t.order_id = @order_id
 and t.order_status in (0, 91)`

const GetFalseSuccDeliveryInfo = `
SELECT 
  t.delivery_id,
  t.order_id,
  t.up_channel_no
FROM
  oms_order_delivery t 
WHERE t.delivery_id = @delivery_id 
  AND t.delivery_status = 0 
  AND t.up_payment_status = 0 `

const LockFalseSuccOrder = `select t.order_id
from oms_order_info t
where t.order_id = @order_id
 and t.order_status in (0, 91)
 for update
`

const CheckFalseSuccReturn = `select count(1)
from oms_refund_up_return t
where t.delivery_id = @delivery_id
and t.delivery_status!=90
`

const CreateFalseSuccReturn = `insert into oms_refund_up_return
(up_channel_no,
 up_product_id,
 order_id,
 delivery_id,
 refund_id,
 down_channel_no,
 down_product_id,
 line_id,
 carrier_no,
 province_no,
 city_no,
 return_status,
 up_refund_status,
 create_time,
 return_face,
 return_num,
 return_total_face,
 return_cost_amount,
 return_commission_amount,
 return_service_amount,
 start_time,
 end_time,
 return_msg)
select t.up_channel_no,
	   t.up_product_id,
	   t.order_id,
	   t.delivery_id,
	   0,
	   t.down_channel_no,
	   t.down_product_id,
	   t.line_id,
	   t.carrier_no,
	   t.province_no,
	   t.city_no,
	   0,
	   20,
	   NOW(),
	   t.face,
	   t.num,
	   t.total_face,
	   t.cost_amount,
	   t.up_commission_amount,
	   t.service_amount,
	   NOW(),
	   NOW(),
	   '假成功退货'
  from oms_order_delivery t
 where t.delivery_id = @delivery_id
 and t.delivery_status=0
 and t.up_payment_status=0`
