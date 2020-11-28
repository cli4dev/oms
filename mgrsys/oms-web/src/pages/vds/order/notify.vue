<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.coop_id" class="input-cos" placeholder="请选择下游商户">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in coopId" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-input clearable v-model="queryData.coop_order_id" placeholder="请输入下游商户订单号">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-input clearable v-model="queryData.order_no" placeholder="请输入订单号">
					</el-input>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.status" class="input-cos" placeholder="请选择通知状态">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in status" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-button type="primary" @click="query" size="small">查询</el-button>
				</el-form-item>
				
		</el-form>
		</div>
    	<!-- query end -->

    	<!-- list start-->
		<el-scrollbar style="height:100%">
			<el-table :data="tableData" border style="width: 100%">
				<el-table-column prop="id" label="通知编号" ></el-table-column>
				<el-table-column prop="coop_id" label="下游商户" >
					<template slot-scope="scope">
						<span>{{scope.row.coop_id | EnumFilter("CoopId")}}</span>
					</template>
				
					<template slot-scope="scope">
						<span>{{scope.row.coop_id | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="coop_order_id" label="下游商户订单号" >
					<template slot-scope="scope">
						<span>{{scope.row.coop_order_id | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="create_time" label="创建时间" ></el-table-column>
				<el-table-column prop="finish_time" label="完成时间" ></el-table-column>
				<el-table-column prop="notify_count" label="通知次数" ></el-table-column>
				<el-table-column prop="notify_limit_count" label="通知限制次数" ></el-table-column>
				<el-table-column prop="notify_url" label="下游通知回调地址" >
					<template slot-scope="scope">
						<span>{{scope.row.notify_url | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="order_no" label="订单号" ></el-table-column>
				<el-table-column prop="status" label="通知状态" >
					<template slot-scope="scope">
						<span>{{scope.row.status | EnumFilter("NotifyStatus")}}</span>
					</template>
				</el-table-column>
				<el-table-column  label="操作">
					<template slot-scope="scope">
						<el-button type="text" size="small" @click="detailShow(scope.row)">详情</el-button>
					</template>
				</el-table-column>
			</el-table>
		</el-scrollbar>
		<!-- list end-->

		

		

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
export default {
  name: "VdsOrderNotify",
  components: {
  },
  data () {
		return {
			pageSizes: [10, 20, 50, 100], 
			params:{pi:1,ps:10},        //页码，页容量控制
			totalcount: 0,              //数据总条数
			editData:{},                //编辑数据对象
			addData:{},                 //添加数据对象 
			queryData:{},
			coopId:this.EnumUtility.Get("CoopId",{},"/oms/down/channel/getdictionary"),
			status:this.EnumUtility.Get("NotifyStatus"),
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
      this.$http.get("/vds/order/notify/query", this.queryData)
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
        getpath: "/vds/order/notify",
        id: val.id,
      }
      this.$emit("addTab","详情"+val.id,"/vds/order/notify.view",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
