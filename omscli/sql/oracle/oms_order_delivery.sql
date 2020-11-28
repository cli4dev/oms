drop table OMS_ORDER_DELIVERY;
drop sequence seq_order_delivery_id;

	create table OMS_ORDER_DELIVERY
(
  DELIVERY_ID          NUMBER(20) default 20000 not null,
  UP_CHANNEL_NO        VARCHAR2(32) not null,
  UP_PRODUCT_ID        NUMBER(10) not null,
  UP_DELIVERY_NO       VARCHAR2(32),
  UP_EXT_PRODUCT_NO    VARCHAR2(32),
  ORDER_ID             NUMBER(20) not null,
  DOWN_CHANNEL_NO      VARCHAR2(32) not null,
  DOWN_PRODUCT_ID      NUMBER(10) not null,
  LINE_ID              NUMBER(10) not null,
  CARRIER_NO           VARCHAR2(8) not null,
  PROVINCE_NO          VARCHAR2(8) not null,
  CITY_NO              VARCHAR2(8) not null,
  INVOICE_TYPE         NUMBER(3) not null,
  DELIVERY_STATUS      NUMBER(3) default 20 not null,
  UP_PAYMENT_STATUS    NUMBER(3) default 10 not null,
  CREATE_TIME          DATE default sysdate not null,
  FACE                 NUMBER(20,5) not null,
  NUM                  NUMBER(10) not null,
  TOTAL_FACE           NUMBER(20,5) not null,
  COST_AMOUNT          NUMBER(20,5) not null,
  UP_COMMISSION_AMOUNT NUMBER(20,5) not null,
  SERVICE_AMOUNT       NUMBER(20,5) not null,
  START_TIME           DATE,
  END_TIME             DATE,
  RETURN_MSG           VARCHAR2(256)
);
	

comment on table OMS_ORDER_DELIVERY is '订单发货表';
comment on column OMS_ORDER_DELIVERY.DELIVERY_ID is '发货编号';
comment on column OMS_ORDER_DELIVERY.UP_CHANNEL_NO is '上游渠道编号';
comment on column OMS_ORDER_DELIVERY.UP_PRODUCT_ID is '上游商品编号';
comment on column OMS_ORDER_DELIVERY.UP_DELIVERY_NO  is '上游发货编号';
comment on column OMS_ORDER_DELIVERY.UP_EXT_PRODUCT_NO is '上游商品请求编号';
comment on column OMS_ORDER_DELIVERY.ORDER_ID  is '订单编号';
comment on column OMS_ORDER_DELIVERY.DOWN_CHANNEL_NO  is '下游渠道编号';
comment on column OMS_ORDER_DELIVERY.DOWN_PRODUCT_ID  is '下游商品编号';
comment on column OMS_ORDER_DELIVERY.LINE_ID  is '产品线';
comment on column OMS_ORDER_DELIVERY.CARRIER_NO  is '运营商';
comment on column OMS_ORDER_DELIVERY.PROVINCE_NO  is '省份';
comment on column OMS_ORDER_DELIVERY.CITY_NO  is '城市';
comment on column OMS_ORDER_DELIVERY.INVOICE_TYPE  is '开票方式（1.不开发票，2.上游开发票）';
comment on column OMS_ORDER_DELIVERY.DELIVERY_STATUS  is '发货状态（0.发货成功，20等待发货，30正在发货，90发货失败）';
comment on column OMS_ORDER_DELIVERY.UP_PAYMENT_STATUS  is '上游支付状态（0支付成功，10未开始,20.等待支付，30.正在支付，99.无需支付）';
comment on column OMS_ORDER_DELIVERY.CREATE_TIME  is '创建时间';
comment on column OMS_ORDER_DELIVERY.FACE  is '商品面值';
comment on column OMS_ORDER_DELIVERY.NUM  is '发货数量';
comment on column OMS_ORDER_DELIVERY.TOTAL_FACE  is '发货总面值';
comment on column OMS_ORDER_DELIVERY.COST_AMOUNT  is '发货成本';
comment on column OMS_ORDER_DELIVERY.UP_COMMISSION_AMOUNT  is '上游佣金';
comment on column OMS_ORDER_DELIVERY.SERVICE_AMOUNT  is '发货服务费';
comment on column OMS_ORDER_DELIVERY.START_TIME  is '开始时间';
comment on column OMS_ORDER_DELIVERY.END_TIME  is '结束时间';
comment on column OMS_ORDER_DELIVERY.RETURN_MSG  is '发货返回信息';
	

 
	alter table oms_order_delivery
	add constraint pk_order_delivery primary key(delivery_id);
	
	create sequence seq_order_delivery_id
	minvalue 20000
	maxvalue 99999999999
	start with 20000
	cache 20;
	