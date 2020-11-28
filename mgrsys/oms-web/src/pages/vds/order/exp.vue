<template>
	<div class="panel panel-default">
    	<!-- query start -->
		<div class="panel-body">
			<el-form ref="form" :inline="true" class="form-inline pull-left">
				<el-form-item>
					<el-select size="medium" v-model="queryData.carrier_no" class="input-cos" placeholder="请选择运营商">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in carrierNo" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.channel_no" class="input-cos" placeholder="请选择上游渠道">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in channelNo" :key="index" :value="item.value" :label="item.name"></el-option>
						</el-select>
				</el-form-item>
			
				<el-form-item>
					<el-select size="medium" v-model="queryData.coop_id" class="input-cos" placeholder="请选择下游商户">
						<el-option value="" label="全部"></el-option>
						<el-option v-for="(item, index) in coopId" :key="index" :value="item.value" :label="item.name"></el-option>
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
				<el-table-column prop="id" label="编号" ></el-table-column>
				<el-table-column prop="carrier_no" label="运营商" >
					<template slot-scope="scope">
						<span>{{scope.row.carrier_no | EnumFilter("CarrierNo")}}</span>
					</template>
				
					<template slot-scope="scope">
						<span>{{scope.row.carrier_no | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="channel_no" label="上游渠道" >
					<template slot-scope="scope">
						<span>{{scope.row.channel_no | EnumFilter("ChannelNo")}}</span>
					</template>
				
					<template slot-scope="scope">
						<span>{{scope.row.channel_no | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
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
				<el-table-column prop="local_ip" label="收单Ip" >
					<template slot-scope="scope">
						<span>{{scope.row.local_ip | EllipsisFilter(20)}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="product_face" label="产品面值" >
					<template slot-scope="scope">
						<span>{{scope.row.product_face/100}}</span>
					</template>
				</el-table-column>
				<el-table-column prop="product_num" label="产品数量" ></el-table-column>
				<el-table-column prop="service_class" label="服务类型 " ></el-table-column>
				<el-table-column prop="user_ip" label="用户Ip" >
					<template slot-scope="scope">
						<span>{{scope.row.user_ip | EllipsisFilter(20)}}</span>
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
  name: "VdsOrderExp",
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
			carrierNo:this.EnumUtility.Get("CarrierNo"),
			channelNo:this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary"),
			coopId:this.EnumUtility.Get("CoopId",{},"/oms/down/channel/getdictionary"),
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
      this.$http.get("/vds/order/exp/query", this.queryData)
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
        getpath: "/vds/order/exp",
        id: val.id,
      }
      this.$emit("addTab","详情"+val.id,"/vds/order/exp.view",data);
		},
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .page-pagination{padding: 10px 15px;text-align: right;}
</style>
