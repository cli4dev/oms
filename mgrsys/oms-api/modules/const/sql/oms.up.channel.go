package sql

//GetOmsUpChannelDictionary  获取数据字典
const GetOmsUpChannelDictionary = `
select channel_no as value,channel_name as name 
from oms_up_channel 
where 1=1 `

const GetUpChannelShelfDictionary = `select shelf_id as value,shelf_name as name 
from oms_up_shelf 
where 1=1 &channel_no`

//InsertOmsUpChannel 添加上游渠道
const InsertOmsUpChannel = `
insert into oms_up_channel
(
	channel_no,
	ext_channel_no,
	channel_name
)
values
(
	@channel_no,
	@ext_channel_no,
	@channel_name
)`

//GetOmsUpChannel 查询单条数据上游渠道
const GetOmsUpChannel = `
select
t.channel_no,
t.channel_name,
t.create_time,
t.status
from oms_up_channel t
where
&channel_no`

//QueryOmsUpChannelCount 获取上游渠道列表条数
const QueryOmsUpChannelCount = `
select count(1)
from oms_up_channel t
INNER JOIN beanpay_account_info i ON t.channel_no = i.eid
where i.groups = 'up_channel'
and t.channel_name like  '%' || @channel_name || '%'
&t.channel_no
&t.status`

//QueryOmsUpChannel 查询上游渠道列表数据
const QueryOmsUpChannel = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.*,
                        i.ident,i.balance,i.credit
                   from oms_up_channel t
                  INNER JOIN beanpay_account_info i ON t.channel_no =
                                                       i.eid
                where i.groups = 'up_channel'
                and t.channel_name like  '%' || @channel_name || '%'
                &t.channel_no
                &t.status
                  order by t.channel_no desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//UpdateOmsUpChannel 更新上游渠道
const UpdateOmsUpChannel = `
update 
oms_up_channel 
set
	channel_name=@channel_name,
	status=@status,
	ext_channel_no=@ext_channel_no
where
	&channel_no`

//DeleteOmsUpChannel 删除上游渠道
const DeleteOmsUpChannel = `
delete from oms_up_channel 
where
&channel_no`
