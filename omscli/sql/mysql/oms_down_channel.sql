
	CREATE TABLE  oms_down_channel (
		channel_no VARCHAR(32)  not null  comment '编号' ,
		channel_name VARCHAR(64)  not null  comment '名称' ,
		status TINYINT(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		ext_channel_no VARCHAR(32)    comment '外部渠道编号' ,
		PRIMARY KEY (channel_no)) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='下游渠道'
