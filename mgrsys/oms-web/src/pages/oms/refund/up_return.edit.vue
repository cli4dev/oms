<template>
	<el-dialog title="编辑上游退货信息表" width="25%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData"  :rules="rules" ref="editForm" label-width="110px">
    </el-form>
		<div slot="footer" class="dialog-footer">
			<el-button size="small" @click="dialogFormVisible = false">取 消</el-button>
			<el-button type="success" size="small" @click="edit">确 定</el-button>
		</div>
	</el-dialog>
</template>

<script>
export default {
	name: "up_return.edit",
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
			rules: {                    //数据验证规则
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
			this.$http.put("/oms/refund/up/return", this.editData)
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
