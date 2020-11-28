// +build !oracle

package sql

const SqlCreateManual = `
insert into oms_audit_info
	(order_id,
	refund_id,
	delivery_id,
	change_type,
	audit_status)
	values
	(@order_id,
	@refund_id,
	@return_id,
	2,
	20)
`
