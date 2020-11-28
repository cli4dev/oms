
	drop table oms_down_shelf;
	drop sequence seq_down_shelf_id;

	create table oms_down_shelf(
		shelf_id  number(10)    default 1 not null ,
		shelf_name  varchar2(64)   not null ,
		channel_no  varchar2(32)   not null ,
		order_overtime  number(10)     not null ,
		refund_overtime  number(10)     not null ,
		status  number(1)     default 0 not null ,
		create_time  date          default sysdate not null 
		);
	

	comment on table oms_down_shelf is '下游货架';
	comment on column oms_down_shelf.shelf_id is '货架编号';	
	comment on column oms_down_shelf.shelf_name is '货架名称';	
	comment on column oms_down_shelf.channel_no is '渠道编号';	
	comment on column oms_down_shelf.order_overtime is '订单超时时长';	
	comment on column oms_down_shelf.refund_overtime is '退款超时时长';
	comment on column oms_down_shelf.status is '状态';	
	comment on column oms_down_shelf.create_time is '创建时间';	
	

 
	alter table oms_down_shelf
	add constraint pk_down_shelf primary key(shelf_id);
	
	create sequence seq_down_shelf_id
	minvalue 1
	maxvalue 99999999999
	start with 1
	cache 20;
	