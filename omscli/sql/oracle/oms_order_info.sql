drop table OMS_ORDER_INFO;
drop sequence seq_order_info_id;


	create table OMS_ORDER_INFO
(
  ORDER_ID              NUMBER(20) default 1100000000 not null,
  DOWN_CHANNEL_NO       VARCHAR2(32) not null,
  REQUEST_NO            VARCHAR2(64) not null,
  DOWN_SHELF_ID         NUMBER(10) not null,
  DOWN_PRODUCT_ID       NUMBER(10) not null,
  EXT_PRODUCT_NO        VARCHAR2(32),
  LINE_ID               NUMBER(10) not null,
  CARRIER_NO            VARCHAR2(8) not null,
  PROVINCE_NO           VARCHAR2(8) not null,
  CITY_NO               VARCHAR2(8) not null,
  INVOICE_TYPE          NUMBER(3) not null,
  FACE                  NUMBER(20,5) not null,
  NUM                   NUMBER(10) not null,
  TOTAL_FACE            NUMBER(20,5) not null,
  SELL_AMOUNT           NUMBER(20,5) not null,
  COMMISSION_AMOUNT     NUMBER(20,5) not null,
  SERVICE_AMOUNT        NUMBER(20,5) not null,
  FEE_AMOUNT            NUMBER(20,5) not null,
  CAN_SPLIT_ORDER       NUMBER(1) not null,
  SPLIT_ORDER_FACE      NUMBER(20,5) not null,
  CREATE_TIME           DATE default sysdate not null,
  ORDER_OVERTIME        DATE not null,
  DELIVERY_PAUSE        NUMBER(1) default 1 not null,
  ORDER_STATUS          NUMBER(3) default 10 not null,
  PAYMENT_STATUS        NUMBER(3) default 20 not null,
  DELIVERY_BIND_STATUS  NUMBER(3) default 10 not null,
  REFUND_STATUS         NUMBER(3) default 10 not null,
  NOTIFY_STATUS         NUMBER(3) default 10 not null,
  IS_REFUND             NUMBER(1) default 1 not null,
  BIND_FACE             NUMBER(20,5) default 0 not null,
  SUCCESS_FACE          NUMBER(20,5) default 0 not null,
  SUCCESS_SELL_AMOUNT   NUMBER(20,5) default 0 not null,
  SUCCESS_COMMISSION    NUMBER(20,5) default 0 not null,
  SUCCESS_SERVICE       NUMBER(20,5) default 0 not null,
  SUCCESS_FEE           NUMBER(20,5) default 0 not null,
  SUCCESS_COST_AMOUNT   NUMBER(20,5) default 0 not null,
  SUCCESS_UP_COMMISSION NUMBER(20,5) default 0 not null,
  SUCCESS_UP_SERVICE    NUMBER(20,5) default 0 not null,
  PROFIT                NUMBER(20,5) default 0 not null,
  RECHARGE_ACCOUNT      VARCHAR2(32),
  COMPLETE_UP_PAY       NUMBER(1) default 1 not null
);
	

comment on table OMS_ORDER_INFO  is '订单记录';
comment on column OMS_ORDER_INFO.ORDER_ID  is '订单编号';
comment on column OMS_ORDER_INFO.DOWN_CHANNEL_NO  is '下游渠道编号';
comment on column OMS_ORDER_INFO.REQUEST_NO  is '下游渠道订单编号';
comment on column OMS_ORDER_INFO.DOWN_SHELF_ID  is '下游货架编号';
comment on column OMS_ORDER_INFO.DOWN_PRODUCT_ID  is '下游商品编号';
comment on column OMS_ORDER_INFO.EXT_PRODUCT_NO  is '外部商品编号';
comment on column OMS_ORDER_INFO.LINE_ID  is '产品线';
comment on column OMS_ORDER_INFO.CARRIER_NO  is '运营商';
comment on column OMS_ORDER_INFO.PROVINCE_NO is '省份';
comment on column OMS_ORDER_INFO.CITY_NO  is '城市';
comment on column OMS_ORDER_INFO.INVOICE_TYPE  is '开票方式（1.不开发票，0.不限制，2.需要发票）';
comment on column OMS_ORDER_INFO.FACE  is '商品面值';
comment on column OMS_ORDER_INFO.NUM  is '商品数量';
comment on column OMS_ORDER_INFO.TOTAL_FACE  is '商品总面值';
comment on column OMS_ORDER_INFO.SELL_AMOUNT  is '总销售金额';
comment on column OMS_ORDER_INFO.COMMISSION_AMOUNT  is '总佣金金额';
comment on column OMS_ORDER_INFO.SERVICE_AMOUNT  is '总服务费金额';
comment on column OMS_ORDER_INFO.FEE_AMOUNT  is '总手续费金额';
comment on column OMS_ORDER_INFO.CAN_SPLIT_ORDER  is '是否拆单';
comment on column OMS_ORDER_INFO.SPLIT_ORDER_FACE  is '拆单面值';
comment on column OMS_ORDER_INFO.CREATE_TIME is '创建时间';
comment on column OMS_ORDER_INFO.ORDER_OVERTIME  is '订单超时时间';
comment on column OMS_ORDER_INFO.DELIVERY_PAUSE  is '发货暂停（0.是，1否）';
comment on column OMS_ORDER_INFO.ORDER_STATUS  is '订单状态（10.支付，20.绑定发货，0.成功，90.失败，91.部分成功）';
comment on column OMS_ORDER_INFO.PAYMENT_STATUS  is '支付状态（0支付成功，10.未开始，20.等待支付，30.正在支付，90.支付超时）';
comment on column OMS_ORDER_INFO.DELIVERY_BIND_STATUS  is '发货绑定状态（0发货成功，10.未开始，20.等待绑定，30.正在发货，90.全部失败）';
comment on column OMS_ORDER_INFO.REFUND_STATUS  is '订单失败退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99.无需退款）';
comment on column OMS_ORDER_INFO.NOTIFY_STATUS  is '订单信息告知状态（0通知成功，100查询成功，10.未开始，20.等待告知，30.正在告知）';
comment on column OMS_ORDER_INFO.IS_REFUND  is '用户退款（0.是，1否）';
comment on column OMS_ORDER_INFO.BIND_FACE  is '成功绑定总面值';
comment on column OMS_ORDER_INFO.SUCCESS_FACE  is '实际成功总面值';
comment on column OMS_ORDER_INFO.SUCCESS_SELL_AMOUNT  is '实际成功总销售金额 （1）';
comment on column OMS_ORDER_INFO.SUCCESS_COMMISSION  is '实际成功总佣金金额 （2）';
comment on column OMS_ORDER_INFO.SUCCESS_SERVICE  is '实际成功总服务费金额 （3）';
comment on column OMS_ORDER_INFO.SUCCESS_FEE  is '实际成功总手续费金额 （4）';
  comment on column OMS_ORDER_INFO.SUCCESS_COST_AMOUNT  is '实际发货成功总成本 （5）';
comment on column OMS_ORDER_INFO.SUCCESS_UP_COMMISSION  is '实际发货成功总上游佣金 （6）';
comment on column OMS_ORDER_INFO.SUCCESS_UP_SERVICE  is '实际发货成功总上游服务费 （7）';
comment on column OMS_ORDER_INFO.PROFIT  is '利润（1-2+3-4-5+6+7）';
comment on column OMS_ORDER_INFO.RECHARGE_ACCOUNT  is '充值账户';
comment on column OMS_ORDER_INFO.COMPLETE_UP_PAY  is '已完成上游支付（0.已完成，1.未完成）';

 
	alter table oms_order_info
	add constraint pk_oms_order_info primary key(order_id);
	alter table oms_order_info
	add constraint oms_order_info_down_channel_no unique(down_channel_no,request_no);
	
	create sequence seq_order_info_id
	minvalue 1100000000
	maxvalue 99999999999
	start with 1100000000
	cache 20;
	