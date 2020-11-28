<template>
  <el-table :data="info" style="width: 100%" :show-header="false" :default-expand-all="true">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand" label-width="120px">
          <el-form-item label="变动编号">
            <span v-text="props.row.record_id"></span>
          </el-form-item>
          <el-form-item label="帐户名称">
            <span>{{props.row.account_id | EnumFilter("AccountId")}}</span>
          </el-form-item>
          <el-form-item label="变动金额">
            <span>{{props.row.amount/100}}</span>
          </el-form-item>
          <el-form-item label="帐户余额">
            <span>{{props.row.balance/100}}</span>
          </el-form-item>
          <el-form-item label="变动类型 ">
            <span>{{props.row.change_type | EnumFilter("ChangeType")}}</span>
          </el-form-item>
          <el-form-item label="创建时间">
            <span v-text="props.row.create_time"></span>
          </el-form-item>
          <el-form-item label="扣款编号">
            <span v-text="props.row.deduct_no"></span>
          </el-form-item>
          <el-form-item label="扩展字段">
            <span v-text="props.row.ext"></span>
          </el-form-item>
          <el-form-item label="交易编号">
            <span v-text="props.row.trade_no"></span>
          </el-form-item>
          <el-form-item label="交易类型 ">
            <span>{{props.row.trade_type | EnumFilter("TradeType")}}</span>
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
    this.EnumUtility.Get(
      "AccountId",
      {},
      "/beanpay/account/info/getdictionary"
    );
    this.EnumUtility.Get("ChangeType");
    this.EnumUtility.Get("TradeType");
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
