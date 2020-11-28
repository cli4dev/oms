package sql

//GetOmsNotifyInfo 查询单条数据订单通知表
const GetOmsNotifyInfo = `
select
t.notify_id,
t.create_time,
t.end_time,
t.max_count,
t.notify_count,
t.notify_msg,
t.notify_status,
t.notify_type,
t.notify_url,
t.order_id,
t.refund_id,
t.start_time
from oms_notify_info t
where
&notify_id`

//QueryOmsNotifyInfoCount 获取订单通知表列表条数
const QueryOmsNotifyInfoCount = `
select count(1)
  from oms_notify_info t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.notify_status &t.notify_type &t.order_id &t.notify_id`

//QueryOmsNotifyInfo 查询订单通知表列表数据
const QueryOmsNotifyInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.notify_id,
                               t.create_time,
                               t.end_time,
                               t.max_count,
                               t.notify_count,
                               t.notify_status,
                               t.notify_type,
                               t.notify_url,
                               t.order_id,
                               t.refund_id,
                               t.start_time
                          from oms_notify_info t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.notify_status &t.notify_type &t.order_id
                         &t.notify_id
                         order by t.notify_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteOmsNotifyInfo 删除订单通知表
const DeleteOmsNotifyInfo = `
delete from oms_notify_info 
where
&notify_id`
