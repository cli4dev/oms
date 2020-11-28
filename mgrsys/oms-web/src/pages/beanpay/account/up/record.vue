<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-date-picker
            class="input-cos"
            v-model="queryData.start_time"
            type="date"
            format="yyyy-MM-dd"
            placeholder="选择开始日期"
          ></el-date-picker>
          <el-date-picker
            class="input-cos"
            v-model="queryData.end_time"
            type="date"
            format="yyyy-MM-dd"
            placeholder="选择结束日期"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.account_id"
            class="input-cos"
            filterable
            placeholder="请选择帐户名称"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in accountList"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.change_type"
            class="input-cos"
            filterable
            placeholder="请选择变动类型 "
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in changeType"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.trade_type"
            class="input-cos"
            filterable
            placeholder="请选择交易类型 "
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in tradeType"
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
        <el-table-column prop="record_id" label="变动编号" align="center"></el-table-column>
        <el-table-column prop="trade_no" label="交易编号" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.trade_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="account_id" label="帐户名称" align="center" min-width="120px">
          <template slot-scope="scope">
            <span>{{accountName(scope.row.account_id )}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="amount" label="变动金额" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.amount/100}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="balance" label="帐户余额" align="center">
			 <template slot-scope="scope">
            <span>{{scope.row.balance/100}}</span>
          </template>
		</el-table-column>
        <el-table-column prop="change_type" label="变动类型 " align="center">
          <template slot-scope="scope">
            <span>{{scope.row.change_type | EnumFilter("ChangeType")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="ext" label="扩展字段" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.ext | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="trade_type" label="交易类型 " align="center">
          <template slot-scope="scope">
            <span>{{scope.row.trade_type | EnumFilter("TradeType")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" align="center"></el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button
              type="text"
              size="small"
              v-if="scope.row.change_type =='1' && scope.row.trade_type != '4'"
              @click="Redrush(scope.row)"
            >红冲</el-button>
            <el-button
              type="text"
              size="small"
              v-if=" scope.row.change_type == '2'&& scope.row.trade_type != '4' "
              @click="RedrushDraw(scope.row)"
            >红冲</el-button>
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
  name: "BeanpayAccountRecord",
  components: {},
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {
        types:"up",
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
      },
      dayformat: "yyyy-MM-dd",
      accountList: [],
      changeType: this.EnumUtility.Get("ChangeType"),
      tradeType: this.EnumUtility.Get("TradeType"),
      tableData: []
    };
  },
  created() {
    this.$http
        .get("/beanpay/account/info/getupdictionary", {})
        .then(res => {
          this.accountList = res
        })
        .catch(err => {
          console.log(err);
        });
  },
  mounted() {
    this.init();
  },
  methods: {
    RedrushDraw(value) {
      this.$confirm("此操作将进行提款红冲, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          value.types = "up_channel";
          this.$http
            .post("/beanpay/account/record/redrushdraw", value)
            .then(res => {
              this.$message({
                type: "success",
                message: "提款红冲成功!"
              });
              this.dialogFormVisible = false;
              this.query();
            })
             .catch(error => {
              this.$message({
                type:"warning",
                message:error.response.data
              })
              console.log(err);
            });
        })
        .catch(() => {
          console.log("9999999")
          this.$message({
            type: "info",
            message: "已取消提款红冲"
          });
        });
    },
    Redrush(value) {
      this.$confirm("此操作将进行加款红冲, 是否继续?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      })
        .then(() => {
          value.types = "up_channel";
          this.$http
            .post("/beanpay/account/record/redrush", value)
            .then(res => {
              this.$message({
                type: "success",
                message: "加款红冲成功!"
              });
              this.dialogFormVisible = false;
              this.query();
            })
             .catch(error => {
              this.$message({
                type:"warning",
                message:error.response.data
              })
              console.log(err);
            });
        })
        .catch(error => {
           console.log("999")
          this.$message({
            type: "info",
            message: "已取消加款红冲"
          });
        });
    },
    accountName(value) {
      var res = value;
     
      this.accountList.forEach(item => {
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
    query() {
      this.queryData.pi = this.params.pi;
      this.queryData.ps = this.params.ps;
      this.queryData.start_time =
        this.DateConvert(this.dayformat, this.queryData.start_time) + " 00:00:00";
      this.queryData.end_time =
        this.DateConvert(this.dayformat, this.queryData.end_time) + " 23:59:59";
      this.$http
        .get("/beanpay/account/record/query", this.queryData)
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
        getpath: "/beanpay/account/record",
        record_id: val.record_id
      };
      this.$emit(
        "addTab",
        "详情" + val.record_id,
        "/beanpay/account/record.view",
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
  width: 160px;
}
</style>
