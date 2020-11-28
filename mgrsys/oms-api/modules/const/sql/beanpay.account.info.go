package sql

//GetBeanpayDownAccountInfoDictionary  获取数据字典
const GetBeanpayDownAccountInfoDictionary = `
select account_id as value,account_name as name 
from beanpay_account_info 
where   groups='down_channel'`

//GetBeanpayUpAccountInfoDictionary  获取数据字典
const GetBeanpayUpAccountInfoDictionary = `
select account_id as value,account_name as name 
from beanpay_account_info 
where   groups='up_channel'`

//InsertBeanpayAccountInfo 添加账户信息
const InsertBeanpayAccountInfo = `
insert into beanpay_account_info
(
	account_name,
	eid,
	groups,
	ident
)
values
(
	@account_name,
	@eid,
	@groups,
	@ident
)`

//GetBeanpayAccountInfo 查询单条数据账户信息
const GetBeanpayAccountInfo = `
select
t.account_id,
t.account_name,
t.balance,
t.credit,
t.create_time,
t.eid,
t.groups,
t.ident,
t.status
from beanpay_account_info t
where
&account_id`

//QueryBeanpayAccountInfoCount 获取账户信息列表条数
const QueryBeanpayAccountInfoCount = `
select count(1)
from beanpay_account_info t
where  t.groups like '%' || @types || '%'
 &t.account_name
 &t.eid
 &t.groups
 &t.ident
 &t.status`

//QueryBeanpayAccountInfo 查询账户信息列表数据
const QueryBeanpayAccountInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.account_id,
                               t.account_name,
                               t.ident,
                               t.groups,
                               t.eid,
                               t.balance,
                               t.credit,
                               t.create_time,
                               t.status
                          from beanpay_account_info t
                         where t.groups like '%' || @types || '%'
                         &t.account_name &t.eid 
                         &t.groups  &t.ident
                         &t.status
                         order by t.account_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

//UpdateBeanpayAccountInfo 更新账户信息
const UpdateBeanpayAccountInfo = `
update 
beanpay_account_info 
set
	account_name=@account_name,
	eid=@eid,
	ident=@ident,
	status=@status
where
	&account_id`

//DeleteBeanpayAccountInfo 删除账户信息
const DeleteBeanpayAccountInfo = `
delete from beanpay_account_info 
where
&account_id`
