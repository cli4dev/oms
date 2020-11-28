package sql

//GetOmsDownChannelDictionary  获取数据字典
const GetOmsDownChannelDictionary = `
select channel_no as value,channel_name as name 
from oms_down_channel 
where 1=1 `

//InsertOmsDownChannel 添加下游渠道
const InsertOmsDownChannel = `
insert into oms_down_channel
(
	channel_no,
	channel_name
)
values
(
	@channel_no,
	@channel_name
)`

//GetOmsDownChannel 查询单条数据下游渠道
const GetOmsDownChannel = `
select
t.channel_no,
t.channel_name,
t.create_time,
t.status
from oms_down_channel t
where
&channel_no`

//QueryOmsDownChannelCount 获取下游渠道列表条数
const QueryOmsDownChannelCount = `
select count(1)
from oms_down_channel t
INNER JOIN beanpay_account_info i ON t.channel_no = i.eid
where
i.groups = 'down_channel'
and t.channel_name  like  '%' || @channel_name || '%' 
&t.channel_no`

//QueryOmsDownChannel 查询下游渠道列表数据
const QueryOmsDownChannel = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.channel_no,
                               t.channel_name,
                               t.create_time,
                               t.status,
                               i.credit,
                               i.balance
                          from oms_down_channel t
                         INNER JOIN beanpay_account_info i ON t.channel_no =
                                                              i.eid
                         where i.groups = 'down_channel'
                           and t.channel_name like
                            '%' || @channel_name || '%'
                         &t.channel_no
                         order by t.channel_no desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateOmsDownChannel 更新下游渠道
const UpdateOmsDownChannel = `
update 
oms_down_channel 
set
	channel_name=@channel_name,
	status=@status
where
	&channel_no`

//DeleteOmsDownChannel 删除下游渠道
const DeleteOmsDownChannel = `
delete from oms_down_channel 
where
&channel_no`
