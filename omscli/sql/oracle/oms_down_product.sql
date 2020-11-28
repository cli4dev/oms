
drop table oms_down_product;
drop sequence seq_down_product_id;


	create table oms_down_product(
		product_id  number(10)    default 300 not null ,
		shelf_id  number(10)     not null ,
		line_id  number(10)     not null ,
		carrier_no  varchar2(8)    not null ,
		province_no  varchar2(8)   default '-' not null ,
		city_no  varchar2(8)   default '-' not null ,
		invoice_type  number(3)      not null ,
		ext_product_no  varchar2(32)    ,
		can_refund  number(1)      not null ,
		face  number(20,5)     not null ,
		sell_discount  number(10,5)   not null ,
		commission_discount  number(10,5)   not null ,
		service_discount  number(10,5)   not null ,
		payment_fee_discount  number(10,5)   not null ,
		can_split_order  number(1)      not null ,
		split_order_face  number(20,5)     not null ,
		limit_count  number(10)     not null ,
		status  number(1)     default 0 not null ,
		create_time  date          default sysdate not null 
		);
	

	comment on table oms_down_product is '下游商品';
	comment on column oms_down_product.product_id is '商品编号';	
	comment on column oms_down_product.shelf_id is '货架编号';	
	comment on column oms_down_product.line_id is '产品线';	
	comment on column oms_down_product.carrier_no is '运营商';	
	comment on column oms_down_product.province_no is '省份';	
	comment on column oms_down_product.city_no is '城市';	
	comment on column oms_down_product.invoice_type is '开票方式（1.不开发票，0.不限制，2.需要发票）';	
	comment on column oms_down_product.ext_product_no is '外部商品编号';	
	comment on column oms_down_product.can_refund is '支持退款';	
	comment on column oms_down_product.face is '面值';	
	comment on column oms_down_product.sell_discount is '销售折扣（以面值算）';	
	comment on column oms_down_product.commission_discount is '佣金折扣（以面值算）';	
	comment on column oms_down_product.service_discount is '服务费折扣';	
	comment on column oms_down_product.payment_fee_discount is '手续费折扣（以销售金额算）';	
	comment on column oms_down_product.can_split_order is '是否拆单';	
	comment on column oms_down_product.split_order_face is '拆单面值';	
	comment on column oms_down_product.limit_count is '单次最大购买数量';	
	comment on column oms_down_product.status is '状态';	
	comment on column oms_down_product.create_time is '创建时间';	
	

 
	alter table oms_down_product
	add constraint pk_down_product primary key(product_id);
	alter table oms_down_product
	add constraint oms_down_product_carrier_no unique(carrier_no,province_no,city_no,invoice_type,can_refund,face);
	
	create sequence seq_down_product_id
	minvalue 300
	maxvalue 99999999999
	start with 300
	cache 20;
	