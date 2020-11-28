<template>
  <!-- Add Form -->
  <el-dialog title="查看秘钥" width="25%" :visible.sync="dialogAddVisible">
    <el-form  ref="addForm" label-width="100px">
      <el-form-item label="秘钥:" prop="secret">{{list.secret}}</el-form-item>
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
      channel_no: "",
      secret: "",
      list:{
          secret:"",
      },
      dialogAddVisible: false,
     
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
      this.add(val);
      this.dialogAddVisible = true;
    },
    add(val) {
      this.$http
        .get("/oms/up/channel/getsecret", { channel_no: val})
        .then(res => {
          this.list = JSON.parse(JSON.parse(res.data));
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
