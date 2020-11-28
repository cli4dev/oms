<template>
  <!-- Add Form -->
  <el-dialog title="设置秘钥" width="20%" :visible.sync="dialogAddVisible">
    <el-form :model="addData" :rules="rules" ref="addForm" label-width="100px">
      <el-form-item label="秘钥" prop="secret">
        <el-input clearable v-model="addData.secret" placeholder="请输入秘钥"></el-input>
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
  name: "channel.add",
  data() {
    return {
      addData: {},
      dialogAddVisible: false,
      rules: {
        //数据验证规则
        secret: [{ required: true, message: "秘钥不能为空", trigger: "blur" }]
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
    show() {
      this.dialogAddVisible = true;
    },
    add(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          var params = {};
          params.secret = this.addData.secret;
          params.channel_no = this.addData.channel_no;
          this.$http
            .get("/oms/down/channel/setsecret", params)
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
