<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.channel_no" class="input-cos" placeholder="请选择上游渠道">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.status" class="input-cos" placeholder="请选择状态">
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
				<el-table-column prop="order_no" label="订单号" ></el-table-column>
				<el-table-column prop="channel_no" label="上游渠道" >
					<template slot-scope="scope">
						<span>{{scope.row.channel_no | EnumFilter("ChannelNo")}}</span>
					</template>
				
					<template slot-scope="scope">
						<span>{{scope.row.channel_no | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="coop_id" label="下游商户编号" >
					<template slot-scope="scope">
						<span>{{scope.row.coop_id | EnumFilter("CoopId")}}</span>
					</template>
				
					<template slot-scope="scope">
						<span>{{scope.row.coop_id | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="create_time" label="创建时间" ></el-table-column>
				<el-table-column prop="last_query_time" label="最近查询时间" ></el-table-column>
				<el-table-column prop="query_count" label="查询次数" ></el-table-column>
				<el-table-column prop="query_result" label="查询请求结果" >
					<template slot-scope="scope">
						<span>{{scope.row.query_result | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="status" label="状态" >
					<template slot-scope="scope">
						<span>{{scope.row.status | EnumFilter("Status")}}</span>
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
  name: "VdsOrderQuery",
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
			channelNo:this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary"),
			status:this.EnumUtility.Get("Status"),
			tableData: [],
		}
  },
  created(){
		this.EnumUtility.Get("CoopId",{},"/oms/down/channel/getdictionary")
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
      this.$http.get("/vds/order/query/query", this.queryData)
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
        getpath: "/vds/order/query",
        order_no: val.order_no,
      }
      this.$emit("addTab","详情"+val.order_no,"/vds/order/query.view",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
