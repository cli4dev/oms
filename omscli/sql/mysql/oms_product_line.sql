
	CREATE TABLE  oms_product_line (
		line_id BIGINT(10)  not null AUTO_INCREMENT comment '产品线编号' ,
		line_name VARCHAR(64)  not null  comment '产品线名称' ,
		payment_queue VARCHAR(32) default 'ebs:order:downpay:sys'   comment '支付队列' ,
		bind_queue VARCHAR(32) default 'ebs:order:bind'   comment '绑定队列' ,
		delivery_queue VARCHAR(32) default 'ebs:order:delivery'   comment '发货队列' ,
		up_payment_queue VARCHAR(32) default 'ebs:order:uppay:sys'   comment '上游支付队列' ,
		notify_queue VARCHAR(32) default 'ebs:order:notify'   comment '通知队列' ,
		return_queue VARCHAR(32) default 'ebs:refund:return'   comment '退货队列' ,
		refund_queue VARCHAR(32) default 'ebs:refund:downpay:sys'   comment '退款队列' ,
		up_refund_queue VARCHAR(32) default 'ebs:refund:uppay:sys'   comment '上游退款队列' ,
		refund_notify_queue VARCHAR(32) default 'ebs:refund:notify'   comment '退款通知队列' ,
		order_refund_queue VARCHAR(32) default 'ebs:refund:orderfail:downpay:sys'   comment '订单失败退款队列' ,
		order_overtime_queue VARCHAR(32) default ' ebs:order:overtime'   comment '订单超时处理队列' ,
		delivery_unknown_queue VARCHAR(32) default 'ebs:order:overtime:delivery'   comment '发货未知处理队列' ,
		refund_overtime_queue VARCHAR(32) default 'ebs:refund:overtime'   comment '退款超时处理队列' ,
		return_unknown_queue VARCHAR(32) default 'ebs:refund:overtime:return'   comment '退货未知处理队列' ,
		PRIMARY KEY (line_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='产品线'
