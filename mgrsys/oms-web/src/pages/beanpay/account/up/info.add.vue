<template>
  <!-- Add Form -->
  <el-dialog title="添加账户信息" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData" :rules="rules" ref="addForm" label-width="120px">
      <el-form-item label="帐户名称" prop="account_name">
        <el-input clearable v-model="addData.account_name" placeholder="请输入帐户名称"></el-input>
      </el-form-item>

      <el-form-item label="外部用户账户" prop="eid">
        <el-select placeholder="---请选择---" clearable v-model="addData.eid" style="width: 100%;">
          <el-option
            v-for="(item, index) in queryList"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="账户类型:" prop="groups">
        <el-select placeholder="---请选择---" clearable v-model="addData.groups" style="width: 100%;">
          <el-option
            v-for="(item, index) in list"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="系统标识" prop="ident">
        <el-input clearable v-model="addData.ident" placeholder="请输入系统标识"></el-input>
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
  name: "info.add",
  data() {
    return {
      addData: {},
      dialogAddVisible: false,
      groups: this.EnumUtility.Get("down"),
      up: this.EnumUtility.Get("up"),
      list: [],
      upList: [],
	  downList: [],
	  queryList:[],
      rules: {
        //数据验证规则
        account_name: [
          { required: true, message: "请输入帐户名称", trigger: "blur" }
        ],
        eid: [
          { required: true, message: "请输入外部用户账户编号", trigger: "blur" }
        ],
        groups: [
          { required: true, message: "请输入账户类型", trigger: "blur" }
        ],
        ident: [{ required: true, message: "请输入系统标识", trigger: "blur" }]
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
    this.$http
      .get("/oms/down/channel/getdictionary", {})
      .then(res => {
        this.downList = res;
      })
      .catch(err => {
        console.log(err);
      });
    this.$http
      .get("/oms/up/channel/getdictionary", {})
      .then(res => {
        this.upList = res;
      })
      .catch(err => {
        console.log(err);
      });
  },
  methods: {
    closed() {
      this.refresh();
    },
    resetForm(formName) {
      this.dialogAddVisible = false;
      this.$refs[formName].resetFields();
    },
    show() {
	  this.queryList =this.upList.concat(this.downList)
      this.dialogAddVisible = true;
    },
    add(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          this.$http
            .post("/beanpay/account/info", this.addData)
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
