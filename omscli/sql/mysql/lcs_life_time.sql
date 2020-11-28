 drop table lcs_life_time;

create table lcs_life_time(
		id        bigint primary key auto_increment  not null    comment 'id',
		order_no  varchar(30)                        not null    comment '业务单据号',
		batch_no   varchar(30)                       null    comment '业务批次号',
		extral_param   varchar(30)                   null    comment '扩展编号',
        content varchar(1000)                        not null    comment '内容',
		ip varchar(20)                               null        comment '服务器ip',
        create_time datetime DEFAULT CURRENT_TIMESTAMP  not null  comment '创建时间' 
  ) comment='生命周期记录表';

alter table lcs_life_time add index index_life_time_orderno(order_no);