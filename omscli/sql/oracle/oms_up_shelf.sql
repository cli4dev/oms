drop table OMS_UP_SHELF;
drop sequence seq_up_shelf_id;

	create table OMS_UP_SHELF
(
  SHELF_ID          NUMBER(10) default 1 not null,
  SHELF_NAME        VARCHAR2(64) not null,
  CHANNEL_NO        VARCHAR2(32) not null,
  DELIVERY_OVERTIME NUMBER(10) default 300 not null,
  STATUS            NUMBER(1) default 0 not null,
  CREATE_TIME       DATE default sysdate not null,
  RETURN_OVERTIME   NUMBER(10) default 300 not null
);

comment on table OMS_UP_SHELF  is '上游货架';
comment on column OMS_UP_SHELF.SHELF_ID  is '货架编号';
comment on column OMS_UP_SHELF.SHELF_NAME  is '货架名称';
comment on column OMS_UP_SHELF.CHANNEL_NO  is '渠道编号';
comment on column OMS_UP_SHELF.DELIVERY_OVERTIME  is '发货超时时间';
comment on column OMS_UP_SHELF.STATUS  is '货架状态';
comment on column OMS_UP_SHELF.CREATE_TIME  is '创建时间';
comment on column OMS_UP_SHELF.RETURN_OVERTIME  is '退货超时时间';
	

 
	alter table oms_up_shelf
	add constraint pk_up_shelf primary key(shelf_id);
	
	create sequence seq_up_shelf_id
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;
	