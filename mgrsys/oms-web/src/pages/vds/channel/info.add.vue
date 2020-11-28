<template>
  <!-- Add Form -->
  <el-dialog title="添加渠道基本信息" width="45%"  :visible.sync="dialogAddVisible">
    <el-form :model="addData" :inline="true" :rules="rules" ref="addForm" label-width="140px">
      
			<el-form-item label="渠道名称:" prop="channel_no">
				<el-select  placeholder="---请选择---" clearable v-model="addData.channel_no" style="width: 100%;">
					<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="预留字段" prop="ext_params">
				<el-input clearable v-model="addData.ext_params" placeholder="请输入预留字段">
				</el-input>
      </el-form-item>
      
      <el-form-item label="首次查询时间" prop="first_query_time">
				<el-input clearable v-model="addData.first_query_time" placeholder="请输入首次查询时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="通知回调地址" prop="notify_url">
				<el-input clearable v-model="addData.notify_url" placeholder="请输入通知回调地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询后补间隔时间" prop="query_replenish_time">
				<el-input clearable v-model="addData.query_replenish_time" placeholder="请输入查询后补间隔时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询地址" prop="query_url">
				<el-input clearable v-model="addData.query_url" placeholder="请输入查询地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="发货后补间隔时间" prop="request_replenish_time">
				<el-input clearable v-model="addData.request_replenish_time" placeholder="请输入发货后补间隔时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="上游请求地址" prop="request_url">
				<el-input clearable v-model="addData.request_url" placeholder="请输入上游请求地址">
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
  name: "info.add",
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
				notify_url: [
					{ required: true, message: "请输入通知回调地址", trigger: "blur" }
				],
				request_replenish_time: [
					{ required: true, message: "请输入发货后补间隔时间", trigger: "blur" }
				],
				request_url: [
					{ required: true, message: "请输入上游请求地址", trigger: "blur" }
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
					this.$http.post("/vds/channel/info", this.addData)
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
