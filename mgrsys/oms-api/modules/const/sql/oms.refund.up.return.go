package sql

//GetOmsRefundUpReturn 查询单条数据上游退货信息表
const GetOmsRefundUpReturn = `
select
t.return_id,
t.carrier_no,
t.province_no,
t.city_no,
t.create_time,
t.delivery_id,
t.down_channel_no,
t.down_product_id,
t.end_time,
t.line_id,
t.order_id,
t.refund_id,
t.return_commission_amount,
t.return_cost_amount,
t.return_face,
t.return_msg,
t.return_num,
t.return_service_amount,
t.return_status,
t.return_total_face,
t.start_time,
t.up_channel_no,
t.up_ext_product_no,
t.up_product_id,
t.up_refund_status,
t.up_return_no
from oms_refund_up_return t
where
&return_id`

//QueryOmsRefundUpReturnCount 获取上游退货信息表列表条数
const QueryOmsRefundUpReturnCount = `
select count(1)
  from oms_refund_up_return t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.carrier_no &t.province_no &t.city_no &t.down_channel_no
 &t.up_refund_status
`

//QueryOmsRefundUpReturn 查询上游退货信息表列表数据
const QueryOmsRefundUpReturn = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.return_id,
                               t.carrier_no,
                               t.province_no,
                               t.city_no,
                               t.create_time,
                               t.delivery_id,
                               t.down_channel_no,
                               t.down_product_id,
                               t.end_time,
                               t.line_id,
                               t.order_id,
                               t.refund_id,
                               t.return_face,
                               t.return_num,
                               t.return_status,
                               t.return_total_face,
                               t.start_time,
                               t.up_channel_no,
                               t.up_ext_product_no,
                               t.up_product_id,
                               t.up_refund_status,
                               t.up_return_no
                          from oms_refund_up_return t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.carrier_no &t.province_no &t.city_no
                         &t.down_channel_no &t.up_refund_status
                         order by t.return_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//DeleteOmsRefundUpReturn 删除上游退货信息表
const DeleteOmsRefundUpReturn = `
delete from oms_refund_up_return 
where
&return_id`
