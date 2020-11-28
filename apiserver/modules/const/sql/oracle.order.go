// +build oracle

package sql

// CheckOrder 检查订单
const CheckOrder = `
SELECT 
  o.order_id,
  o.request_no,
  o.down_channel_no channel_no,
  o.down_account_no account_no,
  o.order_status,
  o.fail_code,
  o.fail_msg,
  o.down_product_id product_id,
  round(((o.order_overtime - sysdate) * 24 * 60 * 60), 0) - 300 flow_overtime 
FROM
  oms_order_info o 
WHERE o.request_no = @request_no 
  AND o.down_channel_no = @channel_no 
`

// QueryProduct 查询下游商品
const QueryProduct = `
SELECT 
  p.product_id,
  (p.sell_price  * @num- @point_num/100 - @amount) / (p.face * @num) diff,
  (p.sell_price  * @num- @point_num/100) sell_amount,
  p.face * @num total_face,
  p.split_order_face,
  p.can_split_order,
  p.face,
  s.order_overtime,
  (s.order_overtime-300) flow_overtime
FROM
  oms_down_channel c 
  INNER JOIN oms_down_shelf s 
    ON c.channel_no = s.channel_no 
  INNER JOIN oms_down_product p 
    ON p.shelf_id = s.shelf_id 
WHERE c.channel_no = @channel_no 
  AND p.line_id = @line_id 
  AND c.status = 0 
  AND s.status = 0 
  AND p.status = 0 
  AND p.limit_count >= @num 
  AND p.face = @face
  &p.can_split_order
  &p.province_no &p.city_no &p.carrier_no &p.invoice_type  
`

// LockProduct 锁产品
const LockProduct = `
SELECT 
  p.product_id
FROM
  oms_down_product p
WHERE p.product_id = @product_id
 FOR UPDATE
 `

// GetNewOrderID 获取新订单ID
const GetNewOrderID = `
select seq_order_info_id.nextval from dual 
`

// OrderCreate 创建订单
const OrderCreate = `
INSERT INTO oms_order_info (
	order_id,
	down_channel_no,
	down_account_no,
	request_no,
	down_shelf_id,
	down_product_id,
	ext_product_no,
	line_id,
	carrier_no,
	province_no,
	city_no,
	invoice_type,
	face,
	num,
	total_face,
	mer_amount,
	courier_amount,
	sell_amount,
	commission_amount,
	service_amount,
	fee_amount,
	can_split_order,
	split_order_face,
	order_overtime,
	delivery_pause,
	order_status,
	payment_status,
	delivery_bind_status,
	refund_status,
	notify_status,
	is_refund,
	recharge_account,
	point_num,
	pay_order_no,
	extend_info
  ) 
  SELECT 
	@id,
	@channel_no,
	@account_no,
	@request_no,
	p.shelf_id,
	p.product_id,
	p.ext_product_no,
	p.line_id,
	p.carrier_no,
	p.province_no,
	p.city_no,
	p.invoice_type,
	p.face,
	@num,
	p.face * @num,
	ROUND(@amount, 5),
	ROUND(@courier_amount,5),
	ROUND(@num * p.sell_price-@point_num/100,5),
	ROUND(p.face * @num * p.commission_discount, 5),
	ROUND(p.face * @num * p.service_discount, 5),
	ROUND((@num * p.sell_price-@point_num/100) * p.payment_fee_discount, 5),
	p.can_split_order,
	p.split_order_face,
	sysdate + s.order_overtime/(24*60*60),
	1,
	10,
	30,
	10,
	10,
	10,
	1,
	@recharge_account,
	@point_num,
	@pay_order_no,
	@extend_info
  FROM
	oms_down_product p 
	INNER JOIN oms_down_shelf s 
	  ON s.shelf_id = p.shelf_id 
  WHERE p.product_id = @product_id 
	AND p.status = 0 
	AND s.status = 0
`

// NotifyCreate 创建通知
const NotifyCreate = `
INSERT INTO oms_notify_info(
	notify_id,
	order_id,
	notify_type,
	notify_status,
	max_count,
	notify_url,
	create_time)
VALUES(
	seq_notify_info_id.nextval,
	@order_id,
	1,
	10,
	10,
	@notify_url,
	sysdate
)
`

// LockOrder 锁定单
const LockOrder = `
SELECT 
  COUNT(1) 
FROM
  oms_order_info o 
WHERE o.request_no = @request_no 
  AND o.down_channel_no = @channel_no 
  AND o.order_status = 0 
  AND o.is_refund = 1 
  AND o.payment_status = 0 
  AND o.delivery_bind_status = 0 
  AND o.refund_status = 10 
  AND o.notify_status = 30
  AND o.order_overtime > sysdate
  FOR UPDATE
`

// SuccessOrderNotify 查询成功订单通知
const SuccessOrderNotify = `
UPDATE 
  oms_order_info o 
SET
  o.notify_status = 100 
WHERE o.request_no = @request_no 
  AND o.down_channel_no = @channel_no  
  AND o.delivery_pause = 1 
  AND o.is_refund = 1 
  AND o.notify_status = 30
`

// QueryOrderInfo 查询订单信息
const QueryOrderInfo = `
SELECT 
  o.order_id,
  o.request_no,
  o.down_channel_no channel_no,
  o.down_account_no account_no,
  o.order_status,
  o.fail_code,
  o.fail_msg,
  o.down_product_id product_id,
  o.notify_status,
  o.point_num,
  r.buy_send_num,
  r.activity_send_num
FROM
oms_order_info o 
left join jf_fd_order_record r
  on o.order_id = r.order_id
 and r.request_type = 1
  WHERE o.request_no = @request_no 
  AND o.down_channel_no = @channel_no 
 `
const GetChannelAccountInfo = `SELECT 
COUNT(1) channel_account_count,
nvl(
  SUM(
	decode(
	  t.account_no, 
	  @down_account_no,
	  1,
	  0
	)
  ),
  0
) find_count 
FROM
oms_down_payment_account t 
WHERE t.down_channel_no = @down_channel_no`
