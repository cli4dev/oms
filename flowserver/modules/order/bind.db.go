package order

import (
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// IDBBind 数据层接口
type IDBBind interface {
	CheckOrderForBindByDB(orderID, taskID int64) (bool, error)
	BindUpProductByDB(trans db.IDBTrans, product types.XMap, orderInfo types.XMap) ([]string, types.XMap, bool, error)
	QueryUpProductByDB(trans db.IDBTrans, data types.XMap) (types.XMap, error)
	LockOrderForBindByDB(trans db.IDBTrans, orderID int64) (types.XMap, error)
	QueryProductByDB(dbTrans db.IDBTrans, data types.XMap) (product types.XMap, err error)
}

// DBBind 绑定数据层
type DBBind struct {
	c component.IContainer
}

// NewDBBind 构建DBBind
func NewDBBind(c component.IContainer) *DBBind {
	return &DBBind{c: c}
}

// CheckOrderForBindByDB 绑定检查订单
func (b *DBBind) CheckOrderForBindByDB(orderID, taskID int64) (bool, error) {
	db := b.c.GetRegularDB()
	datas, sqlStr, args, err := db.Query(sql.CheckOrderForBind, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return false, fmt.Errorf("绑定检查订单发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return true, fmt.Errorf("绑定时,订单不存在或已完成绑定,order_id:%dsql:%v,args:%v", orderID, sqlStr, args)
	}
	return false, nil
}

// BindUpProductByDB 绑定上游商品
func (b *DBBind) BindUpProductByDB(trans db.IDBTrans, product types.XMap, orderInfo types.XMap) ([]string, types.XMap, bool, error) {

	// 1.获取发货记录id
	deliveryID, row, sqlStr, args, err := utils.Insert(trans, sql.GetNewDeliveryID, sql.CreateOrderDelivery, map[string]interface{}{
		"need_bind_num": product.GetString("need_bind_num"),
		"extend_info":   orderInfo.GetString("extend_info"),
		"order_id":      orderInfo.GetInt64("order_id"),
		"product_id":    product.GetString("product_id"),
	})
	if err != nil || row != 1 {
		return nil, nil, false, fmt.Errorf("创建发货记录发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 2.修改订单信息
	row, sqlStr, args, err = trans.Execute(sql.UpdateOrderForBind, map[string]interface{}{
		"order_id":       orderInfo.GetInt64("order_id"),
		"need_bind_face": product.GetString("need_bind_face"),
		"need_bind_num":  product.GetString("need_bind_num"),
	})
	if err != nil || row != 1 {
		return nil, nil, false, fmt.Errorf("修改订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	// 3.判断是否完全绑定
	datas, sqlStr, args, err := trans.Query(sql.CheckCompleteBind, map[string]interface{}{
		"order_id": orderInfo.GetInt64("order_id"),
	})
	if err != nil || datas.IsEmpty() {
		return nil, nil, false, fmt.Errorf("判断完全绑定发生异常,cnt:%d,err:%v,sql:%v,args:%v", types.GetInt(deliveryID), err, sqlStr, args)
	}

	info := types.NewXMapByMap(map[string]interface{}{
		"delivery_id":       deliveryID,
		"delivery_overtime": product.GetInt("delivery_overtime"),
		"order_id":          orderInfo.GetInt64("order_id"),
		"flow_overtime":     orderInfo.GetInt("flow_overtime"),
	})

	return []string{task.TaskType.DeliveryTask}, info, types.GetBool(datas.Get(0).GetString("complete_bind")), nil
}

// QueryUpProductByDB 获取上游商品
func (b *DBBind) QueryUpProductByDB(trans db.IDBTrans, data types.XMap) (types.XMap, error) {

	products, sqlStr, args, err := trans.Query(sql.QueryUpProduct, data)
	if err != nil {
		return nil, fmt.Errorf("查询上游商品发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if products.IsEmpty() {
		return nil, context.NewErrorf(context.ERR_NO_CONTENT, "上游不存在或已禁用面额为%v，数量%d的商品, sql:%v,args:%v", data.GetFloat64("need_bind_face"), data.GetInt("num"), sqlStr, args)
	}
	product := products.Get(0)
	if product.GetInt("limit_count") >= data.GetInt("need_num") {
		product["need_bind_num"] = data.GetInt("need_num")
		return product, nil
	}
	product["need_bind_num"] = product.GetInt("limit_count")
	product.Merge(data)
	return product, nil
}

// LockOrderForBindByDB 绑定锁定订单
func (b *DBBind) LockOrderForBindByDB(trans db.IDBTrans, orderID int64) (types.XMap, error) {
	// 1.锁订单
	datas, sqlStr, args, err := trans.Query(sql.LockOrderForBind, map[string]interface{}{"order_id": orderID})
	if err != nil {
		return nil, fmt.Errorf("绑定锁订单失败,cnt:%d,err:%v,sql:%v,args:%v", datas.Len(), err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil
	}
	return datas.Get(0), nil
}

// QueryProductByDB 查询产品
func (b *DBBind) QueryProductByDB(dbTrans db.IDBTrans, data types.XMap) (product types.XMap, err error) {

	products, sqlStr, args, err := dbTrans.Query(sql.QueryUpProduct, map[string]interface{}{
		"product_id":     data.GetInt64("down_product_id"),
		"order_id":       data.GetString("order_id"),
		"need_bind_face": getNeedBindFace(data),
		"need_num":       data.GetString("need_num"),
		"num":            data.GetString("num"),
	})
	if err != nil {
		return nil, fmt.Errorf("查询上游商品发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if products.IsEmpty() {
		return nil, context.NewErrorf(context.ERR_NO_CONTENT, "上游不存在面额为%v，数量%d的商品", getNeedBindFace(data), data.GetInt("num"))
	}
	product = products.Get(0)

	if product.GetInt("limit_count") >= data.GetInt("need_num") {
		product.SetValue("need_bind_num", data.GetInt("need_num"))
		return product, nil
	}
	product.SetValue("need_bind_num", product.GetInt("limit_count"))
	return
}

// GetNeedBindFace 获取需要绑定面值
func getNeedBindFace(data types.XMap) float64 {
	if data.GetInt("can_split_order") == enum.CanSplitOrder.NoSplitOrder || data.GetInt("num") > 1 {
		return data.GetFloat64("face")
	}
	splitOrderFace := data.GetFloat64("split_order_face")
	totalFace := data.GetFloat64("total_face")
	bindFace := data.GetFloat64("bind_face")
	if splitOrderFace != 0 && totalFace-bindFace > splitOrderFace {
		return splitOrderFace
	}
	return totalFace - bindFace
}
