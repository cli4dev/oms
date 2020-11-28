// +build !oracle

package sql

const NotifyCheckInvoiceApplyInfo = `
SELECT 
  t.invoice_id,
  t.request_no,
  t.invoice_no,
  t.channel_no,
  t.order_id,
  t.fail_code,
  t.fail_msg,
  t.notify_url,
  t.can_red,
  t.orig_invoice_id,
  t.invoice_status
FROM
  oms_invoice_info t 
WHERE t.invoice_id = @invoice_id 
  AND t.invoice_status IN (0, 90) 
  AND t.notify_status = 20 
  AND t.notify_count < t.notify_max_count`

const StartInvoiceNotify = `
UPDATE 
  oms_invoice_info t 
SET
  t.notify_status = 30,
  t.start_time = IFNULL(t.start_time, now()) 
WHERE t.invoice_id = @invoice_id 
  AND t.invoice_status IN (0, 90)
  AND t.notify_status = 20
`

const FailedInvoiceNotify = `
UPDATE 
oms_invoice_info n
SET
  n.notify_count = n.notify_count + 1,
  n.notify_msg = @msg,
  n.notify_status = 20 
  WHERE n.invoice_id = @invoice_id
  AND n.notify_status = 30 `

const SuccessInvoiceNotify = `
UPDATE 
oms_invoice_info n 
SET
  n.notify_count = n.inform_count + 1,
  n.notify_msg = @msg,
  n.notify_status = 0,
  n.end_time = NOW()
WHERE n.invoice_id = @invoice_id 
  AND n.notify_status = 30 `

const NotifyCheckOrderInfo = `
  SELECT 
    t.request_no
  FROM
    oms_order_info t 
    INNER JOIN crp_order_info_extend c 
      ON c.order_id = t.order_id 
  WHERE t.order_id = @order_id 
    AND t.order_status IN (0,91) 
    AND c.invoice_status IN (0, 10)`
