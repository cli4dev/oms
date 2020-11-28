// +build oracle

package sql

// StartNotify 开始通知
const StartNotify = `
UPDATE oms_notify_info n SET 
n.notify_status = 30,
n.start_time = nvl(n.start_time, sysdate)
WHERE n.notify_id = @notify_id
AND n.notify_type = 1
AND n.notify_status = 20
`

// SuccessNotify 成功通知
const SuccessNotify = `
UPDATE 
oms_notify_info n
 SET 
n.notify_status = 0,
n.end_time = sysdate,
n.notify_count = n.notify_count + 1,
n.notify_msg = @msg 
WHERE n.notify_id = @notify_id
AND n.notify_type = 1
AND n.notify_status = 30
`

// SuccessOrderNotify 成功订单通知
const SuccessOrderNotify = `
UPDATE 
  oms_order_info o 
SET
  o.notify_status = 0,
  o.order_overtime = to_date('2099-12-31 23:59:59','yyyy-mm-dd hh24:mi:ss')
WHERE o.order_id = @order_id 
  AND o.order_status in (0, 90, 91)  
  AND o.notify_status = 30 
`

// FailedNotify 失败通知
const FailedNotify = `
UPDATE 
  oms_notify_info n 
SET
  n.notify_count = n.notify_count + 1,
  n.notify_msg = @msg, 
  n.notify_status = 20
WHERE n.notify_id = @notify_id 
  AND n.notify_type = 1 
  AND n.notify_status = 30 
`

// ChechOrderForNotify 检查订单对于通知
const ChechOrderForNotify = `
SELECT o.order_id,
       o.order_status        status,
       o.request_no,
       o.down_product_id     product_id,
       o.down_channel_no     channel_no,
       o.down_account_no account_no,
       o.point_num,
       o.fail_code,
       o.fail_msg,
       r.buy_send_num,
       r.activity_send_num
  FROM oms_order_info o
  left join jf_fd_order_record r
    on o.order_id = r.order_id
   and r.request_type = 1
 WHERE o.order_id = @order_id
   AND o.order_status IN (0, 90, 91)
   AND o.notify_status = 30
`

// CheckNotify 检查通知
const CheckNotify = `
SELECT
  n.notify_id,
  n.order_id,
  n.notify_url
  FROM
  oms_notify_info n 
WHERE n.notify_id = @notify_id 
  AND n.notify_status IN (20, 30) 
  AND n.notify_type = 1
  AND n.notify_count < n.max_count 
`

// QueryNotifyInfo 查询通知信息
const QueryNotifyInfo = `
SELECT 
  o.order_id,
  o.order_status STATUS,
  o.request_no,
  o.down_product_id product_id,
  o.down_channel_no channel_no,
  o.down_account_no account_no,
  n.notify_url
FROM
  oms_order_info o 
  INNER JOIN oms_notify_info n 
    ON o.order_id = n.order_id 
    WHERE o.order_id = @order_id
    AND n.notify_type = 1
`
