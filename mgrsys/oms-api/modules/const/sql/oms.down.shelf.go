package sql

//GetOmsDownShelfDictionary  获取数据字典
const GetOmsDownShelfDictionary = `
select shelf_id as value,shelf_name as name 
from oms_down_shelf 
where 1=1 `

const GetChannelShelfDictionary = `select shelf_id as value,shelf_name as name 
from oms_down_shelf 
where 1=1 &channel_no`

//InsertOmsDownShelf 添加下游货架
const InsertOmsDownShelf = `
insert into oms_down_shelf
(
	shelf_name,
	channel_no,
	order_overtime,
	refund_overtime
)
values
(
	@shelf_name,
	@channel_no,
	@order_overtime,
	@refund_overtime
)`

//GetOmsDownShelf 查询单条数据下游货架
const GetOmsDownShelf = `
select
t.shelf_id,
t.shelf_name,
t.channel_no,
t.order_overtime,
t.refund_overtime,
t.status,
t.create_time
from oms_down_shelf t
where
&shelf_id`

//QueryOmsDownShelfCount 获取下游货架列表条数
const QueryOmsDownShelfCount = `
select count(1)
  from oms_down_shelf t
 where t.shelf_name like  '%' || @shelf_name || '%' 
 &t.channel_no
 &t.status
`

//QueryOmsDownShelf 查询下游货架列表数据
const QueryOmsDownShelf = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.shelf_id,
                               t.shelf_name,
                               t.channel_no,
                               t.order_overtime,
                               t.refund_overtime,
                               t.status,
                               t.create_time
                          from oms_down_shelf t
                         where  t.shelf_name like  '%' || @shelf_name || '%' 
                         &t.channel_no &t.status
                         order by t.shelf_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateOmsDownShelf 更新下游货架
const UpdateOmsDownShelf = `
update 
oms_down_shelf 
set
	shelf_name=@shelf_name,
	channel_no=@channel_no,
	order_overtime=@order_overtime,
	refund_overtime=@refund_overtime,
	status=@status
where
	&shelf_id`

//DeleteOmsDownShelf 删除下游货架
const DeleteOmsDownShelf = `
delete from oms_down_shelf 
where
&shelf_id`
