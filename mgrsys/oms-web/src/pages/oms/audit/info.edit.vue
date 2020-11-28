<template>
  <el-dialog title="人工审核" width="25%" @closed="closed" :visible.sync="dialogFormVisible">
    <el-form :model="editData" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="审核类型">
        <span>{{editData.change_type | EnumFilter("AuditType")}}</span>
      </el-form-item>

      <el-form-item label="审核状态:" prop="audit_status">
        <el-select placeholder="---请选择---" clearable v-model="status" style="width: 100%;">
          <el-option
            v-for="(item, index) in audit_status"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="审核信息" prop="audit_msg">
        <el-input type="textarea" clearable v-model="editData.audit_msg" placeholder="请输入审核信息"></el-input>
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
  name: "info.edit",
  data() {
    return {
      status: "",
      dialogFormVisible: false, //编辑表单显示隐藏
      editData: {}, //编辑数据对象
      audit_status: [],
      deliveryAuditStatus: [
        { name: "审核成功", value: "0" },
        { name: "审核失败", value: "90" }
      ],
      returnAuditStatus: [
        { name: "审核成功", value: "0" },
        { name: "审核失败", value: "90" }
      ],
      orderAuditStatus: [{ name: "审核部分成功", value: "91" }],
      refundAuditStatus: [{ name: "审核审核部分成功", value: "91" }],
      rules: {
        //数据验证规则
        audit_status: [
          { required: true, message: "请输入审核状态", trigger: "blur" }
        ],
        audit_msg: [
          { required: true, message: "请输入审核信息", trigger: "blur" }
        ]
      }
    };
  },
  props: {
    refresh: {
      type: Function,
      default: () => {}
    }
  },
  created() {},
  methods: {
    closed() {
      this.status = "";
      this.audit_msg = "";
      this.refresh();
    },
    show() {
      this.dialogFormVisible = true;
      switch (this.editData.change_type) {
        case "1":
          return (this.audit_status = this.deliveryAuditStatus);
          break;
        case "2":
          this.audit_status = this.returnAuditStatus;
          break;
        case "3":
          this.audit_status = this.orderAuditStatus;
          break;
        case "4":
          this.audit_status = this.refundAuditStatus;
          break;
      }
    },
    edit() {
      this.editData.audit_time = this.DateConvert(
        "yyyy-MM-dd hh:mm:ss",
        this.editData.audit_time
      );
      this.editData.audit_status = this.status;

      this.$http
        .put("/oms/audit/info", this.editData)
        .then(res => {
          this.$message({
            type: "success",
            message: "修改成功!"
          });
          this.status = "";
          this.audit_msg = "";
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
