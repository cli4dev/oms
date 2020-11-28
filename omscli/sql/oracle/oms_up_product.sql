drop table OMS_UP_PRODUCT;
drop sequence seq_up_product_id;

	create table OMS_UP_PRODUCT
(
  PRODUCT_ID          NUMBER(10) default 300 not null,
  SHELF_ID            NUMBER(10) not null,
  LINE_ID             NUMBER(10) not null,
  CARRIER_NO          VARCHAR2(8) not null,
  PROVINCE_NO         VARCHAR2(8) default '-' not null,
  CITY_NO             VARCHAR2(8) default '-' not null,
  INVOICE_TYPE        NUMBER(3) not null,
  EXT_PRODUCT_NO      VARCHAR2(32),
  CAN_REFUND          NUMBER(1) not null,
  FACE                NUMBER(20,5) not null,
  COST_DISCOUNT       NUMBER(10,5) not null,
  COMMISSION_DISCOUNT NUMBER(10,5) not null,
  SERVICE_DISCOUNT    NUMBER(10,5) not null,
  LIMIT_COUNT         NUMBER(10) not null,
  STATUS              NUMBER(1) default 0 not null,
  CREATE_TIME         DATE default sysdate not null
);


comment on table OMS_UP_PRODUCT  is '上游商品';
comment on column OMS_UP_PRODUCT.PRODUCT_ID  is '商品编号';
comment on column OMS_UP_PRODUCT.SHELF_ID  is '货架编号';
comment on column OMS_UP_PRODUCT.LINE_ID  is '产品线';
comment on column OMS_UP_PRODUCT.CARRIER_NO  is '运营商';
comment on column OMS_UP_PRODUCT.PROVINCE_NO  is '省份';
comment on column OMS_UP_PRODUCT.CITY_NO  is '城市';
comment on column OMS_UP_PRODUCT.INVOICE_TYPE  is '开票方式（1.不开发票，2.上游开发票）';
comment on column OMS_UP_PRODUCT.EXT_PRODUCT_NO  is '外部商品编号';
comment on column OMS_UP_PRODUCT.CAN_REFUND  is '支持退货';
comment on column OMS_UP_PRODUCT.FACE  is '面值';
comment on column OMS_UP_PRODUCT.COST_DISCOUNT  is '成本折扣（以面值算）';
comment on column OMS_UP_PRODUCT.COMMISSION_DISCOUNT  is '佣金折扣（以面值算）';
comment on column OMS_UP_PRODUCT.SERVICE_DISCOUNT  is '服务费折扣';
comment on column OMS_UP_PRODUCT.LIMIT_COUNT  is '单次最大发货数量';
comment on column OMS_UP_PRODUCT.STATUS  is '状态';
comment on column OMS_UP_PRODUCT.CREATE_TIME  is '创建时间';
 
	alter table oms_up_product
	add constraint pk_up_product primary key(product_id);
	alter table oms_up_product
	add constraint oms_up_product_line_id unique(line_id,carrier_no,province_no,city_no,invoice_type,can_refund,face);
	
	create sequence seq_up_product_id
	minvalue 300
	maxvalue 99999999999
	start with 300
	cache 20;
	