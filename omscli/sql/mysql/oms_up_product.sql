
	CREATE TABLE  oms_up_product (
		product_id BIGINT(10)  not null AUTO_INCREMENT comment '商品编号' ,
		shelf_id BIGINT(10)  not null  comment '货架编号' ,
		line_id BIGINT(10)  not null  comment '产品线' ,
		carrier_no VARCHAR(8)  not null  comment '运营商' ,
		province_no VARCHAR(8) default '-' not null  comment '省份' ,
		city_no VARCHAR(8) default '-' not null  comment '城市' ,
		invoice_type SMALLINT(3)  not null  comment '开票方式（1.不开发票，2.上游开发票）' ,
		ext_product_no VARCHAR(32)    comment '外部商品编号' ,
		can_refund TINYINT(1)  not null  comment '支持退货' ,
		face DECIMAL(20,5)  not null  comment '面值' ,
		cost_discount DECIMAL(10,5)  not null  comment '成本折扣（以面值算）' ,
		commission_discount DECIMAL(10,5)  not null  comment '佣金折扣（以面值算）' ,
		service_discount DECIMAL(10,5)  not null  comment '服务费折扣' ,
		limit_count BIGINT(10)  not null  comment '单次最大发货数量' ,
		status TINYINT(1) default 0 not null  comment '状态' ,
		create_time DATETIME default current_timestamp not null  comment '创建时间' ,
		extend_info LONGTEXT    comment '扩展信息' ,
		PRIMARY KEY (product_id),
		UNIQUE KEY UNQ_UP_PRODUCT (can_refund,carrier_no,city_no,face,invoice_type,line_id,province_no)
		) ENGINE=InnoDB AUTO_INCREMENT=300 DEFAULT CHARSET=utf8 COMMENT='上游商品'
