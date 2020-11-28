package down

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateProduct 创建下游商品
type CreateProduct struct {
	ShelfId            string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id" valid:"required"`                                     //ShelfId 货架编号
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                                        //LineId 产品线
	CanRefund          string `json:"can_refund" form:"can_refund" m2s:"can_refund" valid:"required"`                               //CanRefund 支持退款
	InvoiceType        string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type" valid:"required"`                         //InvoiceType 开票方式（1.不开发票，0.不限制，2.需要发票）
	CanSplitOrder      string `json:"can_split_order" form:"can_split_order" m2s:"can_split_order" valid:"required"`                //CanSplitOrder 是否允许拆单（0.允许，1不允许）
	ExtProductNo       string `json:"ext_product_no" form:"ext_product_no" m2s:"ext_product_no"`                                    //ExtProductNo 外部商品编号
	Face               string `json:"face" form:"face" m2s:"face" valid:"required"`                                                 //Face 面值
	LimitCount         string `json:"limit_count" form:"limit_count" m2s:"limit_count" valid:"required"`                            //LimitCount 单次最大购买数量
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no" valid:"required"`                               //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no" valid:"required"`                            //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no" valid:"required"`                                        //CityNo 城市
	PaymentFeeDiscount string `json:"payment_fee_discount" form:"payment_fee_discount" m2s:"payment_fee_discount" valid:"required"` //PaymentFeeDiscount 手续费折扣（以销售金额算）
	CommissionDiscount string `json:"commission_discount" form:"commission_discount" m2s:"commission_discount" valid:"required"`    //CommissionDiscount 佣金折扣（以面值算）
	SellDiscount       string `json:"sell_discount" form:"sell_discount" m2s:"sell_discount" valid:"required"`                      //SellDiscount 销售折扣（以面值算）
	ServiceDiscount    string `json:"service_discount" form:"service_discount" m2s:"service_discount" valid:"required"`             //ServiceDiscount 服务费折扣
	SplitOrderFace     string `json:"split_order_face" form:"split_order_face" m2s:"split_order_face" `                             //SplitOrderFace 拆单面值
	Status             string `json:"status" form:"status" m2s:"status" valid:"required"`                                           //Status 状态

}

//UpdateProduct 添加下游商品
type UpdateProduct struct {
	ProductId          string `json:"product_id" form:"product_id" m2s:"product_id" valid:"required"`                               //ProductId 商品编号
	ShelfId            string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id" valid:"required"`                                     //ShelfId 货架编号
	LineId             string `json:"line_id" form:"line_id" m2s:"line_id" valid:"required"`                                        //LineId 产品线
	CanRefund          string `json:"can_refund" form:"can_refund" m2s:"can_refund" valid:"required"`                               //CanRefund 支持退款
	InvoiceType        string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type" valid:"required"`                         //InvoiceType 开票方式（1.不开发票，0.不限制，2.需要发票）
	CanSplitOrder      string `json:"can_split_order" form:"can_split_order" m2s:"can_split_order" valid:"required"`                //CanSplitOrder 是否允许拆单（0.允许，1不允许）
	ExtProductNo       string `json:"ext_product_no" form:"ext_product_no" m2s:"ext_product_no"`                                    //ExtProductNo 外部商品编号
	Face               string `json:"face" form:"face" m2s:"face" valid:"required"`                                                 //Face 面值
	LimitCount         string `json:"limit_count" form:"limit_count" m2s:"limit_count" valid:"required"`                            //LimitCount 单次最大购买数量
	CarrierNo          string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no" valid:"required"`                               //CarrierNo 运营商
	ProvinceNo         string `json:"province_no" form:"province_no" m2s:"province_no" valid:"required"`                            //ProvinceNo 省份
	CityNo             string `json:"city_no" form:"city_no" m2s:"city_no" valid:"required"`                                        //CityNo 城市
	PaymentFeeDiscount string `json:"payment_fee_discount" form:"payment_fee_discount" m2s:"payment_fee_discount" valid:"required"` //PaymentFeeDiscount 手续费折扣（以销售金额算）
	CommissionDiscount string `json:"commission_discount" form:"commission_discount" m2s:"commission_discount" valid:"required"`    //CommissionDiscount 佣金折扣（以面值算）
	SellDiscount       string `json:"sell_discount" form:"sell_discount" m2s:"sell_discount" valid:"required"`                      //SellDiscount 销售折扣（以面值算）
	ServiceDiscount    string `json:"service_discount" form:"service_discount" m2s:"service_discount" valid:"required"`             //ServiceDiscount 服务费折扣
	SplitOrderFace     string `json:"split_order_face" form:"split_order_face" m2s:"split_order_face" valid:"required"`             //SplitOrderFace 拆单面值
	Status             string `json:"status" form:"status" m2s:"status" valid:"required"`                                           //Status 状态

}

//QueryProduct 查询下游商品
type QueryProduct struct {
	ShelfId       string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id"`                      //ShelfId 货架编号
	LineId        string `json:"line_id" form:"line_id" m2s:"line_id"`                         //LineId 产品线
	CanRefund     string `json:"can_refund" form:"can_refund" m2s:"can_refund"`                //CanRefund 支持退款
	InvoiceType   string `json:"invoice_type" form:"invoice_type" m2s:"invoice_type"`          //InvoiceType 开票方式（1.不开发票，0.不限制，2.需要发票）
	CanSplitOrder string `json:"can_split_order" form:"can_split_order" m2s:"can_split_order"` //CanSplitOrder 是否允许拆单（0.允许，1不允许）
	CarrierNo     string `json:"carrier_no" form:"carrier_no" m2s:"carrier_no"`                //CarrierNo 运营商
	ProvinceNo    string `json:"province_no" form:"province_no" m2s:"province_no"`             //ProvinceNo 省份
	CityNo        string `json:"city_no" form:"city_no" m2s:"city_no"`                         //CityNo 城市
	Status        string `json:"status" form:"status" m2s:"status"`                            //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbProduct  下游商品接口
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

//DbProduct 下游商品对象
type DbProduct struct {
	c component.IContainer
}

//NewDbProduct 创建下游商品对象
func NewDbProduct(c component.IContainer) *DbProduct {
	return &DbProduct{
		c: c,
	}
}

//Create 添加下游商品
func (d *DbProduct) Create(input *CreateProduct) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertOmsDownProduct, map[string]interface{}{
		"shelf_id":             input.ShelfId,
		"line_id":              input.LineId,
		"can_refund":           input.CanRefund,
		"invoice_type":         input.InvoiceType,
		"can_split_order":      input.CanSplitOrder,
		"ext_product_no":       input.ExtProductNo,
		"face":                 input.Face,
		"limit_count":          input.LimitCount,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"payment_fee_discount": input.PaymentFeeDiscount,
		"commission_discount":  input.CommissionDiscount,
		"sell_discount":        input.SellDiscount,
		"service_discount":     input.ServiceDiscount,
		"split_order_face":     input.SplitOrderFace,
		"status":               input.Status,
	})
	if err != nil {
		return fmt.Errorf("添加下游商品数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据下游商品
func (d *DbProduct) Get(productId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsDownProduct, map[string]interface{}{
		"product_id": productId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取下游商品数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取下游商品列表
func (d *DbProduct) Query(input *QueryProduct) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsDownProductCount, map[string]interface{}{
		"shelf_id":        input.ShelfId,
		"line_id":         input.LineId,
		"can_refund":      input.CanRefund,
		"invoice_type":    input.InvoiceType,
		"can_split_order": input.CanSplitOrder,
		"carrier_no":      input.CarrierNo,
		"province_no":     input.ProvinceNo,
		"city_no":         input.CityNo,
		"status":          input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游商品列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsDownProduct, map[string]interface{}{
		"shelf_id":        input.ShelfId,
		"line_id":         input.LineId,
		"can_refund":      input.CanRefund,
		"invoice_type":    input.InvoiceType,
		"can_split_order": input.CanSplitOrder,
		"carrier_no":      input.CarrierNo,
		"province_no":     input.ProvinceNo,
		"city_no":         input.CityNo,
		"status":          input.Status,
		"pi":              input.Pi,
		"ps":              input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游商品数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新下游商品
func (d *DbProduct) Update(input *UpdateProduct) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsDownProduct, map[string]interface{}{
		"product_id":           input.ProductId,
		"shelf_id":             input.ShelfId,
		"line_id":              input.LineId,
		"can_refund":           input.CanRefund,
		"invoice_type":         input.InvoiceType,
		"can_split_order":      input.CanSplitOrder,
		"ext_product_no":       input.ExtProductNo,
		"face":                 input.Face,
		"limit_count":          input.LimitCount,
		"carrier_no":           input.CarrierNo,
		"province_no":          input.ProvinceNo,
		"city_no":              input.CityNo,
		"payment_fee_discount": input.PaymentFeeDiscount,
		"commission_discount":  input.CommissionDiscount,
		"sell_discount":        input.SellDiscount,
		"service_discount":     input.ServiceDiscount,
		"split_order_face":     input.SplitOrderFace,
		"status":               input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新下游商品数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
