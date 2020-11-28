<template>
  <el-dialog title="编辑上游渠道" width="20%" @closed="closed" :visible.sync="dialogFormVisible">
    <el-form :model="editData" :rules="rules" ref="editForm" label-width="100px">
      <el-form-item label="名称" prop="channel_name">
        <el-input clearable v-model="editData.channel_name" placeholder="请输入名称"></el-input>
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
      <el-form-item label="外部渠道编号" prop="ext_channel_no">
        <el-input clearable v-model="editData.ext_channel_no" placeholder="请输入外部渠道编号"></el-input>
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
  name: "channel.edit",
  data() {
    return {
      dialogFormVisible: false, //编辑表单显示隐藏
      editData: {}, //编辑数据对象
      status: this.EnumUtility.Get("Status"),
      rules: {
        //数据验证规则
        channel_name: [
          { required: true, message: "请输入名称", trigger: "blur" }
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
  created() {},
  methods: {
    closed() {
      this.refresh();
    },
    show() {
      this.dialogFormVisible = true;
    },
    edit() {
      this.$http
        .put("/oms/up/channel", this.editData)
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
