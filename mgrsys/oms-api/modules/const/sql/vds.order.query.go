package sql

//GetVdsOrderQuery 查询单条数据发货结果查询记录表
const GetVdsOrderQuery = `
select
t.order_no,
t.channel_no,
t.coop_id,
t.create_time,
t.last_query_time,
t.query_count,
t.query_result,
t.status
from vds_order_query t
where
&order_no`

//QueryVdsOrderQueryCount 获取发货结果查询记录表列表条数
const QueryVdsOrderQueryCount = `
select count(1)
from vds_order_query t
where
and &t.channel_no
&t.status`

//QueryVdsOrderQuery 查询发货结果查询记录表列表数据
const QueryVdsOrderQuery = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.order_no,
                               t.channel_no,
                               t.coop_id,
                               t.create_time,
                               t.last_query_time,
                               t.query_count,
                               t.query_result,
                               t.status
                          from vds_order_query t
                         where 1=1 &t.channel_no &t.status) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//DeleteVdsOrderQuery 删除发货结果查询记录表
const DeleteVdsOrderQuery = `
delete from vds_order_query 
where
&order_no`
