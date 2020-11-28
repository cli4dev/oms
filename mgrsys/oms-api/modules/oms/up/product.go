package up

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateProduct 创建上游商品
type CreateProduct struct {
	CanRefund          string `json:"can_refund" form:"can_refund" m2s:"can_refund" valid:"required"`                            //CanRefund 支持退货
	InvoiceType        string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type" valid:"required"`                      //InvoiceType 开票方式（1.不开发票，2.上游开发票）
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no" valid:"required"`                            //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no" valid:"required"`                         //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no" valid:"required"`                                     //CityNo 城市
	CostDiscount       string `json:"cost_discount" form:"cost_discount" m2s:"cost_discount" valid:"required"`                   //CostDiscount 成本折扣（以面值算）
	CommissionDiscount string `json:"commission_discount" form:"commission_discount" m2s:"commission_discount" valid:"required"` //CommissionDiscount 佣金折扣（以面值算）
	ServiceDiscount    string `json:"service_discount" form:"service_discount" m2s:"service_discount" valid:"required"`          //ServiceDiscount 服务费折扣
	ExtProductNo       string `json:"ext_product_no" form:"ext_product_no" m2s:"ext_product_no"`                                 //ExtProductNo 外部商品编号
	Face               string `json:"face" form:"face" m2s:"face" valid:"required"`                                              //Face 面值
	LimitCount         string `json:"limit_count" form:"limit_count" m2s:"limit_count" valid:"required"`                         //LimitCount 单次最大发货数量
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                                     //LineId 产品线
	ShelfId            string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id" valid:"required"`                                  //ShelfId 货架名称

}

//UpdateProduct 添加上游商品
type UpdateProduct struct {
	ProductId          string `json:"product_id" form:"product_id" m2s:"product_id" valid:"required"`                            //ProductId 商品编号
	CanRefund          string `json:"can_refund" form:"can_refund" m2s:"can_refund" valid:"required"`                            //CanRefund 支持退货
	InvoiceType        string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type" valid:"required"`                      //InvoiceType 开票方式（1.不开发票，2.上游开发票）
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no" valid:"required"`                            //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no" valid:"required"`                         //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no" valid:"required"`                                     //CityNo 城市
	CostDiscount       string `json:"cost_discount" form:"cost_discount" m2s:"cost_discount" valid:"required"`                   //CostDiscount 成本折扣（以面值算）
	CommissionDiscount string `json:"commission_discount" form:"commission_discount" m2s:"commission_discount" valid:"required"` //CommissionDiscount 佣金折扣（以面值算）
	ServiceDiscount    string `json:"service_discount" form:"service_discount" m2s:"service_discount" valid:"required"`          //ServiceDiscount 服务费折扣
	ExtProductNo       string `json:"ext_product_no" form:"ext_product_no" m2s:"ext_product_no"`                                 //ExtProductNo 外部商品编号
	Face               string `json:"face" form:"face" m2s:"face" valid:"required"`                                              //Face 面值
	LimitCount         string `json:"limit_count" form:"limit_count" m2s:"limit_count" valid:"required"`                         //LimitCount 单次最大发货数量
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                                     //LineId 产品线
	ShelfId            string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id" valid:"required"`                                  //ShelfId 货架名称
	Status             string `json:"status" form:"status" m2s:"status" valid:"required"`                                        //Status 状态

}

//QueryProduct 查询上游商品
type QueryProduct struct {
	CanRefund   string `json:"can_refund" form:"can_refund" m2s:"can_refund"`       //CanRefund 支持退货
	InvoiceType string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type"` //InvoiceType 开票方式（1.不开发票，2.上游开发票）
	CarrierNo   string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`       //CarrierNo 运营商
	ProvinceNo  string `json:"province_no" form:"province_no" m2s:"province_no"`    //ProvinceNo 省份
	CityNo      string `json:"city_no" form:"city_no" m2s:"city_no"`                //CityNo 城市
	LineId      string `json:"line_id" form:"line_id" m2s:"line_id"`                //LineId 产品线
	ShelfId     string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id"`             //ShelfId 货架名称
	Status      string `json:"status" form:"status" m2s:"status"`                   //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbProduct  上游商品接口
type IDbProduct interface {
	//Create 创建
	Create(input *CreateProduct) error
	//Get 单条查询
	Get(productId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryProduct) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateProduct) (err error)
}

//DbProduct 上游商品对象
type DbProduct struct {
	c component.IContainer
}

//NewDbProduct 创建上游商品对象
func NewDbProduct(c component.IContainer) *DbProduct {
	return &DbProduct{
		c: c,
	}
}

//Create 添加上游商品
func (d *DbProduct) Create(input *CreateProduct) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertOmsUpProduct, map[string]interface{}{
		"can_refund":          input.CanRefund,
		"invoice_type":        input.InvoiceType,
		"carrier_no":          input.CarrierNo,
		"province_no":         input.ProvinceNo,
		"city_no":             input.CityNo,
		"cost_discount":       input.CostDiscount,
		"commission_discount": input.CommissionDiscount,
		"service_discount":    input.ServiceDiscount,
		"ext_product_no":      input.ExtProductNo,
		"face":                input.Face,
		"limit_count":         input.LimitCount,
		"line_id":             input.LineId,
		"shelf_id":            input.ShelfId,
	})
	if err != nil {
		return fmt.Errorf("添加上游商品数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据上游商品
func (d *DbProduct) Get(productId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsUpProduct, map[string]interface{}{
		"product_id": productId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取上游商品数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取上游商品列表
func (d *DbProduct) Query(input *QueryProduct) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsUpProductCount, map[string]interface{}{
		"can_refund":   input.CanRefund,
		"invoice_type": input.InvoiceType,
		"carrier_no":   input.CarrierNo,
		"province_no":  input.ProvinceNo,
		"city_no":      input.CityNo,
		"line_id":      input.LineId,
		"shelf_id":     input.ShelfId,
		"status":       input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游商品列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsUpProduct, map[string]interface{}{
		"can_refund":   input.CanRefund,
		"invoice_type": input.InvoiceType,
		"carrier_no":   input.CarrierNo,
		"province_no":  input.ProvinceNo,
		"city_no":      input.CityNo,
		"line_id":      input.LineId,
		"shelf_id":     input.ShelfId,
		"status":       input.Status,
		"pi":           input.Pi,
		"ps":           input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取上游商品数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新上游商品
func (d *DbProduct) Update(input *UpdateProduct) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsUpProduct, map[string]interface{}{
		"product_id":          input.ProductId,
		"can_refund":          input.CanRefund,
		"invoice_type":        input.InvoiceType,
		"carrier_no":          input.CarrierNo,
		"province_no":         input.ProvinceNo,
		"city_no":             input.CityNo,
		"cost_discount":       input.CostDiscount,
		"commission_discount": input.CommissionDiscount,
		"service_discount":    input.ServiceDiscount,
		"ext_product_no":      input.ExtProductNo,
		"face":                input.Face,
		"limit_count":         input.LimitCount,
		"line_id":             input.LineId,
		"shelf_id":            input.ShelfId,
		"status":              input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新上游商品数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
