// +build !oracle

package sql

// SQLCheckOrderInfo 检查订单信息
const SQLCheckOrderInfo = `
SELECT o.order_id, 
       p.sell_price, 
       o.num, 
  FROM oms_order_info o
 inner join oms_down_product p 
 on p.product_id = o.down_product_id
 WHERE o.order_id = @order_id
   AND o.order_status IN (0, 91)
`

// SQLGetDeliveryInfo 获取发货记录
const SQLGetDeliveryInfo = ` 
select c.invoice_channel_no, t.order_id, p.sell_price, t.num,t.up_product_id,t.down_channel_no,t.carrier_no
  from oms_order_delivery t
 inner join oms_up_channel c on t.up_channel_no = c.channel_no
 inner join oms_down_product p on p.product_id = t.down_product_id
 where t.order_id = @order_id
   and t.delivery_status = 0
   and t.up_payment_status in (0, 99)
`

// SQLStartInvoice 开始开票
const SQLStartInvoice = `
UPDATE 
  oms_invoice_info 
SET
  invoice_status = 30 
WHERE invoice_id = @invoice_id
AND invoice_status = 20
AND payment_status = 10
AND notify_status = 10
`

// SQLCheckInvoiceInfo 检查开票信息
const SQLCheckInvoiceInfo = `
SELECT t.invoice_id,
       t.order_id,
       t.invoice_title,
       t.tax_no,
       t.address,
       t.push_phone_no,
       t.bank_name,
       t.push_type,
       t.tele_phone,
       t.bank_account,
       t.bank_name,
       t.amount,
       t.orig_invoice_id,
       t.push_email,
       t.invoice_method,
       t.deduct_amount
  FROM oms_invoice_info t
 WHERE t.invoice_id = @invoice_id
 AND t.invoice_status = 20
 AND t.payment_status = 10
 AND t.notify_status = 10
 `
