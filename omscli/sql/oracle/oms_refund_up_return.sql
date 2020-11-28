drop table OMS_REFUND_UP_RETURN;
drop sequence seq_refund_up_return_id;

	create table OMS_REFUND_UP_RETURN
(
  RETURN_ID                NUMBER(20) default 20000 not null,
  UP_CHANNEL_NO            VARCHAR2(32) not null,
  UP_PRODUCT_ID            NUMBER(10) not null,
  UP_RETURN_NO             VARCHAR2(32),
  UP_EXT_PRODUCT_NO        VARCHAR2(32),
  ORDER_ID                 NUMBER(20) not null,
  DELIVERY_ID              NUMBER(20) not null,
  REFUND_ID                NUMBER(20) not null,
  DOWN_CHANNEL_NO          VARCHAR2(32) not null,
  DOWN_PRODUCT_ID          NUMBER(10) not null,
  LINE_ID                  NUMBER(10) not null,
  CARRIER_NO               VARCHAR2(8) not null,
  PROVINCE_NO              VARCHAR2(8) not null,
  CITY_NO                  VARCHAR2(8) not null,
  RETURN_STATUS            NUMBER(3) default 20 not null,
  UP_REFUND_STATUS         NUMBER(3) default 10 not null,
  CREATE_TIME              DATE default sysdate not null,
  RETURN_FACE              NUMBER(20,5) not null,
  RETURN_NUM               NUMBER(10) not null,
  RETURN_TOTAL_FACE        NUMBER(20,5) not null,
  RETURN_COST_AMOUNT       NUMBER(20,5) not null,
  RETURN_COMMISSION_AMOUNT NUMBER(20,5) not null,
  RETURN_SERVICE_AMOUNT    NUMBER(20,5) not null,
  START_TIME               DATE,
  END_TIME                 DATE,
  RETURN_MSG               VARCHAR2(256)
);
comment on table OMS_REFUND_UP_RETURN  is '上游退货信息表';
comment on column OMS_REFUND_UP_RETURN.RETURN_ID  is '退货编号';
comment on column OMS_REFUND_UP_RETURN.UP_CHANNEL_NO  is '上游渠道编号';
comment on column OMS_REFUND_UP_RETURN.UP_PRODUCT_ID  is '上游商品编号';
comment on column OMS_REFUND_UP_RETURN.UP_RETURN_NO  is '上游退货编号';
comment on column OMS_REFUND_UP_RETURN.UP_EXT_PRODUCT_NO  is '上游商品请求编号';
comment on column OMS_REFUND_UP_RETURN.ORDER_ID  is '订单编号';
comment on column OMS_REFUND_UP_RETURN.DELIVERY_ID  is '发货编号';
comment on column OMS_REFUND_UP_RETURN.REFUND_ID  is '退款编号';
comment on column OMS_REFUND_UP_RETURN.DOWN_CHANNEL_NO  is '下游渠道编号';
comment on column OMS_REFUND_UP_RETURN.DOWN_PRODUCT_ID  is '下游商品编号';
comment on column OMS_REFUND_UP_RETURN.LINE_ID  is '产品线';
comment on column OMS_REFUND_UP_RETURN.CARRIER_NO  is '运营商';
comment on column OMS_REFUND_UP_RETURN.PROVINCE_NO  is '省份';
comment on column OMS_REFUND_UP_RETURN.CITY_NO  is '城市';
comment on column OMS_REFUND_UP_RETURN.RETURN_STATUS  is '退货状态（0.退货成功，20等待退货，30正在退货，90退货失败）';
comment on column OMS_REFUND_UP_RETURN.UP_REFUND_STATUS  is '退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99无需退款）';
comment on column OMS_REFUND_UP_RETURN.CREATE_TIME  is '创建时间';
comment on column OMS_REFUND_UP_RETURN.RETURN_FACE  is '商品面值';
comment on column OMS_REFUND_UP_RETURN.RETURN_NUM  is '退货数量';
comment on column OMS_REFUND_UP_RETURN.RETURN_TOTAL_FACE  is '退货总面值';
comment on column OMS_REFUND_UP_RETURN.RETURN_COST_AMOUNT  is '退回成本';
comment on column OMS_REFUND_UP_RETURN.RETURN_COMMISSION_AMOUNT  is '退回佣金';
comment on column OMS_REFUND_UP_RETURN.RETURN_SERVICE_AMOUNT  is '退回服务费';
comment on column OMS_REFUND_UP_RETURN.START_TIME  is '开始时间';
comment on column OMS_REFUND_UP_RETURN.END_TIME  is '结束时间';
comment on column OMS_REFUND_UP_RETURN.RETURN_MSG  is '退货返回信息';
	

 
	alter table oms_refund_up_return
	add constraint pk_refund_up_return primary key(return_id);
	
	create sequence seq_refund_up_return_id
	minvalue 20000
	maxvalue 99999999999
	start with 20000
	cache 20;
	