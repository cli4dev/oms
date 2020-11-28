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
                    <div class="pull-right" style="margin-right:10px">退款编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_id}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">订单编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.order_id}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">产品线:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.line_id | EnumFilter("LineId")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">运营商:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.carrier_no | EnumFilter("CarrierNo")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">省份:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{ProvinceName(item.province_no)}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">城市:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{CityName(item.city_no)}}</div>
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
                    <div class="pull-right" style="margin-right:10px">已完成上游退款:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      :class="[item.complete_up_refund === '0' ? 'text-success' : 'text-danger']"
                    >{{item.complete_up_refund | EnumFilter("CompleteUpRefund")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游渠道名称:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_channel_no | EnumFilter("DownChannelNo")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_product_id}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游退款状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.down_refund_status)"
                    >{{item.down_refund_status | EnumFilter("DownRefundStatus")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游货架:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.down_shelf_id | EnumFilter("DownShelfId")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">外部商品编号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.ext_product_no|StringFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">商品面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.face|AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>

              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款总佣金金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_commission_amount|AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款商品总面值:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_face |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款总手续费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_fee_amount |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款通知状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.refund_notify_status)"
                    >{{item.refund_notify_status | EnumFilter("RefundNotifyStatus")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款商品数量:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_num}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款总销售金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_sell_amount |AmountFilter}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款总服务费金额:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_service_amount |AmountFilter}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.refund_status)"
                    >{{item.refund_status | EnumFilter("RefundStatus")}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退款方式:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.refund_type | EnumFilter("RefundType")}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">下游渠道订单号:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.request_no}}</div>
                  </el-col>
                </td>
              </tr>
              <tr>
                <td>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">退货超时时间:</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="grid-content">{{item.return_overtime}}</div>
                  </el-col>
                  <el-col :span="6">
                    <div class="pull-right" style="margin-right:10px">上游退货状态:</div>
                  </el-col>
                  <el-col :span="6">
                    <div
                      class="grid-content"
                      :class="generalStatusClassFilter(item.up_return_status)"
                    >{{item.up_return_status | EnumFilter("UpReturnStatus")}}</div>
                  </el-col>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </el-tab-pane>
      <el-tab-pane label="退款记录" name="2">
        <div class="table-responsive">
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
      <el-tab-pane label="退款通知" name="3">
        <div class="table-responsive">
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
      params: {
        pi: 1,
        ps: 10
	  },
	  format: "yyyy/MM/dd hh:mm:ss",
      pageSizes: [10, 20, 50, 100],
      tabName: "1",
      RefundNotifyList: [],
      totalcount: 0,
      RefundList: [],
      tabName: "1",
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],
      provinceNo: this.EnumUtility.Get(
        "ProvinceNo",
        { grade: 1 },
        "/oms/canton/info/getlist"
      ),
      cityNo: this.EnumUtility.Get("CityNo", {}, "/oms/canton/info/getlist"),
      item: {}
    };
  },
  created() {
    this.EnumUtility.Get("CarrierNo");
    this.EnumUtility.Get("CompleteUpRefund");
    this.EnumUtility.Get(
      "DownChannelNo",
      {},
      "/oms/down/channel/getdictionary"
    );
    this.EnumUtility.Get("DownRefundStatus");
    this.EnumUtility.Get("DownShelfId", {}, "/oms/down/shelf/getdictionary");
    this.EnumUtility.Get("LineId", {}, "/oms/product/line/getdictionary");
    this.EnumUtility.Get("RefundNotifyStatus");
    this.EnumUtility.Get("RefundStatus");
    this.EnumUtility.Get("RefundType");
    this.EnumUtility.Get("UpReturnStatus");
  },
  mounted() {
    this.Init();
  },
  methods: {
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
      this.$http.get(this.$route.query.getpath, this.$route.query).then(res => {
        this.item = res;
      });
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
    handleClick(tab) {
      switch (tab.name) {
        case "1":
          this.QueryDataList();
          break;
        case "2":
          this.QueryRefund();
          break;
        case "3":
          this.QueryRefundNotify();
          break;

        default:
          this.$notify({
            title: "警告",
            message: "选项卡错误！"
          });
          return;
      }
    },
    queryList(tabName) {
      switch (tabName) {
        case "1":
          this.QueryDataList();
          break;
        case "2":
          this.QueryRefund();
          break;
        case "3":
          this.QueryRefundNotify();
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
