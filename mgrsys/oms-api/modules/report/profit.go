package report

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryReportProfit 利润报表
type QueryReportProfit struct {
	CarrierNO string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`
	LineID    string `json:"line_id" form:"line_id" m2s:"line_id"`
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no"`
	StartTime string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime   string `json:"end_time" form:"end_time" m2s:"end_time"`
	TotalFace string `json:"total_face" form:"total_face" m2s:"total_face"`
	Pi        string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps        string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//ExportReportProfit 利润报表
type ExportReportProfit struct {
	CarrierNO string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`
	LineID    string `json:"line_id" form:"line_id" m2s:"line_id"`
	ChannelNO string `json:"channel_no" form:"channel_no" m2s:"channel_no"`
	StartTime string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime   string `json:"end_time" form:"end_time" m2s:"end_time"`
	TotalFace string `json:"total_face" form:"total_face" m2s:"total_face"`
}

//IDbReportProfit  利润报表接口
type IDbReportProfit interface {
	//Query 列表查询
	Query(input *QueryReportProfit) (data db.QueryRows, count int, total db.QueryRow, err error)

	Export(input *ExportReportProfit) (datas [][]string, header []string, err error)

	Query4Export(input *ExportReportProfit) (data db.QueryRows, err error)
}

//DbReportProfit 利润报表对象
type DbReportProfit struct {
	c component.IContainer
}

//NewDbReportProfit 利润报表对象
func NewDbReportProfit(c component.IContainer) *DbReportProfit {
	return &DbReportProfit{
		c: c,
	}
}

//Query4Export 获取利润报表列表
func (d *DbReportProfit) Query4Export(input *ExportReportProfit) (data db.QueryRows, err error) {
	params, err := types.Struct2Map(input)
	if err != nil {
		return nil, fmt.Errorf("Struct2Map Error(err:%v)", err)
	}
	db := d.c.GetRegularDB()

	data, q, a, err := db.Query(sql.Reportprofit4Export, params)
	if err != nil {
		return nil, fmt.Errorf("获取利润报表Excel数据失败(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}

func (d *DbReportProfit) Export(input *ExportReportProfit) (data [][]string, header []string, err error) {
	datas, err := d.Query4Export(input)
	if err != nil {
		return nil, nil, err
	}

	header = []string{"完成时间", "下游渠道", "总笔数", "总面值", "成功面值", "销售总金额", "系统成本", "手续费", "利润"}
	for _, item := range datas {
		list := []string{}
		list = append(list, item.GetString("create_date"))
		list = append(list, item.GetString("channel_name"))
		list = append(list, item.GetString("order_count"))
		list = append(list, item.GetString("total_face"))
		list = append(list, item.GetString("success_face"))
		list = append(list, item.GetString("success_sell"))
		list = append(list, item.GetString("success_cost"))
		list = append(list, item.GetString("success_fee"))
		list = append(list, item.GetString("profit"))
		data = append(data, list)
	}

	return
}

//Query 获取利润列表
func (d *DbReportProfit) Query(input *QueryReportProfit) (data db.QueryRows, count int, total db.QueryRow, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryReportProfitCount, map[string]interface{}{
		"carrier_no": input.CarrierNO,
		"line_id":    input.LineID,
		"channel_no": input.ChannelNO,
		"start_time": input.StartTime,
		"end_time":   input.EndTime,
		"total_face": input.TotalFace,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取利润条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryReportProfitList, map[string]interface{}{
		"carrier_no": input.CarrierNO,
		"line_id":    input.LineID,
		"channel_no": input.ChannelNO,
		"start_time": input.StartTime,
		"end_time":   input.EndTime,
		"total_face": input.TotalFace,
		"pi":         input.Pi,
		"ps":         input.Ps,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取利润数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	tot, q, a, err := db.Query(sql.QueryReportProfitTotal, map[string]interface{}{
		"carrier_no": input.CarrierNO,
		"line_id":    input.LineID,
		"channel_no": input.ChannelNO,
		"start_time": input.StartTime,
		"end_time":   input.EndTime,
		"total_face": input.TotalFace,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取利润合计表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.GetInt(c, 0), tot.Get(0), nil
}
