package report

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryReportDownchannel 下游交易报表
type QueryReportDownchannel struct {
	CarrierNO  string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`
	LineID     string `json:"line_id" form:"line_id" m2s:"line_id"`
	ChannelNO  string `json:"channel_no" form:"channel_no" m2s:"channel_no"`
	ProvinceNO string `json:"province_no" form:"province_no" m2s:"province_no"`
	StartTime  string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime    string `json:"end_time" form:"end_time" m2s:"end_time"`
	TotalFace  string `json:"total_face" form:"total_face" m2s:"total_face"`
	Pi         string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps         string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

type ExportReportDownchannel struct {
	CarrierNO  string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`
	LineID     string `json:"line_id" form:"line_id" m2s:"line_id"`
	ChannelNO  string `json:"channel_no" form:"channel_no" m2s:"channel_no"`
	ProvinceNO string `json:"province_no" form:"province_no" m2s:"province_no"`
	StartTime  string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime    string `json:"end_time" form:"end_time" m2s:"end_time"`
	TotalFace  string `json:"total_face" form:"total_face" m2s:"total_face"`
}

//IDbReportDownchannel  下游交易报表接口
type IDbReportDownchannel interface {
	//Query 列表查询
	Query(input *QueryReportDownchannel) (data db.QueryRows, count int, total db.QueryRow, err error)

	Export(input *ExportReportDownchannel) (datas [][]string, header []string, err error)

	Query4Export(input *ExportReportDownchannel) (data db.QueryRows, err error)
}

//DbReportDownchannel 下游交易报表对象
type DbReportDownchannel struct {
	c component.IContainer
}

//NewDbReportDownchannel 下游交易报表对象
func NewDbReportDownchannel(c component.IContainer) *DbReportDownchannel {
	return &DbReportDownchannel{
		c: c,
	}
}

//Query4Export 获取下游交易列表
func (d *DbReportDownchannel) Query4Export(input *ExportReportDownchannel) (data db.QueryRows, err error) {
	params, err := types.Struct2Map(input)
	if err != nil {
		return nil, fmt.Errorf("Struct2Map Error(err:%v)", err)
	}
	db := d.c.GetRegularDB()

	data, q, a, err := db.Query(sql.ReportDownchannel4Export, params)
	if err != nil {
		return nil, fmt.Errorf("获取下游交易Excel数据失败(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, nil
}

func (d *DbReportDownchannel) Export(input *ExportReportDownchannel) (data [][]string, header []string, err error) {
	datas, err := d.Query4Export(input)
	if err != nil {
		return nil, nil, err
	}

	header = []string{"下游渠道", "总笔数", "失败笔数", "失败金额", "成功笔数", "总金额", "成功金额", "成功率", "实际扣款金额", "系统成本金额", "利润"}
	for _, item := range datas {
		list := []string{}
		list = append(list, item.GetString("channel_name"))
		list = append(list, item.GetString("order_count"))
		list = append(list, item.GetString("fail_count"))
		list = append(list, item.GetString("fail_face"))
		list = append(list, item.GetString("success_count"))
		list = append(list, item.GetString("total_face"))
		list = append(list, item.GetString("success_face"))
		list = append(list, item.GetString("success_ratio"))
		list = append(list, item.GetString("success_sell"))
		list = append(list, item.GetString("success_cost"))
		list = append(list, item.GetString("profit"))
		data = append(data, list)
	}

	return
}

//Query 获取下游交易列表
func (d *DbReportDownchannel) Query(input *QueryReportDownchannel) (data db.QueryRows, count int, total db.QueryRow, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryReportDownchannelCount, map[string]interface{}{
		"carrier_no":  input.CarrierNO,
		"line_id":     input.LineID,
		"channel_no":  input.ChannelNO,
		"province_no": input.ProvinceNO,
		"start_time":  input.StartTime,
		"end_time":    input.EndTime,
		"total_face":  input.TotalFace,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取下游交易条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryReportDownchannelList, map[string]interface{}{
		"carrier_no":  input.CarrierNO,
		"line_id":     input.LineID,
		"channel_no":  input.ChannelNO,
		"province_no": input.ProvinceNO,
		"start_time":  input.StartTime,
		"end_time":    input.EndTime,
		"total_face":  input.TotalFace,
		"pi":          input.Pi,
		"ps":          input.Ps,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取下游交易数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	tot, q, a, err := db.Query(sql.QueryReportDownchannelTotal, map[string]interface{}{
		"carrier_no":  input.CarrierNO,
		"line_id":     input.LineID,
		"channel_no":  input.ChannelNO,
		"province_no": input.ProvinceNO,
		"start_time":  input.StartTime,
		"end_time":    input.EndTime,
		"total_face":  input.TotalFace,
	})
	if err != nil {
		return nil, 0, nil, fmt.Errorf("获取下游交易数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}
	return data, types.GetInt(c, 0), tot.Get(0), nil
}
