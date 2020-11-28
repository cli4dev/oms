// +build oracle

package sql

// SQLInvoicePaymentSuccess 开票支付成功
const SQLInvoicePaymentSuccess = `
update oms_invoice_info i set
i.payment_status = 0,
i.payment_time = sysdate
where i.invoice_id = @invoice_id
and i.invoice_status in (0,90)
and i.payment_status = 20
and i.invoice_method = 1`

// SQLCheckInvoiceForPayment 开票支付检查
const SQLCheckInvoiceForPayment = `
select t.invoice_id
  from oms_invoice_info t
 where t.invoice_id = @invoice_id
   and t.invoice_status = 0 
   and t.payment_status = 20
   and t.invoice_method = 1
   and t.invoice_type = 1
`
