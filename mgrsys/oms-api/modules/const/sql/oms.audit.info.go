package sql

//GetOmsAuditInfo 查询单条数据发货人工审核表
const GetOmsAuditInfo = `
select
t.audit_id,
t.audit_by,
t.audit_msg,
t.audit_status,
t.audit_time,
t.change_type,
t.create_time,
t.delivery_id,
t.order_id,
t.refund_id
from oms_audit_info t
where
&audit_id`

//QueryOmsAuditInfoCount 获取发货人工审核表列表条数
const QueryOmsAuditInfoCount = `
select count(1)
  from oms_audit_info t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.change_type &t.delivery_id &t.order_id &t.refund_id`

//QueryOmsAuditInfo 查询发货人工审核表列表数据
const QueryOmsAuditInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.audit_id,
                               t.audit_by,
                               t.audit_msg,
                               t.audit_status,
                               t.audit_time,
                               t.change_type,
                               t.create_time,
                               t.delivery_id,
                               t.order_id,
                               t.refund_id
                          from oms_audit_info t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.change_type &t.delivery_id &t.order_id
                         &t.refund_id
                         order by t.audit_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateOmsAuditInfo 更新发货人工审核表
const UpdateOmsAuditInfo = `
update 
oms_audit_info 
set
	audit_by=@audit_by,
	audit_msg=@audit_msg,
	audit_status=@audit_status,
	audit_time=@audit_time
where
	&audit_id`

//DeleteOmsAuditInfo 删除发货人工审核表
const DeleteOmsAuditInfo = `
delete from oms_audit_info 
where
&audit_id`
