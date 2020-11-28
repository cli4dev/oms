package sql

const QueryReportProfitList = `select TAB1.*
from (select L.*
        from (select rownum as rn, R.*
                from (SELECT to_char(t.create_time, 'yyyy-mm-dd') create_date,
                             t.down_channel_no,
                             d.channel_name,
                             COUNT(1) order_count,
                             SUM(t.total_face) total_face,
                             SUM(t.success_face) success_face,
                             SUM(t.success_sell_amount) success_sell,
                             SUM(t.success_cost_amount) success_cost,
                             SUM(t.success_fee) success_fee,
                             SUM(t.success_sell_amount -
                                 t.success_commission + t.success_service -
                                 t.success_cost_amount +
                                 t.success_up_commission +
                                 t.success_up_service) profit
                        FROM oms_order_info t
                        left join oms_down_channel d on d.channel_no =
                                                        t.down_channel_no
                       WHERE t.create_time >=
                             to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                         AND t.create_time <
                             to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         AND t.order_status IN (0, 90, 91)
                       &t.carrier_no &t.line_id &t.down_channel_no
                       &t.total_face
                       GROUP BY to_date(t.create_time, 'yyyy-mm-dd'),
                                t.down_channel_no,
                                d.channel_name) R
               where rownum <= @pi * @ps) L
       where L.rn > (@pi - 1) * @ps) TAB1

`
const QueryReportProfitCount = `SELECT count(1)
FROM (SELECT to_char(t.create_time, 'yyyy-mm-dd') create_date,
             t.down_channel_no,
             d.channel_name,
             COUNT(1) order_count,
             SUM(t.total_face) total_face,
             SUM(t.success_face) success_face,
             SUM(t.success_sell_amount) success_sell,
             SUM(t.success_cost_amount) success_cost,
             SUM(t.success_fee) success_fee,
             SUM(t.success_sell_amount - t.success_commission +
                 t.success_service - t.success_cost_amount +
                 t.success_up_commission + t.success_up_service) profit
        FROM oms_order_info t
        left join oms_down_channel d on d.channel_no = t.down_channel_no
       WHERE t.create_time >=
             to_date(@start_time, 'yyyy-mm-dd hh24:mi:ss')
         AND t.create_time < to_date(@end_time, 'yyyy-mm-dd hh24:mi:ss')
         AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
       &t.down_channel_no &t.total_face
       GROUP BY to_date(t.create_time, 'yyyy-mm-dd'),
                t.down_channel_no,
                d.channel_name) o`

const Reportprofit4Export = `SELECT to_char(t.create_time, 'yyyy-mm-dd') create_date,
t.down_channel_no,
d.channel_name,
COUNT(1) order_count,
SUM(t.total_face) total_face,
SUM(t.success_face) success_face,
SUM(t.success_sell_amount) success_sell,
SUM(t.success_cost_amount) success_cost,
SUM(t.success_fee) success_fee,
SUM(t.success_sell_amount - t.success_commission + t.success_service -
    t.success_cost_amount + t.success_up_commission +
    t.success_up_service) profit
FROM oms_order_info t
left join oms_down_channel d on d.channel_no = t.down_channel_no
WHERE t.create_time >= to_date(@start_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.create_time < to_date(@end_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
&t.down_channel_no &t.total_face
GROUP BY to_date(t.create_time, 'yyyy-mm-dd'), t.down_channel_no`

const QueryReportProfitTotal = `SELECT to_char(t.create_time, 'yyyy-mm-dd') create_date,
        COUNT(1) order_count,
        SUM(t.total_face) total_face,
        SUM(t.success_face) success_face,
        SUM(t.success_sell_amount) success_sell,
        SUM(t.success_cost_amount) success_cost,
        SUM(t.success_fee) success_fee,
        SUM(t.success_sell_amount - t.success_commission +
            t.success_service - t.success_cost_amount +
            t.success_up_commission + t.success_up_service) profit
   FROM oms_order_info t
   WHERE t.create_time >= to_date(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_date(@end_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
   &t.down_channel_no &t.total_face
  GROUP BY to_date(t.create_time, 'yyyy-mm-dd')`
