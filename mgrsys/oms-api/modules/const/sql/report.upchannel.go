package sql

const QueryReportUpchannelList = `select TAB1.*
from (select L.*
		from (select rownum as rn, R.*
				from (SELECT t.up_channel_no,
							 u.channel_name,
							 COUNT(1) delivery_count,
							 SUM(IF(t.delivery_status = 0, 1, 0)) success_count,
							 SUM(t.total_face) total_face,
							 SUM(IF(t.delivery_status = 0, t.total_face, 0)) success_face,
							 (SUM(IF(t.delivery_status = 0, 1, 0)) /
							 COUNT(1)) success_ratio,
							 SUM(IF(t.delivery_status = 0, t.cost_amount, 0)) success_cost,
							 SUM(IF(t.delivery_status = 0,
									t.up_commission_amount,
									0)) success_up_commi
						FROM oms_order_delivery t
						left join oms_up_channel u on u.channel_no =
													  t.up_channel_no
					   WHERE t.create_time >=
							 to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
						 AND t.create_time <
							 to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
						 AND t.delivery_status IN (0, 90)
					   &t.carrier_no &t.line_id &t.up_channel_no
					   &t.province_no &t.total_face
					   GROUP BY t.up_channel_no, u.channel_name) R
			   where rownum <= @pi * @ps) L
	   where L.rn > (@pi - 1) * @ps) TAB1

`
const QueryReportUpchannelCount = `SELECT count(*)
FROM (SELECT t.up_channel_no,
			 u.channel_name,
			 COUNT(1) delivery_count,
			 SUM(IF(t.delivery_status = 0, 1, 0)) success_count,
			 SUM(t.total_face) total_face,
			 SUM(IF(t.delivery_status = 0, t.total_face, 0)) success_face,
			 (SUM(IF(t.delivery_status = 0, 1, 0)) / COUNT(1)) success_ratio,
			 SUM(IF(t.delivery_status = 0, t.cost_amount, 0)) success_cost,
			 SUM(IF(t.delivery_status = 0, t.up_commission_amount, 0)) success_up_commi
		FROM oms_order_delivery t
		left join oms_up_channel u on u.channel_no = t.up_channel_no
	   WHERE t.create_time >=
			 to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
		 AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
		 AND t.delivery_status IN (0, 90) &t.carrier_no &t.line_id
	   &t.up_channel_no &t.province_no &t.total_face
	   GROUP BY t.up_channel_no, u.channel_name) f
`

const ReportUpchannel4Export = `SELECT t.up_channel_no,
COUNT(1) delivery_count,
u.channel_name,
SUM(IF(t.delivery_status = 0, 1, 0)) success_count,
SUM(t.total_face) total_face,
SUM(IF(t.delivery_status = 0, t.total_face, 0)) success_face,
(SUM(IF(t.delivery_status = 0, 1, 0)) / COUNT(1)) success_ratio,
SUM(IF(t.delivery_status = 0, t.cost_amount, 0)) success_cost,
SUM(IF(t.delivery_status = 0, t.up_commission_amount, 0)) success_up_commi
FROM oms_order_delivery t
left join oms_up_channel u on u.channel_no = t.up_channel_no
WHERE t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.delivery_status IN (0, 90) &t.carrier_no &t.line_id
&t.up_channel_no &t.province_no &t.total_face
GROUP BY t.up_channel_no`

const QueryReportUpchannelTotal = `SELECT to_char(t.create_time, 'yyyy-mm-dd') create_time,
COUNT(1) delivery_count,
SUM(IF(t.delivery_status = 0, 1, 0)) success_count,
SUM(t.total_face) total_face,
SUM(IF(t.delivery_status = 0, t.total_face, 0)) success_face,
(SUM(IF(t.delivery_status = 0, 1, 0)) / COUNT(1)) success_ratio,
SUM(IF(t.delivery_status = 0, t.cost_amount, 0)) success_cost,
SUM(IF(t.delivery_status = 0, t.up_commission_amount, 0)) success_up_commi
FROM oms_order_delivery t
WHERE t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
AND t.delivery_status IN (0, 90) &t.carrier_no &t.line_id
&t.up_channel_no &t.province_no &t.total_face
GROUP BY to_char(t.create_time, 'yyyy-mm-dd')
`
