// +build oracle

package sql

const CheckRefundFDRecord = `select t.request_type, t.refund_id
from jf_fd_order_record t
where t.order_id = @order_id
and t.refund_id=decode(@refund_id,0,t.refund_id,@refund_id)
 and t.request_type in (2, 3)`

const QueryRefundFDOrder = `select t.order_id,t.point_num
from oms_order_info t
where t.order_id = @order_id
 and t.order_status in (90,91)
 and t.is_refund = @is_refund
`

const QueryRefundFDRefundRecord = `select t.refund_id,t.refund_point_num
from oms_refund_info t
where t.refund_id = @refund_id
 and t.order_id = @order_id
 and t.refund_status in (20, 0)
`
const GetJFUpChannel = `select s.channel_no
from oms_up_product t
inner join oms_up_shelf s
  on t.shelf_id = s.shelf_id
where t.line_id = 4
group by s.channel_no`

const LockRefundFDOrder = `select t.order_id
from oms_order_info t
where t.order_id = @order_id
 and t.order_status = 90
 for update
`
const GetNewRefundFDID = `select seq_jf_fd_id.nextval from dual`

const CreateRefundFDRecord = `insert into jf_fd_order_record
(fd_id,
 down_channel_no,
 fd_channel_no,
 request_type,
 order_id,
 refund_id,
 point_num,
 fd_status,
 create_time,
 buy_send_num,
 activity_send_num)
select @id,
	   t.down_channel_no,
	   @fd_channel_no,
	   @request_type,
	   @order_id,
	   @refund_id,
	   @refund_point_num,
	   20,
	   sysdate,
	   @refund_buy_num,
	   @refund_activity_num
  from jf_fd_order_record t
 where t.order_id = @order_id
   and t.request_type = 1
`

const QueryJFOrderFDRecord = ` select count(1) cn,
sum(decode(t.request_type, 1, 1, -1) * t.buy_send_num) last_buy_num,
sum(decode(t.request_type, 1, 1, -1) * t.activity_send_num) last_activity_num
from jf_fd_order_record t
where t.order_id = @order_id
and t.request_type in (1, 2, 3)
`

const RefundNotifyInfo = `select t.order_id, t.notify_id, t.refund_id
from oms_notify_info t
where t.refund_id = @refund_id
 and t.notify_type = 2
 and t.notify_status = 10
`
