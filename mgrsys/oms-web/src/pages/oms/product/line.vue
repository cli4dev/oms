<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-input clearable v-model="queryData.line_name" placeholder="请输入产品线名称">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-button type="primary" @click="query" size="small">查询</el-button>
				</el-form-item>
				
				<el-form-item>
					<el-button type="success" size="small" @click="addShow">添加</el-button>
				</el-form-item>
				
		</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="tableData" border style="width: 100%">
				<el-table-column prop="line_id" label="编号" align="center" min-width="60px"></el-table-column>
				<el-table-column prop="line_name" label="名称" align="center" min-width="100px" >
					<template slot-scope="scope">
						<span>{{scope.row.line_name | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="bind_queue" label="绑定队列"  align="center">
					<template slot-scope="scope">
						<span>{{scope.row.bind_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="delivery_queue"  align="center"  label="发货队列" >
					<template slot-scope="scope">
						<span>{{scope.row.delivery_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="delivery_unknown_queue"  align="center"  label="发货未知处理队列" >
					<template slot-scope="scope">
						<span>{{scope.row.delivery_unknown_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="notify_queue"  align="center" label="通知队列" >
					<template slot-scope="scope">
						<span>{{scope.row.notify_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="order_overtime_queue"  align="center"  label="订单超时队列" >
					<template slot-scope="scope">
						<span>{{scope.row.order_overtime_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="order_refund_queue"  align="center"  label="订单失败退款队列" >
					<template slot-scope="scope">
						<span>{{scope.row.order_refund_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="payment_queue"   align="center"  label="支付队列" >
					<template slot-scope="scope">
						<span>{{scope.row.payment_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="refund_notify_queue"  align="center"  label="退款通知队列" >
					<template slot-scope="scope">
						<span>{{scope.row.refund_notify_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="refund_overtime_queue"  align="center"  label="退款超时处理队列" >
					<template slot-scope="scope">
						<span>{{scope.row.refund_overtime_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="refund_queue"  align="center" label="退款队列" >
					<template slot-scope="scope">
						<span>{{scope.row.refund_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="return_queue"  align="center"  label="退货队列" >
					<template slot-scope="scope">
						<span>{{scope.row.return_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="return_unknown_queue"  align="center"  label="退货未知处理队列" >
					<template slot-scope="scope">
						<span>{{scope.row.return_unknown_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="up_payment_queue"  align="center"  label="上游支付队列" >
					<template slot-scope="scope">
						<span>{{scope.row.up_payment_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="up_refund_queue"  align="center"  label="上游退款队列" >
					<template slot-scope="scope">
						<span>{{scope.row.up_refund_queue | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column  label="操作"  align="center">
					<template slot-scope="scope">
						<el-button type="text" size="small" @click="editShow(scope.row)">编辑</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		<!-- Add Form -->
		<Add ref="Add" :refresh="query"></Add>
		<!--Add Form -->

		<!-- edit Form start-->
		<Edit ref="Edit" :refresh="query"></Edit>
		<!-- edit Form end-->

		<!-- pagination start -->
		<div class="page-pagination">
		<el-pagination
			@size-change="handleSizeChange"
			@current-change="handleCurrentChange"
			:current-page="params.pi"
			:page-size="params.ps"
			:page-sizes="pageSizes"
			layout="total, sizes, prev, pager, next, jumper"
			:total="totalcount">
		</el-pagination>
		</div>
		<!-- pagination end -->

	</div>
</template>

<script>
import Add from "./line.add"
import Edit from "./line.edit"
export default {
  name: "OmsProductLine",
  components: {
		Add,
		Edit
  },
  data () {
		return {
			pageSizes: [10, 20, 50, 100], 
			params:{pi:1,ps:10},        //页码，页容量控制
			totalcount: 0,              //数据总条数
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
			queryData:{},
			tableData: [],
		}
  },
  created(){
  },
  mounted(){
	  
    this.init()
  },
	methods:{
    /**初始化操作**/
    init(){
      this.query()
		},
    /**查询数据并赋值*/
    query(){
      this.queryData.pi = this.params.pi
			this.queryData.ps = this.params.ps
      this.$http.get("/oms/product/line/query", this.queryData)
  			.then(res => {
          this.tableData = res.data
          this.totalcount = res.count
        })
        .catch(err => {
          console.log(err)
      	})
    },
    /**改变页容量*/
		handleSizeChange(val) {
      this.params.ps = val
      this.query()
    },
    /**改变当前页码*/
    handleCurrentChange(val) {
      this.params.pi = val
      this.query()
    },
    /**重置添加表单*/
    resetForm(formName) {
      this.dialogAddVisible = false
      this.$refs[formName].resetFields();
		},
		detailShow(val){
			var data = {
        getpath: "/oms/product/line",
        line_id: val.line_id,
      }
      this.$emit("addTab","详情"+val.line_id,"/oms/product/line.view",data);
		},
    addShow(){
      this.$refs.Add.show();
		},
    editShow(val) {
      this.$refs.Edit.editData = val
      this.$refs.Edit.show();
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
