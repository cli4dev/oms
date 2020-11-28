<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-input clearable v-model="queryData.account_name" placeholder="请输入帐户名称"></el-input>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.groups"
            class="input-cos"
            filterable
            placeholder="请选择账户类型"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in list"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.ident" placeholder="请输入系统标识"></el-input>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.status"
            class="input-cos"
            placeholder="请选择账户状态"
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

        <el-form-item>
          <el-button type="success" size="small" @click="addShow">添加</el-button>
        </el-form-item>
      </el-form>
    </div>
    <!-- query end -->

    <!-- list start-->
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="account_id" align="center" label="帐户编号"></el-table-column>
        <el-table-column prop="account_name" align="center" label="帐户名称">
          <template slot-scope="scope">
            <span>{{scope.row.account_name | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="eid" align="center" label="外部用户账户">
          <template slot-scope="scope">
              <span>{{scope.row.eid | EnumFilter("DownChannelNO")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="groups" align="center" label="账户类型">
          <template slot-scope="scope">
            <span>{{scope.row.groups | EnumFilter("down")| EnumFilter("up")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" align="center" label="账户状态">
          <template slot-scope="scope">
            <span :class="[scope.row.status === '0' ? 'text-success' : 'text-danger']">{{scope.row.status | EnumFilter("Status")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="balance" align="center" label="帐户余额">
          <template slot-scope="scope">
            <span>{{scope.row.balance/100}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="credit" align="center" label="信用余额">
          <template slot-scope="scope">
            <span>{{scope.row.credit/100}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" align="center" label="创建时间"></el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>
    <!-- list end-->

    <!-- Add Form -->
    <Add ref="Add" :refresh="query"></Add>
    <!--Add Form -->

    <!-- edit Form start-->
    <Edit ref="Edit" :refresh="query"></Edit>
    <!-- edit Form end-->

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
import Add from "./info.add";
import Edit from "./info.edit";
export default {
  name: "BeanpayAccountInfo",
  components: {
    Add,
    Edit
  },
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: { types: "down" },
      channelNo: this.EnumUtility.Get(
        "DownChannelNO",
        {},
        "/oms/down/channel/getdictionary"
      ),

      groups: this.EnumUtility.Get("down"),
      up: this.EnumUtility.Get("up"),
      list: [],
      status: this.EnumUtility.Get("Status"),
      tableData: [],
      upList: [],
      downList: [],
      queryList: []
    };
  },
  created() {
    this.$http
      .get("/oms/down/channel/getdictionary", {})
      .then(res => {
        this.downList = res;
      })
      .catch(err => {
        console.log(err);
      });
    this.$http
      .get("/oms/up/channel/getdictionary", {})
      .then(res => {
        this.upList = res;
      })
      .catch(err => {
        console.log(err);
      });
  },
  mounted() {
    this.init();
    this.queryList = this.upList.concat(this.downList);
    this.list = this.groups;
  },
  methods: {
    name(value) {
      var result = value;
      this.queryList.forEach(item => {
        if (result == item.value) {
          result = item.name;
        }
      });
      return result;
    },
    /**初始化操作**/
    init() {
      this.query();
    },
    /**查询数据并赋值*/
    query() {
      this.queryData.pi = this.params.pi;
      this.queryData.ps = this.params.ps;
      this.$http
        .get("/beanpay/account/info/query", this.queryData)
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
        getpath: "/beanpay/account/info",
        account_id: val.account_id
      };
      this.$emit(
        "addTab",
        "详情" + val.account_id,
        "/beanpay/account/info.view",
        data
      );
    },
    addShow() {
      this.$refs.Add.list = this.list;
      this.$refs.Add.show();
    },
    editShow(val) {
      this.$refs.Edit.editData = val;
      this.$refs.Edit.show();
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
