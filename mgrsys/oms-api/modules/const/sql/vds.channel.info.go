package sql

//InsertVdsChannelInfo 添加渠道基本信息
const InsertVdsChannelInfo = `
insert into vds_channel_info
(
	channel_no,
	ext_params,
	first_query_time,
	notify_url,
	query_replenish_time,
	query_url,
	request_replenish_time,
	request_url,
	service_class
)
values
(
	@channel_no,
	@ext_params,
	@first_query_time,
	@notify_url,
	@query_replenish_time,
	@query_url,
	@request_replenish_time,
	@request_url,
	@service_class
)`

//GetVdsChannelInfo 查询单条数据渠道基本信息
const GetVdsChannelInfo = `
select
t.id,
t.can_query,
t.channel_no,
t.create_time,
t.ext_params,
t.first_query_time,
t.notify_url,
t.query_replenish_time,
t.query_url,
t.request_replenish_time,
t.request_url,
t.service_class,
t.status
from vds_channel_info t
where
&id`

//QueryVdsChannelInfoCount 获取渠道基本信息列表条数
const QueryVdsChannelInfoCount = `
select count(1)
  from vds_channel_info t
 where 1=1 &t.can_query &t.channel_no &t.service_class &t.status`

//QueryVdsChannelInfo 查询渠道基本信息列表数据
const QueryVdsChannelInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.can_query,
                               t.channel_no,
                               t.create_time,
                               t.ext_params,
                               t.first_query_time,
                               t.notify_url,
                               t.query_replenish_time,
                               t.query_url,
                               t.request_replenish_time,
                               t.request_url,
                               t.service_class,
                               t.status
                          from vds_channel_info t
                         where 1=1 &t.can_query &t.channel_no &t.service_class
                          &t.status
                         order by t.id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateVdsChannelInfo 更新渠道基本信息
const UpdateVdsChannelInfo = `
update 
vds_channel_info 
set
	can_query=@can_query,
	channel_no=@channel_no,
	ext_params=@ext_params,
	first_query_time=@first_query_time,
	notify_url=@notify_url,
	query_replenish_time=@query_replenish_time,
	query_url=@query_url,
	request_replenish_time=@request_replenish_time,
	request_url=@request_url,
	service_class=@service_class,
	status=@status
where
	&id`

//DeleteVdsChannelInfo 删除渠道基本信息
const DeleteVdsChannelInfo = `
delete from vds_channel_info 
where
&id`
