package sql

//InsertOmsUpProduct 添加上游商品
const InsertOmsUpProduct = `
insert into oms_up_product
(
	can_refund,
	invoice_type,
	carrier_no,
	province_no,
	city_no,
	cost_discount,
	commission_discount,
	service_discount,
	ext_product_no,
	face,
	limit_count,
	line_id,
	shelf_id
)
values
(
	@can_refund,
	@invoice_type,
	@carrier_no,
	@province_no,
	@city_no,
	@cost_discount,
	@commission_discount,
	@service_discount,
	@ext_product_no,
	@face,
	@limit_count,
	@line_id,
	@shelf_id
	)`

//GetOmsUpProduct 查询单条数据上游商品
const GetOmsUpProduct = `
select
t.product_id,
t.can_refund,
t.invoice_type,
t.carrier_no,
t.province_no,
t.city_no,
t.cost_discount,
t.commission_discount,
t.service_discount,
t.create_time,
t.ext_product_no,
t.face,
t.limit_count,
t.line_id,
t.shelf_id,
t.status
from oms_up_product t
where
&product_id`

//QueryOmsUpProductCount 获取上游商品列表条数
const QueryOmsUpProductCount = `
select count(1)
  from oms_up_product t
 where 1=1 &t.can_refund &t.invoice_type &t.carrier_no &t.province_no
 &t.city_no &t.line_id &t.shelf_id &t.status`

//QueryOmsUpProduct 查询上游商品列表数据
const QueryOmsUpProduct = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.product_id,
                               t.can_refund,
                               t.invoice_type,
                               t.carrier_no,
                               t.province_no,
                               t.city_no,
                               round(t.cost_discount, 3) cost_discount,
                               round(t.commission_discount, 3) commission_discount,
                               round(t.service_discount, 3) service_discount,
                               t.create_time,
                               t.ext_product_no,
                               round(t.face, 3) face,
                               t.limit_count,
                               t.line_id,
                               t.shelf_id,
                               t.status
                          from oms_up_product t
                         where 1=1 &t.can_refund &t.invoice_type &t.carrier_no
                          &t.province_no &t.city_no &t.line_id
                          &t.shelf_id &t.status
                         order by t.product_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//UpdateOmsUpProduct 更新上游商品
const UpdateOmsUpProduct = `
update 
oms_up_product 
set
	can_refund=@can_refund,
	invoice_type=@invoice_type,
	carrier_no=@carrier_no,
	province_no=@province_no,
	city_no=@city_no,
	cost_discount=@cost_discount,
	commission_discount=@commission_discount,
	service_discount=@service_discount,
	ext_product_no=@ext_product_no,
	face=@face,
	limit_count=@limit_count,
	line_id=@line_id,
	shelf_id=@shelf_id,
	status=@status
where
	&product_id`

//DeleteOmsUpProduct 删除上游商品
const DeleteOmsUpProduct = `
delete from oms_up_product 
where
&product_id`
