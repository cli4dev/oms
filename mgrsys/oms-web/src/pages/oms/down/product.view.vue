<template>
  <div class="table-responsive">
    <table :data="item" class="table table-striped m-b-none">
      <tbody class="table-border">
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">商品编号:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.product_id}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">货架编号:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.shelf_id | EnumFilter("ShelfId")}}</div>
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
              <div class="pull-right" style="margin-right:10px">支持退款:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.can_refund | EnumFilter("CanRefund")}}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">开票方式:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.invoice_type | EnumFilter("InvoiceType")}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">是否允许拆单:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.can_split_order | EnumFilter("CanSplitOrder")}}</div>
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
              <div class="pull-right" style="margin-right:10px">外部商品编号:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.ext_product_no}}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">面值:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.face | AmountFilter}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">单次最大购买数量:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.limit_count}}</div>
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
              <div class="pull-right" style="margin-right:10px">省份:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{ProvinceName(item.province_no) }}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">城市:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{CityName(item.city_no )}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">手续费折扣:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.payment_fee_discount|FeeFilter}}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">佣金折扣:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.commission_discount |FeeFilter}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">销售折扣:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.sell_discount |FeeFilter}}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">服务费折扣:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.service_discount|AmountFilter}}</div>
            </el-col>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">拆单面值:</div>
            </el-col>
            <el-col :span="6">
              <div class="grid-content">{{item.split_order_face |AmountFilter}}</div>
            </el-col>
          </td>
        </tr>
        <tr>
          <td>
            <el-col :span="6">
              <div class="pull-right" style="margin-right:10px">状态:</div>
            </el-col>
            <el-col :span="6">
              <div
                class="grid-content"
                :class="[item.status === '0' ? 'text-success' : 'text-danger']"
              >{{item.status | EnumFilter("Status")}}</div>
            </el-col>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>


<script>
export default {
  data() {
    return {
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
    this.EnumUtility.Get("ShelfId", {}, "/oms/down/shelf/getdictionary");
    this.EnumUtility.Get("LineId", {}, "/oms/product/line/getdictionary");
    this.EnumUtility.Get("CanRefund");
    this.EnumUtility.Get("InvoiceType");
    this.EnumUtility.Get("CanSplitOrder");
    this.EnumUtility.Get("CarrierNo");

    this.EnumUtility.Get("Status");
  },
  mounted() {
    this.Init();
  },
  watch: {
    "$route.fullPath"(val) {
      if (this.$route.query.getpath != null) {
        this.Init();
      }
    }
  },
  methods: {
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

    handleClick(tab) {}
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
</style>
