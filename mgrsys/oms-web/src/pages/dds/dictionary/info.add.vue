<template>
  <!-- Add Form -->
  <el-dialog title="添加字典表" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="名称" prop="name">
				<el-input clearable v-model="addData.name" placeholder="请输入名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="排序值" prop="sort_no">
				<el-input clearable v-model="addData.sort_no" placeholder="请输入排序值">
				</el-input>
      </el-form-item>
      
      <el-form-item label="类型" prop="type">
				<el-input clearable v-model="addData.type" placeholder="请输入类型">
				</el-input>
      </el-form-item>
      
      <el-form-item label="值" prop="value">
				<el-input clearable v-model="addData.value" placeholder="请输入值">
				</el-input>
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
			dialogAddVisible:false,
			rules: {                    //数据验证规则
				name: [
					{ required: true, message: "请输入名称", trigger: "blur" }
				],
				sort_no: [
					{ required: true, message: "请输入排序值", trigger: "blur" }
				],
				type: [
					{ required: true, message: "请输入类型", trigger: "blur" }
				],
				value: [
					{ required: true, message: "请输入值", trigger: "blur" }
				],
			},
		}
	},
	props: {
		refresh: {
			type: Function,
				default: () => {
				},
		}
	},
	created(){
	},
	methods: {
		closed() {
			this.refresh()
		},
		resetForm(formName) {
			this.dialogAddVisible = false;
			this.$refs[formName].resetFields();
		},
		show(){
			this.dialogAddVisible = true;
		},
		add(formName) {
			this.$refs[formName].validate((valid) => {
				if (valid) {
					this.$http.post("/dds/dictionary/info", this.addData)
						.then(res => {
							this.$refs[formName].resetFields()
							this.dialogAddVisible = false
							this.refresh()
						})
						.catch(err => {
							this.$message({
								type: "error",
								message: err.response.data
							});
						})
				} else {
						console.log("error submit!!");
						return false;
				}
			});
		},
	}

}
</script>

<style scoped>
</style>
