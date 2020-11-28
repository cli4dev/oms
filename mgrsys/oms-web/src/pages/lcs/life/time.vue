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
          <el-input clearable v-model="queryData.order_no" placeholder="请输入业务单据号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.batch_no" placeholder="请输入业务批次号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.content" placeholder="请输入内容"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input clearable v-model="queryData.ip" placeholder="请输入服务器ip"></el-input>
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
        <el-table-column prop="id" label="id"></el-table-column>
        <el-table-column prop="order_no" label="业务单据号">
          <template slot-scope="scope">
            <span>{{scope.row.order_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="batch_no" label="业务批次号">
          <template slot-scope="scope">
            <span>{{scope.row.batch_no | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="extral_param" label="扩展编号">
          <template slot-scope="scope">
            <span>{{scope.row.extral_param | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="content" label="内容">
          <template slot-scope="scope">
            <span>{{scope.row.content | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间"></el-table-column>
        <el-table-column prop="ip" label="服务器ip">
          <template slot-scope="scope">
            <span>{{scope.row.ip | EllipsisFilter(20)}}</span>
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
  name: "LcsLifeTime",
  components: {},
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {},
      dtCreateTime: this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
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
        .get("/lcs/life/time/query", this.queryData)
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
        getpath: "/lcs/life/time",
        id: val.id
      };
      this.$emit("addTab", "详情" + val.id, "/lcs/life/time.view", data);
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
