// +build !oracle

package sql

// CheckOrderForBind 绑定检查订单
const CheckOrderForBind = `
SELECT 
  p.can_refund
FROM
  oms_order_info o 
  INNER JOIN oms_down_product p 
    ON p.product_id = o.down_product_id 
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status = 0 
  AND o.delivery_bind_status IN (20, 30) 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.delivery_pause = 1 
  AND o.is_refund = 1 
  AND o.complete_up_pay = 1 
  AND o.order_overtime > NOW() 
  AND o.bind_face < o.total_face 
`

// LockOrderForBind 绑定时锁订单
const LockOrderForBind = `
SELECT 
  o.order_id,
  o.can_split_order,
  o.split_order_face,
  o.face,
  o.num,
  if(o.num = 1, o.num, (o.total_face - o.bind_face)/o.face) need_num,
  o.total_face,
  o.bind_face,
  o.carrier_no,
  o.province_no,
  o.city_no,
  o.invoice_type,
  o.extend_info,
  o.line_id,
  o.down_product_id,
  round(UNIX_TIMESTAMP(o.order_overtime) - UNIX_TIMESTAMP(NOW()), 0) - 300 flow_overtime
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
  AND o.order_status = 20 
  AND o.payment_status in (0,99) 
  AND o.delivery_bind_status IN (20, 30) 
  AND o.refund_status = 10 
  AND o.notify_status = 10 
  AND o.delivery_pause = 1 
  AND o.is_refund = 1 
  AND o.complete_up_pay = 1 
  AND o.order_overtime > NOW() 
  AND o.bind_face < o.total_face 
  FOR UPDATE 
`

// CreateOrderDelivery 创建订单发货记录
const CreateOrderDelivery = `
INSERT INTO oms_order_delivery (
  up_channel_no,
  up_product_id,
  up_ext_product_no,
  order_id,
  down_channel_no,
  down_product_id,
  line_id,
  carrier_no,
  province_no,
  city_no,
  invoice_type,
  delivery_status,
  up_payment_status,
  create_time,
  face,
  num,
  total_face,
  cost_amount,
  up_commission_amount,
  service_amount,
  extend_info
) 
SELECT 
  s.channel_no,
  p.product_id,
  p.ext_product_no,
  o.order_id,
  o.down_channel_no,
  o.down_product_id,
  o.line_id,
  o.carrier_no,
  o.province_no,
  o.city_no,
  p.invoice_type,
  20,
  10,
  NOW(),
  p.face,
  @need_bind_num,
  p.face * @need_bind_num,
  round(@need_bind_num * p.cost_price, 5),
  round(p.face * @need_bind_num * p.commission_discount, 5),
  round(p.face * @need_bind_num * p.service_discount, 5),
  @extend_info
FROM
  oms_order_info o 
  INNER JOIN oms_up_product p 
    ON p.line_id = o.line_id 
    AND p.carrier_no = o.carrier_no 
    AND p.province_no = o.province_no 
    AND p.city_no = o.city_no 
  INNER JOIN oms_up_shelf s 
    ON p.shelf_id = s.shelf_id 
WHERE o.order_id = @order_id 
  AND p.product_id = @product_id 
`

// UpdateOrderForBind 绑定修改订单信息
const UpdateOrderForBind = `
UPDATE 
oms_order_info o 
SET
o.delivery_bind_status = 30,
o.bind_face = o.bind_face + @need_bind_face * @need_bind_num
WHERE o.order_id = @order_id 
AND o.bind_face + @need_bind_face * @need_bind_num <= o.total_face
`

// QueryUpProduct 查询上游商品
const QueryUpProduct = `
SELECT 
  p.product_id,
  p.face,
  s.delivery_overtime,
  p.limit_count,
  @need_bind_face need_bind_face
FROM
  oms_up_product p 
  INNER JOIN oms_up_shelf s 
    ON s.shelf_id = p.shelf_id 
  INNER JOIN oms_up_channel c 
    ON c.channel_no = s.channel_no 
  INNER JOIN oms_down_product d 
    ON d.line_id = p.line_id 
WHERE d.product_id = @product_id 
  AND c.status = 0 
  AND s.status = 0 
  AND p.status = 0 
  AND p.carrier_no = d.carrier_no 
  AND p.province_no = d.province_no 
  AND p.city_no = d.city_no 
  AND IF(
    d.can_split_order = 0,
    1 = 1,
    p.limit_count >= @need_num
  ) 
  AND IF(
    d.can_split_order = 1 
    OR @num > 1,
    p.face = @need_bind_face,
    p.face <= @need_bind_face
  ) 
  AND IF(
    p.can_refund = 1,
    1 = 1,
    p.can_refund = d.can_refund
  ) 
  AND IF(
    d.invoice_type = 2,
    d.invoice_type = p.invoice_type,
    1 = 1
  ) 
  AND p.product_id NOT IN 
  (SELECT 
    o.up_product_id 
  FROM
    oms_order_delivery o 
  WHERE o.order_id = @order_id 
    AND o.delivery_status = 90 
    AND o.end_time > DATE_ADD(NOW(), INTERVAL -1 MINUTE)) 
ORDER BY p.face DESC,
  p.cost_price ASC 
`

// QueryManyUpProduct 查询上游商品
const QueryManyUpProduct = `
SELECT 
  p.product_id,
  p.face,
  s.delivery_overtime,
  p.limit_count 
FROM
  oms_up_product p 
  INNER JOIN oms_up_shelf s 
    ON s.shelf_id = p.shelf_id 
  INNER JOIN oms_up_channel c 
    ON c.channel_no = s.channel_no 
WHERE p.line_id = @line_id 
  AND c.status = 0 
  AND s.status = 0 
  AND p.status = 0 
  AND p.carrier_no = @carrier_no 
  AND p.province_no = @province_no 
  AND p.city_no = @city_no 
  AND IF(@can_split_order=0,1=1,p.limit_count>=@need_num)
  AND IF(
    @can_split_order = 1 OR @num > 1,
    p.face = @need_bind_face,
    p.face <= @need_bind_face
  )
  AND IF(
    @can_refund = 1,
    1 = 1,
    p.can_refund = @can_refund
  ) 
  AND IF(
    @invoice_type = 2,
    @invoice_type = p.invoice_type,
    1 = 1
  ) 
ORDER BY p.face DESC, p.cost_price ASC 
`

// CheckCompleteBind 检查是否完成绑定
const CheckCompleteBind = `
SELECT 
  IF(o.total_face = o.bind_face, 1, 0) complete_bind
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
`
