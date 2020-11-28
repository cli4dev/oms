package sql

//GetTskSystemTask 查询单条数据任务表
const GetTskSystemTask = `
select
t.task_id,
t.batch_id,
t.count,
t.create_time,
t.last_execute_time,
t.max_execute_time,
t.msg_content,
t.name,
t.next_execute_time,
t.next_interval,
t.queue_name,
t.status
from tsk_system_task t
where
&task_id`

//QueryTskSystemTaskCount 获取任务表列表条数
const QueryTskSystemTaskCount = `
select count(1)
  from tsk_system_task t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
   and t.queue_name like '%' || @queue_name || '%' &t.name
 &t.status`

//QueryTskSystemTask 查询任务表列表数据
const QueryTskSystemTask = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.task_id,
                               t.batch_id,
                               t.count,
                               t.create_time,
                               t.last_execute_time,
                               t.max_execute_time,
                               t.msg_content,
                               t.name,
                               t.next_execute_time,
                               t.next_interval,
                               t.queue_name,
                               t.status
                          from tsk_system_task t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                           and t.queue_name like '%' || @queue_name || '%'
                         &t.name &t.status) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteTskSystemTask 删除任务表
const DeleteTskSystemTask = `
delete from tsk_system_task 
where
&task_id`
