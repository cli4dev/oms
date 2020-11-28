<template>
<el-table
	:data="info"
	style="width: 100%" :show-header="false" :default-expand-all="true">
	<el-table-column type="expand">
		<template slot-scope="props">
			<el-form label-position="left" inline class="demo-table-expand" label-width="120px">
			<el-form-item label="人工审核编号">
				<span v-text="props.row.audit_id"></span>
			</el-form-item>
			<el-form-item label="审核人">
				<span v-text="props.row.audit_by"></span>
			</el-form-item>
			<el-form-item label="审核信息">
				<span v-text="props.row.audit_msg"></span>
			</el-form-item>
			<el-form-item label="审核状态">
				<span > {{props.row.audit_status| EnumFilter("AuditStatus") }}</span>
			</el-form-item>
			<el-form-item label="审核时间">
				<span v-text="props.row.audit_time"></span>
			</el-form-item>
			<el-form-item label="变动类型">
				<span>{{props.row.change_type | EnumFilter("ChangeType")}}</span>
			</el-form-item>
			<el-form-item label="创建时间">
				<span v-text="props.row.create_time"></span>
			</el-form-item>
			<el-form-item label="发货记录编号">
				<span v-text="props.row.delivery_id"></span>
			</el-form-item>
			<el-form-item label="订单编号">
				<span v-text="props.row.order_id"></span>
			</el-form-item>
			<el-form-item label="退款编号">
				<span v-text="props.row.refund_id"></span>
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
		this.EnumUtility.Get("ChangeType"),
		this.EnumUtility.Get("AuditStatus")
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
