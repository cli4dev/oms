package sql

//GetOmsUpShelfDictionary  获取数据字典
const GetOmsUpShelfDictionary = `
select shelf_id as value,shelf_name as name 
from oms_up_shelf 
where 1=1 `

//InsertOmsUpShelf 添加上游货架
const InsertOmsUpShelf = `
insert into oms_up_shelf
(
	shelf_name,
	channel_no,
	delivery_overtime,
	return_overtime
)
values
(
	@shelf_name,
	@channel_no,
	@delivery_overtime,
	@return_overtime
)`

//GetOmsUpShelf 查询单条数据上游货架
const GetOmsUpShelf = `
select
t.shelf_id,
t.shelf_name,
t.channel_no,
t.create_time,
t.delivery_overtime,
t.return_overtime,
t.status
from oms_up_shelf t
where
&shelf_id`

//QueryOmsUpShelfCount 获取上游货架列表条数
const QueryOmsUpShelfCount = `
select count(1)
  from oms_up_shelf t
 where t.shelf_name like '%' || @shelf_name || '%' &t.channel_no
 &t.status`

//QueryOmsUpShelf 查询上游货架列表数据
const QueryOmsUpShelf = `
select
t.shelf_id,
t.shelf_name,
t.channel_no,
t.create_time,
t.delivery_overtime,
t.return_overtime,
t.status 
from oms_up_shelf t
where t.shelf_name like '%' || @shelf_name || '%' &t.channel_no
 &t.status
order by t.shelf_id desc
`

//UpdateOmsUpShelf 更新上游货架
const UpdateOmsUpShelf = `
update 
oms_up_shelf 
set
	shelf_name=@shelf_name,
	channel_no=@channel_no,
	delivery_overtime=@delivery_overtime,
	return_overtime=@return_overtime,
	status=@status
where
	&shelf_id`

//DeleteOmsUpShelf 删除上游货架
const DeleteOmsUpShelf = `
delete from oms_up_shelf 
where
&shelf_id`
