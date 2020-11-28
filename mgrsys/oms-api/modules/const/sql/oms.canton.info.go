package sql

//GetOmsCantonInfoDictionary  获取数据字典
const GetOmsCantonInfoDictionary = `
select simple_spell as value,chinese_name as name 
from oms_canton_info 
where  1=1
and grade = @grade
order by standard_code desc
 `

const GetOmsCantonInfoDictionaryQuery = `
select simple_spell as value,chinese_name as name 
from oms_canton_info 
where  1=1
order by standard_code desc
 `

//InsertOmsCantonInfo 添加省市信息
const InsertOmsCantonInfo = `
insert into oms_canton_info
(
	canton_code,
	chinese_name,
	spell,
	grade,
	parent,
	simple_spell,
	area_code,
	standard_code
)
values
(
	@canton_code,
	@chinese_name,
	@spell,
	@grade,
	@parent,
	@simple_spell,
	@area_code,
	@standard_code
)`

//GetOmsCantonInfo 查询单条数据省市信息
const GetOmsCantonInfo = `
select
t.canton_code,
t.chinese_name,
t.spell,
t.grade,
t.parent,
t.simple_spell,
t.area_code,
t.standard_code
from oms_canton_info t
where
&canton_code`

//GetOmsCantonInfoDictionaryByProvince  获取数据字典
const GetOmsCantonInfoDictionaryByProvince = `
select t.simple_spell as value,t.chinese_name as name 
from oms_canton_info t
where  1=1
and t.grade = @grade
and t.parent = @parent
`

//QueryOmsCantonInfoCount 获取省市信息列表条数
const QueryOmsCantonInfoCount = `
select count(1)
  from oms_canton_info t
 where 1=1 &t.canton_code &t.chinese_name &t.parent &t.simple_spell`

//QueryOmsCantonInfo 查询省市信息列表数据
const QueryOmsCantonInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.canton_code,
                               t.chinese_name,
                               t.spell,
                               t.grade,
                               t.parent,
                               t.simple_spell,
                               t.area_code,
                               t.standard_code
                          from oms_canton_info t
                         where 1=1 &t.canton_code &t.chinese_name &t.parent
                          &t.simple_spell) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateOmsCantonInfo 更新省市信息
const UpdateOmsCantonInfo = `
update 
oms_canton_info 
set
	chinese_name=@chinese_name,
	spell=@spell,
	grade=@grade,
	parent=@parent,
	simple_spell=@simple_spell,
	area_code=@area_code,
	standard_code=@standard_code
where
	&canton_code`

//DeleteOmsCantonInfo 删除省市信息
const DeleteOmsCantonInfo = `
delete from oms_canton_info 
where
&canton_code`
