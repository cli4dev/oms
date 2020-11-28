// +build oracle

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
  AND o.order_overtime > sysdate 
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
  decode(o.num, 1, o.num, (o.total_face - o.bind_face)/o.face) need_num,
  o.total_face,
  o.bind_face,
  o.carrier_no,
  o.extend_info,
  o.province_no,
  o.city_no,
  o.down_product_id,
  o.invoice_type,
  o.line_id,
  round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) - 300 flow_overtime 
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
  AND o.order_overtime > sysdate 
  AND o.bind_face < o.total_face 
  FOR UPDATE 
`

// CreateOrderDelivery 创建订单发货记录
const CreateOrderDelivery = `
INSERT INTO oms_order_delivery (
  delivery_id,
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
  @id,
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
  sysdate,
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
    INNER JOIN oms_down_product d ON d.line_id = p.line_id
WHERE d.product_id = @product_id
   AND c.status = 0 
   AND s.status = 0 
   AND p.status = 0 
   AND p.carrier_no = d.carrier_no
   AND p.province_no = d.province_no
   AND p.city_no = d.city_no
   AND (d.can_split_order=0 OR p.limit_count>=@need_num)
   AND ((d.can_split_order = 1 OR @num > 1) AND p.face = @need_bind_face OR
       (d.can_split_order = 0 AND @num = 1 AND p.face <= @need_bind_face))
       AND ((p.can_refund = 1 AND d.can_refund = p.can_refund) OR (p.can_refund = 0))
       AND ((d.invoice_type = 2 AND d.invoice_type = p.invoice_type) OR (d.invoice_type != 2))
   AND p.product_id NOT IN 
  (SELECT 
    o.up_product_id 
  FROM
    oms_order_delivery o 
  WHERE o.order_id = @order_id 
    AND o.delivery_status = 90 
    AND o.end_time > SYSDATE -1/24/60)
 ORDER BY p.face DESC, p.cost_price ASC
`

// GetNewDeliveryID 获取新发货ID
const GetNewDeliveryID = `
select seq_order_delivery_id.nextval from dual
`

// CheckCompleteBind 检查是否完成绑定
const CheckCompleteBind = `
SELECT 
  DECODE(o.total_face, o.bind_face, 1, 0) complete_bind
FROM
  oms_order_info o 
WHERE o.order_id = @order_id 
`
