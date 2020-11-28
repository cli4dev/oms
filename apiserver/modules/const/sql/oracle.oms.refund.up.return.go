// +build oracle

package sql

const SQLCreateUpReturn = `
insert into oms_refund_up_return
	   (return_id,
		up_channel_no,
		up_product_id,
		order_id,
		refund_id,
		delivery_id,
		down_channel_no,
		down_product_id,
		line_id,
		carrier_no,
		province_no,
		city_no,
		return_status,
		up_refund_status,
		return_face,
		return_num,
		return_total_face,
		return_cost_amount,
		return_commission_amount,
		return_service_amount,
		extend_info
		)
		(select
		seq_refund_up_return_id.nextval,	 
		a.up_channel_no,
		a.up_product_id,
		a.order_id,
		@refund_id,
		a.delivery_id,
		a.down_channel_no,
		a.down_product_id,
		a.line_id,
		a.carrier_no,
		a.province_no,
		a.city_no,
		20,
		10,
		a.face,
		@ref_num,
		a.face * @ref_num,
		round(@ref_num * b.cost_price,5),
		round(a.face * @ref_num * b.commission_discount,5),
		round(a.face * @ref_num * b.service_discount, 5),
		@extend_info
		from oms_order_delivery a
		inner join oms_up_product b on a.up_product_id = b.product_id
		where a.delivery_id=@delivery_id
		  and a.delivery_status = 0)	
`

const SQLGetCreatedReturnList = `
SELECT 
  a.return_id,
  s.return_overtime 
FROM
  oms_refund_up_return a 
  INNER JOIN oms_up_product p 
    ON a.up_product_id = p.product_id 
  INNER JOIN oms_up_shelf s 
    ON s.shelf_id = p.shelf_id 
WHERE a.refund_id = @refund_id 
  AND a.return_status = 20 
  AND a.up_refund_status = 10 
`
