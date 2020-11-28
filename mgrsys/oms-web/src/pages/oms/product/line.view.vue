<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="产品线编号">
				<span v-text="props.row.line_id"></span>
			</el-form-item>
			<el-form-item label="产品线名称">
				<span v-text="props.row.line_name"></span>
			</el-form-item>
			<el-form-item label="绑定队列">
				<span v-text="props.row.bind_queue"></span>
			</el-form-item>
			<el-form-item label="发货队列">
				<span v-text="props.row.delivery_queue"></span>
			</el-form-item>
			<el-form-item label="发货未知处理队列">
				<span v-text="props.row.delivery_unknown_queue"></span>
			</el-form-item>
			<el-form-item label="通知队列">
				<span v-text="props.row.notify_queue"></span>
			</el-form-item>
			<el-form-item label="订单超时队列">
				<span v-text="props.row.order_overtime_queue"></span>
			</el-form-item>
			<el-form-item label="订单失败退款队列">
				<span v-text="props.row.order_refund_queue"></span>
			</el-form-item>
			<el-form-item label="支付队列">
				<span v-text="props.row.payment_queue"></span>
			</el-form-item>
			<el-form-item label="退款通知队列">
				<span v-text="props.row.refund_notify_queue"></span>
			</el-form-item>
			<el-form-item label="退款超时处理队列">
				<span v-text="props.row.refund_overtime_queue"></span>
			</el-form-item>
			<el-form-item label="退款队列">
				<span v-text="props.row.refund_queue"></span>
			</el-form-item>
			<el-form-item label="退货队列">
				<span v-text="props.row.return_queue"></span>
			</el-form-item>
			<el-form-item label="退货未知处理队列">
				<span v-text="props.row.return_unknown_queue"></span>
			</el-form-item>
			<el-form-item label="上游支付队列">
				<span v-text="props.row.up_payment_queue"></span>
			</el-form-item>
			<el-form-item label="上游退款队列">
				<span v-text="props.row.up_refund_queue"></span>
			</el-form-item>
			</el-form>
		</template>
	</el-table-column>
	<el-table-column
	label=""
	>
	</el-table-column>
	<el-table-column
	label=""
	>
	</el-table-column>
	<el-table-column
	label=""
	>
	</el-table-column>
	</el-table>
</template>


<script>
export default {

  data(){
		return {
      info:[],
    }
	},
	created(){
  },
  mounted() {
    this.Init()
  },
  methods: {
    Init(){
        this.QueryDataList()
    },
    QueryDataList(){
      this.$http.get(this.$route.query.getpath, 
          this.$route.query
        ).then(res => {
			this.info.push(res)
      });
    },
  
    handleClick(tab){
     
    },
   
  },

 
}
</script>

<style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
