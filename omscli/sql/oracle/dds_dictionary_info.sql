drop table dds_dictionary_info;
drop sequence seq_dds_dictionary_info_id;

create table dds_dictionary_info(
    id      number(10)      not null,
    name    varchar2(64)    not null,
    value   varchar2(32)  not null,
    type    varchar2(32)    not null,
    sort_no number(2)       default 0 not null,
    status  number(1)       not null
);
	
comment on table dds_dictionary_info is '字典表';
comment on column dds_dictionary_info.id is '编号 ';	
comment on column dds_dictionary_info.name is '名称';	
comment on column dds_dictionary_info.value is '值';	
comment on column dds_dictionary_info.type is '类型';	
comment on column dds_dictionary_info.sort_no is '排序值';	
comment on column dds_dictionary_info.status is '状态(0:启用,1:禁用)';	
	
alter table dds_dictionary_info
add constraint pk_dds_dictionary_info primary key(id);

create index idx_dictionary_info_type on dds_dictionary_info(type);

create sequence seq_dds_dictionary_info_id
minvalue 1
maxvalue 99999999999
start with 1
cache 20;