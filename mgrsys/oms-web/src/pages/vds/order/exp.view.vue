<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="编号">
				<span v-text="props.row.id"></span>
			</el-form-item>
			<el-form-item label="运营商">
				<span>{{props.row.carrier_no | EnumFilter("CarrierNo")}}</span>
			</el-form-item>
			<el-form-item label="上游渠道">
				<span>{{props.row.channel_no | EnumFilter("ChannelNo")}}</span>
			</el-form-item>
			<el-form-item label="下游商户">
				<span>{{props.row.coop_id | EnumFilter("CoopId")}}</span>
			</el-form-item>
			<el-form-item label="下游商户订单号">
				<span v-text="props.row.coop_order_id"></span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="订单结果消息">
				<span v-text="props.row.error_msg"></span>
			</el-form-item>
			<el-form-item label="原串">
				<span v-text="props.row.ext_params"></span>
			</el-form-item>
			<el-form-item label="收单Ip">
				<span v-text="props.row.local_ip"></span>
			</el-form-item>
			<el-form-item label="产品面值">
				<span>{{props.row.product_face/100}}</span>
			</el-form-item>
			<el-form-item label="产品数量">
				<span v-text="props.row.product_num"></span>
			</el-form-item>
			<el-form-item label="服务类型 ">
				<span v-text="props.row.service_class"></span>
			</el-form-item>
			<el-form-item label="用户Ip">
				<span v-text="props.row.user_ip"></span>
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
		this.EnumUtility.Get("CarrierNo")
		this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary")
		this.EnumUtility.Get("CoopId",{},"/oms/down/channel/getdictionary")
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
