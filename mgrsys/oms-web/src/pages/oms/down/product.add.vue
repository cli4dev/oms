<template>
  <!-- Add Form -->
  <el-dialog title="添加下游商品" width="45%" :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="140px">
      <el-form-item label="渠道名称:" prop="channel_no">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="addData.channel_no"
          filterable
          @change="ShelfList(addData.channel_no)"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in channelNo"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="货架:" prop="shelf_id">
        <el-select
          placeholder="---请选择---"
          clearable
          filterable
          v-model="addData.shelf_id"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in List"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="产品线:" prop="line_id">
        <el-select placeholder="---请选择---" clearable v-model="addData.line_id" style="width: 100%;">
          <el-option
            v-for="(item, index) in lineId"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="支持退款:" prop="can_refund">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="addData.can_refund"
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
          v-model="addData.invoice_type"
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

      <el-form-item label="是否允许拆单:" prop="can_split_order">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="addData.can_split_order"
          style="width: 100%;"
        >
          <el-option
            v-for="(item, index) in canSplitOrder"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="拆单面值" prop="split_order_face" v-if="addData.can_split_order == '0'">
        <el-input clearable v-model="addData.split_order_face" placeholder="请输入拆单面值"></el-input>
      </el-form-item>
      <el-form-item label="面值" prop="face">
        <el-input clearable v-model="addData.face" placeholder="请输入面值"></el-input>
      </el-form-item>

      <el-form-item label="单次最大购买数量" prop="limit_count">
        <el-input clearable v-model="addData.limit_count" placeholder="请输入单次最大购买数量"></el-input>
      </el-form-item>

      <el-form-item label="运营商:" prop="carrier_no">
        <el-select
          placeholder="---请选择---"
          clearable
          v-model="addData.carrier_no"
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
          @change="setCityNO(addData.province_no)"
          v-model="addData.province_no"
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
        <el-select placeholder="---请选择---" clearable v-model="addData.city_no" style="width: 100%;">
          <el-option value="*" label="全部"></el-option>
          <el-option
            v-for="(item, index) in CityList"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="手续费折扣" prop="payment_fee_discount">
        <el-input clearable v-model="addData.payment_fee_discount" placeholder="请输入手续费折扣"></el-input>
      </el-form-item>

      <el-form-item label="佣金折扣" prop="commission_discount">
        <el-input clearable v-model="addData.commission_discount" placeholder="请输入佣金折扣"></el-input>
      </el-form-item>

      <el-form-item label="销售折扣" prop="sell_discount">
        <el-input clearable v-model="addData.sell_discount" placeholder="请输入销售折扣"></el-input>
      </el-form-item>

      <el-form-item label="服务费折扣" prop="service_discount">
        <el-input clearable v-model="addData.service_discount" placeholder="请输入服务费折扣"></el-input>
      </el-form-item>
      <el-form-item label="状态:" prop="status">
        <el-select placeholder="---请选择---" clearable v-model="addData.status" style="width: 100%;">
          <el-option
            v-for="(item, index) in status"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="外部商品编号" prop="ext_product_no">
        <el-input clearable v-model="addData.ext_product_no" placeholder="请输入外部商品编号"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
  <!--Add Form -->
</template>

<script>
export default {
  name: "product.add",
  data() {
    return {
      addData: { shelf_id: "" },
      channelNo: [],
      ProvinceList: [],
      dialogAddVisible: false,
      shelfId: this.EnumUtility.Get(
        "DownShelfId",
        {},
        "/oms/down/shelf/getdictionary"
      ),
      lineId: this.EnumUtility.Get(
        "LineId",
        {},
        "/oms/product/line/getdictionary"
      ),
      canRefund: this.EnumUtility.Get("CanRefund"),
      invoiceType: this.EnumUtility.Get("InvoiceType"),
      canSplitOrder: this.EnumUtility.Get("CanSplitOrder"),
      carrierNo: this.EnumUtility.Get("CarrierNo"),

      city_no: [],
      CityList: [],
      List: [],
      status: this.EnumUtility.Get("Status"),
      rules: {
        channel_no: [
          { required: true, message: "请输入渠道编号", trigger: "blur" }
        ],
        //数据验证规则
        shelf_id: [
          { required: true, message: "请输入货架编号", trigger: "blur" }
        ],
        line_id: [{ required: true, message: "请输入产品线", trigger: "blur" }],
        can_refund: [
          { required: true, message: "请输入支持退款", trigger: "blur" }
        ],
        invoice_type: [
          { required: true, message: "请输入开票方式", trigger: "blur" }
        ],
        can_split_order: [
          { required: true, message: "请输入是否允许拆单", trigger: "blur" }
        ],
        create_time: [
          { required: true, message: "请输入创建时间", trigger: "blur" }
        ],
        face: [{ required: true, message: "请输入面值", trigger: "blur" }],
        limit_count: [
          { required: true, message: "请输入单次最大购买数量", trigger: "blur" }
        ],
        carrier_no: [
          { required: true, message: "请输入运营商", trigger: "blur" }
        ],
        province_no: [
          { required: true, message: "请输入省份", trigger: "blur" }
        ],
        city_no: [{ required: true, message: "请输入城市", trigger: "blur" }],
        payment_fee_discount: [
          { required: true, message: "请输入手续费折扣", trigger: "blur" }
        ],
        commission_discount: [
          { required: true, message: "请输入佣金折扣", trigger: "blur" }
        ],
        sell_discount: [
          { required: true, message: "请输入销售折扣", trigger: "blur" }
        ],
        service_discount: [
          { required: true, message: "请输入服务费折扣", trigger: "blur" }
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
  mounted() {
    this.city();
  },
  created() {},
  methods: {
    channelList() {
      this.$http.get("/oms/down/channel/getdictionary", {}).then(res => {
        this.channelNo = res;
      });
    },
    ShelfList(value) {
      this.addData.shelf_id = "";
      this.$http
        .post("/oms/down/shelf/getchannel", { channel_no: value })
        .then(res => {
          this.List = res;
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
    resetForm(formName) {
      this.dialogAddVisible = false;
      this.$refs[formName].resetFields();
    },
    show() {
      this.channelList();
      this.dialogAddVisible = true;
    },
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
      this.$http
        .post("/oms/canton/info/getdictionarybyprovince", {
          grade: 2,
          canton_code: value
        })
        .then(res => {
          this.CityList = res;
        })
        .catch(err => {
          this.$message({
            type: "error",
            message: err.response.data
          });
        });
    },

    add(formName) {
      this.addData.create_time = this.DateConvert(
        "yyyy-MM-dd hh:mm:ss",
        this.addData.create_time
      );
      this.$refs[formName].validate(valid => {
        if (valid) {
          if (this.addData.can_split_order == "1") {
            this.addData.split_order_face = "0";
          }
          this.$http
            .post("/oms/down/product", this.addData)
            .then(res => {
              this.$refs[formName].resetFields();
              this.dialogAddVisible = false;
              this.refresh();
            })
            .catch(err => {
              this.$message({
                type: "error",
                message: err.response.data
              });
            });
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    }
  }
};
</script>

<style scoped>
</style>
