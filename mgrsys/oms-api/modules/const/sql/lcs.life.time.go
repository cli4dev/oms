package sql

//GetLcsLifeTime 查询单条数据生命周期记录表
const GetLcsLifeTime = `
select
t.id,
t.order_no,
t.batch_no,
t.extral_param,
t.content,
t.create_time,
t.ip
from lcs_life_time t
where
&id`

//QueryLcsLifeTimeCount 获取生命周期记录表列表条数
const QueryLcsLifeTimeCount = `
select count(1)
  from lcs_life_time t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
   and t.content like   '%' || @content || '%' 
   &t.order_no
 &t.batch_no &t.ip = @ip`

//QueryLcsLifeTime 查询生命周期记录表列表数据
const QueryLcsLifeTime = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.id,
                               t.order_no,
                               t.batch_no,
                               t.extral_param,
                               t.content,
                               t.create_time,
                               t.ip
                          from lcs_life_time t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                           and t.content like   '%' || @content || '%' 
                         &t.order_no &t.batch_no &t.ip = @ip) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteLcsLifeTime 删除生命周期记录表
const DeleteLcsLifeTime = `
delete from lcs_life_time 
where
&id`
