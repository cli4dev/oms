// +build !oracle

package sql

const SqlUpdateRefundNotifyToWait = `
update oms_notify_info t
set t.notify_status = 20
where t.refund_id = @refund_id
  and t.notify_type = 2
  and t.notify_status = 10
`

const SqlUpdateNotifyToWait = `
update oms_notify_info t
set t.notify_status = 20
where t.refund_id = @refund_id
  and t.notify_status = 10
`
const SqlUpdateFailNotify = `
update oms_notify_info t
set t.notify_msg = @notify_msg,
    t.notify_status = 20
where t.notify_id = @notify_id
  and t.notify_status = 30
`

const SqlUpdateNotifyToSuccess = `
update oms_notify_info t
set t.notify_status = 0,
    t.end_time = now(),
    t.notify_msg = @notify_msg
where t.notify_id = @notify_id
  and t.notify_status = 30
`

const SqlGetNotifyCount = `
select 
    a.notify_count, 
    a.max_count 
from oms_notify_info a
where a.notify_id = @notify_id
  and a.notify_status = 20
`

const SqlGetWaitNotifyInfo = `
select 
    a.notify_url,
    a.notify_id,
    a.refund_id
from  oms_notify_info a
where a.notify_id = @notify_id
  and a.notify_status = 20
`
const SqlGetRequestNotifyInfo = `
select 
    a.notify_url,
    a.notify_id,
    a.refund_id
from  oms_notify_info a
where a.notify_id = @notify_id
`

const SqlUpdateNotifyToProcess = `
update oms_notify_info a set 
      a.notify_status = 30,
      a.notify_count = a.notify_count + 1,
      a.start_time = IFNULL(a.start_time,NOW())
where a.notify_id = @notify_id
  and a.notify_status = 20
  and a.notify_count <= a.max_count  
`

const SqlGetNotifyInfo = `
select 
    a.notify_id
from oms_notify_info a
where a.refund_id = @refund_id
  and a.notify_type = 2
  and a.notify_status = 10
`
