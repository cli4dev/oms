package sql

//GetVdsOrderExp 查询单条数据发货异常订单记录表
const GetVdsOrderExp = `
select
t.id,
t.carrier_no,
t.channel_no,
t.coop_id,
t.coop_order_id,
t.create_time,
t.error_msg,
t.ext_params,
t.local_ip,
t.product_face,
t.product_num,
t.service_class,
t.user_ip
from vds_order_exp t
where
&id`

//QueryVdsOrderExpCount 获取发货异常订单记录表列表条数
const QueryVdsOrderExpCount = `
select count(1)
  from vds_order_exp t
 where 1=1 &t.carrier_no &t.channel_no &t.coop_id`

//QueryVdsOrderExp 查询发货异常订单记录表列表数据
const QueryVdsOrderExp = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.carrier_no,
                               t.channel_no,
                               t.coop_id,
                               t.coop_order_id,
                               t.create_time,
                               t.local_ip,
                               t.product_face,
                               t.product_num,
                               t.service_class,
                               t.user_ip
                          from vds_order_exp t
                         where 1=1 &t.carrier_no &t.channel_no &t.coop_id) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteVdsOrderExp 删除发货异常订单记录表
const DeleteVdsOrderExp = `
delete from vds_order_exp 
where
&id`
