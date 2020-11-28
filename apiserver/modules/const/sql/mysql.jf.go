// +build !oracle

package sql

// JFOrderCreate 创建积分订单
const JFOrderCreate = `
INSERT INTO oms_order_info (
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
	extend_info
  ) 
  SELECT 
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
	ROUND(@num * p.sell_price,5),
	ROUND(p.face * @num * p.commission_discount, 5),
	ROUND(p.face * @num * p.service_discount, 5),
	ROUND(@num * p.sell_price * p.payment_fee_discount, 5),
	p.can_split_order,
	p.split_order_face,
	DATE_ADD(
	  NOW(),
	  INTERVAL s.order_overtime SECOND
	),
	1,
	20,
	99,
	20,
	10,
	10,
	1,
	ifnull(@recharge_account,'-'),
	@extend_info
  FROM
	oms_down_product p 
	INNER JOIN oms_down_shelf s 
	  ON s.shelf_id = p.shelf_id 
  WHERE p.product_id = @product_id 
	AND p.status = 0 
	AND s.status = 0
`

//CheckJFPreOrder 检查积分预下单
const CheckJFPreOrder = `select t.pre_order_id, t.down_channel_no, t.pre_request_no
from jf_pre_order_info t
where t.down_channel_no = @channel_no
 and t.pre_request_no = @pre_request_no
`

//CheckJFProduct 检查积分产品配置
const CheckJFProduct = `select t.channel_no,
s.shelf_id,
p.product_id,
p.face,
p.sell_discount,
p.sell_price
from oms_down_channel t
inner join oms_down_shelf s
on t.channel_no = s.channel_no
inner join oms_down_product p
on s.shelf_id = p.shelf_id
where t.channel_no = @channel_no
and p.line_id=@line_id
and p.line_id = 4
and p.face = 0.01
`

//LockJFProduct 锁积分产品
const LockJFProduct = `select t.product_id
from oms_down_product t
where t.product_id = @product_id
 for update`

//CreateJFPreOrder 创建积分预习单订单
const CreateJFPreOrder = `insert into jf_pre_order_info
(down_channel_no,
 pre_request_no,
 line_id,
 down_shelf_id,
 down_product_id,
 point_type,
 num,
 activate_num,
 create_time,
 user_no,
 over_time)
values
(@channel_no,
 @pre_request_no,
 @line_id,
 @down_shelf_id,
 @down_product_id,
 @point_type,
 @num,
 0,
 sysdate,
 @user_no,
 @over_time)
`

const GetNewJFPreOrderID = ``

const QueryJFPreOrderInfo = `select t.pre_order_id,
t.down_channel_no,
t.pre_request_no,
t.line_id,
t.down_shelf_id,
t.down_product_id,
t.point_type,
t.num,
t.activate_num,
t.create_time,
t.user_no,
t.over_time
from jf_pre_order_info t
where t.down_channel_no = @channel_no
and t.pre_request_no = @pre_request_no
and t.has_cancel=1
`
const CheckUseFdOrders = `select t.order_id, t.point_num
from oms_order_info t
where t.down_channel_no = @down_channel_no
 and t.pay_order_no = @channel_pay_no`

const CheckFdRequestRecord = `select t.request_id, t.down_channel_no, t.request_no
from jf_fd_request_info t
where t.down_channel_no = @down_channel_no
 and t.request_type = @request_type
 and t.request_no = @request_no
`
const CheckRelationNo = `select count(1)
from jf_fd_request_info t
where t.down_channel_no = @down_channel_no
 and t.request_type = @request_type
 and t.relation_no = @relation_no`

const GetNewFDRequestID = ``

const CreateJFFDRequestRecord = `insert into jf_fd_request_info
(down_channel_no,
 request_no,
 request_type,
 relation_no,
 create_time,
 buy_send_num,
 activity_send_num)
values
(@down_channel_no,
 @request_no,
 @request_type,
 @relation_no,
 now(),
 @buy_send_num,
 @activity_send_num)`

const GetNewFDOrderID = ``

const CreateJFFDOrderRecord = `insert into jf_fd_order_record
(down_channel_no,
 fd_channel_no,
 request_type,
 order_id,
 request_id,
 point_num,
 fd_status,
 create_time,
 buy_send_num,
 activity_send_num)
values
(@down_channel_no,
 @fd_channel_no,
 @request_type,
 @order_id,
 @request_id,
 @point_num,
 20,
 now(),
 @buy_send_num,
 @activity_send_num)
`
const GetJFUpChannel = `select s.channel_no up_channel_no, max(c.ext_channel_no) ext_channel_no
from oms_up_product t
inner join oms_up_shelf s
  on t.shelf_id = s.shelf_id
inner join oms_up_channel c
  on s.channel_no = c.channel_no
where t.line_id = 4
group by s.channel_no
`

const GetGainPointsInfo = `select t.order_id,t.num,t.extend_info
from oms_order_delivery t
where t.delivery_id = @gain_no
 and t.delivery_status = 0`

const GetHasVoidpointInfo = `select sum(t.buy_send_num) buy_send_num, sum(t.activity_send_num) activity_send_num
from jf_fd_request_info t
where t.down_channel_no = @down_channel_no
 and t.request_type = @request_type
 and t.relation_no = @gain_no
`

const CancelJFPreOrder = `update jf_pre_order_info t
set t.has_cancel = 0
where t.down_channel_no = @channel_no
and t.pre_request_no = @pre_request_no
and t.has_cancel = 1
`
const CancelPreOrderQuery = `select t.pre_order_id, t.user_no,date_format(t.create_time,'%Y-%m-%d %H:%mi:%s') create_time
from jf_pre_order_info t
where t.down_channel_no = @channel_no
 and t.pre_request_no = @pre_request_no
 and t.has_cancel = 1
`
const CancelOrderCheck = `select t.order_id from oms_order_info t
where t.create_time > STR_TO_DATE(@create_time, '%Y-%m-%d %H:%mi:%s')
and t.extend_info like '%#pre_order_id%'`
