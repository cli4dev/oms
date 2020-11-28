
	CREATE TABLE  oms_audit_info (
		audit_id BIGINT(10)  not null AUTO_INCREMENT comment '人工审核编号' ,
		order_id BIGINT(20)  not null  comment '订单编号' ,
		refund_id BIGINT(20)    comment '退款编号' ,
		delivery_id BIGINT(20)    comment '发货记录编号' ,
		change_type SMALLINT(3)  not null  comment '变动类型（1.发货，2.退货，3.订单，4.退款，5假成功）' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		audit_status SMALLINT(3)  not null  comment '审核状态' ,
		audit_by BIGINT(10)    comment '审核人' ,
		audit_time DATETIME    comment '审核时间' ,
		audit_msg VARCHAR(256)    comment '审核信息' ,
		PRIMARY KEY (audit_id),
		KEY key_oms_audit_info_create_time (create_time)
		) ENGINE=InnoDB AUTO_INCREMENT=2000 DEFAULT CHARSET=utf8 COMMENT='发货人工审核表'
