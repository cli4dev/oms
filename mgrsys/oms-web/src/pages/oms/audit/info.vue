<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-date-picker
            clearable
            filterable
            v-model="queryData.start_time"
            @change="TimeShow(queryData.start_time,queryData.end_time)"
            value-format="yyyy-MM-dd"
            placeholder="请输入开始时间"
            style="width:180px"
            size="medium"
          ></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-date-picker
            clearable
            v-model="queryData.end_time"
            @change="TimeShow(queryData.start_time,queryData.end_time)"
            value-format="yyyy-MM-dd"
            placeholder="请选择结束时间"
            style="width:180px"
            size="medium"
          >结束时间</el-date-picker>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.delivery_id" placeholder="请输入发货记录编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.order_id" placeholder="请输入订单编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.refund_id" placeholder="请输入退款编号"></el-input>
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
        <el-table-column prop="create_time" label="创建时间" align="center" min-width="140px">
          <template slot-scope="scope">
            <span v-if="dayNum !=0">{{scope.row.create_time |DateFilter("yyyy-MM-dd hh:mm:ss") }}</span>
            <span v-else>{{scope.row.create_time | DateFilter("hh:mm:ss")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="audit_id" label="审核编号" align="center"></el-table-column>
        <el-table-column prop="order_id" label="订单编号" align="center">
          <template slot-scope="scope">
            <el-button
              type="text"
              size="small"
              @click="detail(scope.row.order_id)"
            >{{scope.row.order_id}}</el-button>
          </template>
        </el-table-column>
        <el-table-column prop="delivery_id" label="发货编号" align="center"></el-table-column>
        <el-table-column prop="audit_by" label="审核人" align="center">
          <!-- <template scope="scope" -->
        </el-table-column>
        <el-table-column prop="audit_msg" label="审核信息" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.audit_msg | EllipsisFilter(20)}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="audit_time" label="审核时间" align="center"></el-table-column>
        <el-table-column prop="change_type" label="审核类型" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.change_type | EnumFilter("AuditType") }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="change_type" label="变动类型" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.change_type | EnumFilter("ChangeType")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="audit_status" label="审核状态" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.audit_status | EnumFilter("AuditStatus") }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button
              type="text"
              size="small"
              v-if="scope.row.audit_status =='20'"
              @click="editShow(scope.row)"
            >审核</el-button>
            <!-- <el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button> -->
          </template>
        </el-table-column>
      </el-table>
    </el-scrollbar>
    <!-- list end-->

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
import Edit from "./info.edit";
export default {
  name: "OmsAuditInfo",
  components: {
    Edit
  },
  data() {
    return {
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      auditStatus: this.EnumUtility.Get("AuditStatus"),
      auditType: this.EnumUtility.Get("AuditType"),
      dayformat: "yyyy-MM-dd",
      dayNum: 0,
      queryData: {
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
          new Date().getDate()
      },
      changeType: this.EnumUtility.Get("ChangeType"),
      dtCreateTime: this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
      tableData: []
    };
  },
  created() {},
  mounted() {
    this.init();
  },
  methods: {
    TimeShow(start, end) {
      var dateSpan, tempDate, iDays;
      start = Date.parse(start);
      end = Date.parse(end);
      dateSpan = end - start;
      dateSpan = Math.abs(dateSpan);
      iDays = Math.floor(dateSpan / (24 * 3600 * 1000));

      this.dayNum = iDays;
    },
    detail(orderId) {
      let val = { order_id: orderId };
      val.getpath = "/oms/order/info";
      this.$emit("addTab", "详情" + orderId, "/oms/order/info.view", val);
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
        this.DateConvert(this.dayformat, this.queryData.start_time) +
        " 00:00:00";
      this.queryData.end_time =
        this.DateConvert(this.dayformat, this.queryData.end_time) + " 23:59:59";
      this.$http
        .get("/oms/audit/info/query", this.queryData)
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
        getpath: "/oms/audit/info",
        audit_id: val.audit_id
      };
      this.$emit("addTab", "详情" + val.audit_id, "/oms/audit/info.view", data);
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
