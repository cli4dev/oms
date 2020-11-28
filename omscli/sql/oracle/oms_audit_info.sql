drop table oms_audit_info;
drop sequence seq_audit_info_id;

	create table oms_audit_info(
		audit_id  number(10)     default 2000 not null ,
		order_id  number(20)      not null ,
		refund_id  number(20)       ,
		delivery_id  number(20)       ,
		change_type  number(3)       not null ,
		create_time  date           default sysdate not null ,
		audit_status  number(3)       not null ,
		audit_by  number(10)       ,
		audit_time  date             ,
		audit_msg  varchar2(256)    
		);
	

	comment on table oms_audit_info is '发货人工审核表';
	comment on column oms_audit_info.audit_id is '人工审核编号';	
	comment on column oms_audit_info.order_id is '订单编号';	
	comment on column oms_audit_info.refund_id is '退款编号';	
	comment on column oms_audit_info.delivery_id is '发货记录编号';	
	comment on column oms_audit_info.change_type is '变动类型（1.发货，2.退货，3.订单，4.退款）';	
	comment on column oms_audit_info.create_time is '创建时间';	
	comment on column oms_audit_info.audit_status is '审核状态';	
	comment on column oms_audit_info.audit_by is '审核人';	
	comment on column oms_audit_info.audit_time is '审核时间';	
	comment on column oms_audit_info.audit_msg is '审核信息';	
	

 
	alter table oms_audit_info
	add constraint pk_audit_info primary key(audit_id);
	
	create sequence seq_audit_info_id
	minvalue 2000
	maxvalue 99999999999
	start with 2000
	cache 20;
	