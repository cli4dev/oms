package sql

//GetOmsOrderDelivery 查询单条数据订单发货表
const GetOmsOrderDelivery = `
select
t.delivery_id,
t.carrier_no,
t.city_no,
t.cost_amount,
t.create_time,
t.delivery_status,
t.down_channel_no,
t.down_product_id,
t.end_time,
t.face,
t.invoice_type,
t.line_id,
t.num,
t.order_id,
t.province_no,
t.return_msg,
t.service_amount,
t.start_time,
t.total_face,
t.up_channel_no,
t.up_commission_amount,
t.up_delivery_no,
t.up_ext_product_no,
t.up_payment_status,
t.up_product_id
from oms_order_delivery t
where
&delivery_id`

//QueryOmsOrderDeliveryCount 获取订单发货表列表条数
const QueryOmsOrderDeliveryCount = `
select count(1)
  from oms_order_delivery t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.delivery_id &t.carrier_no &t.city_no &t.delivery_status
 &t.down_channel_no &t.invoice_type &t.line_id &t.province_no
 &t.up_payment_status`

//QueryOmsOrderDelivery 查询订单发货表列表数据
const QueryOmsOrderDelivery = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.delivery_id,
                               t.carrier_no,
                               t.city_no,
                               t.cost_amount,
                               t.create_time,
                               t.delivery_status,
                               t.down_channel_no,
                               t.down_product_id,
                               t.end_time,
                               t.face,
                               t.invoice_type,
                               t.line_id,
                               t.num,
                               t.order_id,
                               t.province_no,
                               t.return_msg,
                               t.service_amount,
                               t.start_time,
                               t.total_face,
                               t.up_channel_no,
                               t.up_commission_amount,
                               t.up_delivery_no,
                               t.up_ext_product_no,
                               t.up_payment_status,
                               t.up_product_id
                          from oms_order_delivery t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.delivery_id &t.carrier_no &t.city_no
                         &t.delivery_status &t.down_channel_no
                         &t.invoice_type &t.line_id &t.province_no
                         &t.up_payment_status
                         order by t.delivery_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//DeleteOmsOrderDelivery 删除订单发货表
const DeleteOmsOrderDelivery = `
delete from oms_order_delivery 
where
&delivery_id`
