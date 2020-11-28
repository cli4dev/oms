<template>
	<el-dialog title="编辑产品线" width="65%" @closed="closed" :visible.sync="dialogFormVisible">
		<el-form :model="editData" :inline="true" :rules="rules" ref="editForm" label-width="110px">
      <el-form-item label="产品线名称" prop="line_name">
				<el-input clearable v-model="editData.line_name" placeholder="请输入产品线名称">
				</el-input>
      </el-form-item>
      
      <el-form-item label="绑定队列" prop="bind_queue">
				<el-input clearable v-model="editData.bind_queue" placeholder="请输入绑定队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="发货队列" prop="delivery_queue">
				<el-input clearable v-model="editData.delivery_queue" placeholder="请输入发货开始队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="发货未知处理队列" prop="delivery_unknown_queue">
				<el-input clearable v-model="editData.delivery_unknown_queue" placeholder="请输入发货未知处理队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="通知队列" prop="notify_queue">
				<el-input clearable v-model="editData.notify_queue" placeholder="请输入通知队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="订单超时队列" prop="order_overtime_queue">
				<el-input clearable v-model="editData.order_overtime_queue" placeholder="请输入订单超时队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="订单失败退款队列" prop="order_refund_queue">
				<el-input clearable v-model="editData.order_refund_queue" placeholder="请输入订单失败退款队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="支付队列" prop="payment_queue">
				<el-input clearable v-model="editData.payment_queue" placeholder="请输入支付队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退款通知队列" prop="refund_notify_queue">
				<el-input clearable v-model="editData.refund_notify_queue" placeholder="请输入退款通知队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退款超时处理队列" prop="refund_overtime_queue">
				<el-input clearable v-model="editData.refund_overtime_queue" placeholder="请输入退款超时处理队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退款队列" prop="refund_queue">
				<el-input clearable v-model="editData.refund_queue" placeholder="请输入退款队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退货队列" prop="return_queue">
				<el-input clearable v-model="editData.return_queue" placeholder="请输入退货队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="退货未知处理队列" prop="return_unknown_queue">
				<el-input clearable v-model="editData.return_unknown_queue" placeholder="请输入退货未知处理队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="上游支付队列" prop="up_payment_queue">
				<el-input clearable v-model="editData.up_payment_queue" placeholder="请输入上游支付队列">
				</el-input>
      </el-form-item>
      
      <el-form-item label="上游退款队列" prop="up_refund_queue">
				<el-input clearable v-model="editData.up_refund_queue" placeholder="请输入上游退款队列">
				</el-input>
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
	name: "line.edit",
	data() {
		return {
			dialogFormVisible: false,    //编辑表单显示隐藏
			editData: {},                //编辑数据对象
			rules: {                    //数据验证规则
				line_name: [
					{ required: true, message: "请输入产品线名称", trigger: "blur" }
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
			this.$http.put("/oms/product/line", this.editData)
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
