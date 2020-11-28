// +build !oracle

package sql

const SqlUpdateRefundAllSuccess = `
update oms_refund_info t
set t.refund_status = 20,  
    t.up_return_status = 0,
    t.down_refund_status = if(t.down_refund_status=10,20,99),
    t.extend_info = if(@extend_info is null or @extend_info ='',t.extend_info,@extend_info),
    t.refund_notify_status = if(@notify_id=0,10,30)
where t.refund_id = @refund_id 
  and t.refund_status = 10
  and t.down_refund_status in (10,99)
  and t.refund_notify_status = 10
`

const SqlGetAllFailReturnList = ` 
select
	   t.return_id
from oms_refund_up_return t
where t.refund_id = @refund_id
  and t.return_status != 90
`

const SqlUpdateRefundToFail = `
update oms_refund_info t
set t.refund_status = 90,
    t.up_return_status = 90,
    t.down_refund_status = 99,
    t.refund_notify_status = if(@notify_id = 0,99,30)
where t.refund_id = @refund_id
  and t.refund_status = 10
  and t.up_return_status in (20,30)
  and t.down_refund_status = 10
`

const SqlUpdateReturnToSuccess = `
update oms_refund_info t
set t.refund_status = 20,
    t.up_return_status = 0,
    t.down_refund_status = 20,
    t.refund_notify_status = 20
where t.refund_id = @refund_id
 and t.down_refund_status = 10
 and t.refund_status = 10
 and t.up_return_status = 30
 and t.refund_notify_status = 10
`

const SqlUpdateDownRefundToSuccess = `
update oms_refund_info t
set t.down_refund_status = 0,
    t.refund_status = if(@code = 0,0,t.refund_status),
    t.complete_up_refund = if(@code = 0,0,t.complete_up_refund),
    t.return_overtime = DATE_FORMAT('2099-01-01 00:00:00','%Y-%m-%d %H:%i:%s')
where t.refund_id = @refund_id
  and t.down_refund_status = 20
  and t.refund_status = 20
`

const SqlCheckRefundStatus = `
SELECT 
  a.refund_id,
  a.refund_notify_status,
  a.refund_face,
  a.order_id,
  a.refund_sell_amount,
  a.refund_commission_amount,
  a.refund_service_amount,
  a.refund_fee_amount,
  a.refund_type,
  a.refund_point_num,
  a.extend_info refund_extend_info,
  TIMESTAMPDIFF(SECOND,a.return_overtime,NOW()) overtime
FROM oms_refund_info a
WHERE a.refund_id = @refund_id
  AND a.up_return_status = 30
  AND a.down_refund_status in (10,99)
  AND a.refund_notify_status = 10
`

const SqlCheckDownRefundStatus = `
select 
  a.refund_id,
  a.order_id,
  a.down_channel_no,
  a.down_account_no,
  a.face,
  a.refund_face,
  ROUND(a.refund_face,5) refund_unit,
  a.create_time,
  round(a.refund_sell_amount ,5) sell_amount,
  round(a.refund_commission_amount ,5)  commission_amount,
  round(a.refund_service_amount ,5) service_amount,
  round(a.refund_fee_amount ,5) fee_amount
from oms_refund_info a
where a.refund_id = @refund_id
  and a.up_return_status in (0,99)
  and a.down_refund_status = 20
  and a.refund_status = 20
`

const SqlUpdateToFail = `
update oms_refund_info t
set t.refund_status = 90,
	  t.return_overtime = DATE_FORMAT('2099-01-01 00:00:00','%Y-%m-%d %H:%i:%s') 
where t.refund_id = @refund_id
`

const SqlUpdateReturnOvertime = `
update oms_refund_info t
set t.return_overtime = DATE_FORMAT('2099-01-01 00:00:00','%Y-%m-%d %H:%i:%s')
where t.refund_id = @refund_id
  and t.refund_status not in (0,90,91)
  and t.return_overtime < DATE_FORMAT('2099-01-01 00:00:00','%Y-%m-%d %H:%i:%s')
`

const SqlCheckDownRefundFinish = `
select 
	a.refund_id
from oms_refund_info a
where a.refund_id = @refund_id 
  and a.down_refund_status = 0
`

const SqlGetTimeOutRefundInfo = `
select 
  a.refund_id,
  a.refund_status,
  a.refund_type
from oms_refund_info a
where a.refund_id = @refund_id 
  and a.refund_status not in (0,90,91)
`

const SqlChangeRefundReturnStart = `
update oms_refund_info a
set a.up_return_status = 30
where a.refund_id = @refund_id
  and a.up_return_status in (20,30)
`

const SqlUpdateDownRefundToProcess = `
update oms_refund_info t
set   t.down_refund_status = 30
where t.refund_id = @refund_id
  and t.down_refund_status = 20
`

const SqlUpdateRefundNotifyToSuccess = `
update oms_refund_info t
set   t.refund_notify_status = if(t.refund_notify_status = 100,t.refund_notify_status,0)
where t.refund_id = @refund_id
  and t.refund_notify_status in (30,100)
`

const SqlCloseRefund = `
update oms_refund_info a
set a.refund_status = 0,
    a.complete_up_refund = 0,
    a.return_overtime = DATE_FORMAT('2099-01-01 00:00:00','%Y-%m-%d %H:%i:%s')
where a.refund_id = @refund_id 
and a.refund_status = 20
`

const SqlGetRefundInfo = `
select b.refund_id,
       b.order_id,
       @notify_id notify_id,
       b.down_channel_no as channel_no,
       b.down_account_no as account_no,
       b.down_refund_no as refund_no,
       b.request_no,
       b.refund_num,
       case b.refund_status
         when 10 THEN
          30
         when 20 THEN
          0
         else
          b.refund_status
       end as status,
       b.fail_code,
       b.fail_msg,
       r.point_num,
       r.buy_send_num,
       r.activity_send_num
  from oms_refund_info b
  left join jf_fd_order_record r
    on b.refund_id = r.refund_id
   and b.order_id = r.order_id
   and r.request_type = 3
 where b.refund_id = @refund_id
   and b.refund_notify_status = 30
`

const SqlLockCurrentRefund = `
select t.refund_id
  from oms_refund_info t
 where t.refund_id = @refund_id
   and t.refund_status = 20
   for update
`

const SqlQueryRefundInfo = `
select t.refund_id,
      t.down_channel_no
  from oms_refund_info t
 where t.refund_id = @refund_id
`
const GetOrderJfInfo = `select t.point_num,
r.fd_id,
r.buy_send_num,
r.activity_send_num,
r.substitute_send_num
from oms_order_info t
left join jf_fd_order_record r
on t.order_id = r.order_id
and r.request_type = 1
where t.order_id = @order_id
`
