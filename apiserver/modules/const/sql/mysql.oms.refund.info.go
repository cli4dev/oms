// +build !oracle

package sql

// SQLQueryRefund 查询退款信息
const SQLQueryRefund = `
select  t.refund_id,
		t.down_refund_no as refund_no,
		t.down_channel_no as channel_no,
		t.down_account_no as account_no,
		t.request_no,
		t.refund_num,
		case t.refund_status
		when 10 THEN
			30
		when 20 THEN
			0
		else
		    t.refund_status
    end as status,
    t.fail_code,
    t.fail_msg
from oms_refund_info t
where t.down_channel_no=@down_channel_no
  and t.down_refund_no=@down_refund_no
`

// SQLQueryProcessRefund 查询正在退款
const SQLQueryProcessRefund = `
select t.refund_id
from oms_refund_info t
where t.order_id = @order_id 
  and t.refund_status in (10,20)
`

// SQLCreateRefund 创建退款
const SQLCreateRefund = `
insert into oms_refund_info
    (order_id,
    down_channel_no,
    down_account_no,
    request_no,
    down_refund_no,
    down_shelf_id,
    down_product_id,
    ext_product_no,
    line_id,
    carrier_no,
    province_no,
    city_no,
    refund_type,
    face,
    refund_num,
    refund_face,
    refund_mer_amount,
    refund_sell_amount,
    refund_commission_amount,
    refund_service_amount,
    refund_fee_amount,
    refund_status,
    up_return_status,
    down_refund_status,
    refund_notify_status,
    return_overtime,
    complete_up_refund,
    refund_point_num,
    extend_info
  )
  (select 
    a.order_id,
    @down_channel_no,
    a.down_account_no,
    @request_no,
    @down_refund_no,
    a.down_shelf_id,
    a.down_product_id,
    a.ext_product_no,
    a.line_id,
    a.carrier_no,
    a.province_no,
    a.city_no,
    @refund_type,
    a.face,
    @refund_num,
    a.face * @refund_num,
    @refund_mer_amount,
    ROUND((@refund_num * c.sell_price-@refund_point_num/100),5),
    ROUND(a.face * @refund_num * c.commission_discount,5),
    ROUND(a.face * @refund_num * c.service_discount,5),
    ROUND((@refund_num * c.sell_price-@refund_point_num/100) * c.payment_fee_discount,5),
    10,
    20,
    10,
    10,
    date_add(now(),interval b.refund_overtime second),
    1,
    @refund_point_num,
    @extend_info
  from oms_order_info a
  inner join oms_down_shelf b on a.down_shelf_id = b.shelf_id
  inner join oms_down_product c on a.down_product_id = c.product_id
  where order_id = @order_id)
`

// SQLCreateJFRefund 创建积分退款
const SQLCreateJFRefund = `
insert into oms_refund_info
		(order_id,
		down_channel_no,
		down_account_no,
		request_no,
		down_refund_no,
		down_shelf_id,
		down_product_id,
		ext_product_no,
		line_id,
		carrier_no,
		province_no,
		city_no,
		refund_type,
		face,
		refund_num,
		refund_face,
		refund_mer_amount,
		refund_sell_amount,
		refund_commission_amount,
		refund_service_amount,
		refund_fee_amount,
		refund_status,
		up_return_status,
		down_refund_status,
		refund_notify_status,
		return_overtime,
		complete_up_refund,
		extend_info
	)
	(select 
		a.order_id,
		@down_channel_no,
		a.down_account_no,
		@request_no,
		@down_refund_no,
		a.down_shelf_id,
		a.down_product_id,
		a.ext_product_no,
		a.line_id,
		a.carrier_no,
		a.province_no,
		a.city_no,
		@refund_type,
		a.face,
		@refund_num,
		a.face * @refund_num,
		@refund_mer_amount,
		ROUND(@refund_num * c.sell_price,5),
		ROUND(a.face * @refund_num * c.commission_discount,5),
		ROUND(a.face * @refund_num * c.service_discount,5),
		ROUND(@refund_num * c.sell_price * c.payment_fee_discount,5),
		10,
		20,
		99,
		10,
		date_add(now(),interval b.refund_overtime second),
		1,
		@extend_info
	from oms_order_info a
	inner join oms_down_shelf b on a.down_shelf_id = b.shelf_id
	inner join oms_down_product c on a.down_product_id = c.product_id
	where order_id = @order_id)
`

const QueryCreateRefund = `SELECT 
t.refund_id,
t.order_id,
t.down_channel_no channel_no,
t.down_account_no account_no,
t.down_refund_no refund_no,
t.request_no,
t.refund_num,
t.refund_notify_status,
ROUND(
  UNIX_TIMESTAMP(t.return_overtime) - UNIX_TIMESTAMP(NOW()),
  0
) - 300 refund_flow_overtime,
CASE
  t.refund_status 
  WHEN 10 
  THEN 30 
  WHEN 20 
  THEN 0 
  ELSE t.refund_status 
END AS status,
t.fail_code,
t.fail_msg,
t.extend_info 
FROM
oms_refund_info t 
WHERE t.refund_id = @refund_id `

// SQLQueryRefundInfo 查询退款
const SQLQueryRefundInfo = `
select t.refund_id,
       t.order_id,
       t.down_channel_no channel_no,
       t.down_account_no account_no,
       t.down_refund_no refund_no,
       t.request_no,
       t.refund_num,
       t.refund_notify_status,
       t.extend_info,
       t.refund_face,
       case t.refund_status
         when 10 THEN
          30
         when 20 THEN
          0
         else
          t.refund_status
       end as status,
       t.fail_code,
       t.fail_msg,
       r.point_num,
       r.buy_send_num,
       r.activity_send_num
  from oms_refund_info t
  left join jf_fd_order_record r
    on t.refund_id = r.refund_id
   and t.order_id = r.order_id
   and r.request_type = 3
 where t.down_channel_no = @down_channel_no
   and t.down_refund_no = @down_refund_no

`

// SQLUpdateRefundNotifyToQuerySuccess 查询成功修改通知状态
const SQLUpdateRefundNotifyToQuerySuccess = `
update oms_refund_info t
set t.refund_notify_status = 100
where t.refund_id = @refund_id
  and t.refund_notify_status in (20,30)
`

// SQLGetCreatedRefund 查询退款信息
const SQLGetCreatedRefund = `
select t.refund_face,
       t.refund_sell_amount,
       t.refund_commission_amount,
       t.refund_service_amount,
	   t.refund_fee_amount
  from oms_refund_info t
where t.refund_id = @refund_id
  and t.refund_status = 20
  and t.up_return_status = 99
  and t.down_refund_status = 20
  and t.refund_notify_status = 30
`
const QueryOrderExtendInfo = `select t.extend_info, d.delivery_id
from oms_order_info t
inner join oms_order_delivery d
  on t.order_id = d.order_id
where t.order_id = @order_id
 and d.delivery_status = 0
`
