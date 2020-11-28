package refund

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryUpReturn 查询上游退货信息表
type QueryUpReturn struct {
	CarrierNo      string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`                   //CarrierNo 运营商
	ProvinceNo     string `json:"province_no" form:"province_no" m2s:"province_no"`                //ProvinceNo 省份
	CityNo         string `json:"city_no" form:"city_no" m2s:"city_no"`                            //CityNo 城市
	DownChannelNo  string `json:"down_channel_no" form:"down_channel_no" m2s:"down_channel_no"`    //DownChannelNo 下游渠道
	LineId         string `json:"line_id" form:"line_id" m2s:"line_id"`                            //LineId 产品线
	ReturnStatus   string `json:"return_status" form:"return_status" m2s:"return_status"`          //ReturnStatus 退货状态（0.退货成功，20等待退货，30正在退货，90退货失败）
	UpChannelNo    string `json:"up_channel_no" form:"up_channel_no" m2s:"up_channel_no"`          //UpChannelNo 上游渠道
	UpRefundStatus string `json:"up_refund_status" form:"up_refund_status" m2s:"up_refund_status"` //UpRefundStatus 退款状态（0退款成功，10.未开始，20.等待退款，30.正在退款，99无需退款）
	CreateTime     string `json:"create_time" form:"create_time" m2s:"create_time"`                //CreateTime 创建时间
	StartTime      string `json:"start_time" form:"start_time" m2s:"start_time"`
	EndTime        string `json:"end_time" form:"end_time" m2s:"end_time"`
	Pi             string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps             string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbUpReturn  上游退货信息表接口
type IDbUpReturn interface {
	//Get 单条查询
	Get(returnId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryUpReturn) (data db.QueryRows, count int, err error)
}

//DbUpReturn 上游退货信息表对象
type DbUpReturn struct {
	c component.IContainer
}

//NewDbUpReturn 创建上游退货信息表对象
func NewDbUpReturn(c component.IContainer) *DbUpReturn {
	return &DbUpReturn{
		c: c,
	}
}

//Get 查询单条数据上游退货信息表
func (d *DbUpReturn) Get(returnId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsRefundUpReturn, map[string]interface{}{
		"return_id": returnId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取上游退货信息表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取上游退货信息表列表
func (d *DbUpReturn) Query(input *QueryUpReturn) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	fmt.Println(input.CreateTime, "shijian1")
	c, q, a, err := db.Scalar(sql.QueryOmsRefundUpReturnCount, map[string]interface{}{
		"carrier_no":       input.CarrierNo,
		"province_no":      input.ProvinceNo,
		"city_no":          input.CityNo,
		"down_channel_no":  input.DownChannelNo,
		"line_id":          input.LineId,
		"return_status":    input.ReturnStatus,
		"up_channel_no":    input.UpChannelNo,
		"up_refund_status": input.UpRefundStatus,
		"end_time":         input.EndTime,
		"start_time":       input.StartTime,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游退货信息表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsRefundUpReturn, map[string]interface{}{
		"carrier_no":       input.CarrierNo,
		"province_no":      input.ProvinceNo,
		"city_no":          input.CityNo,
		"down_channel_no":  input.DownChannelNo,
		"line_id":          input.LineId,
		"return_status":    input.ReturnStatus,
		"end_time":         input.EndTime,
		"start_time":       input.StartTime,
		"up_channel_no":    input.UpChannelNo,
		"up_refund_status": input.UpRefundStatus,
		"pi":               input.Pi,
		"ps":               input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游退货信息表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
