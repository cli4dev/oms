// +build oracle

package sql

// SQLCheckOrderByInvoice 开票检查订单
const SQLCheckOrderByInvoice = `
select round(t.sell_amount/t.num,2) sell_amount,
       decode(round(t.sell_amount/t.num,1),round(@amount,1),'true','false') can_sell_amount,
       t.order_id,
       t.invoice_type,
       t.down_channel_no,
       t.down_product_id,
       t.line_id,
       round(t.point_num * 0.01/t.num,2) deduct_amount,
       decode(round(t.point_num * 0.01/t.num,2),@deduct_amount,'true','false') can_deduct_amount,
       case when to_char(add_months(t.create_time,1),'yyyy-mm') >= to_char(sysdate,'yyyy-mm') then 'true' else 'false' end can_invoice,
       p.invoice_obj_type
       FROM oms_order_info t
       inner join oms_down_product p on p.product_id = t.down_product_id
 where t.down_channel_no = @channel_no
   and t.request_no = @request_no
   and t.order_status in (0, 91)
   and t.is_refund = 1
`

// SQLCheckDeliveryByInvoice 开票检查订单
const SQLCheckDeliveryByInvoice = `
select t.up_product_id, t.invoice_type, t.delivery_id,t.num
  from oms_order_delivery t
 where t.order_id = @order_id
   and t.delivery_status = 0
   and t.up_payment_status in (0,99)
`

// SQLCheckConsume 检查销券
const SQLCheckConsume = `
select count(1)
  from ebs_order_delivery_coupon t
 where t.detail_id = @detail_id
   and t.consume_status = 0
   and t.revoke_status = 1
`

// SQLCheckInvoice 检查是否开票
const SQLCheckInvoice = `
select t.invoice_id,
       t.invoice_no,
       t.request_no,
       t.fail_msg,
       t.fail_code,
       case when t.invoice_status = 20 then 30 else t.invoice_status end status,
       t.channel_no
  from oms_invoice_info t
 where t.channel_no = @channel_no
   and t.invoice_no = @invoice_no
`

// SQLCheckRefundByInvoice 检查是否存在退款
const SQLCheckRefundByInvoice = `
select t.order_id
  from oms_refund_info t
 where t.order_id = @order_id
 and t.refund_status != 90
`

// LockOrderByInvoice 锁定单
const SQLLockOrderByInvoice = `
SELECT 
o.sell_amount 
FROM
  oms_order_info o 
WHERE o.order_id = @order_id
  AND o.order_status = 0 
  AND o.is_refund = 1 
  AND o.payment_status = 0 
  FOR UPDATE
`

// SQLCreateInvocieInfo 创建开票
const SQLCreateInvocieInfo = `
insert into oms_invoice_info
  (invoice_id,
   invoice_no,
   channel_no,
   request_no,
   order_id,
   invoice_title,
   tax_no,
   address,
   tele_phone,
   bank_name,
   bank_account,
   push_type,
   push_phone_no,
   push_email,
   amount,
   deduct_amount,
   remark,
   create_time,
   invoice_method,
   invoice_type,
   invoice_status,
   payment_status,
   notify_status,
   notify_count,
   notify_max_count,
   notify_url,
   can_red,
   invoice_obj_type
   )values(
   seq_invoice_info_id.nextval,
   @invoice_no,
   @channel_no,
   @request_no,
   @order_id,
   @invoice_title,
   @tax_no,
   @address,
   @tele_phone,
   @bank_name,
   @bank_account,
   @push_type,
   @push_phone_no,
   @push_email,
   @sell_amount,
   @deduct_amount,
   @remark,
   sysdate,
   decode(@invoice_type,3,2,1),
   1,
   20,
   10,
   10,
   0,
   10,
   @notify_url,
   1,
   @invoice_obj_type
   )
`
