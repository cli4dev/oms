// +build !oracle

package sql

// CheckOrderForDownPay 下游支付检查订单
const CheckOrderForDownPay = `
SELECT 
o.order_id,
o.request_no,
o.carrier_no,
o.province_no,
o.city_no,
round(o.total_face,5) total_face,
o.recharge_account,
o.sell_amount,
o.create_time,
o.line_id,
ROUND(o.sell_amount,5) sell_amount,
ROUND(o.commission_amount,5) commission_amount,
ROUND(o.service_amount,5) service_amount,
ROUND(o.fee_amount,5) fee_amount,
o.down_channel_no channel_no,
o.down_account_no account_no,
ROUND(UNIX_TIMESTAMP(o.order_overtime)-UNIX_TIMESTAMP(NOW()), 0) order_overtime,
ROUND(UNIX_TIMESTAMP(o.order_overtime)-UNIX_TIMESTAMP(NOW()), 0) - 300 flow_overtime 
FROM
oms_order_info o 
WHERE o.order_id = @order_id 
AND o.order_status = 10 
AND o.payment_status = 30 
AND o.delivery_bind_status = 10 
AND o.refund_status = 10 
AND o.notify_status = 10 
AND o.is_refund = 1 
AND o.complete_up_pay = 1
AND o.delivery_pause = 1 
AND o.order_overtime > NOW() 
`

// OrderDownPaySuccess 订单下游支付成功
const OrderDownPaySuccess = `
UPDATE 
oms_order_info o 
SET
o.payment_status = 0,
o.order_status = 20,
o.delivery_bind_status = 20 
WHERE o.order_id = @order_id 
  AND o.order_status = 10 
  AND o.payment_status = 30 
  AND o.delivery_bind_status = 10 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.complete_up_pay = 1
  AND o.is_refund = 1 
  AND o.delivery_pause = 1
  AND o.order_overtime > NOW()
`

// CheckOrderForUpPay 上游支付检查订单
const CheckOrderForUpPay = `
SELECT 
   o.down_account_no,
   o.total_face,
   round(o.total_face,5) down_order_face,
   o.request_no,
   o.recharge_account,
   o.line_id,
   o.face,
   o.create_time as order_time,
   round(o.success_sell_amount ,5) down_sell_amount,
   round(o.success_commission ,5) down_commission_amount,
   round(o.success_service ,5) down_service_amount,
   round(o.success_fee,5) down_fee_amount
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
  AND o.order_status in (0, 20)
  AND o.payment_status in (0,99) 
  AND o.delivery_bind_status in (0, 30)
  AND o.refund_status = 10 
  AND o.is_refund = 1 
  AND o.delivery_pause = 1 
  AND o.complete_up_pay = 1 
`

// DeliveryUpPaySuccess 订单发货上游支付成功
const DeliveryUpPaySuccess = `
UPDATE 
  oms_order_delivery d 
SET
  d.up_payment_status = 0 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status = 0 
  AND d.up_payment_status = 30
`

// OrderCompleteUpPay 完成上游支付
const OrderCompleteUpPay = `
UPDATE 
  oms_order_info o 
SET
  o.complete_up_pay = 0 
WHERE o.order_id = @order_id 
  AND o.order_status = 0 
  AND o.payment_status in (0,99) 
  AND o.delivery_bind_status = 0 
  AND o.refund_status = 10 
  AND o.delivery_pause = 1 
  AND o.complete_up_pay = 1
  AND o.is_refund = 1   
`

// CheckDeliveryForUpPay 上游支付检查发货
const CheckDeliveryForUpPay = `
SELECT 
d.order_id,
d.up_channel_no,
d.down_channel_no,
d.carrier_no,
d.province_no,
d.city_no,
d.face,
d.delivery_id,
round(d.cost_amount ,5) cost_amount,
round(d.up_commission_amount ,5) up_commission_amount,
round(d.service_amount ,5) up_service_amount,
d.total_face delivery_face,
round(d.total_face,5) up_face,
d.up_channel_no  
FROM
  oms_order_delivery d 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status = 0 
  AND d.up_payment_status = 30  
`

// QuerySuccessDeliveryFace 查询发货成功总面值
const QuerySuccessDeliveryFace = `
SELECT 
 SUM(d.total_face) delivery_total_face
FROM
  oms_order_delivery d 
WHERE d.order_id = @order_id
  AND d.delivery_status = 0 
  AND d.up_payment_status = 0
`
