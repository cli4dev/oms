<template>
<div class="table-responsive">
  <table :data="item" class="table table-striped m-b-none">
    <tbody class="table-border">
      <tr>
        <td>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知编号:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.notify_id}}</div>
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
            <div class="pull-right" style="margin-right:10px">退款编号:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.refund_id}}</div>
          </el-col>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">创建时间:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.create_time}}</div>
          </el-col>
        </td>
      </tr>
      <tr>
        <td>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">开始时间:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.start_time}}</div>
          </el-col>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">结束时间:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.end_time}}</div>
          </el-col>
        </td>
      </tr>
      <tr>
        <td>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知次数:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.notify_count}}</div>
          </el-col>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知结果信息:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.notify_msg}}</div>
          </el-col>
        </td>
      </tr>
      <tr>
        <td>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知状态:</div>
          </el-col>
          <el-col :span="6">
            <div
              class="grid-content"
              :class="generalStatusClassFilter(item.notify_status)"
            >{{item.notify_status | EnumFilter("NotifyStatus")}}</div>
          </el-col>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知类型:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.notify_type | EnumFilter("NotifyType")}}</div>
          </el-col>
        </td>
      </tr>
      <tr>
        <td>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">通知地址:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.notify_url}}</div>
          </el-col>
          <el-col :span="6">
            <div class="pull-right" style="margin-right:10px">最大通知次数:</div>
          </el-col>
          <el-col :span="6">
            <div class="grid-content">{{item.max_count}}</div>
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
      colorClass: ["text-success", "text-danger", "text-muted", "text-primary"],
      item: {}
    };
  },
  created() {
    this.EnumUtility.Get("NotifyStatus");
    this.EnumUtility.Get("NotifyType");
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
