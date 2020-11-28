drop table lcs_life_time;
drop sequence seq_lcs_life_time_id;

create table lcs_life_time(
    id              NUMBER(20) not null,
    order_no        VARCHAR2(30) not null,
    batch_no        VARCHAR2(30)  null,
    extral_param    VARCHAR2(30)  null,
    content         VARCHAR2(1000) not null,
    ip              VARCHAR2(20)  null,
    create_time     DATE default sysdate not null 
);


comment on table lcs_life_time is '生命周期记录表';
comment on column lcs_life_time.id is '编号';
comment on column lcs_life_time.order_no is '业务单据号';
comment on column lcs_life_time.batch_no is '业务批次号';
comment on column lcs_life_time.extral_param is '扩展编号';
comment on column lcs_life_time.content is '操作内容';
comment on column lcs_life_time.ip is '服务器ip';
comment on column lcs_life_time.create_time is '创建时间';


alter table lcs_life_time add constraint pk_lcs_life_time_id primary key (id);

create sequence seq_lcs_life_time_id
    minvalue 1
    maxvalue 9999999999999
    start with 1
    increment by 1
    cache 20;

create index index_life_time_orderno on lcs_life_time(order_no);