// +build !oracle

package sql

const SqlCheckReturnInfo = `
SELECT 
  a.return_id,
  a.refund_id,
  a.up_ext_product_no AS product_no,
  a.return_status,
  a.up_refund_status,
  a.down_channel_no,
  a.up_channel_no,
  a.carrier_no,
  ROUND(a.return_face, 5) return_face,
  a.return_num,
  c.return_overtime,
  a.extend_info,
  d.channel_no,
  d.ext_channel_no,
  c.pre_tag
FROM
  oms_refund_up_return a 
  INNER JOIN oms_up_product b 
    ON a.up_product_id = b.product_id 
  INNER JOIN oms_up_shelf c 
    ON b.shelf_id = c.shelf_id 
  INNER JOIN oms_up_channel d 
    ON a.up_channel_no = d.channel_no 
WHERE a.return_id = @return_id 
  AND a.return_status = 20 
  AND a.up_refund_status = 10 
`

const SqlCheckRefundInfo = `
select 
	a.refund_id
from oms_refund_info a
where a.refund_id = @refund_id
  and a.return_overtime >= now()
  and a.refund_status = 10
  and a.down_refund_status in (10,99)	
  and a.up_return_status in (20,30)
`

const SqlChangeReturnToProgress = `	
update oms_refund_up_return t
set t.return_status = 30,	
	  t.start_time = now()
where t.return_id = @return_id
and t.return_status = 20
`

const SqlCheckReturnStatus = `
select 
  t.return_id,
  t.refund_id,
  t.return_cost_amount,
  t.return_commission_amount,
  t.return_service_amount,
  t.up_channel_no,
  t.up_product_id,
  t.extend_info return_extend_info,
  s.pre_tag
from oms_refund_up_return t
inner join oms_up_product p on t.up_product_id=p.product_id
inner join oms_up_shelf s on p.shelf_id=s.shelf_id
where t.return_id = @return_id
  and t.return_status = 30
  and t.up_refund_status = 10
`

const SqlLockReturn = `
SELECT 
  a.return_id,
  a.up_channel_no,
  a.up_ext_product_no,
  a.up_return_no,
  a.up_product_id,
  a.refund_id,
  a.return_total_face,
  a.extend_info return_extend_info 
FROM
  oms_refund_up_return a 
WHERE a.return_id = @return_id 
  AND a.return_status = 30 
  AND a.up_refund_status = 10 
for update 
`

const GetRefundNotifyInfo = `select 
a.notify_id
from oms_notify_info a
where a.refund_id = @refund_id
and a.notify_type = 2
and a.notify_status = 10`

const SqlLockRefund = `
select 
     t.refund_id,
     t.down_channel_no,
     t.order_id,
     t.extend_info refund_extend_info,
     t.refund_type,
     t.refund_sell_amount,
     t.refund_commission_amount,
     t.refund_service_amount,
     t.refund_fee_amount,
     t.refund_point_num
from oms_refund_info t
where t.refund_id = @refund_id
  and t.refund_status = 10
  and t.up_return_status = 30
  and t.down_refund_status in (10,99)
  and t.refund_notify_status = 10
for update
`

const SqlUpdateCurrentReturnToSuccess = `
update oms_refund_up_return t
set t.return_status = 0,
    t.up_refund_status = 20,
    t.end_time = now(),
    t.return_msg = @return_msg,
    t.extend_info = @extend_info,
    t.courier_refund_amount = if(@courier_refund_amount='',0,@courier_refund_amount)
where t.return_id = @return_id
  and t.return_status = 30
  and t.up_refund_status = 10	
`

const SqlUpdateCurrentReturnToFail = `
update oms_refund_up_return t
set t.return_status = if(@code = 1,90,t.return_status),
    t.return_msg = @return_msg,
    t.up_refund_status = if(@code = 1,99,t.up_refund_status),
    t.end_time = if(@code = 1,now(),t.end_time)
where t.return_id = @return_id
  and t.return_status = 30
  and t.up_refund_status = 10
`

const SqlUpdateUpRefundToNoneed = `
update oms_refund_up_return t
set t.up_refund_status = 99
where t.return_id = @return_id
  and t.up_refund_status = 10
`

const SqlGetAllFailUpRefund = `
select 
	count(0)
from oms_refund_up_return t
where t.refund_id = @refund_id
  and t.up_refund_status != 0 
`

const SqlCheckUpRefundStatus = `
SELECT 
  a.return_id,
  a.refund_id,
  a.delivery_id,
  a.up_channel_no,
  a.down_channel_no,
  r.down_account_no,
  a.order_id,
  a.carrier_no,
  a.province_no,
  a.line_id,
  a.return_face,
  ROUND(a.return_total_face, 5) up_refund_face,
  a.return_total_face,
  a.create_time,
  ROUND(a.return_cost_amount, 5) return_cost_amount,
  ROUND(a.return_commission_amount, 5) return_commission_amount,
  ROUND(a.return_service_amount, 5) return_service_amount 
FROM
  oms_refund_up_return a 
  INNER JOIN oms_refund_info r 
    ON a.refund_id = r.refund_id 
from oms_refund_up_return a
where a.return_id = @return_id
  and a.up_refund_status = 20
  and a.return_status in (0,99)
`
const SqlUpdateUpRefundToSuccess = `
update oms_refund_up_return t
set t.up_refund_status = 0	
where t.return_id = @return_id
  and t.up_refund_status = 20
`

const UpdateWaitReturnToFail = `
update oms_refund_up_return a
set	a.return_status = 90,
    a.up_refund_status = 99
where a.refund_id = @refund_id
  and a.return_status = 20
`

const SqlUpdateReturnToProcess = `
update oms_refund_up_return t
set t.up_refund_status = 30
where t.return_id = @return_id
  and t.up_refund_status = 20
`

const SqlGetProcessReturnList = `
select 
    t.return_id,
    t.return_status,
    t.refund_id,
    t.order_id
from oms_refund_up_return t
where t.refund_id = @refund_id
  and t.return_status = 30
  and t.up_refund_status = 10 
`

const SqlGetReturnInfo = `
select 
    t.return_id,
    t.up_channel_no,
    t.up_ext_product_no,  
    t.up_return_no,
    t.refund_id 
from oms_refund_up_return t
where t.return_id = @return_id 
  and t.return_status = 30
  and t.up_refund_status = 10  
`

const SqlCheckUpRefundAllFinish = `
select 
     t.return_id 
from oms_refund_up_return t 
where t.refund_id = @refund_id
  and t.up_refund_status != 0
`

const SqlGetReturnSuccessAmount = ` 
select  
    if((sum(a.return_total_face) - @refund_face) =0,true,false) finish,
    sum(a.return_total_face) face,
    sum(a.return_cost_amount) return_cost_amount,
    sum(a.return_commission_amount) return_commission_amount,
    sum(a.return_service_amount) return_service_amount
from oms_refund_up_return a 
where a.refund_id = @refund_id
  and a.return_status = 0
`
