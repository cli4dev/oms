<template>
	<el-dialog title="编辑上游货架" width="25%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="货架名称" prop="shelf_name">
				<el-input clearable v-model="editData.shelf_name" placeholder="请输入货架名称">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="渠道名称:" prop="channel_no">
				<el-select  placeholder="---请选择---" clearable v-model="editData.channel_no" filterable style="width: 100%;">
					<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name" ></el-option>
				</el-select>
			</el-form-item>
      
      <el-form-item label="发货超时时间" prop="delivery_overtime">
				<el-input clearable v-model="editData.delivery_overtime" placeholder="请输入发货超时时间">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退货超时时间" prop="return_overtime">
				<el-input clearable v-model="editData.return_overtime" placeholder="请输入退货超时时间">
				</el-input>
      </el-form-item>
      
      
			<el-form-item label="货架状态:" prop="status">
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
	name: "shelf.edit",
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
			channelNo:this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary"),
			status:this.EnumUtility.Get("Status"),
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
				status: [
					{ required: true, message: "请输入货架状态", trigger: "blur" }
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
			this.$http.put("/oms/up/shelf", this.editData)
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
