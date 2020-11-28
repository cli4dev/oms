<template>
  <!-- Add Form -->
  <el-dialog title="添加上游货架" width="25%" :visible.sync="dialogAddVisible">
    <el-form :model="addData"  :rules="rules" ref="addForm" label-width="110px">
      <el-form-item label="货架名称" prop="shelf_name">
				<el-input clearable v-model="addData.shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="渠道名称:" prop="channel_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.channel_no" filterable style="width: 100%;">
					<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="发货超时时间" prop="delivery_overtime">
				<el-input clearable v-model="addData.delivery_overtime" placeholder="请输入发货超时时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退货超时时间" prop="return_overtime">
				<el-input clearable v-model="addData.return_overtime" placeholder="请输入退货超时时间">
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
  name: "shelf.add",
	data() {
		return {
			addData: {},
			dialogAddVisible:false,
			channelNo:this.EnumUtility.Get("UpChannelNo",{},"/oms/up/channel/getdictionary"),
			rules: {                    //数据验证规则
				shelf_name: [
					{ required: true, message: "请输入货架名称", trigger: "blur" }
				],
				channel_no: [
					{ required: true, message: "请输入渠道名称", trigger: "blur" }
				],
				delivery_overtime: [
					{ required: true, message: "请输入发货超时时间", trigger: "blur" }
				],
				return_overtime: [
					{ required: true, message: "请输入退货超时时间", trigger: "blur" }
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
					this.$http.post("/oms/up/shelf", this.addData)
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
