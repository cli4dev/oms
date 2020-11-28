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
			<el-form-item label="渠道名称">
				<span>{{props.row.channel_no | EnumFilter("ChannelNo")}}</span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="处理码">
				<span v-text="props.row.deal_code"></span>
			</el-form-item>
			<el-form-item label="错误码">
				<span v-text="props.row.error_code"></span>
			</el-form-item>
			<el-form-item label="错误码描述">
				<span v-text="props.row.error_code_desc"></span>
			</el-form-item>
			<el-form-item label="服务类型 ">
				<span>{{props.row.service_class | EnumFilter("ServiceClass")}}</span>
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
		this.EnumUtility.Get("ChannelNo",{},"/oms/up/channel/getdictionary")
		this.EnumUtility.Get("ServiceClass")
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
