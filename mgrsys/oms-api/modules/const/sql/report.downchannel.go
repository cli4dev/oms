package sql

const QueryReportDownchannelList = `select TAB1.*
from (select L.*
		from (select rownum as rn, R.*
				from (SELECT t.down_channel_no,
							 d.channel_name,
							 COUNT(1) order_count,
							 SUM(IF(order_status != 0, 1, 0)) fail_count,
							 SUM(IF(order_status != 0, t.total_face, 0)) fail_face,
							 SUM(IF(order_status = 0, 1, 0)) success_count,
							 SUM(IF(order_status = 0, t.total_face, 0)) success_face,
							 SUM(IF(order_status = 0, 1, 0)) / COUNT(1) success_ratio,
							 SUM(t.success_sell_amount) success_sell,
							 sum(t.total_face) total_face,
							 SUM(t.success_cost_amount) success_cost,
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
					   &t.province_no &t.total_face
					   GROUP BY t.down_channel_no, d.channel_name) R
			   where rownum <= @pi * @ps) L
	   where L.rn > (@pi - 1) * @ps) TAB1

`
const QueryReportDownchannelCount = `SELECT count(1)
FROM (SELECT t.down_channel_no,
			 COUNT(1) order_count,
			 SUM(IF(order_status != 0, 1, 0)) fail_count,
			 SUM(IF(order_status != 0, t.total_face, 0)) fail_face,
			 SUM(IF(order_status = 0, 1, 0)) success_count,
			 SUM(IF(order_status = 0, t.total_face, 0)) success_face,
			 SUM(IF(order_status = 0, 1, 0)) / COUNT(1) success_ratio,
			 SUM(t.success_sell_amount) success_sell,
			 SUM(t.success_cost_amount) success_cost,
			 SUM(t.success_sell_amount - t.success_commission +
				 t.success_service - t.success_cost_amount +
				 t.success_up_commission + t.success_up_service) profit
		FROM oms_order_info t
	   WHERE t.create_time >=
			 to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
		 AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
		 AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
	   &t.down_channel_no &t.province_no &t.total_face
	   GROUP BY t.down_channel_no) O
`

const ReportDownchannel4Export = `SELECT t.down_channel_no,
d.channel_name,
COUNT(1) order_count,
SUM(IF(order_status != 0, 1, 0)) fail_count,
SUM(IF(order_status != 0, t.total_face, 0)) fail_face,
SUM(IF(order_status = 0, 1, 0)) success_count,
SUM(IF(order_status = 0, t.total_face, 0)) success_face,
SUM(IF(order_status = 0, 1, 0)) / COUNT(1) success_ratio,
SUM(t.success_sell_amount) success_sell,
sum(t.total_face) total_face,
SUM(t.success_cost_amount) success_cost,
SUM(t.success_sell_amount - t.success_commission + t.success_service -
	t.success_cost_amount + t.success_up_commission +
	t.success_up_service) profit
FROM oms_order_info t
left join oms_down_channel d on d.channel_no = t.down_channel_no
WHERE t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
&t.down_channel_no &t.province_no &t.total_face
GROUP BY t.down_channel_no, d.channel_name `

const QueryReportDownchannelTotal = `SELECT 
		to_char(t.create_time, 'yyyy-mm-dd') create_time,
		COUNT(1) order_count,
		SUM(IF(order_status != 0, 1, 0)) fail_count,
		SUM(IF(order_status != 0, t.total_face, 0)) fail_face,
		SUM(IF(order_status = 0, 1, 0)) success_count,
		SUM(IF(order_status = 0, t.total_face, 0)) success_face,
		SUM(IF(order_status = 0, 1, 0)) / COUNT(1) success_ratio,
		SUM(t.success_sell_amount) success_sell,
		sum(t.total_face) total_face,
		SUM(t.success_cost_amount) success_cost,
		SUM(t.success_sell_amount - t.success_commission +
			t.success_service - t.success_cost_amount +
			t.success_up_commission + t.success_up_service) profit
   FROM oms_order_info t
   WHERE t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.order_status IN (0, 90, 91) &t.carrier_no &t.line_id
   &t.down_channel_no &t.province_no &t.total_face
  GROUP BY to_char(t.create_time, 'yyyy-mm-dd')
`
