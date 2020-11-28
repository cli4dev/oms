<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="通知编号">
				<span v-text="props.row.id"></span>
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
			<el-form-item label="完成时间">
				<span v-text="props.row.finish_time"></span>
			</el-form-item>
			<el-form-item label="通知内容">
				<span v-text="props.row.notify_content"></span>
			</el-form-item>
			<el-form-item label="通知次数">
				<span v-text="props.row.notify_count"></span>
			</el-form-item>
			<el-form-item label="通知限制次数">
				<span v-text="props.row.notify_limit_count"></span>
			</el-form-item>
			<el-form-item label="下游通知回调地址">
				<span v-text="props.row.notify_url"></span>
			</el-form-item>
			<el-form-item label="订单号">
				<span v-text="props.row.order_no"></span>
			</el-form-item>
			<el-form-item label="通知结果信息">
				<span v-text="props.row.result_msg"></span>
			</el-form-item>
			<el-form-item label="通知状态">
				<span>{{props.row.status | EnumFilter("NotifyStatus")}}</span>
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
		this.EnumUtility.Get("CoopId",{},"/oms/down/channel/getdictionary")
		this.EnumUtility.Get("NotifyStatus")
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
