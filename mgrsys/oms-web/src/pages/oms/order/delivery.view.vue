<template>
  <div>
    <el-tabs v-model="tabName" type="border-card" @tab-click="handleClick">
      <el-tab-pane label="发货详情" name="1">
        <div class="table-responsive">
          <table :date="info" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.delivery_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.create_time}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">省份:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{ProvinceName(info.province_no)}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">城市:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{CityName(info.city_no) }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货成本:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.cost_amount|AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">运营商:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.carrier_no | EnumFilter("CarrierNo")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(info.delivery_status)"
                    >{{info.delivery_status | EnumFilter("DeliveryStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货服务费:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.service_amount |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游渠道:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.down_channel_no | EnumFilter("DownChannelNo")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.down_product_id }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商品面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.face|AmountFilter }}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">开票方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.invoice_type | EnumFilter("InvoiceType")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.line_id | EnumFilter("LineId")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.num}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.order_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货返回信息:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.return_msg}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">开始时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.start_time}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">结束时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.end_time}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.total_face |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游渠道编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.up_channel_no | EnumFilter("UpChannelNo")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游佣金:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.up_commission_amount |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游发货编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.up_delivery_no }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游商品请求编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.up_ext_product_no |StringFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游支付状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(info.up_payment_status)"
                    >{{info.up_payment_status | EnumFilter("UpPaymentStatus") }}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{info.up_product_id}}</div>
                  </el-col>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
       <el-tab-pane label="上游发货详情" name="2">
        <div class="table-responsive">
          <table :date="item" class="table table-striped m-b-none">
            <tbody class="table-border">
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游发货编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.order_no}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游商户编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.coop_id}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游商户订单号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.coop_order_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游渠道编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.channel_name}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">服务类型:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.service_class| EnumFilter("ServiceClass")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">运营商:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.carrier_no| EnumFilter("CarrierNo")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.product_face/100}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.product_num}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content"  :class="generalStatusClassFilter(item.status)">{{item.status| EnumFilter("DeliveryStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游通知回调地址:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.notify_url}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">请求开始时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.request_start_time}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">请求结束时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.request_finish_time}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">创建时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.create_time}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货结果来源:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.result_source}}</div>
                  </el-col>
                </td>
              </tr>
          
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货结果码:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.result_code}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">结果描述:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.result_desc}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游发货订单号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.up_order_no |StringFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">成功面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.succ_face/100}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">最后更新时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.last_update_time}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">流程超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.flow_timeout}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">发货信息参数:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.request_params}}</div>
                  </el-col>
                
                </td>
              </tr>
              <tr>
                <td>
                 
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游返回数据:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.result_params}}</div>
                  </el-col>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="发货扣款" name="3">
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
      </el-tab-pane>
      <el-tab-pane label="退货记录" name="4">
        <div class="table-responsive">
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
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],
      tabName: "1",
      provinceNo: this.EnumUtility.Get(
        "ProvinceNo",
        { grade: 1 },
        "/oms/canton/info/getlist"
      ),
      cityNo: this.EnumUtility.Get("CityNo", {}, "/oms/canton/info/getlist"),
      info: {},
      item: {},
      UpPayList: [],
      ReturnList: [],
      format: "yyyy/MM/dd hh:mm:ss",
      totalcount: 0,
      pageSizes: [10, 20, 50, 100],
      params: {
        pi: 1,
        ps: 10
      }
    };
  },
  created() {
    this.EnumUtility.Get("CarrierNo");
    this.EnumUtility.Get("DeliveryStatus");
    this.EnumUtility.Get(
      "DownChannelNo",
      {},
      "/oms/down/channel/getdictionary"
    );
    this.EnumUtility.Get("InvoiceType");
    this.EnumUtility.Get("LineId", {}, "/oms/product/line/getdictionary");
    this.EnumUtility.Get("UpChannelNo", {}, "/oms/up/channel/getdictionary");
    this.EnumUtility.Get("UpPaymentStatus");
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
      this.QueryDataList();
    },
    QueryDataList() {
      this.info = {};
      this.$http.get(this.$route.query.getpath, this.$route.query).then(res => {
        this.info = res;
      });
    },
    QueryVdsOrderInfo(){
      this.item ={}
      console.log(this.$route.query,"详情参数")
      this.$http.get("/oms/order/info/queryvds",this.$route.query).then(res =>{
        this.item = res
      })
    },
    handleSizeChange(val) {
      this.params.ps = val;
      this.params.pi = 1;
      this.refreshDataList();
    },
    handleCurrentChange(val) {
      this.params.pi = val;
      this.queryList(this.tabName);
    },
    handleClick(tab) {
      switch (tab.name) {
        case "1":
          this.QueryDataList();
          break;
        case "2":
          this.QueryVdsOrderInfo();
          break;
        case "3":
           this.QueryUpPay();
          break;
        case "4":
           this.QueryReturn();
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
