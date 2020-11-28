package product

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateLine 创建产品线
type CreateLine struct {
	LineName string `json:"line_name" form:"line_name" m2s:"line_name" valid:"required"` //LineName 产品线名称

}

//UpdateLine 添加产品线
type UpdateLine struct {
	LineId               string `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                             //LineId 产品线编号
	LineName             string `json:"line_name" form:"line_name" m2s:"line_name" valid:"required"`                       //LineName 产品线名称
	BindQueue            string `json:"bind_queue" form:"bind_queue" m2s:"bind_queue"`                                     //BindQueue 绑定队列
	DeliveryQueue        string `json:"delivery_queue" form:"delivery_queue" m2s:"delivery_queue"`                         //DeliveryFinishQueue 发货结束队列
	DeliveryUnknownQueue string `json:"delivery_unknown_queue" form:"delivery_unknown_queue" m2s:"delivery_unknown_queue"` //DeliveryUnknownQueue 发货未知处理队列
	NotifyQueue          string `json:"notify_queue" form:"notify_queue" m2s:"notify_queue"`                               //NotifyQueue 通知队列
	OrderOvertimeQueue   string `json:"order_overtime_queue" form:"order_overtime_queue" m2s:"order_overtime_queue"`       //OrderOvertimeQueue 订单超时队列
	OrderRefundQueue     string `json:"order_refund_queue" form:"order_refund_queue" m2s:"order_refund_queue"`             //OrderRefundQueue 订单失败退款队列
	PaymentQueue         string `json:"payment_queue" form:"payment_queue" m2s:"payment_queue"`                            //PaymentQueue 支付队列
	RefundNotifyQueue    string `json:"refund_notify_queue" form:"refund_notify_queue" m2s:"refund_notify_queue"`          //RefundNotifyQueue 退款通知队列
	RefundOvertimeQueue  string `json:"refund_overtime_queue" form:"refund_overtime_queue" m2s:"refund_overtime_queue"`    //RefundOvertimeQueue 退款超时处理队列
	RefundQueue          string `json:"refund_queue" form:"refund_queue" m2s:"refund_queue"`                               //RefundQueue 退款队列
	ReturnQueue          string `json:"return_queue" form:"return_queue" m2s:"return_queue"`                               //ReturnQueue 退货队列
	ReturnUnknownQueue   string `json:"return_unknown_queue" form:"return_unknown_queue" m2s:"return_unknown_queue"`       //ReturnUnknownQueue 退货未知处理队列
	UpPaymentQueue       string `json:"up_payment_queue" form:"up_payment_queue" m2s:"up_payment_queue"`                   //UpPaymentQueue 上游支付队列
	UpRefundQueue        string `json:"up_refund_queue" form:"up_refund_queue" m2s:"up_refund_queue"`                      //UpRefundQueue 上游退款队列

}

//QueryLine 查询产品线
type QueryLine struct {
	LineName string `json:"line_name" form:"line_name" m2s:"line_name"` //LineName 产品线名称

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbLine  产品线接口
type IDbLine interface {
	//Create 创建
	Create(input *CreateLine) error
	//Get 单条查询
	Get(lineId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryLine) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateLine) (err error)
	//GetLineDictionary 获取数据字典
	GetLineDictionary() (db.QueryRows, error)
}

//DbLine 产品线对象
type DbLine struct {
	c component.IContainer
}

//NewDbLine 创建产品线对象
func NewDbLine(c component.IContainer) *DbLine {
	return &DbLine{
		c: c,
	}
}

//GetLineDictionary 获取数据字典
func (d *DbLine) GetLineDictionary() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsProductLineDictionary, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取产品线数据字典发生错误")
	}
	return data, nil
}

//Create 添加产品线
func (d *DbLine) Create(input *CreateLine) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertOmsProductLine, map[string]interface{}{
		"line_name": input.LineName,
	})
	if err != nil {
		return fmt.Errorf("添加产品线数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据产品线
func (d *DbLine) Get(lineId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsProductLine, map[string]interface{}{
		"line_id": lineId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取产品线数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取产品线列表
func (d *DbLine) Query(input *QueryLine) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsProductLineCount, map[string]interface{}{
		"line_name": input.LineName,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取产品线列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsProductLine, map[string]interface{}{
		"line_name": input.LineName,
		"pi":        input.Pi,
		"ps":        input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取产品线数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新产品线
func (d *DbLine) Update(input *UpdateLine) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsProductLine, map[string]interface{}{
		"line_id":                input.LineId,
		"line_name":              input.LineName,
		"bind_queue":             input.BindQueue,
		"delivery_queue":         input.DeliveryQueue,
		"delivery_unknown_queue": input.DeliveryUnknownQueue,
		"notify_queue":           input.NotifyQueue,
		"order_overtime_queue":   input.OrderOvertimeQueue,
		"order_refund_queue":     input.OrderRefundQueue,
		"payment_queue":          input.PaymentQueue,
		"refund_notify_queue":    input.RefundNotifyQueue,
		"refund_overtime_queue":  input.RefundOvertimeQueue,
		"refund_queue":           input.RefundQueue,
		"return_queue":           input.ReturnQueue,
		"return_unknown_queue":   input.ReturnUnknownQueue,
		"up_payment_queue":       input.UpPaymentQueue,
		"up_refund_queue":        input.UpRefundQueue,
	})
	if err != nil {
		return fmt.Errorf("更新产品线数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
