
	drop table oms_notify_info;
	drop sequence seq_notify_info_id;

create table oms_notify_info(
		notify_id  number(20)     default 2000 not null ,
		order_id  number(20)      not null ,
		refund_id  number(20)       ,
		notify_type  number(3)       not null ,
		notify_status  number(3)      default 10 not null ,
		notify_count  number(3)      default 0 not null ,
		max_count  number(3)       not null ,
		create_time  date           default sysdate not null ,
		start_time  date             ,
		end_time  date             ,
		notify_url  varchar2(128)   not null ,
		notify_msg  varchar2(256)    
		);
	

	comment on table oms_notify_info is '订单通知表';
	comment on column oms_notify_info.notify_id is '通知编号';	
	comment on column oms_notify_info.order_id is '订单编号';	
	comment on column oms_notify_info.refund_id is '退款编号';	
	comment on column oms_notify_info.notify_type is '通知类型（1.订单通知，2.退款通知）';	
	comment on column oms_notify_info.notify_status is '通知状态（0成功，10未开始,20等待通知，30正在通知）';	
	comment on column oms_notify_info.notify_count is '通知次数';	
	comment on column oms_notify_info.max_count is '最大通知次数';	
	comment on column oms_notify_info.create_time is '创建时间';	
	comment on column oms_notify_info.start_time is '开始时间';	
	comment on column oms_notify_info.end_time is '结束时间';	
	comment on column oms_notify_info.notify_url is '通知地址';	
	comment on column oms_notify_info.notify_msg is '通知结果信息';	
	

 
	alter table oms_notify_info
	add constraint pk_notify_info primary key(notify_id);
	
	create sequence seq_notify_info_id
	minvalue 2000
	maxvalue 99999999999
	start with 2000
	cache 20;
	