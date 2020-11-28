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
          <el-select
            size="medium"
            v-model="queryData.notify_status"
            class="input-cos"
            placeholder="请选择通知状态"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in notifyStatus"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.order_id" placeholder="请输入订单编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-input clearable v-model="queryData.notify_id" placeholder="请输入通知编号"></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="query(1)" size="small">查询</el-button>
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
        <el-table-column prop="notify_id" label="通知编号" align="center"></el-table-column>
        <el-table-column prop="order_id" label="订单编号" align="center">
           <template slot-scope="scope">
						<el-button type="text" size="small" @click="detail(scope.row.order_id)">{{scope.row.order_id}}</el-button>
					</template>
        </el-table-column>

        <el-table-column prop="start_time" label="开始时间" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.start_time |StringFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="end_time" label="结束时间" align="center">
           <template slot-scope="scope">
            <span>{{scope.row.end_time |StringFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="notify_count" label="通知次数" align="center"></el-table-column>
        <el-table-column prop="max_count" label="最大通知次数" align="center"></el-table-column>
        <el-table-column prop="notify_status" label="通知状态" align="center">
          <template slot-scope="scope">
            <span :class="generalStatusClassFilter(scope.row.notify_status)">{{scope.row.notify_status | EnumFilter("NotifyStatus")}}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" align="center">
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
  name: "OmsNotifyInfo",
  components: {},
  data() {
    return {
      colorClass: [
          'text-success',
          'text-danger',
          'text-muted',
          'text-primary',
        ],
        dayNum: 0,
      pageSizes: [10, 20, 50, 100],
      params: { pi: 1, ps: 10 }, //页码，页容量控制
      totalcount: 0, //数据总条数
      editData: {}, //编辑数据对象
      addData: {}, //添加数据对象
      dayformat:"yyyy-MM-dd",
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
          new Date().getDate(),
      },
      dtCreateTime: this.DateConvert("yyyy-MM-dd 00:00:00", new Date()),
      notifyStatus: this.EnumUtility.Get("NotifyStatus"),
      notifyType: this.EnumUtility.Get("NotifyType"),
      tableData: []
    };
  },
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
    detail(orderId){
      let val = {order_id:orderId}
      val.getpath ="/oms/order/info"
      this.$emit("addTab","详情"+orderId,"/oms/order/info.view",val);
      },
    generalStatusClassFilter(item) {
        switch (item) {
          case '0':
            return this.colorClass[0]
            break;
          case '90':
            return this.colorClass[1]
            break;
          case '20':
            return this.colorClass[2]
            break;
          default:
            return this.colorClass[3]
            break;
        }
      },
    /**初始化操作**/
    init() {
      this.query();
    },
    /**查询数据并赋值*/
    query(value) {
       if (value == 1) {
        this.queryData.pi = 1;
      } else {
        this.queryData.pi = this.params.pi;
      }
      this.queryData.ps = this.params.ps;
      this.queryData.start_time =
        this.DateConvert(this.dayformat, this.queryData.start_time) +
        " 00:00:00";
      this.queryData.end_time =
        this.DateConvert(this.dayformat, this.queryData.end_time) + " 23:59:59";
      this.$http
        .get("/oms/notify/info/query", this.queryData)
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
        getpath: "/oms/notify/info",
        notify_id: val.notify_id
      };
      this.$emit(
        "addTab",
        "详情" + val.notify_id,
        "/oms/notify/info.view",
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
</style>
