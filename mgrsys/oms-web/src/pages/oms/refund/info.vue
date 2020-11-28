<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-date-picker
            clearable
            filterable
            v-model="queryData.start_time"
            @change="TimeShow(queryData.start_time,queryData.end_time)"
            value-format="yyyy-MM-dd"
            placeholder="请输入开始时间"
            style="width:180px"
            size="medium"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-date-picker
            clearable
            v-model="queryData.end_time"
            @change="TimeShow(queryData.start_time,queryData.end_time)"
            value-format="yyyy-MM-dd"
            placeholder="请选择结束时间"
            style="width:180px"
            size="medium"
          >结束时间</el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.carrier_no"
            class="input-cos"
            filterable
            placeholder="请选择运营商"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in carrierNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            filterable
            v-model="queryData.province_no"
            class="input-cos"
            placeholder="请选择省份"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in provinceNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.city_no"
            class="input-cos"
            filterable
            placeholder="请选择城市"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in cityNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.down_channel_no"
            class="input-cos"
            filterable
            placeholder="请选择下游渠道名称"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in downChannelNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.refund_status"
            class="input-cos"
            placeholder="请选择状态"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in refundStatus"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-radio-group v-model="radioOption">
          <el-radio
            v-for="(item, index) in radioList"
            :key="index"
            :label="item.value"
          >{{item.label}}</el-radio>
        </el-radio-group>
        <el-input
          v-model="keyword"
          size="medium"
          style="margin-left:10px;width:150px"
          placeholder="请输入关键字"
        ></el-input>
        <el-form-item>
          <el-button type="primary" @click="query(1)" size="small">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- query end -->

    <!-- list start-->
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="create_time" label="创建时间" align="center" min-width="140px">
          <template slot-scope="scope">
            <span v-if="dayNum !=0">{{scope.row.create_time |DateFilter("yyyy-MM-dd hh:mm:ss") }}</span>
            <span v-else>{{scope.row.create_time | DateFilter("hh:mm:ss")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="refund_id" label="退款编号" align="center"></el-table-column>

        <el-table-column prop="order_id" label="订单编号" align="center">
          <template slot-scope="scope">
            <el-button
              type="text"
              size="small"
              @click="detail(scope.row.order_id)"
            >{{scope.row.order_id}}</el-button>
          </template>
        </el-table-column>
        <el-table-column prop="request_no" label="下游订单号" align="center" min-width="120px">
          <template slot-scope="scope">
            <span>{{scope.row.request_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="down_channel_no" label="下游渠道" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.down_channel_no | EnumFilter("DownChannelNo")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="line_id" label="产品线" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.line_id | EnumFilter("LineId")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="carrier_no" label="运营商" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.carrier_no | EnumFilter("CarrierNo")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="province_no" label="省份" align="center">
          <template slot-scope="scope">
            <span>{{ProvinceName(scope.row.province_no) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="city_no" label="城市" align="center">
          <template slot-scope="scope">
            <span>{{CityName(scope.row.city_no) }}</span>
          </template>
        </el-table-column>

        <el-table-column prop="face" label="面值" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.face |AmountFilter}}</span>
          </template>
        </el-table-column>

        <el-table-column prop="refund_num" label="数量" align="center"></el-table-column>
        <el-table-column prop="refund_face" label="总面值" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.refund_face |AmountFilter}}</span>
          </template>
        </el-table-column>

        <el-table-column prop="refund_type" label="退款方式" align="center">
          <template slot-scope="scope">
            <span
              :class="generalStatusClassFilter(scope.row.return_status)"
            >{{scope.row.refund_type | EnumFilter("RefundType")}}</span>
          </template>
        </el-table-column>

        <el-table-column prop="refund_status" label="状态" align="center">
          <template slot-scope="scope">
            <span
              :class="generalStatusClassFilter(scope.row.refund_status)"
            >{{scope.row.refund_status | EnumFilter("RefundStatus")}}</span>
          </template>
        </el-table-column>

        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>
    <!-- list end-->

    <!-- pagination start -->
    <div class="page-pagination">
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
    <!-- pagination end -->
  </div>
</template>

<script>
export default {
  name: "OmsRefundInfo",
  components: {},
  data() {
    return {
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],
      radioList: [
        { value: 1, label: "退款编号" },
        { value: 2, label: "订单编号" }
      ],
      keyword: "",
      radioOption: 1,
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {
        order_id: "",
        refund_id: "",
        start_time:
          new Date().getFullYear() +
          "-" +
          (new Date().getMonth() + 1) +
          "-" +
          new Date().getDate(),
        end_time:
          new Date().getFullYear() +
          "-" +
          (new Date().getMonth() + 1) +
          "-" +
          new Date().getDate()
      },
      dayNum: 0,
      dayformat: "yyyy-MM-dd",
      carrierNo: this.EnumUtility.Get("CarrierNo"),
      provinceNo: this.EnumUtility.Get(
        "ProvinceNo",
        { grade: 1 },
        "/oms/canton/info/getlist"
      ),
      cityNo: this.EnumUtility.Get("CityNo", {}, "/oms/canton/info/getlist"),
      completeUpRefund: this.EnumUtility.Get("CompleteUpRefund"),
      dtCreateTime: this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
      downChannelNo: this.EnumUtility.Get(
        "DownChannelNo",
        {},
        "/oms/down/channel/getdictionary"
      ),
      downRefundStatus: this.EnumUtility.Get("DownRefundStatus"),
      downShelfId: this.EnumUtility.Get(
        "DownShelfId",
        {},
        "/oms/down/shelf/getdictionary"
      ),
      lineId: this.EnumUtility.Get(
        "LineId",
        {},
        "/oms/product/line/getdictionary"
      ),
      refundNotifyStatus: this.EnumUtility.Get("RefundNotifyStatus"),
      refundStatus: this.EnumUtility.Get("RefundStatus"),
      refundType: this.EnumUtility.Get("RefundType"),
      upReturnStatus: this.EnumUtility.Get("UpReturnStatus"),
      tableData: []
    };
  },
  created() {},
  mounted() {
    this.init();
  },
  methods: {
    TimeShow(start, end) {
      var dateSpan, tempDate, iDays;
      start = Date.parse(start);
      end = Date.parse(end);
      dateSpan = end - start;
      dateSpan = Math.abs(dateSpan);
      iDays = Math.floor(dateSpan / (24 * 3600 * 1000));

      this.dayNum = iDays;
    },
    detail(orderId) {
      let val = { order_id: orderId };
      val.getpath = "/oms/order/info";
      this.$emit("addTab", "详情" + orderId, "/oms/order/info.view", val);
    },
    generalStatusClassFilter(item) {
      switch (item) {
        case "0":
          return this.colorClass[0];
          break;
        case "90":
          return this.colorClass[1];
          break;
        case "30":
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
    /**初始化操作**/
    init() {
      this.query();
    },
    /**查询数据并赋值*/
    query(value) {
      this.queryData["refund_id"] = "";
      this.queryData["order_id"] = "";

      switch (this.radioOption) {
        case 1:
          this.queryData["refund_id"] = this.keyword;
          break;
        case 2:
          this.queryData["order_id"] = this.keyword;
          break;
        default:
          Notification.warning("输入参数错误");
          return;
      }
      if (value == 1) {
        this.queryData.pi = 1;
      } else {
        this.queryData.pi = this.params.pi;
      }
      this.queryData.ps = this.params.ps;

      this.queryData.start_time =
        this.DateConvert(this.dayformat, this.queryData.start_time) +
        " 00:00:00";
      this.queryData.end_time =
        this.DateConvert(this.dayformat, this.queryData.end_time) + " 23:59:59";

      this.$http
        .get("/oms/refund/info/query", this.queryData)
        .then(res => {
          this.tableData = res.data;
          this.totalcount = res.count;
        })
        .catch(err => {
          console.log(err);
        });
    },
    /**改变页容量*/
    handleSizeChange(val) {
      this.params.ps = val;
      this.query();
    },
    /**改变当前页码*/
    handleCurrentChange(val) {
      this.params.pi = val;
      this.query();
    },
    /**重置添加表单*/
    resetForm(formName) {
      this.dialogAddVisible = false;
      this.$refs[formName].resetFields();
    },
    detailShow(val) {
      var data = {
        getpath: "/oms/refund/info",
        refund_id: val.refund_id,
        order_id: val.order_id
      };
      this.$emit(
        "addTab",
        "详情" + val.refund_id,
        "/oms/refund/info.view",
        data
      );
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}
.input-cos {
  width: 150px;
}
</style>
