<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-input clearable v-model="queryData.channel_no" placeholder="请输入编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.channel_name" placeholder="请输入名称"></el-input>
        </el-form-item>

        <el-form-item>
          <el-select size="medium" v-model="queryData.status" class="input-cos" placeholder="请选择状态">
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
        <el-table-column prop="channel_no" align="center" label="编号">
          <template slot-scope="scope">
            <span>{{scope.row.channel_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="channel_name" align="center" label="名称">
          <template slot-scope="scope">
            <span>{{scope.row.channel_name | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="balance" align="center" label="帐户余额">
          <template slot-scope="scope">
            <span>{{scope.row.balance/100}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="credit" align="center" label="授信金额">
          <template slot-scope="scope">
            <span>{{scope.row.credit/100}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="ext_channel_no" label="外部渠道编号" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.ext_channel_no |StringFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" align="center" label="状态">
          <template slot-scope="scope">
            <span
              :class="[scope.row.status === '0' ? 'text-success' : 'text-danger']"
            >{{scope.row.status | EnumFilter("Status")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" align="center" label="创建时间"></el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="addBalanceShow(scope.row.channel_no,1)">加款</el-button>
            <el-button type="text" size="small" @click="drawMoney(scope.row.channel_no)">提款</el-button>
            <el-button type="text" size="small" @click="addBalanceShow(scope.row.channel_no,2)">交易平账</el-button>
            <el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="setSecret(scope.row)">设置秘钥</el-button>
            <el-button type="text" size="small" @click="checkSecret(scope.row.channel_no)">查看秘钥</el-button>
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

    <!-- AddBalance Form -->
    <AddBalance ref="AddBalance" :refresh="query"></AddBalance>
    <DrawMoney ref="DrawMoney" :refresh="query"></DrawMoney>
    <!--AddBalance Form -->
    <SetSecret ref="SetSecret" :refresh="query"></SetSecret>
    <CheckSecret ref="CheckSecret" :refresh="query"></CheckSecret>
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
import Add from "./channel.add";
import Edit from "./channel.edit";
import SetSecret from "./channel.secret.vue";
import CheckSecret from "./channel.check.secret.vue";
import AddBalance from "./channel.add.balance";
import DrawMoney from "./channel.draw.money.vue";
export default {
  name: "OmsUpChannel",
  components: {
    Add,
    Edit,
    DrawMoney,
    SetSecret,
    CheckSecret,
    AddBalance
  },
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {},
      status: this.EnumUtility.Get("Status"),
      tableData: []
    };
  },
  created() {},
  mounted() {
    this.init();
  },
  methods: {
    checkSecret(val) {
      // this.$refs.CheckSecret.addData = val;
      this.$refs.CheckSecret.show(val);
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
        .get("/oms/up/channel/query", this.queryData)
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
        getpath: "/oms/up/channel",
        channel_no: val.channel_no
      };
      this.$emit(
        "addTab",
        "详情" + val.channel_no,
        "/oms/up/channel.view",
        data,
        false
      );
    },
    addShow() {
      this.$refs.Add.show();
    },
    addBalanceShow(val,v) {
      this.$refs.AddBalance.show(val,v);
    },
    drawMoney(val) {
      this.$refs.DrawMoney.show(val);
    },
    editShow(val) {
      this.$refs.Edit.editData = val;
      this.$refs.Edit.show();
    },
    setSecret(val) {
      this.$refs.SetSecret.addData = val;
      this.$refs.SetSecret.show();
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
