package sql
//GetVdsOrderInfo 查询单条数据发货订单信息表
const GetVdsOrderInfo = `
select
t.order_no,
t.carrier_no,
t.channel_no,
t.coop_id,
t.coop_order_id,
t.create_time,
t.flow_timeout,
t.last_update_time,
t.notify_url,
t.product_face,
t.product_num,
t.request_finish_time,
t.request_params,
t.request_start_time,
t.result_code,
t.result_desc,
t.result_params,
t.result_source,
t.service_class,
t.status,
t.succ_face,
t.up_order_no
from vds_order_info t
where
&order_no`

//QueryVdsOrderInfoCount 获取发货订单信息表列表条数
const QueryVdsOrderInfoCount = `
select count(1)
  from vds_order_info t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.carrier_no &t.channel_no &t.coop_id &t.coop_order_id
 &t.result_source &t.service_class &t.status`

//QueryVdsOrderInfo 查询发货订单信息表列表数据
const QueryVdsOrderInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.order_no,
                               t.carrier_no,
                               t.channel_no,
                               t.coop_id,
                               t.coop_order_id,
                               t.create_time,
                               t.flow_timeout,
                               t.notify_url,
                               t.product_face,
                               t.product_num,
                               t.request_finish_time,
                               t.request_start_time,
                               t.result_params,
                               t.result_source,
                               t.service_class,
                               t.status,
                               t.succ_face,
                               t.up_order_no
                          from vds_order_info t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.carrier_no &t.channel_no &t.coop_id
                         &t.coop_order_id &t.result_source
                         &t.service_class &t.status) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`
//DeleteVdsOrderInfo 删除发货订单信息表
const DeleteVdsOrderInfo = `
delete from vds_order_info 
where
&order_no`

