// +build !oracle

package sql

const GetBindInfo = `select count(1)
from oms_order_info t
where t.order_id = @order_id
and t.order_status=20
and t.delivery_bind_status in (20,30)
`

const UpdateDeliveryUpPayStatus = `update oms_order_delivery t
set t.up_payment_status = 20
where t.delivery_id = @delivery_id
and t.up_payment_status = 30
`

const UpdateOrderNotifyStatus = `update oms_notify_info t
set t.notify_status = 20, t.max_count = t.max_count + 3
where t.order_id = @order_id
and t.notify_status = 30
and t.notify_type = 1
`

const UpdateReturnOvertime = `update oms_refund_info t
set t.return_overtime =
    date_add(now(),interval 15 minute)
where t.refund_id = @refund_id
`

const UpdateReturnStatus = `update oms_refund_up_return t
set t.return_status = 20
where t.return_id = @return_id
and t.refund_id=@refund_id
and t.return_status = 90
`

const UpdateUpRefundStatus = `update oms_refund_up_return t
set t.up_refund_status = 20
where t.return_id = @return_id
and t.up_refund_status = 30`

const UpdateDownRefundStatus = `update oms_refund_info t
set t.down_refund_status = 20
where t.refund_id = @refund_id
and t.down_refund_status = 30`

const UpdateRefundNotifyStatus = `update oms_notify_info t
set t.notify_status = 20, t.max_count = t.max_count + 3
where t.order_id = @order_id
and t.refund_id = @refund_id
and t.notify_type = 2
`
