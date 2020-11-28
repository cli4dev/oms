// +build !oracle

package sql

const CheckFinishOrderInfo = `
SELECT 
t.order_id,
t.line_id
FROM
oms_order_info t 
WHERE t.order_id =@order_id
AND t.order_status in (0,91)  `

const CheckInvoicing = `
SELECT t.order_id,
t.invoice_method,
t.orig_invoice_id,
t.notify_url
  FROM oms_invoice_info t
 WHERE t.invoice_id = @invoice_id
   and t.invoice_status = 30
   and t.notify_status = 10
   and t.payment_status = 10
 `

const UpdateInvoiceFinish = `
update oms_invoice_info i set
i.invoice_status = if(@status='SUCCESS',0,if(@status='FAILED',90,i.invoice_status)),
i.fail_msg = @fail_msg,
i.fail_code = @fail_code,
i.notify_status = if(i.notify_url is null or i.notify_url ='',99,20),
i.payment_status = if(@status='SUCCESS' and i.invoice_method = 1,20,i.payment_status)
where i.invoice_id = @invoice_id
and i.invoice_status = 30
and i.notify_status = 10
and i.payment_status = 10 `

const UpdateOrderExtendInfo = `UPDATE 
crp_order_info_extend 
SET
invoice_status = if(@status='SUCCESS',0,10) 
WHERE order_id = @order_id
and invoice_status = 20 `

// SQLInvoiceRed 开票冲红
const SQLInvoiceRed = `
update oms_invoice_info i set
i.can_red = 0
where i.invoice_id = @orig_invoice_id
and i.invoice_status =0
and i.invoice_type = 1
and i.payment_status = 0
`
