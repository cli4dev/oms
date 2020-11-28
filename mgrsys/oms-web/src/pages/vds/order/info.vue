<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-date-picker
            class="input-cos"
            v-model="dtCreateTime"
            popper-class="datetime-to-date"
            type="datetime"
            value-format="yyyy-MM-dd HH:mm:ss"
            placeholder="选择日期"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.carrier_no"
            class="input-cos"
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
            v-model="queryData.channel_no"
            class="input-cos"
            placeholder="请选择上游渠道编号"
          >
            <el-option value label="全部"></el-option>
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
            size="medium"
            v-model="queryData.coop_id"
            class="input-cos"
            placeholder="请选择下游商户编号"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in coopId"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.coop_order_id" placeholder="请输入下游商户订单号"></el-input>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.result_source"
            class="input-cos"
            placeholder="请选择发货结果来源"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in resultSource"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.service_class"
            class="input-cos"
            placeholder="请选择服务类型 "
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in serviceClass"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.status"
            class="input-cos"
            placeholder="请选择发货状态"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in status"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="query" size="small">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- query end -->

    <!-- list start-->
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="order_no" label="发货编号"></el-table-column>
        <el-table-column prop="carrier_no" label="运营商">
          <template slot-scope="scope">
            <span>{{scope.row.carrier_no | EnumFilter("CarrierNo")}}</span>
          </template>

          <template slot-scope="scope">
            <span>{{scope.row.carrier_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="channel_no" label="上游渠道编号">
          <template slot-scope="scope">
            <span>{{scope.row.channel_no | EnumFilter("ChannelNo")}}</span>
          </template>

          <template slot-scope="scope">
            <span>{{scope.row.channel_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="coop_id" label="下游商户编号">
          <template slot-scope="scope">
            <span>{{scope.row.coop_id | EnumFilter("CoopId")}}</span>
          </template>

          <template slot-scope="scope">
            <span>{{scope.row.coop_id | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="coop_order_id" label="下游商户订单号">
          <template slot-scope="scope">
            <span>{{scope.row.coop_order_id | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间"></el-table-column>
        <el-table-column prop="flow_timeout" label="流程超时时间"></el-table-column>
        <el-table-column prop="notify_url" label="下游通知回调地址">
          <template slot-scope="scope">
            <span>{{scope.row.notify_url | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="product_face" label="产品面值">
          <template slot-scope="scope">
						<span>{{scope.row.product_face/100}}</span>
					</template>
        </el-table-column>
        <el-table-column prop="product_num" label="产品数量"></el-table-column>
        <el-table-column prop="request_finish_time" label="请求完成时间"></el-table-column>
        <el-table-column prop="request_start_time" label="请求开始时间"></el-table-column>
        <el-table-column prop="result_params" label="记录上游返回的业务数据">
          <template slot-scope="scope">
            <span>{{scope.row.result_params | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="result_source" label="发货结果来源">
          <template slot-scope="scope">
            <span>{{scope.row.result_source | EnumFilter("ResultSource")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="service_class" label="服务类型 ">
          <template slot-scope="scope">
            <span>{{scope.row.service_class | EnumFilter("ServiceClass")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="发货状态">
          <template slot-scope="scope">
            <span>{{scope.row.status | EnumFilter("DeliveryStatus")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="succ_face" label="成功面值">
           <template slot-scope="scope">
						<span>{{scope.row.succ_face/100}}</span>
					</template>
        </el-table-column>
        <el-table-column prop="up_order_no" label="上游发货订单号">
          <template slot-scope="scope">
            <span>{{scope.row.up_order_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作">
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
  name: "VdsOrderInfo",
  components: {},
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {},
      carrierNo: this.EnumUtility.Get("CarrierNo"),
      channelNo: this.EnumUtility.Get(
        "ChannelNo",
        {},
        "/oms/up/channel/getdictionary"
      ),
      coopId: this.EnumUtility.Get(
        "CoopId",
        {},
        "/oms/down/channel/getdictionary"
      ),
      dtCreateTime: this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
      resultSource: this.EnumUtility.Get("ResultSource"),
      serviceClass: this.EnumUtility.Get("ServiceClass"),
      status: this.EnumUtility.Get("DeliveryStatus"),
      tableData: []
    };
  },
  created() {},
  mounted() {
    this.init();
  },
  methods: {
    /**初始化操作**/
    init() {
      this.query();
    },
    /**查询数据并赋值*/
    query() {
      this.queryData.pi = this.params.pi;
      this.queryData.ps = this.params.ps;
      this.queryData.create_time = this.DateConvert(
        "yyyy-MM-dd hh:mm:ss",
        this.dtCreateTime
      );
      this.$http
        .get("/vds/order/info/query", this.queryData)
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
        getpath: "/vds/order/info",
        order_no: val.order_no
      };
      this.$emit("addTab", "详情" + val.order_no, "/vds/order/info.view", data);
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
</style>
