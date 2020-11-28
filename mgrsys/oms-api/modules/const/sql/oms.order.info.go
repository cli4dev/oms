package sql

//GetOmsOrderInfo 查询单条数据订单记录
const GetOmsOrderInfo = `
select
t.order_id,
t.bind_face,
t.can_split_order,
t.carrier_no,
t.province_no,
t.city_no,
t.commission_amount,
t.complete_up_pay,
t.create_time,
t.delivery_bind_status,
t.delivery_pause,
t.down_channel_no,
t.down_product_id,
t.down_shelf_id,
t.ext_product_no,
t.face,
t.fee_amount,
t.invoice_type,
t.is_refund,
t.line_id,
t.notify_status,
t.num,
t.order_overtime,
t.order_status,
t.payment_status,
t.profit,
t.rechage_account,
t.refund_status,
t.request_no,
t.sell_amount,
t.service_amount,
t.split_order_face,
t.success_commission,
t.success_cost_amount,
t.success_face,
t.success_fee,
t.success_sell_amount,
t.success_service,
t.success_up_commission,
t.success_up_service,
o.sell_discount,
o.commission_discount,
o.service_discount,
o.payment_fee_discount,
t.total_face
from oms_order_info t
left join oms_down_product o on o.product_id = t.down_product_id
where
&order_id`

const GetVdsOrderInfoList = `select
t.order_no,
t.carrier_no,
t.channel_no,
t.coop_id,
t.coop_order_id,
t.create_time,
t.flow_timeout,
t.last_update_time,
t.notify_url,
t.product_face,
t.product_num,
t.request_finish_time,
t.request_params,
t.request_start_time,
t.result_code,
t.result_desc,
t.result_params,
t.result_source,
t.service_class,
t.status,
t.succ_face,
t.up_order_no,
o.channel_name
from vds_order_info t
left join oms_up_channel o on o.channel_no =  t.channel_no
where
t.coop_order_id = @coop_order_id`

//QueryOmsOrderInfoCount 获取订单记录列表条数
const QueryOmsOrderInfoCount = `
select count(1)
  from oms_order_info t
 where t.create_time >= to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
   AND t.create_time < to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
 &t.can_split_order &t.carrier_no &t.province_no &t.order_id
 &t.city_no &t.complete_up_pay &t.delivery_bind_status
 &t.delivery_pause &t.down_channel_no &t.down_product_id
 &t.down_shelf_id &t.invoice_type &t.is_refund &t.line_id
 &t.notify_status &t.order_status &t.payment_status
 &t.rechage_account &t.refund_status &t.request_no
`

//QueryOmsOrderInfo 查询订单记录列表数据
const QueryOmsOrderInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select t.order_id,
                               t.bind_face,
                               t.can_split_order,
                               t.carrier_no,
                               t.province_no,
                               t.city_no,
                               t.commission_amount,
                               t.complete_up_pay,
                               t.create_time,
                               t.delivery_bind_status,
                               t.delivery_pause,
                               t.down_channel_no,
                               t.down_product_id,
                               t.down_shelf_id,
                               t.ext_product_no,
                               t.face,
                               t.invoice_type,
                               t.is_refund,
                               t.line_id,
                               t.notify_status,
                               t.num,
                               t.order_overtime,
                               t.order_status,
                               t.payment_status,
                               t.rechage_account,
                               t.refund_status,
                               t.request_no,
                               t.total_face
                          from oms_order_info t
                         where t.create_time >=
                               to_char(@start_time, 'yyyy-mm-dd hh24:mi:ss')
                           AND t.create_time <
                               to_char(@end_time, 'yyyy-mm-dd hh24:mi:ss')
                         &t.can_split_order &t.carrier_no &t.province_no
                         &t.order_id &t.city_no &t.complete_up_pay
                         &t.delivery_bind_status &t.delivery_pause
                         &t.down_channel_no &t.down_product_id
                         &t.down_shelf_id &t.invoice_type &t.is_refund
                         &t.line_id &t.notify_status &t.order_status
                         &t.payment_status &t.rechage_account
                         &t.refund_status &t.request_no
                         order by order_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

//DeleteOmsOrderInfo 删除订单记录
const DeleteOmsOrderInfo = `
delete from oms_order_info 
where
&order_id`

// GetOrderDeliveryInfoCount 获取订单发货信息
const GetOrderDeliveryInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_order_delivery d on o.order_id = d.order_id
  where o.order_id = @order_id
`

// GetOrderDeliveryInfo 获取订单发货信息
const GetOrderDeliveryInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_order_delivery d on o.order_id =
                                                            d.order_id
                         where o.order_id = @order_id
                         order by d.delivery_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

// GetOrderNotifyInfoCount 获取订单通知信息
const GetOrderNotifyInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_notify_info d on o.order_id = d.order_id and d.notify_type ="1"
  where o.order_id = @order_id
`

// GetOrderNotifyInfo 获取订单通知信息
const GetOrderNotifyInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_notify_info d on o.order_id =
                                                         d.order_id
                                                     and d.notify_type = "1"
                         where o.order_id = @order_id
                         order by d.notify_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

// GetOrderRefundNotifyInfoCount 获取订单退款通知信息
const GetOrderRefundNotifyInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_notify_info d on o.order_id = d.order_id and d.notify_type ="2"
  where o.order_id = @order_id
`

// GetOrderRefundNotifyInfo 获取订单退款通知信息
const GetOrderRefundNotifyInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_notify_info d on o.order_id =
                                                         d.order_id
                                                     and d.notify_type = "2"
                         where o.order_id = @order_id
                         order by d.notify_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

// GetOrderAuditInfoCount 获取订单审核信息
const GetOrderAuditInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_audit_info d on o.order_id = d.order_id
  where o.order_id = @order_id
`

// GetOrderAuditInfo 获取订单审核信息
const GetOrderAuditInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_audit_info d on o.order_id =
                                                        d.order_id
                         where o.order_id = @order_id
                         order by d.audit_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1
`

// GetOrderRefundInfoCount 获取订单退款信息
const GetOrderRefundInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_refund_info d on o.order_id = d.order_id
  where o.order_id = @order_id
`

// GetOrderRefundInfo 获取订单退款信息
const GetOrderRefundInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_refund_info d on o.order_id =
                                                         d.order_id
                         where o.order_id = @order_id
                         order by d.refund_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

// GetOrderReturnInfoCount 获取订单退货信息
const GetOrderReturnInfoCount = `
select count(1)
  from oms_order_info o
  inner join oms_refund_up_return d on o.order_id = d.order_id
  where o.order_id = @order_id
`

// GetOrderReturnInfo 获取订单退货信息
const GetOrderReturnInfo = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select d.*
                          from oms_order_info o
                         inner join oms_refund_up_return d on o.order_id =
                                                              d.order_id
                         where o.order_id = @order_id
                         order by d.return_id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1

`

// GetOrderLifeTimeCount 获取订单生命周期
const GetOrderLifeTimeCount = `
select count(1) from (select t.*
	from oms_order_info o
	inner join lcs_life_time t on o.order_id = t.order_no
   where o.order_id = @order_id
  union
  select t.*
	from oms_order_info o
	inner join oms_order_delivery d on o.order_id = d.order_id
	inner join lcs_life_time t on t.order_no = d.delivery_id
   where o.order_id =  @order_id
  ) l order by l.create_time desc
`

// GetOrderLifeTime 获取订单生命周期
const GetOrderLifeTime = `
select TAB1.*
  from (select L.*
          from (select rownum as rn, R.*
                  from (select l.*
                          from (select t.*
                                  from oms_order_info o
                                 inner join lcs_life_time t on o.order_id =
                                                               t.order_no
                                 where o.order_id = @order_id
                                union
                                select t.*
                                  from oms_order_info o
                                 inner join oms_order_delivery d on o.order_id =
                                                                    d.order_id
                                 inner join lcs_life_time t on t.order_no =
                                                               d.delivery_id
                                 where o.order_id = @order_id) l
                         order by l.id desc) R
                 where rownum <= @pi * @ps) L
         where L.rn > (@pi - 1) * @ps) TAB1


`

const GetDownPayTime = `select TAB1.*
from (select L.*
        from (select rownum as rn, R.*
                from (select d.*
                        from oms_order_info o
                       inner join beanpay_account_record d on o.order_id =
                                                              d.trade_no
                       where o.order_id = @order_id
                       order by d.record_id desc) R
               where rownum <= @pi * @ps) L
       where L.rn > (@pi - 1) * @ps) TAB1`

const GetDownPayCount = `select count(1)
from oms_order_info o
inner join beanpay_account_record d on o.order_id = d.trade_no
where o.order_id = @order_id`

const GetUpPayTime = `select TAB1.*
from (select L.*
        from (select rownum as rn, R.*
                from (select d.*
                        from oms_order_delivery o
                       inner join beanpay_account_record d on o.delivery_id =
                                                              d.trade_no
                       where o.order_id = @order_id
                       order by d.record_id desc) R
               where rownum <= @pi * @ps) L
       where L.rn > (@pi - 1) * @ps) TAB1
`

const GetUpPayCount = `select count(1)
from oms_order_delivery o
inner join beanpay_account_record d on o.delivery_id = d.trade_no
where o.order_id = @order_id`
