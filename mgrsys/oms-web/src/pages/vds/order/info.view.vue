<template>
  <el-table :data="info" style="width: 100%" :show-header="false" :default-expand-all="true">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand" label-width="120px">
          <el-form-item label="发货编号">
            <span v-text="props.row.order_no"></span>
          </el-form-item>
          <el-form-item label="运营商">
            <span>{{props.row.carrier_no | EnumFilter("CarrierNo")}}</span>
          </el-form-item>
          <el-form-item label="上游渠道编号">
            <span>{{props.row.channel_no | EnumFilter("ChannelNo")}}</span>
          </el-form-item>
          <el-form-item label="下游商户编号">
            <span>{{props.row.coop_id | EnumFilter("CoopId")}}</span>
          </el-form-item>
          <el-form-item label="下游商户订单号">
            <span v-text="props.row.coop_order_id"></span>
          </el-form-item>
          <el-form-item label="创建时间">
            <span v-text="props.row.create_time"></span>
          </el-form-item>
          <el-form-item label="流程超时时间">
            <span v-text="props.row.flow_timeout"></span>
          </el-form-item>
          <el-form-item label="最后更新时间">
            <span v-text="props.row.last_update_time"></span>
          </el-form-item>
          <el-form-item label="下游通知回调地址">
            <span v-text="props.row.notify_url"></span>
          </el-form-item>
          <el-form-item label="产品面值">
            <span>{{props.row.product_face/100}}</span>
          </el-form-item>
          <el-form-item label="产品数量">
            <span v-text="props.row.product_num"></span>
          </el-form-item>
          <el-form-item label="请求完成时间">
            <span v-text="props.row.request_finish_time"></span>
          </el-form-item>
          <el-form-item label="发货信息参数json">
            <span v-text="props.row.request_params"></span>
          </el-form-item>
          <el-form-item label="请求开始时间">
            <span v-text="props.row.request_start_time"></span>
          </el-form-item>
          <el-form-item label="发货结果码">
            <span v-text="props.row.result_code"></span>
          </el-form-item>
          <el-form-item label="结果描述">
            <span v-text="props.row.result_desc"></span>
          </el-form-item>
          <el-form-item label="记录上游返回的业务数据">
            <span v-text="props.row.result_params"></span>
          </el-form-item>
          <el-form-item label="发货结果来源">
            <span>{{props.row.result_source | EnumFilter("ResultSource")}}</span>
          </el-form-item>
          <el-form-item label="服务类型 ">
            <span>{{props.row.service_class | EnumFilter("ServiceClass")}}</span>
          </el-form-item>
          <el-form-item label="发货状态">
            <span>{{props.row.status | EnumFilter("DeliveryStatus")}}</span>
          </el-form-item>
          <el-form-item label="成功面值">
            <span>{{props.row.succ_face/100}}</span>
          </el-form-item>
          <el-form-item label="上游发货订单号">
            <span v-text="props.row.up_order_no"></span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column label></el-table-column>
    <el-table-column label></el-table-column>
    <el-table-column label></el-table-column>
  </el-table>
</template>


<script>
export default {
  data() {
    return {
      info: []
    };
  },
  created() {
    this.EnumUtility.Get("CarrierNo");
    this.EnumUtility.Get("ChannelNo", {}, "/oms/up/channel/getdictionary");
    this.EnumUtility.Get("CoopId", {}, "/oms/down/channel/getdictionary");
    this.EnumUtility.Get("ResultSource");
    this.EnumUtility.Get("ServiceClass");
    this.EnumUtility.Get("DeliveryStatus");
  },
  mounted() {
    this.Init();
  },
  methods: {
    Init() {
      this.QueryDataList();
    },
    QueryDataList() {
      this.$http.get(this.$route.query.getpath, this.$route.query).then(res => {
        this.info.push(res);
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
