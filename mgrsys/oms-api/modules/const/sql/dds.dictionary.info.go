package sql

//InsertDdsDictionaryInfo 添加字典表
const InsertDdsDictionaryInfo = `
insert into dds_dictionary_info
(	id,
	name,
	sort_no,
	type,
	value,
	status
)
values
(	seq_dds_dictionary_info_id.nextval,
	@name,
	@sort_no,
	@type,
	@value,
	0
)`

//GetDdsDictionaryInfo 查询单条数据字典表
const GetDdsDictionaryInfo = `
select
t.id,
t.name,
t.sort_no,
t.status,
t.type,
t.value
from dds_dictionary_info t
where
&id`

//QueryDdsDictionaryInfoCount 获取字典表列表条数
const QueryDdsDictionaryInfoCount = `
select count(1)
from dds_dictionary_info t
where
 t.name like '%' ||@name || '%'
&t.status
&t.type`

//QueryDdsDictionaryInfo 查询字典表列表数据
const QueryDdsDictionaryInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.name,
                               t.sort_no,
                               t.status,
                               t.type,
                               t.value
                          from dds_dictionary_info t
                         where t.name like '%' ||@name || '%'
                          &t.type
                         order by t.id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateDdsDictionaryInfo 更新字典表
const UpdateDdsDictionaryInfo = `
update 
dds_dictionary_info 
set
	name=@name,
	sort_no=@sort_no,
	status=@status,
	type=@type,
	value=@value
where
	&id`

//DeleteDdsDictionaryInfo 删除字典表
const DeleteDdsDictionaryInfo = `
delete from dds_dictionary_info 
where
&id`
