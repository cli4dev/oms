package sql

//GetBeanpayAccountRecord 查询单条数据账户余额变动信息
const GetBeanpayAccountRecord = `
select
t.record_id,
t.account_id,
t.amount,
t.balance,
t.change_type,
t.create_time,
t.ext_no,
t.ext,
t.trade_no,
t.trade_type
from beanpay_account_record t
where
&record_id`

//QueryBeanpayAccountRecordCount 获取账户余额变动信息列表条数
const QueryBeanpayAccountRecordCount = `
SELECT COUNT(1)
  FROM beanpay_account_record t
 INNER JOIN beanpay_account_info i ON i.account_id = t.account_id
 WHERE t.create_time >= to_date(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_date(@end_time, 'yyyy-mm-dd hh24:mi:ss')
   and i.groups like '%' || @types || '%' 
   &t.account_id
 &t.change_type &t.trade_type
`

//QueryBeanpayAccountRecord 查询账户余额变动信息列表数据
const QueryBeanpayAccountRecord = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.record_id,
                               t.account_id,
                               t.amount,
                               t.balance,
                               t.change_type,
                               t.create_time,
                               t.deduct_no,
                               t.ext,
                               i.eid channel_no,
                               t.trade_no,
                               t.trade_type
                          from beanpay_account_record t
                         INNER JOIN beanpay_account_info i ON i.account_id =
                                                              t.account_id
                         WHERE t.create_time >=
                         to_date(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                           to_date(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                           and i.groups like '%' || @types || '%'
                         &t.account_id &t.change_type &t.trade_type
                         order by t.record_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//DeleteBeanpayAccountRecord 删除账户余额变动信息
const DeleteBeanpayAccountRecord = `
delete from beanpay_account_record 
where
&record_id`
