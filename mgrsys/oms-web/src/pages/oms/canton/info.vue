<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-input clearable v-model="queryData.canton_code" placeholder="请输入区域编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.chinese_name" placeholder="请输入中文名称"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.parent" placeholder="请输入父级"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.simple_spell" placeholder="请输入简拼"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="query" size="small">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-scrollbar style="height:100%">
      <el-table :data="tableData" border style="width: 100%">
        <el-table-column prop="canton_code" align="center" label="区域编号">
          <template slot-scope="scope">
            <span>{{scope.row.canton_code | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="chinese_name" align="center" label="中文名称">
          <template slot-scope="scope">
            <span>{{scope.row.chinese_name | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="spell" align="center" label="英文或全拼">
          <template slot-scope="scope">
            <span>{{scope.row.spell | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="grade" align="center" label="行政级别"></el-table-column>
        <el-table-column prop="parent" align="center" label="父级">
          <template slot-scope="scope">
            <span>{{scope.row.parent | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="simple_spell" align="center" label="简拼">
          <template slot-scope="scope">
            <span>{{scope.row.simple_spell | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="area_code" align="center" label="区号">
          <template slot-scope="scope">
            <span>{{scope.row.area_code | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="standard_code" align="center" label="行政编码">
          <template slot-scope="scope">
            <span>{{scope.row.standard_code | EllipsisFilter(20)}}</span>
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
  name: "OmsCantonInfo",
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      queryData: {},
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
      this.$http
        .get("/oms/canton/info/query", this.queryData)
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
