package sql

//InsertVdsChannelErrorCode 添加渠道错误码
const InsertVdsChannelErrorCode = `
insert into vds_channel_error_code
(
	channel_no,
	deal_code,
	error_code,
	error_code_desc,
	service_class
)
values
(
	@channel_no,
	@deal_code,
	@error_code,
	@error_code_desc,
	@service_class
)`

//GetVdsChannelErrorCode 查询单条数据渠道错误码
const GetVdsChannelErrorCode = `
select
t.id,
t.channel_no,
t.create_time,
t.deal_code,
t.error_code,
t.error_code_desc,
t.service_class
from vds_channel_error_code t
where
&id`

//QueryVdsChannelErrorCodeCount 获取渠道错误码列表条数
const QueryVdsChannelErrorCodeCount = `
select count(1)
from vds_channel_error_code t
where 1=1 &t.channel_no
&t.service_class`

//QueryVdsChannelErrorCode 查询渠道错误码列表数据
const QueryVdsChannelErrorCode = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.channel_no,
                               t.create_time,
                               t.deal_code,
                               t.error_code,
                               t.error_code_desc,
                               t.service_class
                          from vds_channel_error_code t
                         where 1=1 &t.channel_no &t.service_class
                         order by t.id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//UpdateVdsChannelErrorCode 更新渠道错误码
const UpdateVdsChannelErrorCode = `
update 
vds_channel_error_code 
set
	channel_no=@channel_no,
	deal_code=@deal_code,
	error_code=@error_code,
	error_code_desc=@error_code_desc,
	service_class=@service_class
where
	&id`

//DeleteVdsChannelErrorCode 删除渠道错误码
const DeleteVdsChannelErrorCode = `
delete from vds_channel_error_code 
where
&id`
