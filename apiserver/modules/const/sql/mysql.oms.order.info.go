// +build !oracle

package sql

// SQLQueryOrderInfo 查询订单状态
const SQLQueryOrderInfo = `
SELECT a.order_id,
	a.down_channel_no AS channel_no,
	a.request_no,
	a.num,
	a.success_face,
	a.face,
	b.can_refund,
	b.status AS product_status,
	c.extend_info,
	c.refund_overtime,
	a.order_status,
	a.can_split_order,
	a.point_num
FROM oms_order_info a 
INNER JOIN oms_down_product b ON a.down_product_id = b.product_id
INNER JOIN oms_down_shelf c ON b.shelf_id = c.shelf_id
INNER JOIN oms_down_channel d ON d.channel_no = a.down_channel_no	
WHERE a.down_channel_no = @down_channel_no
  AND a.request_no = @request_no
  AND a.num >= @num
  AND a.order_status IN (0,91)
  AND b.status = 0
  AND c.status = 0
  AND d.status = 0
`

// SQLLockOrder 锁定单
const SQLLockOrder = `
select 
	  t.order_id,
	  CASE WHEN t.success_face - t.face * @num >= 0 THEN 'true' ELSE 'false' END can_refund
from oms_order_info t
where t.down_channel_no = @channel_no
  and t.request_no = @request_no
  and t.order_status IN (0,91)
for update
`

// SQLGetDeliveryDetailList 查询发货详情
const SQLGetDeliveryDetailList = `
SELECT c.delivery_id,c.num FROM (
	SELECT a.delivery_id,
		 (a.dnum - IFNULL(b.rnum, 0)) num
	  FROM (SELECT t.delivery_id, t.num AS dnum
			  FROM oms_order_delivery t
			 WHERE t.order_id = @order_id
			   AND t.delivery_status = 0
			   AND t.up_payment_status = 0) a
	  LEFT JOIN (SELECT t.delivery_id, IFNULL(SUM(t.return_num), 0) rnum
				   FROM oms_refund_up_return t
				  WHERE t.order_id = @order_id
					AND t.return_status = 0
			GROUP BY t.delivery_id) b 
	  ON a.delivery_id = b.delivery_id) c 
	WHERE c.num > 0
`

// SQLUpdateOrderInfo 更新订单
const SQLUpdateOrderInfo = `
update oms_order_info t
set t.success_face = t.success_face - @refund_face,
	t.success_sell_amount = t.success_sell_amount - @refund_sell_amount,
	t.success_commission = t.success_commission - @refund_commission_amount,
	t.success_service = t.success_service - @refund_service_amount,
	t.success_fee = t.success_fee - @refund_fee_amount,
	t.profit = t.profit - @refund_sell_amount + @refund_commission_amount - @refund_service_amount + @refund_fee_amount,
	t.is_refund = 0
where t.order_id = @order_id
  and t.order_status = 0
  and t.payment_status = 0
`
