###  产品线[oms_product_line]

| 字段名                 | 类型         |            默认值             | 为空  |     约束     | 描述             |
| ---------------------- | ------------ | :---------------------------: | :---: | :----------: | :--------------- |
| line_id                | number(10)   |               1               |  否   | PK,SEQ,LR,DI | 产品线编号       |
| line_name              | varchar2(64) |                               |  否   |   LCRUQ,DN   | 产品线名称       |
| can_package_delivery   | number(1)    |               1               |  否   |   LCRUQ,SL   | 支持打包发货     |
| payment_queue          | varchar2(32) |         oms:order:pay         |  是   |     LCRU     | 支付队列         |
| bind_queue             | varchar2(32) |        oms:order:bind         |  是   |     LCRU     | 绑定队列         |
| refund_queue           | varchar2(32) |        oms:refund:pay         |  是   |     LCRU     | 退款队列         |
| notify_queue           | varchar2(32) |       oms:order:notify        |  是   |     LCRU     | 通知队列         |
| up_payment_queue       | varchar2(32) |       oms:order:up_pay        |  是   |     LCRU     | 上游支付队列     |
| up_refund_queue        | varchar2(32) |       oms:refund:up_pay       |  是   |     LCRU     | 上游退款队列     |
| refund_notify_queue    | varchar2(32) |       oms:refund:notify       |  是   |     LCRU     | 退款通知队列     |
| order_refund_queue     | varchar2(32) |      oms:overtime:refund      |  是   |     LCRU     | 订单失败退款队列 |
| order_overtime_queue   | varchar2(32) |    oms:overtime:order_deal    |  是   |     LCRU     | 订单超时处理队列 |
| refund_overtime_queue  | varchar2(32) |   oms:overtime:refund_deal    |  是   |     LCRU     | 退款超时处理队列 |
| delivery_unknown_queue | varchar2(32) | oms:overtime:delivery_unknown |  是   |     LCRU     | 发货未知处理队列 |
| return_unknown_queue   | varchar2(32) |  oms:overtime:return_unknown  |  是   |     LCRU     | 退货未知处理队列 |
| delivery_start_queue   | varchar2(32) |      oms:order:delivery       |  是   |     LCRU     | 发货开始队列     |
| delivery_finish_queue  | varchar2(32) |   oms:order:delivery_finish   |  是   |     LCRU     | 发货结束队列     |
| return_queue           | varchar2(32) |       oms:refund:return       |  是   |     LCRU     | 退货队列         |
| return_finish_queue    | varchar2(32) |  oms:refund:return_complete   |  是   |     LCRU     | 退货结束队列     |

###  字典表[dds_dictionary_info]

| 字段名  | 类型         | 默认值 | 为空  |                 约束                  | 描述   |
| ------- | ------------ | :----: | :---: | :-----------------------------------: | :----- |
| id      | number(10)   |        |  否   |               PK,SEQ,LR               | 编号   |
| name    | varchar2(64) |        |  否   |                 LCRUQ                 | 名称   |
| value   | varchar2(32) |        |  否   |                 LCRU                  | 值     |
| type    | varchar2(32) |        |  否   | LCRUQ,IDX(IDX_DICTIONARY_INFO_TYPE,1) | 类型   |
| sort_no | number(2)    |   0    |  否   |                 LCRU                  | 排序值 |
| status  | number(1)    |        |  否   |                LRUQ,SL                | 状态   |

###  下游货架[oms_down_shelf]

| 字段名          | 类型         | 默认值  | 为空  |            约束            | 描述         |
| --------------- | ------------ | :-----: | :---: | :------------------------: | :----------- |
| shelf_id        | number(10)   |    1    |  否   |        PK,SEQ,RL,DI        | 货架编号     |
| shelf_name      | varchar2(64) |         |  否   |          CRUQL,DN          | 货架名称     |
| channel_no      | varchar2(32) |         |  否   | CRUQL,SL(oms_down_channel) | 渠道编号     |
| order_overtime  | number(10)   |         |  否   |            RUCL            | 订单超时时长 |
| refund_overtime | number(10)   |         |  否   |            RCUL            | 退款超时时长 |
| status          | number(1)    |    0    |  否   |          RUQL,SL           | 状态         |
| create_time     | date         | sysdate |  否   |             RL             | 创建时间     |

###  上游货架[oms_up_shelf]

| 字段名            | 类型         | 默认值  | 为空  |           约束           | 描述         |
| ----------------- | ------------ | :-----: | :---: | :----------------------: | :----------- |
| shelf_id          | number(10)   |    1    |  否   |       PK,SEQ,RL,DI       | 货架编号     |
| shelf_name        | varchar2(64) |         |  否   |         CRUQL,DN         | 货架名称     |
| channel_no        | varchar2(32) |         |  否   | CRQUL,SL(oms_up_channel) | 渠道编号     |
| status            | number(1)    |    0    |  否   |         RUQL,SL          | 货架状态     |
| delivery_overtime | number(10)   |   300   |  否   |           RUCL           | 发货超时时间 |
| return_overtime   | number(10)   |   300   |  否   |           CULR           | 退货超时时间 |
| create_time       | date         | sysdate |  否   |            RL            | 创建时间     |

###  账户信息[beanpay_account_info]

| 字段名       | 类型         | 默认值  | 为空  |     约束     | 描述                     |
| ------------ | ------------ | :-----: | :---: | :----------: | :----------------------- |
| account_id   | number(20)   |         |  否   | PK,SEQ,RL,DI | 帐户编号                 |
| account_name | varchar2(32) |         |  否   |   CRUQL,DN   | 帐户名称                 |
| ident        | varchar2(32) |         |  否   |    CRUQL     | 系统标识                 |
| groups       | varchar2(32) |         |  否   |    CRUQL     | 用户分组                 |
| eid          | varchar2(64) |         |  否   |    CRUQL     | 外部用户账户编号         |
| balance      | number(20)   |    0    |  否   |     CRUL     | 帐户余额(单位：分)       |
| credit       | number(20)   |    0    |  否   |     CRUL     | 信用余额(单位：分)       |
| status       | number(1)    |    0    |  否   |   RUQL,SL    | 账户状态(0：正常 1:锁定) |
| create_time  | date         | sysdate |  否   |      RL      | 创建时间                 |
###  渠道基本信息[vds_channel_info]

|         字段名         |     类型      | 默认值  | 为空  | 约束                     | 描述                                                             |
| :--------------------: | :-----------: | :-----: | :---: | :----------------------- | :--------------------------------------------------------------- |
|           id           |  number(20)   |         |  否   | PK,SEQ,RL                | id                                                               |
|       channel_no       | varchar2(32)  |         |  否   | CRULQ,SL(oms_up_channel) | 渠道编号                                                         |
|     service_class      |   number(8)   |         |  否   | CRULQ,SL                 | 服务类型 (10:加油卡充值 20:电子券发货 21:电子券退货 30:话费充值) |
|      request_url       | varchar2(128) |         |  否   | CRUL                     | 上游请求地址                                                     |
|       notify_url       | varchar2(128) |         |  否   | CRUL                     | 通知回调地址                                                     |
| request_replenish_time |  number(10)   |         |  否   | CRUL                     | 发货后补间隔时间                                                 |
|         status         |   number(1)   |    0    |  否   | RULQ,SL                  | 状态                                                             |
|       can_query        |   number(1)   |    0    |  否   | RUL                      | 是否支持查询                                                     |
|       query_url        | varchar2(128) |         |  是   | RCUL                     | 查询地址                                                         |
|    first_query_time    |  number(10)   |         |  是   | RCUL                     | 首次查询时间                                                     |
|  query_replenish_time  |  number(10)   |         |  是   | RCUL                     | 查询后补间隔时间                                                 |
|       ext_params       | varchar2(256) |         |  是   | RCUL                     | 预留字段                                                         |
|      create_time       |     date      | sysdate |  否   | RL                       | 创建时间                                                         |

###  发货异常订单记录表[vds_order_exp]

| 字段名        | 类型           | 默认值  | 为空  |           约束           | 描述                                                            |
| ------------- | -------------- | :-----: | :---: | :----------------------: | :-------------------------------------------------------------- |
| id            | number(20)     |         |  否   |          PK,RL           | 编号                                                            |
| coop_id       | varchar2(32)   |         |  否   | QRL,SL(oms_down_channel) | 下游商户编号                                                    |
| coop_order_id | varchar2(32)   |         |  否   |            RL            | 下游商户订单号                                                  |
| channel_no    | varchar2(32)   |         |  否   |  QRL,SL(oms_up_channel)  | 上游渠道编号                                                    |
| create_time   | date           | sysdate |  否   |          RL,DL           | 创建时间                                                        |
| service_class | number(8)      |         |  否   |          QRL,SL          | 服务类型(10:加油卡充值 20:电子券发货 21:电子券退货 30:话费充值) |
| carrier_no    | varchar2(8)    |         |  是   |          RLQ,SL          | 运营商                                                          |
| product_face  | number(10)     |         |  是   |            RL            | 产品面值                                                        |
| product_num   | number(8)      |         |  是   |            RL            | 产品数量                                                        |
| error_msg     | varchar2(256)  |         |  是   |            RL            | 订单结果消息                                                    |
| user_ip       | varchar2(32)   |         |  否   |            RL            | 用户Ip                                                          |
| local_ip      | varchar2(32)   |         |  否   |            RL            | 收单Ip                                                          |
| ext_params    | varchar2(4000) |         |  否   |            RL            | 原串                                                            |

###  生命周期记录表[lcs_life_time]

| 字段名       | 类型           | 默认值  | 为空  | 约束  | 描述       |
| ------------ | -------------- | :-----: | :---: | :---: | :--------- |
| id           | number(20)     |         |  否   | PK,RL | 编号       |
| order_no     | varchar2(30)   |         |  否   |  RQL  | 业务单据号 |
| batch_no     | varchar2(30)   |         |  是   |  RQL  | 业务批次号 |
| extral_param | varchar2(30)   |         |  是   |  RL   | 扩展编号   |
| content      | varchar2(1000) |         |  否   |  RQL  | 操作内容   |
| ip           | varchar2(20)   |         |  是   |  RQL  | 服务器ip   |
| create_time  | date           | sysdate |  否   | RL,DT | 创建时间   |

###  订单发货表[oms_order_delivery]

| 字段名               | 类型          | 默认值  | 为空  |           约束           | 描述                                                                      |
| -------------------- | ------------- | :-----: | :---: | :----------------------: | :------------------------------------------------------------------------ |
| delivery_id          | number(20)    |  20000  |  否   |          PK,QRL          | 发货编号                                                                  |
| up_channel_no        | varchar2(32)  |         |  否   |  RL,SL(oms_up_channel)   | 上游渠道编号                                                              |
| up_product_id        | number(10)    |         |  否   |            RL            | 上游商品编号                                                              |
| up_delivery_no       | varchar2(32)  |         |  是   |            RL            | 上游发货编号                                                              |
| up_ext_product_no    | varchar2(32)  |         |  是   |            RL            | 上游商品请求编号                                                          |
| order_id             | number(20)    |         |  否   |            RL            | 订单编号                                                                  |
| down_channel_no      | varchar2(32)  |         |  否   | RQL,SL(oms_down_channel) | 下游渠道编号                                                              |
| down_product_id      | number(10)    |         |  否   |            RL            | 下游商品编号                                                              |
| line_id              | number(10)    |         |  否   | RQL,SL(oms_product_line) | 产品线                                                                    |
| carrier_no           | varchar2(8)   |         |  否   |          QRL,SL          | 运营商                                                                    |
| province_no          | varchar2(8)   |         |  否   | QRL,SL(oms_canton_info)  | 省份                                                                      |
| city_no              | varchar2(8)   |         |  否   | QRL,SL(oms_canton_info)  | 城市                                                                      |
| invoice_type         | number(3)     |         |  否   |          RQL,SL          | 开票方式（1.不开发票，2.上游开发票）                                      |
| delivery_status      | number(3)     |   20    |  否   |          RQL,SL          | 发货状态（0.发货成功，20等待发货，30正在发货，90发货失败）                |
| up_payment_status    | number(3)     |   10    |  否   |          RQL,SL          | 上游支付状态（0支付成功，10未开始,20.等待支付，30.正在支付，99.无需支付） |
| create_time          | date          | sysdate |  否   |          RL,DT           | 创建时间                                                                  |
| face                 | number(20,5)  |         |  否   |            RL            | 商品面值                                                                  |
| num                  | number(10)    |         |  否   |            RL            | 发货数量                                                                  |
| total_face           | number(20,5)  |         |  否   |            RL            | 发货总面值                                                                |
| cost_amount          | number(20,5)  |         |  否   |            RL            | 发货成本                                                                  |
| up_commission_amount | number(20,5)  |         |  否   |            RL            | 上游佣金                                                                  |
| service_amount       | number(20,5)  |         |  否   |            RL            | 发货服务费                                                                |
| start_time           | date          |         |  是   |            RL            | 开始时间                                                                  |
| end_time             | date          |         |  是   |            RL            | 结束时间                                                                  |
| return_msg           | varchar2(256) |         |  是   |            RL            | 发货返回信息                                                              |

###  订单记录[oms_order_info]

| 字段名                | 类型         |   默认值   | 为空  |           约束           | 描述                                                                            |
| --------------------- | ------------ | :--------: | :---: | :----------------------: | :------------------------------------------------------------------------------ |
| order_id              | number(20)   | 1100000000 |  否   |          PK,RL           | 订单编号                                                                        |
| down_channel_no       | varchar2(32) |            |  否   | QRL,SL(oms_down_channel) | 下游渠道编号                                                                    |
| request_no            | varchar2(64) |            |  否   |           QRL            | 下游渠道订单编号                                                                |
| down_shelf_id         | number(10)   |            |  否   |  QRL,SL(oms_down_shelf)  | 下游货架编号                                                                    |
| down_product_id       | number(10)   |            |  否   |           QRL            | 下游商品编号                                                                    |
| ext_product_no        | varchar2(32) |            |  是   |           QRL            | 外部商品编号                                                                    |
| line_id               | number(10)   |            |  否   | QRL,SL(oms_product_line) | 产品线                                                                          |
| carrier_no            | varchar2(8)  |            |  否   |          QRL,SL          | 运营商                                                                          |
| province_no           | varchar2(8)  |            |  否   | QRL,SL(oms_canton_info)  | 省份                                                                            |
| city_no               | varchar2(8)  |            |  否   | QRL,SL(oms_canton_info)  | 城市                                                                            |
| invoice_type          | number(3)    |            |  否   |          QRL,SL          | 开票方式（1.不开发票，0.不限制，2.需要发票）                                    |
| face                  | number(20,5) |            |  否   |            RL            | 商品面值                                                                        |
| num                   | number(10)   |            |  否   |            RL            | 商品数量                                                                        |
| total_face            | number(20,5) |            |  否   |            RL            | 商品总面值                                                                      |
| sell_amount           | number(20,5) |            |  否   |            RL            | 总销售金额                                                                      |
| commission_amount     | number(20,5) |            |  否   |            RL            | 总佣金金额                                                                      |
| service_amount        | number(20,5) |            |  否   |            RL            | 总服务费金额                                                                    |
| fee_amount            | number(20,5) |            |  否   |            RL            | 总手续费金额                                                                    |
| can_split_order       | number(1)    |            |  否   |          QRL,SL          | 是否拆单（0.是，1否）                                                           |
| split_order_face      | number(20,5) |            |  否   |            RL            | 拆单面值                                                                        |
| create_time           | date         |  sysdate   |  否   |          RL,DT           | 创建时间                                                                        |
| order_overtime        | date         |            |  否   |          RL,DT           | 订单超时时间                                                                    |
| delivery_pause        | number(1)    |     1      |  否   |          QRL,SL          | 发货暂停（0.是，1否）                                                           |
| order_status          | number(3)    |     10     |  否   |          QRL,SL          | 订单状态（10.支付，20.绑定发货，0.成功，90.失败，91.部分成功）                  |
| payment_status        | number(3)    |     20     |  否   |          QRL,SL          | 支付状态（0支付成功，10.未开始，20.等待支付，30.正在支付，90.支付超时）         |
| delivery_bind_status  | number(3)    |     10     |  否   |          QRL,SL          | 发货绑定状态（0发货成功，10.未开始，20.等待绑定，30.正在发货，90.全部失败）     |
| refund_status         | number(3)    |     10     |  否   |          QRL,SL          | 订单失败退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99.无需退款） |
| notify_status         | number(3)    |     10     |  否   |          QRL,SL          | 订单信息告知状态（0通知成功，100查询成功，10.未开始，20.等待告知，30.正在告知） |
| is_refund             | number(1)    |     1      |  否   |          QRL,SL          | 用户退款（0.是，1否）                                                           |
| bind_face             | number(20,5) |     0      |  否   |            RL            | 成功绑定总面值                                                                  |
| success_face          | number(20,5) |     0      |  否   |            RL            | 实际成功总面值                                                                  |
| success_sell_amount   | number(20,5) |     0      |  否   |            RL            | 实际成功总销售金额 （1）                                                        |
| success_commission    | number(20,5) |     0      |  否   |            RL            | 实际成功总佣金金额 （2）                                                        |
| success_service       | number(20,5) |     0      |  否   |            RL            | 实际成功总服务费金额 （3）                                                      |
| success_fee           | number(20,5) |     0      |  否   |            RL            | 实际成功总手续费金额 （4）                                                      |
| success_cost_amount   | number(20,5) |     0      |  否   |            RL            | 实际发货成功总成本 （5）                                                        |
| success_up_commission | number(20,5) |     0      |  否   |            RL            | 实际发货成功总上游佣金 （6）                                                    |
| success_up_service    | number(20,5) |     0      |  否   |            RL            | 实际发货成功总上游服务费 （7）                                                  |
| profit                | number(20,5) |     0      |  否   |            RL            | 利润（1-2+3-4-5+6+7）                                                           |
| recharge_account      | varchar2(32) |            |  是   |            RL            | 充值账户                                                                        |
| complete_up_pay       | number(1)    |     1      |  否   |          QRL,SL          | 已完成上游支付（0.已完成，1.未完成）                                            |

###  下游渠道[oms_down_channel]

| 字段名       | 类型         | 默认值  | 为空  |     约束     | 描述     |
| ------------ | ------------ | :-----: | :---: | :----------: | :------- |
| channel_no   | varchar2(32) |         |  否   | PK,SEQ,RL,DI | 编号     |
| channel_name | varchar2(64) |         |  否   |   CRUQL,DN   | 名称     |
| status       | number(1)    |    0    |  否   |   RUQL,SL    | 状态     |
| create_time  | date         | sysdate |  否   |    RL,DT     | 创建时间 |

###  账户余额变动信息[beanpay_account_record]

| 字段名      | 类型           | 默认值  | 为空  |             约束             | 描述                                     |
| ----------- | -------------- | :-----: | :---: | :--------------------------: | :--------------------------------------- |
| record_id   | number(20)     |         |  否   |            PK,LR             | 变动编号                                 |
| account_id  | number(20)     |         |  否   | LRQ,SL(beanpay_account_info) | 帐户编号                                 |
| trade_no    | varchar2(32)   |         |  否   |             LRQ              | 交易编号                                 |
| deduct_no   | varchar2(32)   |    0    |  是   |             LRQ              | 扣款编号                                 |
| change_type | number(1)      |         |  否   |            LRQ,SL            | 变动类型 (1:加款 2:提款 3：扣款 4：退款) |
| trade_type  | number(1)      |    1    |  否   |            LRQ,SL            | 交易类型 (1:交易 2：手续费 3:佣金)       |
| amount      | number(20)     |         |  否   |              LR              | 变动金额 (单位：分)                      |
| balance     | number(20)     |         |  否   |              LR              | 帐户余额 (单位：分)                      |
| create_time | date           | sysdate |  否   |              LR              | 创建时间                                 |
| ext         | varchar2(1024) |         |  否   |              LR              | 扩展字段                                 |

###  发货订单信息表[vds_order_info]

| 字段名              | 类型           | 默认值  | 为空  |           约束           | 描述                                                            |
| ------------------- | -------------- | :-----: | :---: | :----------------------: | :-------------------------------------------------------------- |
| order_no            | number(20)     |         |  否   |          PK,LR           | 发货编号                                                        |
| coop_id             | varchar2(32)   |         |  否   | QRL,SL(oms_down_channel) | 下游商户编号                                                    |
| coop_order_id       | varchar2(32)   |         |  否   |           QRL            | 下游商户订单号                                                  |
| channel_no          | varchar2(32)   |         |  否   |  QRL,SL(oms_up_channel)  | 上游渠道编号                                                    |
| service_class       | number(8)      |         |  否   |          QRL,SL          | 服务类型(10:加油卡充值 20:电子券发货 21:电子券退货 30:话费充值) |
| carrier_no          | varchar2(8)    |         |  否   |          LRQ,SL          | 运营商                                                          |
| product_face        | number(10)     |         |  否   |            LR            | 产品面值                                                        |
| product_num         | number(8)      |         |  否   |            LR            | 产品数量                                                        |
| status              | number(2)      |   20    |  否   |          QRL,SL          | 发货状态                                                        |
| notify_url          | varchar2(128)  |         |  否   |            RL            | 下游通知回调地址                                                |
| request_params      | varchar2(2000) |         |  是   |            RL            | 发货信息参数json                                                |
| request_start_time  | date           |         |  是   |            RL            | 请求开始时间                                                    |
| request_finish_time | date           |         |  是   |            RL            | 请求完成时间                                                    |
| create_time         | date           | sysdate |  否   |          RL,DT           | 创建时间                                                        |
| result_source       | number(1)      |         |  是   |          RQL,SL          | 发货结果来源（1：通知，2：查询，3：同步返回）                   |
| result_code         | varchar2(32)   |         |  是   |            LR            | 发货结果码                                                      |
| result_desc         | varchar2(64)   |         |  是   |            LR            | 结果描述                                                        |
| result_params       | varchar2(256)  |         |  是   |            LR            | 记录上游返回的业务数据（如：折扣）                              |
| up_order_no         | varchar2(32)   |         |  是   |            RL            | 上游发货订单号                                                  |
| succ_face           | number(10)     |         |  是   |            RL            | 成功面值                                                        |
| last_update_time    | date           | sysdate |  否   |            RL            | 最后更新时间                                                    |
| flow_timeout        | date           |         |  否   |            RL            | 流程超时时间                                                    |

### 任务表[tsk_system_task]

| 字段名            | 类型          | 默认值  | 为空  |  约束  | 描述         |
| ----------------- | ------------- | :-----: | :---: | :----: | :----------- |
| task_id           | number(20)    |         |  否   | PK,LR  | 编号         |
| name              | varchar2(32)  |         |  否   |  RLQ   | 名称         |
| create_time       | date          | sysdate |  否   | RLQ,DT | 创建时间     |
| last_execute_time | date          |         |  是   |   RL   | 上次执行时间 |
| next_execute_time | date          |         |  否   |   RL   | 下次执行时间 |
| max_execute_time  | date          |         |  否   | RLQ,DT | 执行期限     |
| next_interval     | number(10)    |         |  否   |   RL   | 时间间隔     |
| count             | number(10)    |    0    |  否   |   RL   | 执行次数     |
| status            | number(2)     |         |  否   | RLQ,SL | 状态         |
| batch_id          | number(20)    |         |  是   |   RL   | 执行批次号   |
| queue_name        | varchar2(64)  |         |  否   |  RLQ   | 消息队列     |
| msg_content       | varchar2(256) |         |  是   |   RL   | 消息内容     |

###  退款记录[oms_refund_info]

| 字段名                   | 类型         | 默认值  | 为空  |           约束           | 描述                                                                           |
| ------------------------ | ------------ | :-----: | :---: | :----------------------: | :----------------------------------------------------------------------------- |
| refund_id                | number(20)   |  20000  |  否   |          PK,QLR          | 退款编号                                                                       |
| order_id                 | number(20)   |         |  否   |           QRL            | 订单编号                                                                       |
| down_channel_no          | varchar2(32) |         |  否   | QRL,SL(oms_down_channel) | 下游渠道编号                                                                   |
| request_no               | varchar2(32) |         |  否   |           QRL            | 下游渠道订单号                                                                 |
| down_refund_no           | varchar2(32) |         |  否   |           QRL            | 下游退款编号                                                                   |
| down_shelf_id            | number(10)   |         |  否   |  QRL,SL(oms_down_shelf)  | 下游货架编号                                                                   |
| down_product_id          | number(10)   |         |  否   |           QRL            | 下游商品编号                                                                   |
| ext_product_no           | varchar2(32) |         |  是   |           QRL            | 外部商品编号                                                                   |
| line_id                  | number(10)   |         |  否   | QRL,SL(oms_product_line) | 产品线                                                                         |
| carrier_no               | varchar2(8)  |         |  否   |          QRL,SL          | 运营商                                                                         |
| province_no              | varchar2(8)  |         |  否   | QRL,SL(oms_canton_info)  | 省份                                                                           |
| city_no                  | varchar2(8)  |         |  否   | QRL,SL(oms_canton_info)  | 城市                                                                           |
| refund_type              | number(1)    |         |  否   |          QRL,SL          | 退款方式（1.普通退款，2.强制退款,3.假成功退款）                                |
| face                     | number(20,5) |         |  否   |            RL            | 商品面值                                                                       |
| refund_num               | number(10)   |         |  否   |            RL            | 退款商品数量                                                                   |
| refund_face              | number(20,5) |         |  否   |            RL            | 退款商品总面值                                                                 |
| refund_sell_amount       | number(20,5) |         |  否   |            RL            | 退款总销售金额                                                                 |
| refund_commission_amount | number(20,5) |         |  否   |            RL            | 退款总佣金金额                                                                 |
| refund_service_amount    | number(20,5) |         |  否   |            RL            | 退款总服务费金额                                                               |
| refund_fee_amount        | number(20,5) |         |  否   |            RL            | 退款总手续费金额                                                               |
| create_time              | date         | sysdate |  否   |          RL,DT           | 创建时间                                                                       |
| refund_status            | number(3)    |   10    |  否   |          QRL,SL          | 状态（10.退货，20.退款，0成功，90失败）                                        |
| up_return_status         | number(3)    |   20    |  否   |          QRL,SL          | 上游退货状态（0.退货成功，20.等待退货，30.正在退货，90.退款失败，91.部分退款） |
| down_refund_status       | number(3)    |   10    |  否   |          QRL,SL          | 下游退款状态（0成功，10.未开始，20.等待，30正在，99无需）                      |
| refund_notify_status     | number(3)    |   10    |  否   |          QRL,SL          | 退款通知状态（0成功，100，查询成功，10.未开始，20.等待，30正在，99无需）       |
| return_overtime          | date         |         |  否   |            RL            | 退货超时时间                                                                   |
| complete_up_refund       | number(1)    |    1    |  否   |          QRL,SL          | 已完成上游退款（0.已完成，1.未完成）                                           |

###  上游退货信息表[oms_refund_up_return]

| 字段名                   | 类型          | 默认值  | 为空  |           约束           | 描述                                                                   |
| ------------------------ | ------------- | :-----: | :---: | :----------------------: | :--------------------------------------------------------------------- |
| return_id                | number(20)    |  20000  |  否   |        PK,SEQ,LR         | 退货编号                                                               |
| up_channel_no            | varchar2(32)  |         |  否   |  LQR,SL(oms_up_channel)  | 上游渠道编号                                                           |
| up_product_id            | number(10)    |         |  否   |           LQR            | 上游商品编号                                                           |
| up_return_no             | varchar2(32)  |         |  是   |            LR            | 上游退货编号                                                           |
| up_ext_product_no        | varchar2(32)  |         |  是   |            LR            | 上游商品请求编号                                                       |
| order_id                 | number(20)    |         |  否   |            LR            | 订单编号                                                               |
| delivery_id              | number(20)    |         |  否   |            LR            | 发货编号                                                               |
| refund_id                | number(20)    |         |  否   |            LR            | 退款编号                                                               |
| down_channel_no          | varchar2(32)  |         |  否   | LRQ,SL(oms_down_channel) | 下游渠道编号                                                           |
| down_product_id          | number(10)    |         |  否   |           LRQ            | 下游商品编号                                                           |
| line_id                  | number(10)    |         |  否   | LRQ,SL(oms_product_line) | 产品线                                                                 |
| carrier_no               | varchar2(8)   |         |  否   |          LRQ,SL          | 运营商                                                                 |
| province_no              | varchar2(8)   |    -    |  否   | LRQ,SL(oms_canton_info)  | 省份                                                                   |
| city_no                  | varchar2(8)   |    -    |  否   | LRQ,SL(oms_canton_info)  | 城市                                                                   |  |
| return_status            | number(3)     |   20    |  否   |          LRQ,SL          | 退货状态（0.退货成功，20等待退货，30正在退货，90退货失败）             |
| up_refund_status         | number(3)     |   10    |  否   |          LRQ,SL          | 退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99无需退款） |
| create_time              | date          | sysdate |  否   |            RL            | 创建时间                                                               |
| return_face              | number(20,5)  |         |  否   |            RL            | 商品面值                                                               |
| return_num               | number(10)    |         |  否   |            RL            | 退货数量                                                               |
| return_total_face        | number(20,5)  |         |  否   |            RL            | 退货总面值                                                             |
| return_cost_amount       | number(20,5)  |         |  否   |            RL            | 退回成本                                                               |
| return_commission_amount | number(20,5)  |         |  否   |            RL            | 退回佣金                                                               |
| return_service_amount    | number(20,5)  |         |  否   |            RL            | 退回服务费                                                             |
| start_time               | date          |         |  是   |            RL            | 开始时间                                                               |
| end_time                 | date          |         |  是   |            RL            | 结束时间                                                               |
| return_msg               | varchar2(256) |         |  是   |            RL            | 退货返回信息                                                           |

###  上游商品[oms_up_product]

| 字段名              | 类型         | 默认值  | 为空  |            约束            | 描述                                 |
| ------------------- | ------------ | :-----: | :---: | :------------------------: | :----------------------------------- |
| product_id          | number(10)   |   300   |  否   |         PK,SEQ,LR          | 商品编号                             |
| shelf_id            | number(10)   |         |  否   |   CRUQL,SL(oms_up_shelf)   | 货架编号                             |
| line_id             | number(10)   |         |  否   | CRUQL,SL(oms_product_line) | 产品线                               |
| carrier_no          | varchar2(8)  |         |  否   |          CRUQL,SL          | 运营商                               |
| province_no         | varchar2(8)  |    -    |  否   | CRUQL,SL(oms_canton_info)  | 省份                                 |
| city_no             | varchar2(8)  |    -    |  否   | CRUQL,SL(oms_canton_info)  | 城市                                 |
| invoice_type        | number(3)    |         |  否   |          CRUQL,SL          | 开票方式（1.不开发票，2.上游开发票） |
| ext_product_no      | varchar2(32) |         |  是   |            CRUL            | 外部商品编号                         |
| can_refund          | number(1)    |         |  否   |          CRUQL,SL          | 支持退货 (0.是,1.否)                 |
| face                | number(20,5) |         |  否   |            CRUL            | 面值                                 |
| cost_discount       | number(10,5) |         |  否   |            CRUL            | 成本折扣（以面值算）                 |
| commission_discount | number(10,5) |         |  否   |            CRUL            | 佣金折扣（以面值算）                 |
| service_discount    | number(10,5) |         |  否   |            CRUL            | 服务费折扣                           |
| limit_count         | number(10)   |         |  否   |            CRUL            | 单次最大发货数量                     |
| status              | number(1)    |    0    |  否   |          RUQL,SL           | 状态                                 |
| create_time         | date         | sysdate |  否   |             RL             | 创建时间                             |

###  发货结果查询记录表[vds_order_query]

| 字段名          | 类型          | 默认值  | 为空  |          约束           | 描述                         |
| --------------- | ------------- | :-----: | :---: | :---------------------: | :--------------------------- |
| order_no        | number(20)    |         |  否   |          PK,LR          | 订单号                       |
| coop_id         | varchar2(32)  |         |  否   | LR,SL(oms_down_channel) | 下游商户编号                 |
| channel_no      | varchar2(32)  |         |  否   | LRQ,SL(oms_up_channel)  | 上游渠道                     |
| query_count     | number(4)     |    0    |  否   |           LR            | 查询次数                     |
| status          | number(2)     |   20    |  否   |         LRQ,SL          | 状态(20-等待查询 0-查询结束) |
| last_query_time | date          |         |  是   |           LR            | 最近查询时间                 |
| query_result    | varchar2(128) |         |  是   |           LR            | 查询请求结果                 |
| create_time     | date          | sysdate |  否   |           RL            | 创建时间                     |

###  发货通知记录表[vds_order_notify]

| 字段名             | 类型          | 默认值  | 为空  |           约束           | 描述                                  |
| ------------------ | ------------- | :-----: | :---: | :----------------------: | :------------------------------------ |
| id                 | number(20)    |         |  否   |        PK,SEQ,LR         | 通知编号                              |
| order_no           | number(20)    |         |  否   |           LRQ            | 订单号                                |
| coop_id            | varchar2(32)  |         |  否   | LRQ,SL(oms_down_channel) | 下游商户编号                          |
| coop_order_id      | varchar2(32)  |         |  否   |           LRQ            | 下游商户订单号                        |
| notify_url         | varchar2(128) |         |  否   |            LR            | 下游通知回调地址                      |
| notify_content     | varchar2(512) |         |  否   |            R             | 通知内容（json）                      |
| status             | number(2)     |   20    |  否   |          LRQ,SL          | 通知状态(20等待通知,0通知成功,90失败) |
| result_msg         | varchar2(64)  |         |  是   |            LR            | 通知结果信息                          |
| finish_time        | date          |         |  是   |            LR            | 完成时间                              |
| notify_count       | number(2)     |    0    |  否   |            LR            | 通知次数                              |
| notify_limit_count | number(2)     |   15    |  否   |            LR            | 通知限制次数                          |
| create_time        | date          | sysdate |  否   |            LR            | 创建时间                              |

###  订单通知表[oms_notify_info]

| 字段名        | 类型          | 默认值  | 为空  |   约束    | 描述                                             |
| ------------- | ------------- | :-----: | :---: | :-------: | :----------------------------------------------- |
| notify_id     | number(20)    |  2000   |  否   | PK,SEQ,LR | 通知编号                                         |
| order_id      | number(20)    |         |  否   |    LQR    | 订单编号                                         |
| refund_id     | number(20)    |         |  是   |    LQR    | 退款编号                                         |
| notify_type   | number(3)     |         |  否   |  LRQ,SL   | 通知类型（1.订单通知，2.退款通知）               |
| notify_status | number(3)     |   10    |  否   |  LQR,SL   | 通知状态（0成功,10未开始,20等待通知,30正在通知） |
| notify_count  | number(3)     |    0    |  否   |    LR     | 通知次数                                         |
| max_count     | number(3)     |         |  否   |    LR     | 最大通知次数                                     |
| create_time   | date          | sysdate |  否   |  LRQ,DT   | 创建时间                                         |
| start_time    | date          |         |  是   |    LR     | 开始时间                                         |
| end_time      | date          |         |  是   |    LR     | 结束时间                                         |
| notify_url    | varchar2(128) |         |  否   |    LR     | 通知地址                                         |
| notify_msg    | varchar2(256) |         |  是   |    LR     | 通知结果信息                                     |

###  发货人工审核表[oms_audit_info]

| 字段名       | 类型          | 默认值  | 为空  |   约束    | 描述                                                                  |
| ------------ | ------------- | :-----: | :---: | :-------: | :-------------------------------------------------------------------- |
| audit_id     | number(10)    |  2000   |  否   | PK,SEQ,LR | 人工审核编号                                                          |
| order_id     | number(20)    |         |  否   |    LQR    | 订单编号                                                              |
| refund_id    | number(20)    |         |  是   |    LQR    | 退款编号                                                              |
| delivery_id  | number(20)    |         |  是   |    LQR    | 发货记录编号                                                          |
| change_type  | number(3)     |         |  否   |  LRQ,SL   | 变动类型（1.发货，2.退货，3.订单，4.退款）                            |
| create_time  | date          | sysdate |  否   |    LQR    | 创建时间                                                              |
| audit_status | number(3)     |         |  否   |  LRUQ,SL  | 审核状态(0.审核为成功，20.等待审核，90.审核为失败，80.审核为部分成功) |
| audit_by     | number(10)    |         |  是   |    LRU    | 审核人                                                                |
| audit_time   | date          |         |  是   |    LRU    | 审核时间                                                              |
| audit_msg    | varchar2(256) |         |  是   |    LRU    | 审核信息                                                              |

###  下游商品[oms_down_product]

| 字段名               | 类型         | 默认值  | 为空  |            约束            | 描述                                         |
| -------------------- | ------------ | :-----: | :---: | :------------------------: | :------------------------------------------- |
| product_id           | number(10)   |   300   |  否   |         PK,SEQ,LR          | 商品编号                                     |
| shelf_id             | number(10)   |         |  否   |  CRUQL,SL(oms_down_shelf)  | 货架编号                                     |
| line_id              | number(10)   |         |  否   | CRULQ,SL(oms_product_line) | 产品线                                       |
| carrier_no           | varchar2(8)  |         |  否   |          CRUQL,SL          | 运营商                                       |
| province_no          | varchar2(8)  |    -    |  否   | CRUQL,SL(oms_canton_info)  | 省份                                         |
| city_no              | varchar2(8)  |    -    |  否   | CRUQL,SL(oms_canton_info)  | 城市                                         |
| invoice_type         | number(3)    |         |  否   |          CRUQL,SL          | 开票方式（1.不开发票，0.不限制，2.需要发票） |
| ext_product_no       | varchar2(32) |         |  是   |           CRUQL            | 外部商品编号                                 |
| can_refund           | number(1)    |         |  否   |          CRUQL,SL          | 支持退款(0.是,1否)                           |
| face                 | number(20,5) |         |  否   |            CRUL            | 面值                                         |
| sell_discount        | number(10,5) |         |  否   |            CRUL            | 销售折扣（以面值算）                         |
| commission_discount  | number(10,5) |         |  否   |            CRUL            | 佣金折扣（以面值算）                         |
| service_discount     | number(10,5) |         |  否   |            CRUL            | 服务费折扣                                   |
| payment_fee_discount | number(10,5) |         |  否   |            CRUL            | 手续费折扣（以销售金额算）                   |
| can_split_order      | number(1)    |         |  否   |          CRULQ,SL          | 是否拆单(0.是,1.否)                          |
| split_order_face     | number(20,5) |         |  否   |            CRUL            | 拆单面值                                     |
| limit_count          | number(10)   |         |  否   |            CRUL            | 单次最大购买数量                             |
| status               | number(1)    |    0    |  否   |          RUQL,SL           | 状态(0.是,1.否)                              |
| create_time          | date         | sysdate |  否   |             RL             | 创建时间                                     |

###  渠道错误码[vds_channel_error_code]

| 字段名          | 类型         | 默认值  | 为空  |           约束           | 描述                                                            |
| --------------- | ------------ | :-----: | :---: | :----------------------: | :-------------------------------------------------------------- |
| id              | number(20)   |         |  否   |        PK,SEQ,LR         | 编号                                                            |
| channel_no      | varchar2(32) |         |  否   | CRUQL,SL(oms_up_channel) | 渠道编号                                                        |
| service_class   | number(8)    |         |  否   |         CRUQL,SL         | 服务类型(10:加油卡充值 20:电子券发货 21:电子券退货 30:话费充值) |
| deal_code       | number(2)    |         |  否   |           CRUL           | 处理码                                                          |
| error_code      | varchar2(32) |         |  否   |           CRUL           | 错误码                                                          |
| error_code_desc | varchar2(64) |         |  否   |           CRUL           | 错误码描述                                                      |
| create_time     | date         | sysdate |  否   |            RL            | 创建时间                                                        |
###  上游渠道[oms_up_channel]

| 字段名       | 类型         | 默认值  | 为空  |    约束     | 描述     |
| ------------ | ------------ | :-----: | :---: | :---------: | :------- |
| channel_no   | varchar2(32) |         |  否   | PK,CRULQ,DI | 编号     |
| channel_name | varchar2(64) |         |  否   |  CRULQ,DN   | 名称     |
| status       | number(1)    |    0    |  否   |   RULQ,SL   | 状态     |
| create_time  | date         | sysdate |  否   |     RL      | 创建时间 |

###  省市信息[oms_canton_info]

| 字段名        | 类型         | 默认值 | 为空  |   约束   | 描述         |
| ------------- | ------------ | :----: | :---: | :------: | :----------- |
| canton_code   | varchar2(32) |        |  否   | PK,CRULQ | 区域编号     |
| chinese_name  | varchar2(32) |        |  否   | CRULQ,DN | 中文名称     |
| spell         | varchar2(64) |        |  否   |   CRUL   | 英文或全拼   |
| grade         | number(1)    |        |  否   |   CRUL   | 行政级别     |
| parent        | varchar2(32) |        |  否   |  CRULQ   | 父级         |
| simple_spell  | varchar2(8)  |        |  否   | CRULQ,DI | 简拼         |
| area_code     | varchar2(8)  |        |  否   |   CRUL   | 区号         |
| standard_code | number(6)    |        |  否   |   CRUL   | 标准行政编码 |