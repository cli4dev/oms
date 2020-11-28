<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="货架编号">
				<span v-text="props.row.shelf_id"></span>
			</el-form-item>
			<el-form-item label="货架名称">
				<span v-text="props.row.shelf_name"></span>
			</el-form-item>
			<el-form-item label="渠道名称">
				<span>{{props.row.channel_no | EnumFilter("UpChannelNo")}}</span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="发货超时时间">
				<span v-text="props.row.delivery_overtime"></span>
			</el-form-item>
			<el-form-item label="退货超时时间">
				<span v-text="props.row.return_overtime"></span>
			</el-form-item>
			<el-form-item label="货架状态">
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
		this.EnumUtility.Get("UpChannelNo",{},"/oms/up/channel/getdictionary")
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
