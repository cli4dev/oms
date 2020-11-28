<template>
  <div class="panel panel-default">
    <!-- query start -->
    <div class="panel-body">
      <el-form ref="form" :inline="true" class="form-inline pull-left">
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.carrier_no"
            class="input-cos"
            placeholder="请选择运营商"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in carrierNo"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.line_id"
            class="input-cos"
            placeholder="请选择产品线"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in lineId"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
            size="medium"
            v-model="queryData.shelf_id"
            class="input-cos"
            filterable
            placeholder="请选择货架名称"
          >
            <el-option value label="全部"></el-option>
            <el-option
              v-for="(item, index) in shelfId"
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
        <el-table-column prop="product_id" label="商品编号" align="center"></el-table-column>
        <el-table-column prop="shelf_id" label="货架名称" align="center" min-width="120px">
          <template slot-scope="scope">
            <span>{{scope.row.shelf_id | EnumFilter("ShelfId")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="line_id" label="产品线" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.line_id | EnumFilter("LineId")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="carrier_no" label="运营商" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.carrier_no | EnumFilter("CarrierNo")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="province_no" label="省份" align="center">
          <template slot-scope="scope">
            <span>{{ProvinceName(scope.row.province_no) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="city_no" label="城市" align="center">
          <template slot-scope="scope">
            <span>{{CityName(scope.row.city_no) }}</span>
          </template>
        </el-table-column>
        <el-table-column prop="cost_discount" label="成本折扣" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.cost_discount |FeeFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="service_discount" label="服务费折扣" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.service_discount |FeeFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="commission_discount" label="佣金折扣" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.commission_discount |FeeFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="face" label="面值" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.face |AmountFilter}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="can_refund" label="支持退货" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.can_refund | EnumFilter("CanRefund")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="invoice_type" label="开票方式" align="center">
          <template slot-scope="scope">
            <span>{{scope.row.invoice_type | EnumFilter("InvoiceType")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="status" label="状态" align="center">
          <template slot-scope="scope">
            <span
              :class="[scope.row.status === '0' ? 'text-success' : 'text-danger']"
            >{{scope.row.status | EnumFilter("Status")}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="create_time" label="创建时间" align="center" min-width="120px"></el-table-column>
        <el-table-column label="操作" align="center">
          <template slot-scope="scope">
            <el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
            <el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button>
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
import Add from "./product.add";
import Edit from "./product.edit";
export default {
  name: "OmsUpProduct",
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
      canRefund: this.EnumUtility.Get("CanRefund"),
      invoiceType: this.EnumUtility.Get("InvoiceType"),
      carrierNo: this.EnumUtility.Get("CarrierNo"),
      provinceNo: this.EnumUtility.Get(
        "ProvinceNo",
        { grade: 1 },
        "/oms/canton/info/getlist"
      ),
      cityNo: this.EnumUtility.Get("CityNo", {}, "/oms/canton/info/getlist"),
      lineId: this.EnumUtility.Get(
        "LineId",
        {},
        "/oms/product/line/getdictionary"
      ),
      shelfId: this.EnumUtility.Get(
        "ShelfId",
        {},
        "/oms/up/shelf/getdictionary"
      ),
      status: this.EnumUtility.Get("Status"),
      tableData: []
    };
  },
  created() {},
  mounted() {
    this.init();
  },
  methods: {
    ProvinceName(value) {
      var res = value;
      if (res == "*") {
        return (res = "全省");
      }
      this.cityNo.forEach(item => {
        if (res == item.value) {
          res = item.name;
        }
      });
      return res;
    },
    CityName(value) {
      var res = value;
      if (res == "*") {
        return (res = "全市");
      }
      this.cityNo.forEach(item => {
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
      this.$http
        .get("/oms/up/product/query", this.queryData)
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
        getpath: "/oms/up/product",
        product_id: val.product_id
      };
      this.$emit(
        "addTab",
        "详情" + val.product_id,
        "/oms/up/product.view/" + val.product_id,
        data
      );
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
