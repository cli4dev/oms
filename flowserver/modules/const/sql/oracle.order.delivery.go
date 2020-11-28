// +build oracle

package sql

// UpdateDeliveryStatus 修改发货记录状态为正在发货
const UpdateDeliveryStatus = `
UPDATE 
  oms_order_delivery d 
SET
  d.delivery_status = 30,
  d.start_time = sysdate
WHERE d.delivery_id = @delivery_id
`

// DeliveryFailed 发货失败
const DeliveryFailed = `
UPDATE 
  oms_order_delivery d 
SET
  d.delivery_status = 90,
  d.end_time = nvl2(@up_delivery_no,sysdate,null),
  d.up_payment_status = 99,
  d.return_msg = @msg,
  d.extend_info = nvl(@extend_info,d.extend_info),
  d.up_delivery_no = nvl(@up_delivery_no,0)
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status in (20, 30) 
  AND d.up_payment_status = 10 
`

// UpdateOrderForDeliveryFailed 发货失败修改订单信息
const UpdateOrderForDeliveryFailed = `
UPDATE oms_order_info o SET
o.bind_face = o.bind_face - @delivery_face
WHERE o.order_id = @order_id
AND o.order_status = 20
AND o.payment_status = 0 
AND o.delivery_bind_status = 30
AND o.delivery_pause = 1
AND o.is_refund = 1
AND o.complete_up_pay = 1
`

// UpdateOrderForDeliverySuccess 发货成功
const UpdateOrderForDeliverySuccess = `
UPDATE 
  oms_order_info o 
SET
  o.order_status = decode(o.total_face, o.success_face + @delivery_face, 0, 20),
  o.delivery_bind_status = decode(o.total_face,o.success_face + @delivery_face, 0, 30),
  o.notify_status = decode(o.total_face, o.success_face + @delivery_face, decode(@notify_id,0,10,30), 10),
  o.order_overtime = decode(o.total_face, o.success_face + @delivery_face, to_date('2099-12-31 23:59:59','yyyy-mm-dd hh24:mi:ss'), o.order_overtime),
  o.success_face = o.success_face + @delivery_face,
  o.success_sell_amount = ROUND(@sell_price * (o.success_face + @delivery_face)/o.face, 5),
  o.success_commission = ROUND(@commission_discount * (o.success_face + @delivery_face), 5),
  o.success_service = ROUND(@service_discount * (o.success_face + @delivery_face), 5),
  o.success_fee = ROUND(@sell_price * (o.success_face + @delivery_face) / o.face * @payment_fee_discount, 5),
  o.success_cost_amount = o.success_cost_amount + @cost_amount,
  o.success_up_commission = o.success_up_commission + @up_commission_amount,
  o.success_up_service = o.success_up_service + @up_service_amount,
  o.profit = ROUND(@sell_price * (o.success_face + @delivery_face)/o.face, 5) 
  - ROUND(@commission_discount * (o.success_face + @delivery_face), 5)
  + ROUND(@service_discount * (o.success_face + @delivery_face), 5) 
  - ROUND(@sell_price * (o.success_face + @delivery_face)/o.face * @payment_fee_discount, 5)
  - (o.success_cost_amount + @cost_amount) 
  + (o.success_up_commission + @up_commission_amount) 
  + (o.success_up_service + @up_service_amount),
  o.extend_info = @extend_info
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status in (0,99) 
  AND o.delivery_bind_status = 30 
  AND o.delivery_pause = 1 
  AND o.is_refund = 1 
  AND o.complete_up_pay = 1
  AND o.refund_status = 10 
  AND o.notify_status = 10 
`

// DeliverySuccess 发货成功
const DeliverySuccess = `
UPDATE oms_order_delivery d SET 
d.delivery_status = 0,
d.up_payment_status = 30,
d.end_time = sysdate,
d.return_msg = @msg,
d.courier_cost_amount = decode(@courier_cost_amount,null,0,@courier_cost_amount),
d.up_delivery_no = @up_delivery_no,
d.extend_info = nvl(@extend_info,d.extend_info)
WHERE d.delivery_id = @delivery_id
AND d.delivery_status = 30
AND d.up_payment_status = 10
`

// UpdateNotifyStatus 修改通知状态
const UpdateNotifyStatus = `
UPDATE 
  oms_notify_info n 
SET
  n.notify_status = 20,
  n.start_time = sysdate 
WHERE n.notify_id = @notify_id 
  AND n.notify_status = 10 
  AND n.notify_type = 1 
`

// CheckOrderForDelivery 发货检查订单
const CheckOrderForDelivery = `
SELECT o.down_channel_no,
       o.bind_face,
       o.total_face,
       o.success_face,
       decode(round(o.total_face - o.success_face,5), round(@delivery_face,5),1,0) all_success,
       decode(round(o.bind_face,5), round(o.total_face,5), 1, 0) complete_bind,
       round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) order_overtime,
       round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) - 300 flow_overtime,
       p.sell_discount,
       p.sell_price,
       p.service_discount,
       p.commission_discount,
       p.payment_fee_discount,
       o.extend_info order_extend_info
  FROM oms_order_info o
 INNER JOIN oms_down_product p ON p.product_id  = o.down_product_id 
 WHERE o.order_id = @order_id
   AND o.order_status = 20
   AND o.payment_status in (0,99)
   AND o.delivery_bind_status = 30
   AND o.refund_status = 10
   AND o.notify_status = 10
   AND o.is_refund = 1
   AND o.complete_up_pay = 1
`

// LockOrderForDelivery 发货锁订单
const LockOrderForDelivery = `
SELECT 
o.recharge_account,
o.bind_face,
o.bind_face,
o.total_face,
o.success_face,
decode(o.bind_face, o.total_face, 1, 0) complete_bind,
round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) order_overtime,
round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) - 300 flow_overtime,
o.extend_info order_extend_info
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status in (0,99) 
  AND o.delivery_bind_status = 30 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.is_refund = 1 
  AND o.delivery_pause = 1 
  AND o.complete_up_pay = 1 
  AND o.order_overtime > sysdate
  FOR UPDATE
`

// CheckDeliveryStatus 检查订单发货状态
const CheckDeliveryStatusForWaiting = `
SELECT 
d.order_id,
s.delivery_overtime,
p.extend_info product_extend_info,
s.extend_info shelf_extend_info,
p.ext_product_no,
c.ext_channel_no,
d.delivery_id,
d.order_id,
d.up_channel_no,
d.down_channel_no,
d.num,
d.face,
d.province_no,
d.city_no,
d.carrier_no,
d.total_face delivery_face,
d.extend_info
FROM
  oms_order_delivery d 
  INNER JOIN oms_up_product p 
    ON p.product_id = d.up_product_id 
  INNER JOIN oms_up_shelf s 
    ON s.shelf_id = p.shelf_id 
  INNER JOIN oms_up_channel c 
    ON c.channel_no = d.up_channel_no 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status = 20 
  AND d.up_payment_status = 10 
`

// LockDelivery 锁发货
const LockDelivery = `
SELECT 
  d.delivery_id,
  d.order_id,
  d.up_channel_no,
  d.down_channel_no,
  d.num,
  d.face,
  d.province_no,
  d.city_no,
  d.carrier_no,
  d.cost_amount,
  d.up_commission_amount,
  d.service_amount,
  d.total_face delivery_face
FROM
  oms_order_delivery d 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status = 20 
  AND d.up_payment_status = 10 
  FOR UPDATE
`

// CheckDeliveryIn 检查订单发货中
const CheckDeliveryIn = `
SELECT 
  d.order_id,
  d.delivery_id,
  d.cost_amount,
  d.up_commission_amount,
  d.service_amount up_service_amount,
  d.total_face delivery_face,
  d.up_channel_no,
  d.up_product_id,
  d.delivery_status,
  d.extend_info delivery_extend_info,
  s.pre_tag
FROM
  oms_order_delivery d 
  INNER JOIN oms_up_product p 
    ON p.product_id = d.up_product_id 
  INNER JOIN oms_up_shelf s 
    ON s.shelf_id = p.shelf_id 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status in (20, 30)
  AND d.up_payment_status = 10  
`

// CheckNotifyExist 检查通知是否存在
const CheckNotifyExist = `
SELECT 
  n.notify_id 
FROM
  oms_notify_info n 
WHERE n.order_id = @order_id 
  AND n.notify_status = 10 
  AND n.notify_type = 1 
`

// QueryNoFailedDelivery 查询非失败发货
const QueryNoFailedDelivery = `
SELECT 
CASE WHEN @bind_face = 0 AND COUNT(1) = 0 THEN 1 ELSE 0 END bind
FROM
  oms_order_delivery d 
WHERE d.order_id = @order_id 
  AND d.delivery_status != 90 
`
