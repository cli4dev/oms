// +build !oracle

package sql

// QueryOrderForOvertime 超时订单信息查询
const QueryOrderForOvertime = `
SELECT 
  o.order_status,
  o.payment_status,
  o.delivery_bind_status,
  o.notify_status,
  o.bind_face,
  o.total_face,
  o.success_face,
  o.order_id,
  o.point_num
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
  AND o.order_status NOT IN (0, 90, 91) 
  AND o.order_overtime <= NOW()
`

// OrderOvertimeFailed 超时失败
const OrderOvertimeFailed = `
UPDATE 
  oms_order_info o 
SET
  o.order_status = 90,
  o.payment_status = 90,
  o.refund_status = 99,
  o.notify_status = 30,
  o.order_overtime = '2099-12-31 23:59:59' 
WHERE o.order_id = @order_id 
  AND o.payment_status = 30 
  AND o.order_status = 10 
`

// QueryDeliveryInfo 查询正在发货信息
const QueryDeliveryInfo = `
SELECT 
  d.delivery_id,
  d.order_id
FROM
  oms_order_delivery d 
WHERE d.order_id = @order_id 
  AND d.delivery_status in (20, 30)
`

// OrderRefund 订单退款
const OrderRefund = `
UPDATE 
  oms_order_info o 
SET
  o.order_status = 90,
  o.delivery_bind_status = 90,
  o.refund_status = 20,
  o.notify_status = 30,
  o.order_overtime = '2099-12-31 23:59:59'
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.bind_face = 0
  AND o.refund_status = 10 
  AND o.payment_status = 0 
  AND o.delivery_bind_status IN (20, 30) 
  AND o.notify_status = 10 
  AND o.delivery_pause = 1 
  AND o.is_refund = 1 
`

// CreateAuditForDelivery 创建人工审核对于发货
const CreateAuditForDelivery = `
INSERT INTO oms_audit_info (
  order_id,
  delivery_id,
  change_type,
  create_time,
  audit_status
) 
VALUES
  (@order_id, @delivery_id, 1, now(), 20)
`

// UpdateOrderOvertime 更新订单超时时间
const UpdateOrderOvertime = `
UPDATE 
  oms_order_info o 
SET
  o.order_overtime = '2099-12-31 23:59:59' 
WHERE o.order_id = @order_id 
`

// LockDeliveryIn 发货中锁发货记录
const LockDeliveryIn = `
SELECT 
d.delivery_id,
d.order_id,
d.up_channel_no,
d.down_channel_no,
d.num,
d.face,
d.province_no,
d.city_no,
d.up_product_id,
d.carrier_no,
d.cost_amount,
d.up_commission_amount,
d.service_amount up_service_amount,
  IF(@total_face - @success_face - d.total_face = 0,TRUE,FALSE) all_success,
  IF(@bind_face - d.total_face = 0, TRUE,FALSE) all_failed,
  d.total_face delivery_face,
  d.delivery_status,
  d.extend_info delivery_extend_info
FROM
  oms_order_delivery d 
WHERE d.delivery_id = @delivery_id 
  AND d.delivery_status = 30 
  AND d.up_payment_status = 10 
  FOR UPDATE
`

// CheckOrderForDeliveryOvertime 发货超时处理
const CheckOrderForDeliveryOvertime = `
SELECT 
  o.down_channel_no,
  o.bind_face,
  IF(o.bind_face = o.total_face, 1, 0) complete_bind,
  round(UNIX_TIMESTAMP(o.order_overtime) - UNIX_TIMESTAMP(NOW()), 0) - 300 overtime
FROM oms_order_info o
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status = 0 
  AND o.delivery_bind_status = 30 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.is_refund = 1 
  AND o.delivery_pause = 1 
  AND o.complete_up_pay = 1 
  AND o.order_overtime <= NOW()
`

// LockOrderForDeliveryOvertime 发货超时锁订单
const LockOrderForDeliveryOvertime = `
SELECT 
o.down_channel_no,
o.bind_face,
o.order_id,
o.total_face,
o.success_face,
o.down_product_id,
o.point_num
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status = 0 
  AND o.delivery_bind_status = 30 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.is_refund = 1 
  AND o.delivery_pause = 1 
  AND o.complete_up_pay = 1 
  AND o.order_overtime <= NOW()
  FOR UPDATE
`

// QueryDownProduct 查询下游产品信息
const QueryDownProduct = `
SELECT 
       p.sell_discount,
       p.sell_price,
       p.service_discount,
       p.commission_discount,
       p.payment_fee_discount
  FROM  oms_down_product p
 WHERE p.product_id = @product_id
`
