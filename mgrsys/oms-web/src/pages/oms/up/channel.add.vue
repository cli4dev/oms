<template>
  <!-- Add Form -->
  <el-dialog title="添加上游渠道" width="20%" :visible.sync="dialogAddVisible">
    <el-form :model="addData" :rules="rules" ref="addForm" label-width="100px">
      <el-form-item label="编号" prop="channel_no">
        <el-input clearable v-model="addData.channel_no" placeholder="请输入编号"></el-input>
      </el-form-item>

      <el-form-item label="名称" prop="channel_name">
        <el-input clearable v-model="addData.channel_name" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item label="外部渠道编号" prop="ext_channel_no">
        <el-input clearable v-model="addData.ext_channel_no" placeholder="请输入外部渠道编号"></el-input>
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
        channel_no: [
          { required: true, message: "请输入编号", trigger: "blur" }
        ],
        channel_name: [
          { required: true, message: "请输入名称", trigger: "blur" }
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
          this.$http
            .post("/oms/up/channel", this.addData)
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
