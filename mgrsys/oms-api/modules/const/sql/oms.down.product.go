package sql

//InsertOmsDownProduct 添加下游商品
const InsertOmsDownProduct = `
insert into oms_down_product
(
	shelf_id,
	line_id,
	can_refund,
	invoice_type,
	can_split_order,
	ext_product_no,
	face,
	limit_count,
	carrier_no,
	province_no,
	city_no,
	payment_fee_discount,
	commission_discount,
	sell_discount,
	service_discount,
	split_order_face,
	status
)
values
(
	@shelf_id,
	@line_id,
	@can_refund,
	@invoice_type,
	@can_split_order,
	@ext_product_no,
	@face,
	@limit_count,
	@carrier_no,
	nvl(@province_no,"*"),
	nvl(@city_no,"*"),
	@payment_fee_discount,
	@commission_discount,
	@sell_discount,
	@service_discount,
	nvl(@split_order_face,"0"),
	@status
)`

//GetOmsDownProduct 查询单条数据下游商品
const GetOmsDownProduct = `
select
t.product_id,
t.shelf_id,
t.line_id,
t.can_refund,
t.invoice_type,
t.can_split_order,
t.create_time,
t.ext_product_no,
t.face,
t.limit_count,
t.carrier_no,
t.province_no,
t.city_no,
t.payment_fee_discount,
t.commission_discount,
t.sell_discount,
t.service_discount,
t.split_order_face,
t.status
from oms_down_product t
where
&product_id`

//QueryOmsDownProductCount 获取下游商品列表条数
const QueryOmsDownProductCount = `
select count(1)
  from oms_down_product t
 where 1=1 &t.shelf_id &t.line_id &t.can_refund &t.invoice_type
  &t.can_split_order &t.carrier_no &t.province_no &t.city_no &t.status
`

//QueryOmsDownProduct 查询下游商品列表数据
const QueryOmsDownProduct = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.product_id,
                               t.shelf_id,
                               t.line_id,
                               t.can_refund,
                               t.invoice_type,
                               t.can_split_order,
                               t.create_time,
                               t.ext_product_no,
                               round(t.face, 3) face,
                               t.limit_count,
                               t.carrier_no,
                               t.province_no,
                               t.city_no,
                               round(t.split_order_face, 3) split_order_face,
                               t.status,
                               round(t.sell_discount, 3) sell_discount,
                               round(t.commission_discount, 3) commission_discount,
                               round(t.service_discount, 3) service_discount,
                               round(t.payment_fee_discount, 3) payment_fee_discount
                          from oms_down_product t
                         where 1=1 &t.shelf_id &t.line_id &t.can_refund
                          &t.invoice_type &t.can_split_order
                          &t.carrier_no &t.province_no &t.city_no
                          &t.status
                         order by t.product_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateOmsDownProduct 更新下游商品
const UpdateOmsDownProduct = `
update 
oms_down_product 
set
	shelf_id=@shelf_id,
	line_id=@line_id,
	can_refund=@can_refund,
	invoice_type=@invoice_type,
	can_split_order=@can_split_order,
	ext_product_no=@ext_product_no,
	face=@face,
	limit_count=@limit_count,
	carrier_no=@carrier_no,
	province_no=@province_no,
	city_no=@city_no,
	payment_fee_discount=@payment_fee_discount,
	commission_discount=@commission_discount,
	sell_discount=@sell_discount,
	service_discount=@service_discount,
	split_order_face=@split_order_face,
	status=@status
where
	&product_id`

//DeleteOmsDownProduct 删除下游商品
const DeleteOmsDownProduct = `
delete from oms_down_product 
where
&product_id`
