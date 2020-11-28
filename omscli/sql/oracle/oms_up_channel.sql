drop table oms_up_channel;

	create table oms_up_channel(
		channel_no  varchar2(32)   not null ,
		channel_name  varchar2(64)   not null ,
		status  number(1)     default 0 not null ,
		create_time  date          default sysdate not null 
		);
	

	comment on table oms_up_channel is '上游渠道';
	comment on column oms_up_channel.channel_no is '编号';	
	comment on column oms_up_channel.channel_name is '名称';	
	comment on column oms_up_channel.status is '状态';	
	comment on column oms_up_channel.create_time is '创建时间';	
	

 
	alter table oms_up_channel
	add constraint pk_up_channel primary key(channel_no);
	