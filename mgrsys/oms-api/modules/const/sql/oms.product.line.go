package sql

//GetOmsProductLineDictionary  获取数据字典
const GetOmsProductLineDictionary = `
select line_id as value,line_name as name 
from oms_product_line 
where 1=1 `

//InsertOmsProductLine 添加产品线
const InsertOmsProductLine = `
insert into oms_product_line
(
	line_name
)
values
(
	@line_name
)`

//GetOmsProductLine 查询单条数据产品线
const GetOmsProductLine = `
select
t.line_id,
t.line_name,
t.bind_queue,
t.delivery_queue,
t.delivery_unknown_queue,
t.notify_queue,
t.order_overtime_queue,
t.order_refund_queue,
t.payment_queue,
t.refund_notify_queue,
t.refund_overtime_queue,
t.refund_queue,
t.return_finish_queue,
t.return_queue,
t.return_unknown_queue,
t.up_payment_queue,
t.up_refund_queue
from oms_product_line t
where
&line_id`

//QueryOmsProductLineCount 获取产品线列表条数
const QueryOmsProductLineCount = `
select count(1)
from oms_product_line t
where t.line_name like  '%' || @line_name || '%' `

//QueryOmsProductLine 查询产品线列表数据
const QueryOmsProductLine = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.*
                          from oms_product_line t
                         where t.line_name like  '%' || @line_name || '%'
                         order by t.line_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//UpdateOmsProductLine 更新产品线
const UpdateOmsProductLine = `
update 
oms_product_line 
set
	line_name=@line_name,
	bind_queue=@bind_queue,
	delivery_queue=@delivery_queue,
	delivery_unknown_queue=@delivery_unknown_queue,
	notify_queue=@notify_queue,
	order_overtime_queue=@order_overtime_queue,
	order_refund_queue=@order_refund_queue,
	payment_queue=@payment_queue,
	refund_notify_queue=@refund_notify_queue,
	refund_overtime_queue=@refund_overtime_queue,
	refund_queue=@refund_queue,
	return_queue=@return_queue,
	return_unknown_queue=@return_unknown_queue,
	up_payment_queue=@up_payment_queue,
	up_refund_queue=@up_refund_queue
where
	&line_id`

//DeleteOmsProductLine 删除产品线
const DeleteOmsProductLine = `
delete from oms_product_line 
where
&line_id`
