<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="帐户编号">
				<span v-text="props.row.account_id"></span>
			</el-form-item>
			<el-form-item label="帐户名称">
				<span v-text="props.row.account_name"></span>
			</el-form-item>
			<el-form-item label="帐户余额">
				<span>{{props.row.balance/100}}</span>
			</el-form-item>
			<el-form-item label="信用余额">
				<span >{{props.row.credit/100}}</span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="外部用户账户编号">
				<span v-text="props.row.eid"></span>
			</el-form-item>
			<el-form-item label="账户类型">
				<span>{{props.row.groups | EnumFilter("Groups")}}</span>
			</el-form-item>
			<el-form-item label="系统标识">
				<span v-text="props.row.ident"></span>
			</el-form-item>
			<el-form-item label="账户状态">
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
		this.EnumUtility.Get("Groups")
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
