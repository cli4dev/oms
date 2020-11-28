// +build !oracle

package sql

// SQLCreateNotifyInfo 创建通知
const SQLCreateNotifyInfo = `
insert into oms_notify_info
	(
	order_id,
	refund_id,
	notify_type,
	notify_status,
	max_count,
	notify_url)
	values
	(
	@order_id,
	@refund_id,
	2,
	10,
	10,
	@notify_url)
`

// SQLGetRefundNotifyInfo 查询通知信息
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
