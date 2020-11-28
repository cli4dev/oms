<template>
	<el-dialog title="编辑渠道基本信息" width="45%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="140px">
      
			<el-form-item label="是否支持查询:" prop="can_query">
				<el-select  placeholder="---请选择---" clearable v-model="editData.can_query" style="width: 100%;">
					<el-option v-for="(item, index) in canQuery" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="渠道名称:" prop="channel_no">
				<el-select  placeholder="---请选择---" clearable v-model="editData.channel_no" style="width: 100%;">
					<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="预留字段" prop="ext_params">
				<el-input clearable v-model="editData.ext_params" placeholder="请输入预留字段">
				</el-input>
      </el-form-item>
      
      <el-form-item label="首次查询时间" prop="first_query_time">
				<el-input clearable v-model="editData.first_query_time" placeholder="请输入首次查询时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="通知回调地址" prop="notify_url">
				<el-input clearable v-model="editData.notify_url" placeholder="请输入通知回调地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询后补间隔时间" prop="query_replenish_time">
				<el-input clearable v-model="editData.query_replenish_time" placeholder="请输入查询后补间隔时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="查询地址" prop="query_url">
				<el-input clearable v-model="editData.query_url" placeholder="请输入查询地址">
				</el-input>
      </el-form-item>
      
      <el-form-item label="发货后补间隔时间" prop="request_replenish_time">
				<el-input clearable v-model="editData.request_replenish_time" placeholder="请输入发货后补间隔时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="上游请求地址" prop="request_url">
				<el-input clearable v-model="editData.request_url" placeholder="请输入上游请求地址">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="服务类型 :" prop="service_class">
				<el-select  placeholder="---请选择---" clearable v-model="editData.service_class" style="width: 100%;">
					<el-option v-for="(item, index) in serviceClass" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      
			<el-form-item label="状态:" prop="status">
				<el-select  placeholder="---请选择---" clearable v-model="editData.status" style="width: 100%;">
					<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name" ></el-option>
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
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
			canQuery:this.EnumUtility.Get("CanQuery"),
			channelNo:this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary"),
			serviceClass:this.EnumUtility.Get("ServiceClass"),
			status:this.EnumUtility.Get("Status"),
			rules: {                    //数据验证规则
				can_query: [
					{ required: true, message: "请输入是否支持查询", trigger: "blur" }
				],
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
				status: [
					{ required: true, message: "请输入状态", trigger: "blur" }
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
		show() {
			this.dialogFormVisible = true;
		},
		edit() {
			this.$http.put("/vds/channel/info", this.editData)
				.then(res => {
					this.$message({
						type: "success",
						message: "修改成功!"
					});
					this.dialogFormVisible = false;
					this.refresh()
				})
				.catch(err => {
					this.$message({
						type: "error",
						message: err.response.data
					});
				})
		},
	}
}
</script>

<style scoped>
</style>
