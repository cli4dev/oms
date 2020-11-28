// +build oracle

package sql

const GetNewID = ``

const SQLCreateNotifyInfo = `
insert into oms_notify_info
	(notify_id, 
	order_id,
	refund_id,
	notify_type,
	notify_status,
	max_count,
	notify_url)
	values
	(seq_notify_info_id.nextval,
	@order_id,
	@refund_id,
	2,
	decode(@code,0,10,20),
	10,
	@notify_url)
`

const GetNewNotifyID = `select seq_notify_info_id.nextval from dual`

const SQLGetRefundNotifyInfo = `
select 
	t.notify_id,
	t.refund_id,
	t.notify_status
from oms_notify_info t
where t.order_id = @order_id
  and t.notify_type = 2
  and t.notify_status = 20
`

const SqlCloseNotify = `
update oms_notify_info t
set   t.notify_status = 0
where t.notify_id = @notify_id
`
