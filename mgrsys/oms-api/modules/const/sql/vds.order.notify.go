package sql

//GetVdsOrderNotify 查询单条数据发货通知记录表
const GetVdsOrderNotify = `
select
t.id,
t.coop_id,
t.coop_order_id,
t.create_time,
t.finish_time,
t.notify_content,
t.notify_count,
t.notify_limit_count,
t.notify_url,
t.order_no,
t.result_msg,
t.status
from vds_order_notify t
where
&id`

//QueryVdsOrderNotifyCount 获取发货通知记录表列表条数
const QueryVdsOrderNotifyCount = `
select count(1)
  from vds_order_notify t
 where 1=1 &t.coop_id &t.coop_order_id &t.order_no &t.status
`

//QueryVdsOrderNotify 查询发货通知记录表列表数据
const QueryVdsOrderNotify = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.coop_id,
                               t.coop_order_id,
                               t.create_time,
                               t.finish_time,
                               t.notify_count,
                               t.notify_limit_count,
                               t.notify_url,
                               t.order_no,
                               t.status
                          from vds_order_notify t
                         where 1=1 &t.coop_id &t.coop_order_id &t.order_no
                          &t.status) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteVdsOrderNotify 删除发货通知记录表
const DeleteVdsOrderNotify = `
delete from vds_order_notify 
where
&id`
