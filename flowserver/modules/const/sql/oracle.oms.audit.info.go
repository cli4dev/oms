// +build oracle

package sql

const SqlCreateManual = `
insert into oms_audit_info
	(audit_id,
	order_id,
	refund_id,
	delivery_id,
	change_type,
	audit_status)
	values
	(seq_audit_info_id.nextval,
	@order_id,
	@refund_id,
	@return_id,
	2,
	20)
`
const GetNewID = ``
