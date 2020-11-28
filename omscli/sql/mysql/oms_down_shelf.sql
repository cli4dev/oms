
	CREATE TABLE  oms_down_shelf (
		shelf_id BIGINT(10)  not null AUTO_INCREMENT comment '货架编号' ,
		shelf_name VARCHAR(64)  not null  comment '货架名称' ,
		channel_no VARCHAR(32)  not null  comment '渠道编号' ,
		order_overtime BIGINT(10)  not null  comment '订单超时时长' ,
		refund_overtime BIGINT(10)  not null  comment '退款超时时长' ,
		status TINYINT(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		extend_info LONGTEXT    comment '扩展信息' ,
		PRIMARY KEY (shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='下游货架'
