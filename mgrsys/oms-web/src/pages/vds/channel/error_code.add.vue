<template>
  <!-- Add Form -->
  <el-dialog title="添加渠道错误码" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      
			<el-form-item label="渠道名称:" prop="channel_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.channel_no" style="width: 100%;">
					<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="处理码" prop="deal_code">
				<el-input clearable v-model="addData.deal_code" placeholder="请输入处理码">
				</el-input>
      </el-form-item>
      
      <el-form-item label="错误码" prop="error_code">
				<el-input clearable v-model="addData.error_code" placeholder="请输入错误码">
				</el-input>
      </el-form-item>
      
      <el-form-item label="错误码描述" prop="error_code_desc">
				<el-input clearable v-model="addData.error_code_desc" placeholder="请输入错误码描述">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="服务类型 :" prop="service_class">
				<el-select  placeholder="---请选择---" clearable v-model="addData.service_class" style="width: 100%;">
					<el-option v-for="(item, index) in serviceClass" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
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
  name: "error_code.add",
	data() {
		return {
			addData: {},
			dialogAddVisible:false,
			channelNo:this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary"),
			serviceClass:this.EnumUtility.Get("ServiceClass"),
			rules: {                    //数据验证规则
				channel_no: [
					{ required: true, message: "请输入渠道名称", trigger: "blur" }
				],
				deal_code: [
					{ required: true, message: "请输入处理码", trigger: "blur" }
				],
				error_code: [
					{ required: true, message: "请输入错误码", trigger: "blur" }
				],
				error_code_desc: [
					{ required: true, message: "请输入错误码描述", trigger: "blur" }
				],
				service_class: [
					{ required: true, message: "请输入服务类型 ", trigger: "blur" }
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
					this.$http.post("/vds/channel/error/code", this.addData)
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
