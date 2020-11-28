<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="订单详情" name="1">
        <div class="table-responsive">
          <table :data="item" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单号：</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.order_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游渠道编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_channel_no | EnumFilter("DownChannelNo")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游渠道订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.request_no}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游货架编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_shelf_id | EnumFilter("DownShelfId")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_product_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.line_id | EnumFilter("LineId")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">运营商:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.carrier_no | EnumFilter("CarrierNo")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">省/市:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{ProvinceName(item.province_no)}} / {{CityName(item.city_no)}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">销售折扣/佣金折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.sell_discount |FeeFilter}}/{{item.commission_discount |FeeFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">服务费折扣/手续费折扣:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.service_discount |FeeFilter}}/{{item.payment_fee_discount |FeeFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">外部商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.ext_product_no |StringFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">开票方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.invoice_type | EnumFilter("InvoiceType")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商品面值/数量/总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.face |AmountFilter}} / {{item.num}} / {{item.total_face |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">总销售/佣金/服务费/手续费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.sell_amount |AmountFilter}} / {{item.commission_amount |AmountFilter}} / {{item.service_amount |AmountFilter}} / {{item.fee_amount |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">是否拆单/拆单面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.can_split_order | EnumFilter("CanSplitOrder")}}/{{item.split_order_face |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.create_time }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.order_overtime}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货暂停:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="[item.delivery_pause === '0' ? 'text-success' : 'text-danger']"
                    >{{item.delivery_pause | EnumFilter("DeliveryPause")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.order_status)"
                    >{{item.order_status | EnumFilter("OrderStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">支付状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.payment_status)"
                    >{{item.payment_status | EnumFilter("PaymentStatus")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货绑定状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.delivery_bind_status)"
                    >{{item.delivery_bind_status | EnumFilter("DeliveryBindStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">已完成上游支付:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.complete_up_pay)"
                    >{{item.complete_up_pay | EnumFilter("CompleteUpPay")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单信息告知状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.notify_status)"
                    >{{item.notify_status | EnumFilter("NotifyStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单失败退款状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.refund_status)"
                    >{{item.refund_status | EnumFilter("OrderFailRefundStatus")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">用户退款:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="[item.status === '0' ? 'text-success' : 'text-danger']"
                    >{{item.is_refund | EnumFilter("IsRefund")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">成功绑定/实际成功总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.bind_face|AmountFilter}}/ {{item.success_face |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际成功总销售/总佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.success_sell_amount |AmountFilter}}/ {{item.success_commission |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">总服务费/总手续费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.success_service |AmountFilter}}/ {{item.success_fee |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">实际发货成功总成本:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.success_cost_amount |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">总上游佣金/总上游服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                    >{{item.success_up_commission |AmountFilter}} / {{item.success_up_service |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">利润:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.profit|AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">充值账户:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.rechage_account}}</div>
                  </el-col>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="付款记录/退款记录 " name="2">
        <div class="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">变动编号</th>
                <th style="text-align:center">账户编号</th>
                <th style="text-align:center">交易编号</th>
                <th style="text-align:center">扣款编号</th>
                <th style="text-align:center">交易类型</th>
                <th style="text-align:center">变动类型</th>
                <th style="text-align:center">变动金额</th>
                <th style="text-align:center">账户余额</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in DownPayList" :key="index">
                <td style="text-align:center;width:90px">{{item.record_id}}</td>
                <td style="text-align:center">{{item.account_id}}</td>
                <td style="text-align:center">{{item.trade_no}}</td>
                <td style="text-align:center">{{item.deduct_no}}</td>
                <td style="text-align:center">{{item.trade_type | EnumFilter("TradeType")}}</td>
                <td style="text-align:center">{{item.change_type | EnumFilter("ChangeType")}}</td>
                <td style="text-align:center">{{item.amount/100 |AmountFilter}}</td>
                <td style="text-align:center">{{item.balance/100 |AmountFilter}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="height-position">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">退款编号</th>
                <th style="text-align:center">订单编号</th>
                <th style="text-align:center">下游渠道</th>
                <th style="text-align:center">下游渠道订单号</th>
                <th style="text-align:center">下游退款编号</th>
                <th style="text-align:center">下游货架编号</th>
                <th style="text-align:center">下游商品编号</th>
                <th style="text-align:center">外部商品编号</th>
                <th style="text-align:center">产品线</th>
                <th style="text-align:center">运营商</th>
                <th style="text-align:center">省/市</th>
                <th style="text-align:center">状态</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in RefundList" :key="index">
                <td style="text-align:center;width:90px">{{item.refund_id}}</td>
                <td style="text-align:center">{{item.order_id}}</td>
                <td style="text-align:center">{{item.down_channel_no | EnumFilter("DownChannelNo")}}</td>
                <td style="text-align:center">{{item.request_no}}</td>
                <td style="text-align:center">{{item.down_refund_no}}</td>
                <td style="text-align:center">{{item.down_shelf_id | EnumFilter("DownShelfId")}}</td>
                <td style="text-align:center">{{item.down_product_id}}</td>
                <td style="text-align:center">{{item.ext_product_no |StringFilter}}</td>
                <td style="text-align:center">{{item.line_id | EnumFilter("LineId")}}</td>
                <td style="text-align:center">{{item.carrier_no | EnumFilter("CarrierNo")}}</td>
                <td
                  style="text-align:center"
                >{{ProvinceName(item.province_no)}} / {{CityName(item.city_no)}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.refund_status)"
                >{{item.refund_status | EnumFilter("RefundStatus")}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="通知记录/退款通知" name="3">
        <div class="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">通知编号</th>
                <th style="text-align:center">订单编号</th>
                <th style="text-align:center">退款编号</th>
                <th style="text-align:center">通知类型</th>
                <th style="text-align:center">通知状态</th>
                <th style="text-align:center">通知次数</th>
                <th style="text-align:center">最大通知次数</th>
                <th style="text-align:center">创建时间</th>
                <th style="text-align:center">开始时间</th>
                <th style="text-align:center">结束时间</th>
                <th style="text-align:center">通知地址</th>
                <th style="text-align:center">通知结果信息</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in NotifyList" :key="index">
                <td style="text-align:center;width:90px">{{item.notify_id}}</td>
                <td style="text-align:center">{{item.order_id}}</td>
                <td style="text-align:center">{{item.refund_id}}</td>
                <td style="text-align:center">{{item.notify_type| EnumFilter("NotifyType")}}</td>
                <td style="text-align:center">{{item.notify_status | EnumFilter("NotifyStatus")}}</td>
                <td style="text-align:center">{{item.notify_count}}</td>
                <td style="text-align:center">{{item.max_count}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.start_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.end_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.notify_url |StringFilter}}</td>
                <td style="text-align:center">{{item.notify_msg |StringFilter}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="height-position">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">通知编号</th>
                <th style="text-align:center">订单编号</th>
                <th style="text-align:center">退款编号</th>
                <th style="text-align:center">通知类型</th>
                <th style="text-align:center">通知次数</th>
                <th style="text-align:center">最大通知次数</th>
                <th style="text-align:center">开始时间</th>
                <th style="text-align:center">结束时间</th>
                <th style="text-align:center">通知地址</th>
                <th style="text-align:center">通知结果信息</th>
                <th style="text-align:center">通知状态</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in RefundNotifyList" :key="index">
                <td style="text-align:center;width:90px">{{item.notify_id}}</td>
                <td style="text-align:center">{{item.order_id}}</td>
                <td style="text-align:center">{{item.refund_id}}</td>
                <td style="text-align:center">{{item.notify_type| EnumFilter("NotifyType")}}</td>
                <td style="text-align:center">{{item.notify_count}}</td>
                <td style="text-align:center">{{item.max_count}}</td>
                <td style="text-align:center">{{item.start_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.end_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.notify_url |StringFilter}}</td>
                <td style="text-align:center">{{item.notify_msg |StringFilter}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.notify_status)"
                >{{item.notify_status | EnumFilter("NotifyStatus")}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="发货记录" name="4">
        <div class="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">发货编号</th>
                <th style="text-align:center">上游渠道</th>
                <th style="text-align:center">上游商品编号</th>
                <th style="text-align:center">上游发货编号</th>
                <th style="text-align:center">上游商品请求编号</th>
                <th style="text-align:center">下游渠道</th>
                <th style="text-align:center">下游商品编号</th>
                <th style="text-align:center">运营商</th>
                <th style="text-align:center">省/市</th>
                <th style="text-align:center">发货状态</th>
                <th style="text-align:center">上游支付状态</th>
                <th style="text-align:center">开始时间</th>
                <th style="text-align:center">结束时间</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in DeliveryList" :key="index">
                <td style="text-align:center;width:90px">{{item.delivery_id}}</td>
                <td style="text-align:center">{{item.up_channel_no| EnumFilter("UpChannelNo")}}</td>
                <td style="text-align:center">{{item.up_product_id}}</td>
                <td style="text-align:center">{{item.up_delivery_no}}</td>
                <td style="text-align:center">{{item.up_ext_product_no |StringFilter}}</td>
                <td style="text-align:center">{{item.down_channel_no | EnumFilter("DownChannelNo")}}</td>
                <td style="text-align:center">{{item.down_product_id}}</td>
                <td style="text-align:center">{{item.carrier_no | EnumFilter("CarrierNo")}}</td>
                <td
                  style="text-align:center"
                >{{ProvinceName(item.province_no) }} / {{CityName(item.city_no) }}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.delivery_status)"
                >{{item.delivery_status | EnumFilter("DeliveryStatus")}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.up_payment_status)"
                >{{item.up_payment_status | EnumFilter("UpPaymentStatus")}}</td>
                <td style="text-align:center">{{item.start_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.end_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="发货扣款/退货记录" name="5">
        <div class="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">变动编号</th>
                <th style="text-align:center">账户编号</th>
                <th style="text-align:center">交易编号</th>
                <th style="text-align:center">扣款编号</th>
                <th style="text-align:center">交易类型</th>
                <th style="text-align:center">变动类型</th>
                <th style="text-align:center">变动金额</th>
                <th style="text-align:center">账户余额</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in UpPayList" :key="index">
                <td style="text-align:center;width:90px">{{item.record_id}}</td>
                <td style="text-align:center">{{item.account_id}}</td>
                <td style="text-align:center">{{item.trade_no}}</td>
                <td style="text-align:center">{{item.deduct_no}}</td>
                <td style="text-align:center">{{item.trade_type | EnumFilter("TradeType")}}</td>
                <td style="text-align:center">{{item.change_type | EnumFilter("ChangeType")}}</td>
                <td style="text-align:center">{{item.amount/100 |AmountFilter}}</td>
                <td style="text-align:center">{{item.balance/100 |AmountFilter}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
        <div class="height-position">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">退货编号</th>
                <th style="text-align:center">上游渠道</th>
                <th style="text-align:center">上游商品编号</th>
                <th style="text-align:center">上游退货编号</th>
                <th style="text-align:center">上游商品请求编号</th>
                <th style="text-align:center">订单编号</th>
                <th style="text-align:center">发货编号</th>
                <th style="text-align:center">退款编号</th>
                <th style="text-align:center">下游渠道</th>
                <th style="text-align:center">下游商品</th>
                <th style="text-align:center">运营商</th>
                <th style="text-align:center">省/市</th>
                <th style="text-align:center">退货状态</th>
                <th style="text-align:center">退款状态</th>
                <th style="text-align:center">创建时间</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in ReturnList" :key="index">
                <td style="text-align:center;width:90px">{{item.return_id}}</td>
                <td style="text-align:center">{{item.up_channel_no | EnumFilter("UpChannelNo")}}</td>
                <td style="text-align:center">{{item.up_product_id}}</td>
                <td style="text-align:center">{{item.up_return_no |StringFilter}}</td>
                <td style="text-align:center">{{item.up_ext_product_no |StringFilter}}</td>
                <td style="text-align:center">{{item.order_id}}</td>
                <td style="text-align:center">{{item.delivery_id}}</td>
                <td style="text-align:center">{{item.refund_id}}</td>
                <td style="text-align:center">{{item.down_channel_no | EnumFilter("DownChannelNo")}}</td>
                <td style="text-align:center">{{item.down_product_id}}</td>
                <td style="text-align:center">{{item.carrier_no | EnumFilter("CarrierNo")}}</td>
                <td
                  style="text-align:center"
                >{{ProvinceName(item.province_no) }} / {{CityName(item.city_no )}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.return_status)"
                >{{item.return_status | EnumFilter("ReturnStatus")}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.up_refund_status)"
                >{{item.up_refund_status | EnumFilter("UpRefundStatus")}}</td>
                <td style="text-align:center">{{item.create_time | DateFilter(format)}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="审核记录" name="6">
        <div clas="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">审核编号</th>
                <th style="text-align:center">订单编号</th>
                <th style="text-align:center">退款编号</th>
                <th style="text-align:center">发货记录编号</th>
                <th style="text-align:center">审核类型</th>
                <th style="text-align:center">创建时间</th>
                <th style="text-align:center">审核状态</th>
                <th style="text-align:center">审核人</th>
                <th style="text-align:center">审核时间</th>
                <th style="text-align:center">审核信息</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in AuditList" :key="index">
                <td style="text-align:center;width:90px">{{item.audit_id}}</td>
                <td style="text-align:center">{{item.order_id}}</td>
                <td style="text-align:center">{{item.refund_id}}</td>
                <td style="text-align:center">{{item.delivery_id}}</td>
                <td style="text-align:center">{{item.change_type|EnumFilter("AuditType")}}</td>
                <td style="text-align:center">{{item.create_time|DateFilter(format)}}</td>
                <td
                  style="text-align:center"
                  :class="generalStatusClassFilter(item.audit_status)"
                >{{item.audit_status | EnumFilter("AuditStatus")}}</td>
                <td style="text-align:center">{{item.audit_by |StringFilter}}</td>
                <td style="text-align:center">{{item.audit_time | DateFilter(format)}}</td>
                <td style="text-align:center">{{item.audit_msg |StringFilter}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="订单生命周期" name="7">
        <div clas="table-responsive">
          <table class="table table-striped m-b-none">
            <thead>
              <tr>
                <th style="text-align:center">操作编号</th>
                <th style="text-align:center">业务单据号</th>
                <th style="text-align:center">服务器ip</th>
                <th style="text-align:center">创建时间</th>
                <th style="text-align:center">操作内容</th>
              </tr>
            </thead>
            <tbody class="table-border">
              <tr v-for="(item, index) in LifetimeList" :key="index">
                <td style="text-align:center;width:90px">{{item.id}}</td>
                <td style="text-align:center">{{item.order_no}}</td>
                <td style="text-align:center">{{item.ip}}</td>
                <td style="text-align:center">{{item.create_time|DateFilter(format)}}</td>
                <td style="text-align:center">{{item.content}}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
    </el-tabs>
    <div class="page-pagination" v-show="tabName!=1">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="params.pi"
        :page-size="params.ps"
        :page-sizes="pageSizes"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalcount"
      ></el-pagination>
    </div>
  </div>
</template>



<script>
export default {
  data() {
    return {
      tabName: "1",
      item: {},
      format: "yyyy/MM/dd hh:mm:ss",

      // 數據接收
      params: {
        pi: 1,
        ps: 10
      },
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],
      DownPayList: [],
      UpPayList: [],
      pageSizes: [10, 20, 50, 100],
      totalcount: 0,
      List: [],
      NotifyList: [],
      RefundNotifyList: [],
      AuditList: [],
      RefundList: [],
      ReturnList: [],
      LifetimeList: [],
      SuitList: [],
      DeliveryList: [],
      provinceNo: this.EnumUtility.Get(
        "ProvinceNo",
        { grade: 1 },
        "/oms/canton/info/getlist"
      ),
      cityNo: this.EnumUtility.Get("CityNo", {}, "/oms/canton/info/getlist"),
      channelNoList: this.EnumUtility.Get(
        "DownChannelNo",
        {},
        "/oms/down/channel/getdictionary"
      ),
      up_channel_no: this.EnumUtility.Get(
        "UpChannelNo",
        {},
        "/oms/up/channel/getdictionary"
      ),
      line_id: this.EnumUtility.Get(
        "LineId",
        {},
        "/oms/product/line/getdictionary"
      ),
      down_shelf_id: this.EnumUtility.Get(
        "DownShelfId",
        {},
        "/oms/down/shelf/getdictionary"
      ),

      merchantList: [],
      supplierList: [],
      provinceList: [],
      cityList: [],
      sysProductList: [],
      operatorList: [],
      stationList: [],
      shelfList: [],
      StatusList: this.EnumUtility.Get("Status")
    };
  },
  mounted() {
    this.Init();
  },
  methods: {
    generalStatusClassFilter(item) {
      switch (item) {
        case "0":
          return this.colorClass[0];
          break;
        case "90":
          return this.colorClass[1];
          break;
        case "20":
          return this.colorClass[2];
          break;
        default:
          return this.colorClass[3];
          break;
      }
    },
    ProvinceName(value) {
      var res = value;
      if (res == "*") {
        return (res = "全省");
      }
      this.cityNo.forEach(item => {
        if (res == item.value) {
          res = item.name;
        }
      });
      return res;
    },
    CityName(value) {
      var res = value;
      if (res == "*") {
        return (res = "全市");
      }
      this.cityNo.forEach(item => {
        if (res == item.value) {
          res = item.name;
        }
      });
      return res;
    },
    Init() {
      this.QueryOrderDetail();
      this.QueryNotify();
      this.QueryRefundNotify();
      this.QueryDelivery();
      this.QueryRefund();
      this.QueryReturn();
      this.QueryAudit();
      this.QueryLifetime();
      this.QueryDownPay();
      this.QueryUpPay();
    },
    handleSizeChange(val) {
      this.params.ps = val;
      this.params.pi = 1;
      this.refreshDataList();
    },
    refreshDataList() {
      this.params.pi = 1;
      this.queryList(this.tabName);
    },
    handleCurrentChange(val) {
      this.params.pi = val;
      this.queryList(this.tabName);
    },
    QueryOrderDetail() {
      this.item = {};
      this.$http.get(this.$route.query.getpath, this.$route.query).then(res => {
        this.item = res;
      });
    },
    QueryNotify() {
      this.$http
        .post("/oms/order/info/notify", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.NotifyList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryRefundNotify() {
      this.$http
        .post("/oms/order/info/refundnotify", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.RefundNotifyList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryDelivery() {
      this.$http
        .post("/oms/order/info/delivery", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.DeliveryList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryRefund() {
      this.$http
        .post("/oms/order/info/refund", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.RefundList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryDownPay() {
      this.$http
        .post("/oms/order/info/downpay", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.DownPayList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryUpPay() {
      this.$http
        .post("/oms/order/info/uppay", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.UpPayList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryReturn() {
      this.$http
        .post("/oms/order/info/return", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.ReturnList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryAudit() {
      this.$http
        .post("/oms/order/info/audit", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.AuditList = response.data;
          this.totalcount = response.count;
        });
    },
    QueryLifetime() {
      this.$http
        .post("/oms/order/info/lifetime", {
          order_id: this.$route.query.order_id,
          pi: this.params.pi,
          ps: this.params.ps
        })
        .then(response => {
          this.LifetimeList = response.data;
          this.totalcount = response.count;
        });
    },
    queryList(tabName) {
      switch (tabName) {
        case "1":
          this.QueryOrderDetail();
          break;
        case "2":
          this.QueryRefund();
          this.QueryDownPay();
          break;
        case "3":
          this.QueryNotify();
          this.QueryRefundNotify();
          break;
        case "4":
          this.QueryDelivery();
          break;
        case "5":
          this.QueryUpPay();
          this.QueryReturn();
          break;
        case "6":
          this.QueryAudit();
          break;
        case "7":
          this.QueryLifetime();
          break;
        default:
          this.$notify({
            title: "警告",
            message: "选项卡错误！"
          });
          return;
      }
    },
    handleClick(tab) {
      switch (tab.name) {
        case "1":
          this.QueryOrderDetail();
          break;
        case "2":
          this.QueryRefund();
          this.QueryDownPay();
          break;
        case "3":
          this.QueryNotify();
          this.QueryRefundNotify();
          break;
        case "4":
          this.QueryDelivery();
          break;
        case "5":
          this.QueryUpPay();
          this.QueryReturn();
          break;
        case "6":
          this.QueryAudit();
          break;
        case "7":
          this.QueryLifetime();
          break;
        default:
          this.$notify({
            title: "警告",
            message: "选项卡错误！"
          });
          return;
      }
    }
  }
};
</script>

<style>
.demo-table-expand {
  font-size: 0;
}
.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}
.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}

.height-position {
  margin-top: 150px;
}
</style>
