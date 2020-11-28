<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="id">
				<span v-text="props.row.id"></span>
			</el-form-item>
			<el-form-item label="是否支持查询">
				<span>{{props.row.can_query | EnumFilter("CanQuery")}}</span>
			</el-form-item>
			<el-form-item label="渠道名称">
				<span>{{props.row.channel_no | EnumFilter("ChannelNo")}}</span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="预留字段">
				<span v-text="props.row.ext_params"></span>
			</el-form-item>
			<el-form-item label="首次查询时间">
				<span v-text="props.row.first_query_time"></span>
			</el-form-item>
			<el-form-item label="通知回调地址">
				<span v-text="props.row.notify_url"></span>
			</el-form-item>
			<el-form-item label="查询后补间隔时间">
				<span v-text="props.row.query_replenish_time"></span>
			</el-form-item>
			<el-form-item label="查询地址">
				<span v-text="props.row.query_url"></span>
			</el-form-item>
			<el-form-item label="发货后补间隔时间">
				<span v-text="props.row.request_replenish_time"></span>
			</el-form-item>
			<el-form-item label="上游请求地址">
				<span v-text="props.row.request_url"></span>
			</el-form-item>
			<el-form-item label="服务类型 ">
				<span>{{props.row.service_class | EnumFilter("ServiceClass")}}</span>
			</el-form-item>
			<el-form-item label="状态">
				<span>{{props.row.status | EnumFilter("Status")}}</span>
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
		this.EnumUtility.Get("CanQuery")
		this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary")
		this.EnumUtility.Get("ServiceClass")
		this.EnumUtility.Get("Status")
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
