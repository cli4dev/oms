drop table OMS_PRODUCT_LINE;
drop sequence seq_product_line_id;

	create table OMS_PRODUCT_LINE
(
  LINE_ID                NUMBER(10) default 1 not null,
  LINE_NAME              VARCHAR2(64) not null,
  CAN_PACKAGE_DELIVERY   NUMBER(1) default 1 not null,
  PAYMENT_QUEUE          VARCHAR2(32) default 'oms:order:pay',
  BIND_QUEUE             VARCHAR2(32) default 'oms:order:bind',
  REFUND_QUEUE           VARCHAR2(32) default 'oms:refund:pay',
  NOTIFY_QUEUE           VARCHAR2(32) default 'oms:order:notify',
  UP_PAYMENT_QUEUE       VARCHAR2(32) default 'oms:order:up_pay',
  UP_REFUND_QUEUE        VARCHAR2(32) default 'oms:refund:up_pay',
  REFUND_NOTIFY_QUEUE    VARCHAR2(32) default 'oms:refund:notify',
  ORDER_REFUND_QUEUE     VARCHAR2(32) default 'oms:overtime:refund',
  ORDER_OVERTIME_QUEUE   VARCHAR2(32) default 'oms:overtime:order_deal',
  REFUND_OVERTIME_QUEUE  VARCHAR2(32) default 'oms:overtime:refund_deal',
  DELIVERY_UNKNOWN_QUEUE VARCHAR2(32) default 'oms:overtime:delivery_unknown',
  RETURN_UNKNOWN_QUEUE   VARCHAR2(32) default 'oms:overtime:return_unknown',
  DELIVERY_START_QUEUE   VARCHAR2(32) default 'oms:order:delivery',
  DELIVERY_FINISH_QUEUE  VARCHAR2(32) default 'oms:order:delivery_finish',
  RETURN_QUEUE           VARCHAR2(32) default 'oms:refund:return',
  RETURN_FINISH_QUEUE    VARCHAR2(32) default 'oms:refund:return_complete'
);
comment on table OMS_PRODUCT_LINE  is '产品线';
comment on column OMS_PRODUCT_LINE.LINE_ID  is '产品线编号';
comment on column OMS_PRODUCT_LINE.LINE_NAME  is '产品线名称';
comment on column OMS_PRODUCT_LINE.CAN_PACKAGE_DELIVERY  is '支持打包发货';
comment on column OMS_PRODUCT_LINE.PAYMENT_QUEUE  is '支付队列';
comment on column OMS_PRODUCT_LINE.BIND_QUEUE  is '绑定队列';
comment on column OMS_PRODUCT_LINE.REFUND_QUEUE  is '退款队列';
comment on column OMS_PRODUCT_LINE.NOTIFY_QUEUE  is '通知队列';
comment on column OMS_PRODUCT_LINE.UP_PAYMENT_QUEUE  is '上游支付队列';
comment on column OMS_PRODUCT_LINE.UP_REFUND_QUEUE  is '上游退款队列';
comment on column OMS_PRODUCT_LINE.REFUND_NOTIFY_QUEUE  is '退款通知队列';
comment on column OMS_PRODUCT_LINE.ORDER_REFUND_QUEUE  is '订单失败退款队列';
comment on column OMS_PRODUCT_LINE.ORDER_OVERTIME_QUEUE  is '订单超时处理队列';
comment on column OMS_PRODUCT_LINE.REFUND_OVERTIME_QUEUE  is '退款超时处理队列';
comment on column OMS_PRODUCT_LINE.DELIVERY_UNKNOWN_QUEUE  is '发货未知处理队列';
comment on column OMS_PRODUCT_LINE.RETURN_UNKNOWN_QUEUE  is '退货未知处理队列';
comment on column OMS_PRODUCT_LINE.DELIVERY_START_QUEUE is '发货开始队列';
comment on column OMS_PRODUCT_LINE.DELIVERY_FINISH_QUEUE  is '发货结束队列';
comment on column OMS_PRODUCT_LINE.RETURN_QUEUE  is '退货队列';
comment on column OMS_PRODUCT_LINE.RETURN_FINISH_QUEUE  is '退货结束队列';

 
	alter table oms_product_line
	add constraint pk_product_line primary key(line_id);
	
	create sequence seq_product_line_id
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;
	