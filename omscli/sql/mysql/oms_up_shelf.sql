
	CREATE TABLE  oms_up_shelf (
		shelf_id BIGINT(10)  not null AUTO_INCREMENT comment '货架编号' ,
		shelf_name VARCHAR(64)  not null  comment '货架名称' ,
		channel_no VARCHAR(32)  not null  comment '渠道编号' ,
		delivery_overtime BIGINT(10) default 300 not null  comment '发货超时时间' ,
		return_overtime BIGINT(10) default 300 not null  comment '退货超时时间' ,
		status TINYINT(1) default 0 not null  comment '货架状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		PRIMARY KEY (shelf_id)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='上游货架'
