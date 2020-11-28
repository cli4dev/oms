package sql

//GetOmsRefundInfo 查询单条数据退款记录
const GetOmsRefundInfo = `
select
t.refund_id,
t.carrier_no,
t.province_no,
t.city_no,
t.create_time,
t.down_channel_no,
t.down_product_id,
t.down_refund_status,
t.down_shelf_id,
t.ext_product_no,
t.face,
t.line_id,
t.order_id,
t.refund_commission_amount,
t.refund_face,
t.refund_fee_amount,
t.refund_notify_status,
t.refund_num,
t.refund_sell_amount,
t.refund_service_amount,
t.refund_status,
t.refund_type,
t.request_no,
t.return_overtime,
t.complete_up_refund,
t.up_return_status
from oms_refund_info t
where
&refund_id`

//QueryOmsRefundInfoCount 获取退款记录列表条数
const QueryOmsRefundInfoCount = `
select count(1)
  from oms_refund_info t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.refund_id &t.carrier_no &t.province_no &t.city_no
 &t.down_channel_no &t.down_refund_status &t.down_shelf_id &t.line_id
 &t.order_id &t.refund_notify_status &t.refund_status &t.refund_type
 &t.up_return_status`

//QueryOmsRefundInfo 查询退款记录列表数据
const QueryOmsRefundInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.refund_id,
                               t.carrier_no,
                               t.province_no,
                               t.city_no,
                               t.create_time,
                               t.down_channel_no,
                               t.down_product_id,
                               t.down_refund_status,
                               t.down_shelf_id,
                               t.ext_product_no,
                               t.face,
                               t.line_id,
                               t.order_id,
                               t.refund_commission_amount,
                               t.refund_face,
                               t.refund_fee_amount,
                               t.refund_notify_status,
                               t.refund_num,
                               t.refund_sell_amount,
                               t.refund_service_amount,
                               t.refund_status,
                               t.refund_type,
                               t.request_no,
                               t.return_overtime,
                               t.complete_up_refund,
                               t.up_return_status
                          from oms_refund_info t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.refund_id &t.carrier_no &t.province_no
                         &t.city_no &t.down_channel_no
                         &t.down_refund_status &t.down_shelf_id
                         &t.line_id &t.order_id &t.refund_notify_status
                         &t.refund_status &t.refund_type
                         &t.up_return_status
                         order by t.refund_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//DeleteOmsRefundInfo 删除退款记录
const DeleteOmsRefundInfo = `
delete from oms_refund_info 
where
&refund_id`
