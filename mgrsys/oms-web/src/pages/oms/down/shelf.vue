<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-input clearable v-model="queryData.shelf_name" placeholder="请输入货架名称"></el-input>
        </el-form-item>

        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.channel_no"
            class="input-cos"
            filterable
            placeholder="请选择渠道名称"
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
        <el-table-column prop="shelf_id" label="货架编号"  @click="ok" align="center">
          
        </el-table-column>
        <el-table-column prop="shelf_name" label="货架名称" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.shelf_name | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="channel_no" label="渠道名称" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.channel_no | EnumFilter("DownChannelNo")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="order_overtime" label="订单超时时长" align="center"></el-table-column>
        <el-table-column prop="refund_overtime" label="退款超时时间" align="center"></el-table-column>
        <el-table-column prop="status" label="状态" align="center">
          <template slot-scope="scope">
            <span
              :class="[scope.row.status === '0' ? 'text-success' : 'text-danger']"
            >{{scope.row.status | EnumFilter("Status")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" align="center"></el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
            <!-- <el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button> -->
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
import Add from "./shelf.add";
import Edit from "./shelf.edit";
export default {
  name: "OmsDownShelf",
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
      queryData: {},
      channelNo: this.EnumUtility.Get(
        "DownChannelNo",
        {},
        "/oms/down/channel/getdictionary"
      ),
      status: this.EnumUtility.Get("Status"),
      tableData: []
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    ok(){

    },
	   shelfList() {
      this.$http.get("/oms/down/channel/getdictionary", {}).then(res => {
        this.channelNo = res;
      });
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
        .get("/oms/down/shelf/query", this.queryData)
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
        getpath: "/oms/down/shelf",
        shelf_id: val.shelf_id
      };
      this.$emit("addTab", "详情" + val.shelf_id, "/oms/down/shelf.view", data);
    },
    addShow() {
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
