
	CREATE TABLE  oms_notify_info (
		notify_id BIGINT(20)  not null AUTO_INCREMENT comment '通知编号' ,
		order_id BIGINT(20)  not null  comment '订单编号' ,
		refund_id BIGINT(20)    comment '退款编号' ,
		notify_type SMALLINT(3)  not null  comment '通知类型（1.订单通知，2.退款通知）' ,
		notify_status SMALLINT(3) default 10 not null  comment '通知状态（0成功，10未开始,20等待通知，30正在通知）' ,
		notify_count SMALLINT(3) default 0 not null  comment '通知次数' ,
		max_count SMALLINT(3)  not null  comment '最大通知次数' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		start_time DATETIME    comment '开始时间' ,
		end_time DATETIME    comment '结束时间' ,
		notify_url VARCHAR(128)  not null  comment '通知地址' ,
		notify_msg VARCHAR(256)    comment '通知结果信息' ,
		PRIMARY KEY (notify_id),
		KEY key_oms_notify_info_create_time (create_time)
		) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=utf8 COMMENT='订单通知表'
