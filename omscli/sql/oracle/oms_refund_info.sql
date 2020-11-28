drop table OMS_REFUND_INFO;
drop sequence seq_refund_info_id;

create table OMS_REFUND_INFO
(
  REFUND_ID                NUMBER(20) default 20000 not null,
  ORDER_ID                 NUMBER(20) not null,
  DOWN_CHANNEL_NO          VARCHAR2(32) not null,
  REQUEST_NO               VARCHAR2(32) not null,
  DOWN_REFUND_NO           VARCHAR2(32) not null,
  DOWN_SHELF_ID            NUMBER(10) not null,
  DOWN_PRODUCT_ID          NUMBER(10) not null,
  EXT_PRODUCT_NO           VARCHAR2(32),
  LINE_ID                  NUMBER(10) not null,
  CARRIER_NO               VARCHAR2(8) not null,
  PROVINCE_NO              VARCHAR2(8) not null,
  CITY_NO                  VARCHAR2(8) not null,
  REFUND_TYPE              NUMBER(1) not null,
  FACE                     NUMBER(20,5) not null,
  REFUND_NUM               NUMBER(10) not null,
  REFUND_FACE              NUMBER(20,5) not null,
  REFUND_SELL_AMOUNT       NUMBER(20,5) not null,
  REFUND_COMMISSION_AMOUNT NUMBER(20,5) not null,
  REFUND_SERVICE_AMOUNT    NUMBER(20,5) not null,
  REFUND_FEE_AMOUNT        NUMBER(20,5) not null,
  CREATE_TIME              DATE default sysdate not null,
  REFUND_STATUS            NUMBER(3) default 10 not null,
  UP_RETURN_STATUS         NUMBER(3) default 20 not null,
  DOWN_REFUND_STATUS       NUMBER(3) default 10 not null,
  REFUND_NOTIFY_STATUS     NUMBER(3) default 10 not null,
  RETURN_OVERTIME          DATE not null,
  COMPLETE_UP_REFUND       NUMBER(1) default 1 not null
);
comment on table OMS_REFUND_INFO  is '退款记录';
comment on column OMS_REFUND_INFO.REFUND_ID  is '退款编号';
comment on column OMS_REFUND_INFO.ORDER_ID  is '订单编号';
comment on column OMS_REFUND_INFO.DOWN_CHANNEL_NO  is '下游渠道编号';
comment on column OMS_REFUND_INFO.REQUEST_NO  is '下游渠道订单号';
comment on column OMS_REFUND_INFO.DOWN_REFUND_NO  is '下游退款编号';
comment on column OMS_REFUND_INFO.DOWN_SHELF_ID  is '下游货架编号';
comment on column OMS_REFUND_INFO.DOWN_PRODUCT_ID  is '下游商品编号';
comment on column OMS_REFUND_INFO.EXT_PRODUCT_NO  is '外部商品编号';
comment on column OMS_REFUND_INFO.LINE_ID  is '产品线';
comment on column OMS_REFUND_INFO.CARRIER_NO  is '运营商';
comment on column OMS_REFUND_INFO.PROVINCE_NO  is '省份';
comment on column OMS_REFUND_INFO.CITY_NO  is '城市';
comment on column OMS_REFUND_INFO.REFUND_TYPE  is '退款方式（1.普通退款，2.强制退款,3.假成功退款）';
comment on column OMS_REFUND_INFO.FACE  is '商品面值';
comment on column OMS_REFUND_INFO.REFUND_NUM  is '退款商品数量';
comment on column OMS_REFUND_INFO.REFUND_FACE  is '退款商品总面值';
comment on column OMS_REFUND_INFO.REFUND_SELL_AMOUNT  is '退款总销售金额';
comment on column OMS_REFUND_INFO.REFUND_COMMISSION_AMOUNT  is '退款总佣金金额';
comment on column OMS_REFUND_INFO.REFUND_SERVICE_AMOUNT  is '退款总服务费金额';
comment on column OMS_REFUND_INFO.REFUND_FEE_AMOUNT  is '退款总手续费金额';
comment on column OMS_REFUND_INFO.CREATE_TIME  is '创建时间';
comment on column OMS_REFUND_INFO.REFUND_STATUS  is '状态（10.退货，20.退款，0成功，90失败）';
comment on column OMS_REFUND_INFO.UP_RETURN_STATUS  is '上游退货状态（0.退货成功，20.等待退货，30.正在退货，90.退款失败，91.部分退款）';
comment on column OMS_REFUND_INFO.DOWN_REFUND_STATUS  is '下游退款状态（0成功，10.未开始，20.等待，30正在，99无需）';
comment on column OMS_REFUND_INFO.REFUND_NOTIFY_STATUS  is '退款通知状态（0成功，100，查询成功，10.未开始，20.等待，30正在，99无需）';
comment on column OMS_REFUND_INFO.RETURN_OVERTIME  is '退货超时时间';
comment on column OMS_REFUND_INFO.COMPLETE_UP_REFUND  is '已完成上游退款（0.已完成，1.未完成）';

 
	alter table oms_refund_info
	add constraint pk_refund_info primary key(refund_id);
	alter table oms_refund_info
	add constraint oms_refund_channel_no unique(down_channel_no,down_refund_no);
	
	create sequence seq_refund_info_id
	minvalue 20000
	maxvalue 99999999999
	start with 20000
	cache 20;
	