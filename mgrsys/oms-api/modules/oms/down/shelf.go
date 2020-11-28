package down

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//CreateShelf 创建下游货架
type CreateShelf struct {
	ShelfName      string `json:"shelf_name" form:"shelf_name" m2s:"shelf_name" valid:"required"`                //ShelfName 货架名称
	ChannelNo      string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                //ChannelNo 渠道名称
	OrderOvertime  string `json:"order_overtime" form:"order_overtime" m2s:"order_overtime" valid:"required"`    //OrderOvertime 订单超时时长
	RefundOvertime string `json:"refund_overtime" form:"refund_overtime" m2s:"refund_overtime" valid:"required"` //RefundOvertime 退款超时时间

}

//UpdateShelf 添加下游货架
type UpdateShelf struct {
	ShelfId        string `json:"shelf_id" form:"shelf_id" m2s:"shelf_id" valid:"required"`                      //ShelfId 货架编号
	ShelfName      string `json:"shelf_name" form:"shelf_name" m2s:"shelf_name" valid:"required"`                //ShelfName 货架名称
	ChannelNo      string `json:"channel_no" form:"channel_no" m2s:"channel_no" valid:"required"`                //ChannelNo 渠道名称
	OrderOvertime  string `json:"order_overtime" form:"order_overtime" m2s:"order_overtime" valid:"required"`    //OrderOvertime 订单超时时长
	RefundOvertime string `json:"refund_overtime" form:"refund_overtime" m2s:"refund_overtime" valid:"required"` //RefundOvertime 退款超时时间
	Status         string `json:"status" form:"status" m2s:"status" valid:"required"`                            //Status 状态

}

//QueryShelf 查询下游货架
type QueryShelf struct {
	ShelfName string `json:"shelf_name" form:"shelf_name" m2s:"shelf_name"` //ShelfName 货架名称
	ChannelNo string `json:"channel_no" form:"channel_no" m2s:"channel_no"` //ChannelNo 渠道名称
	Status    string `json:"status" form:"status" m2s:"status"`             //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbShelf  下游货架接口
type IDbShelf interface {
	//Create 创建
	Create(input *CreateShelf) error
	//Get 单条查询
	Get(shelfId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryShelf) (data db.QueryRows, count int, err error)
	//Update 更新
	Update(input *UpdateShelf) (err error)
	//GetShelfDictionary 获取数据字典
	GetShelf(Channel string) (db.QueryRows, error)
	GetShelfDictionary() (db.QueryRows, error)
}

//DbShelf 下游货架对象
type DbShelf struct {
	c component.IContainer
}

//NewDbShelf 创建下游货架对象
func NewDbShelf(c component.IContainer) *DbShelf {
	return &DbShelf{
		c: c,
	}
}

//GetShelfDictionary 获取数据字典
func (d *DbShelf) GetShelfDictionary() (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetOmsDownShelfDictionary, map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("获取下游货架数据字典发生错误")
	}
	return data, nil
}

func (d *DbShelf) GetShelf(Channel string) (db.QueryRows, error) {

	db := d.c.GetRegularDB()
	data, _, _, err := db.Query(sql.GetChannelShelfDictionary, map[string]interface{}{
		"channel_no": Channel,
	})
	if err != nil {
		return nil, fmt.Errorf("获取下游货架数据字典发生错误")
	}
	return data, nil
}

//Create 添加下游货架
func (d *DbShelf) Create(input *CreateShelf) error {

	db := d.c.GetRegularDB()
	lastInsertID, affectedRow, q, a, err := db.Executes(sql.InsertOmsDownShelf, map[string]interface{}{
		"shelf_name":      input.ShelfName,
		"channel_no":      input.ChannelNo,
		"order_overtime":  input.OrderOvertime,
		"refund_overtime": input.RefundOvertime,
	})
	if err != nil {
		return fmt.Errorf("添加下游货架数据发生错误(err:%v),sql:%s,参数：%v,lastInsertID:%v,受影响的行数：%v", err, q, a, lastInsertID, affectedRow)
	}
	return nil
}

//Get 查询单条数据下游货架
func (d *DbShelf) Get(shelfId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetOmsDownShelf, map[string]interface{}{
		"shelf_id": shelfId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取下游货架数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取下游货架列表
func (d *DbShelf) Query(input *QueryShelf) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryOmsDownShelfCount, map[string]interface{}{
		"shelf_name": input.ShelfName,
		"channel_no": input.ChannelNo,
		"status":     input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游货架列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryOmsDownShelf, map[string]interface{}{
		"shelf_name": input.ShelfName,
		"channel_no": input.ChannelNo,
		"status":     input.Status,
		"pi":         input.Pi,
		"ps":         input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取下游货架数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}

//Update 更新下游货架
func (d *DbShelf) Update(input *UpdateShelf) error {

	db := d.c.GetRegularDB()
	affectedRow, q, a, err := db.Execute(sql.UpdateOmsDownShelf, map[string]interface{}{
		"shelf_id":        input.ShelfId,
		"shelf_name":      input.ShelfName,
		"channel_no":      input.ChannelNo,
		"order_overtime":  input.OrderOvertime,
		"refund_overtime": input.RefundOvertime,
		"status":          input.Status,
	})
	if err != nil {
		return fmt.Errorf("更新下游货架数据发生错误(err:%v),sql:%s,参数：%v,受影响的行数：%v", err, q, a, affectedRow)
	}
	return nil
}
