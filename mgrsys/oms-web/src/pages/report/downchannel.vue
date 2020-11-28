<template>
  <div class="panel panel-default">
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-date-picker
            class="input-cos"
            v-model="params.start_time"
            type="date"
            format="yyyy-MM-dd"
            placeholder="选择开始日期"
          ></el-date-picker>
          <el-date-picker
            class="input-cos"
            v-model="params.end_time"
            type="date"
            format="yyyy-MM-dd"
            placeholder="选择结束日期"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-select
            v-model="params.channel_no"
            placeholder="下游渠道"
            class="input-cos"
            filterable
            clearable
          >
            <el-option value label="全部下游渠道"></el-option>
            <el-option
              v-for="(item, index) in channelNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            v-model="params.province_no"
            placeholder="请选择省份"
            class="input-cos"
            filterable
            clearable
          >
            <el-option value label="全部省份"></el-option>
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
            v-model="params.line_id"
            placeholder="请选择产品线"
            size="medium"
            style="width:200px"
            filterable
            clearable
          >
            <el-option value label="全部产品线"></el-option>

            <el-option
              v-for="(item, index) in lineId"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            v-model="params.carrier_no"
            placeholder="请选择运营商"
            size="medium"
            style="width:200px"
            filterable
            clearable
          >
            <el-option value label="全部运营商"></el-option>
            <el-option
              v-for="(item, index) in carrierNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-input placeholder="请输入总面值" v-model="params.total_face"></el-input>
        </el-form-item>
        <div class="form-group">
          <button class="btn btn-primary" type="button" @click="QueryDataList">查询</button>
          <button class="btn btn-success" type="button" @click="Export">导出到Excel</button>
        </div>
      </el-form>
    </div>
    <div class="table-responsive">
      <table class="table table-striped m-b-none th-bg">
        <thead>
          <tr>
            <th style="text-align:center">下游渠道</th>
            <th style="text-align:center">总笔数</th>
            <th style="text-align:center">失败笔数</th>
            <th style="text-align:center">失败金额</th>
            <th style="text-align:center">成功笔数</th>
            <th style="text-align:center">总金额</th>
            <th style="text-align:center">成功金额</th>
            <th style="text-align:center">成本率</th>
            <th style="text-align:center">实际扣款金额</th>
            <th style="text-align:center">实际发货成功总成本</th>
            <th style="text-align:center">利润</th>
          </tr>
        </thead>
        <tbody v-if="dataList == '' && totalcount == 0" class="table-border">
          <tr>
            <td colspan="12">
              <div
                style="text-align:center;min-width:1355px;height:494px;line-height: 494px;color: #909399;"
              >暂无数据</div>
            </td>
          </tr>
        </tbody>
        <tbody class="table-border" else>
          <tr v-for="(item, index) in dataList" :key="index">
            <td style="text-align:center">{{item.channel_name}}</td>
            <td style="text-align:center">{{item.order_count |AmountFilter}}</td>
            <td style="text-align:center">{{item.fail_count |AmountFilter}}</td>
            <td style="text-align:center">{{item.fail_face |AmountFilter}}</td>
            <td style="text-align:center">{{item.success_count |AmountFilter}}</td>
            <td style="text-align:center">{{item.total_face |AmountFilter}}</td>
            <td style="text-align:center">{{item.success_face |AmountFilter}}</td>
            <td style="text-align:center">{{item.success_ratio |FeeFilter}}</td>
            <td style="text-align:center">{{item.success_sell |AmountFilter}}</td>
            <td style="text-align:center">{{item.success_cost |AmountFilter}}</td>
            <td style="text-align:center">{{item.profit |AmountFilter}}</td>
          </tr>

          <tr>
            <td style="text-align:center;padding: 14px 15px" colspan="1">总计</td>
            <td style="text-align:center">{{totaldata.order_count |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.fail_count |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.fail_face |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.success_count |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.total_face |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.success_face |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.success_ratio |FeeFilter}}</td>
            <td style="text-align:center">{{totaldata.success_sell |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.success_cost |AmountFilter}}</td>
            <td style="text-align:center">{{totaldata.profit |AmountFilter}}</td>
          </tr>
        </tbody>
      </table>
    </div>

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
  </div>
</template>
<script>
export default {
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      status: this.EnumUtility.Get("Status"),
      channelNo: this.EnumUtility.Get(
        "DownChannelNo",
        {},
        "/oms/down/channel/getdictionary"
      ),
      provinceNo: [],
      lineId: this.EnumUtility.Get(
        "LineId",
        {},
        "/oms/product/line/getdictionary"
      ),
      carrierNo: this.EnumUtility.Get("CarrierNo"),
      params: {
        pi: 1,
        ps: 10,
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
          new Date().getDate(),
        province_no: "",
        channel_no: "",
        line_id: "",
        carrier_no: "",
        total_face: ""
      },
      dayformat: "yyyy-MM-dd",
      totalcount: 0,
      dataList: [],
      totaldata: {
        order_count: 0,
        fail_count: 0,
        fail_face: 0,
        success_count: 0,
        total_face: 0,
        success_face: 0,
        success_ratio: 0,
        success_sell: 0,
        success_cost: 0,
        profit: 0
      },
      ProvinceList: []
    };
  },
  mounted() {
    this.Init();
    this.city();
  },
  methods: {
    total() {
      for (var key in this.totaldata) {
        this.totaldata[key] = 0;
      }
      this.dataList.forEach(item => {
        for (var key in this.totaldata) {
          this.totaldata[key] = this.totaldata[key] + parseFloat(item[key]);
        }
      });
    },
    city() {
      this.$http
        .post("/oms/canton/info/getdictionary", { grade: 1 })
        .then(res => {
          this.provinceNo = res;
        })
        .catch(err => {
          this.$message({
            type: "error",
            message: err.response.data
          });
        });
    },
    Export() {
      var oReq = new XMLHttpRequest();
      oReq.open(
        "POST",
        process.env.VUE_APP_API_URL + "/oms/report/downchannel/export",
        true
      );
      oReq.responseType = "blob";
      oReq.setRequestHeader("Content-Type", "application/json");

      oReq.onload = function(oEvent) {
        var content = oReq.response;
        var elink = document.createElement("a");
        elink.download = "下游渠道统计报表.xlsx"; //xlsx    因为后台输入是scv格式，用xlsx显示的不理想
        elink.style.display = "none";
        var blob = new Blob([content]);
        elink.href = URL.createObjectURL(blob);
        document.body.appendChild(elink);
        elink.click();
        document.body.removeChild(elink);
      };
      oReq.send(JSON.stringify(this.params));
    },
    Init() {
      this.params.pi = 1;
      this.params.ps = 10;
      this.QueryDataList();
    },

    QueryDataList() {
      this.params.start_time =
        this.DateConvert(this.dayformat, this.params.start_time) + " 00:00:00";
      this.params.end_time =
        this.DateConvert(this.dayformat, this.params.end_time) + " 23:59:59";
      this.$http
        .get("/oms/report/downchannel/query", this.params)
        .then(response => {
          if (!response.data || response.count === 0) {
            this.dataList = [];
            this.totalcount = 0;
            this.total();
            return;
          }
          this.dataList = response.data;
          this.totalcount = response.count;
          this.total();
        });
    },
    handleSizeChange(val) {
      this.params.ps = val;
      this.QueryDataList();
    },
    handleCurrentChange(val) {
      this.params.pi = val;
      this.QueryDataList();
    }
  }
};
</script>

<style scoped>
.page-pagination {
  padding: 10px 15px;
  text-align: right;
}
.input-cos {
  width: 160px;
}
</style>