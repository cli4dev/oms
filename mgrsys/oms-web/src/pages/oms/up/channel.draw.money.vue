<template>
  <el-dialog title="上游渠道加款" width="20%" :visible.sync="dialogAddVisible">
    <el-form :model="addData" :rules="rules" ref="addForm" label-width="100px">
      <el-form-item label="金额" prop="amount">
        <el-input clearable v-model="addData.amount" placeholder="请输入金额"></el-input>
      </el-form-item>
    </el-form>
    <div slot="footer" class="dialog-footer">
      <el-button size="small" @click="resetForm('addForm')">取 消</el-button>
      <el-button size="small" type="success" @click="add('addForm')">确 定</el-button>
    </div>
  </el-dialog>
</template>

<script>
export default {
  name: "channel.add.balance",
  data() {
    return {
      addData: {
        channel_no: ""
      },
      dialogAddVisible: false,
      rules: {
        //数据验证规则
        amount: [{ required: true, message: "请输入金额", trigger: "blur" }]
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
      this.refresh();
    },
    resetForm(formName) {
      this.dialogAddVisible = false;
      this.$refs[formName].resetFields();
    },
    show(val) {
      this.dialogAddVisible = true;
      this.addData.channel_no = val;
    },
    add(formName) {
      this.$refs[formName].validate(valid => {
        this.addData.types = "up_channel";
        if (valid) {
          this.$http
            .post("/beanpay/account/info/draw", this.addData)
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
