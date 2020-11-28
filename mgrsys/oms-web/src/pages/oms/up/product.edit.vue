<template>
  <el-dialog title="编辑上游商品" width="45%" @closed="closed" :visible.sync="dialogFormVisible">
    <el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="140px">
      <el-form-item label="货架名称:" prop="shelf_id">
        <el-select
          placeholder="---请选择---"
          clearable
          filterable
          v-model="editData.shelf_id"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in shelfId"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="产品线:" prop="line_id">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="editData.line_id"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in lineId"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="运营商:" prop="carrier_no">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="editData.carrier_no"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in carrierNo"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="省份:" prop="province_no">
        <el-select
          placeholder="---请选择---"
          clearable
          @change="setCityNO(editData.province_no)"
          v-model="editData.province_no"
          style="width: 100%;"
        >
          <el-option value="*" label="全部"></el-option>
          <el-option
            v-for="(item, index) in ProvinceList"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="城市:" prop="city_no">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="editData.city_no"
          style="width: 100%;"
        >
          <el-option value="*" label="全部"></el-option>
          <el-option
            v-for="(item, index) in cityNo"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="成本折扣" prop="cost_discount">
        <el-input clearable v-model="editData.cost_discount" placeholder="请输入成本折扣"></el-input>
      </el-form-item>

      <el-form-item label="佣金折扣" prop="commission_discount">
        <el-input clearable v-model="editData.commission_discount" placeholder="请输入佣金折扣"></el-input>
      </el-form-item>

      <el-form-item label="服务费折扣" prop="service_discount">
        <el-input clearable v-model="editData.service_discount" placeholder="请输入服务费折扣"></el-input>
		</el-form-item>
        <el-form-item label="支持退货:" prop="can_refund">
          <el-select
            placeholder="---请选择---"
            clearable
            v-model="editData.can_refund"
            style="width: 100%;"
          >
            <el-option
              v-for="(item, index) in canRefund"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="开票方式:" prop="invoice_type">
          <el-select
            placeholder="---请选择---"
            clearable
            v-model="editData.invoice_type"
            style="width: 100%;"
          >
            <el-option
              v-for="(item, index) in invoiceType"
              :key="index"
              :value="item.value"
              :label="item.name"
            ></el-option>
          </el-select>
        </el-form-item>
      <el-form-item label="面值" prop="face">
        <el-input clearable v-model="editData.face" placeholder="请输入面值"></el-input>
      </el-form-item>

      <el-form-item label="单次最大发货数量" prop="limit_count">
        <el-input clearable v-model="editData.limit_count" placeholder="请输入单次最大发货数量"></el-input>
      </el-form-item>

      <el-form-item label="状态:" prop="status">
        <el-select placeholder="---请选择---" clearable v-model="editData.status" style="width: 100%;">
          <el-option
            v-for="(item, index) in status"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>
	  <el-form-item label="外部商品编号" prop="ext_product_no">
        <el-input clearable v-model="editData.ext_product_no" placeholder="请输入外部商品编号"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="dialogFormVisible = false">取 消</el-button>
      <el-button type="success" size="small" @click="edit">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  name: "product.edit",
  data() {
    return {
      ProvinceList: [],
      CityList: [],
      dialogFormVisible: false, //编辑表单显示隐藏
      editData: {}, //编辑数据对象
      canRefund: this.EnumUtility.Get("CanRefund"),
      invoiceType: this.EnumUtility.Get("InvoiceType"),
      carrierNo: this.EnumUtility.Get("CarrierNo"),
cityNo:this.EnumUtility.Get("CityNo",{},"/oms/canton/info/getlist"),
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
      rules: {
        //数据验证规则
        can_refund: [
          { required: true, message: "请输入支持退货", trigger: "blur" }
        ],
        invoice_type: [
          { required: true, message: "请输入开票方式", trigger: "blur" }
        ],
        carrier_no: [
          { required: true, message: "请输入运营商", trigger: "blur" }
        ],
        province_no: [
          { required: true, message: "请输入省份", trigger: "blur" }
        ],
        city_no: [{ required: true, message: "请输入城市", trigger: "blur" }],
        cost_discount: [
          { required: true, message: "请输入成本折扣", trigger: "blur" }
        ],
        commission_discount: [
          { required: true, message: "请输入佣金折扣", trigger: "blur" }
        ],
        service_discount: [
          { required: true, message: "请输入服务费折扣", trigger: "blur" }
        ],
        face: [{ required: true, message: "请输入面值", trigger: "blur" }],
        limit_count: [
          { required: true, message: "请输入单次最大发货数量", trigger: "blur" }
        ],
        line_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
        shelf_id: [
          { required: true, message: "请输入货架名称", trigger: "blur" }
        ],
        status: [{ required: true, message: "请输入状态", trigger: "blur" }]
      }
    };
  },
  props: {
    refresh: {
      type: Function,
      default: () => {}
    }
  },
  created() {

  },
  mounted() {
    this.city();
  },
  methods: {
    city() {
      this.$http
        .post("/oms/canton/info/getdictionary", { grade: 1 })
        .then(res => {
          this.ProvinceList = res;
        })
        .catch(err => {
          this.$message({
            type: "error",
            message: err.response.data
          });
        });
    },
    setCityNO(value) {
		this.editData.city_no="",
      this.$http
        .post("/oms/canton/info/getdictionarybyprovince", {
          grade: 2,
          canton_code: value
        })
        .then(res => {
          this.cityNo = res;
        })
        .catch(err => {
          this.$message({
            type: "error",
            message: err.response.data
          });
        });
    },
    closed() {
      this.refresh();
    },
    show() {
      this.dialogFormVisible = true;
    },
    edit() {
      this.$http
        .put("/oms/up/product", this.editData)
        .then(res => {
          this.$message({
            type: "success",
            message: "修改成功!"
          });
          this.dialogFormVisible = false;
          this.refresh();
        })
        .catch(err => {
          this.$message({
            type: "error",
            message: err.response.data
          });
        });
    }
  }
};
</script>

<style scoped>
</style>
