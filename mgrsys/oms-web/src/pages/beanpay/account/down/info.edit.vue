<template>
  <el-dialog title="编辑账户信息" width="25%" @closed="closed" :visible.sync="dialogFormVisible">
    <el-form :model="editData" :rules="rules" ref="editForm" label-width="130px">
      <el-form-item label="帐户名称" prop="account_name">
        <el-input clearable v-model="editData.account_name" placeholder="请输入帐户名称"></el-input>
      </el-form-item>

      <el-form-item label="外部用户账户" prop="eid">
		   <el-select placeholder="---请选择---" clearable v-model="editData.eid" style="width: 100%;">
          <el-option
            v-for="(item, index) in queryList"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
        <!-- <el-input clearable v-model="editData.eid" placeholder="请输入外部用户账户编号"></el-input> -->
      </el-form-item>

      <el-form-item label="系统标识" prop="ident">
        <el-input clearable v-model="editData.ident" placeholder="请输入系统标识"></el-input>
      </el-form-item>

      <el-form-item label="账户状态:" prop="status">
        <el-select placeholder="---请选择---" clearable v-model="editData.status" style="width: 100%;">
          <el-option
            v-for="(item, index) in status"
            :key="index"
            :value="item.value"
            :label="item.name"
          ></el-option>
        </el-select>
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
      dialogFormVisible: false, //编辑表单显示隐藏
      editData: {}, //编辑数据对象
	  status: this.EnumUtility.Get("Status"),
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
        ident: [{ required: true, message: "请输入系统标识", trigger: "blur" }],
        status: [{ required: true, message: "请输入账户状态", trigger: "blur" }]
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
    show() {
		 this.queryList =this.upList.concat(this.downList)
      this.dialogFormVisible = true;
    },
    edit() {
      this.$http
        .put("/beanpay/account/info", this.editData)
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
