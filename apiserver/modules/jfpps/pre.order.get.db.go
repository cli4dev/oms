package jfpps

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//DBJFPreOrder 积分预下单db操作
type DBJFPreOrder struct {
	c component.IContainer
}

//NewDBJFPreOrder 初始化DBJFPreOrder
func NewDBJFPreOrder(c component.IContainer) *DBJFPreOrder {
	return &DBJFPreOrder{
		c: c,
	}
}

//CheckJFPreOrder 积分预下单检查
func (o *DBJFPreOrder) CheckJFPreOrder(req *JFPreRequestInfo) (types.XMap, error) {
	db := o.c.GetRegularDB()

	rows, _, _, err := db.Query(sql.CheckJFPreOrder, map[string]interface{}{
		"channel_no":     req.ChannelNO,
		"pre_request_no": req.PreRequestNO,
	})
	if err != nil {
		return nil, err
	}
	if !rows.IsEmpty() {
		return rows.Get(0), nil
	}

	return nil, nil
}

//CheckJFProduct 检查积分产品配置
func (o *DBJFPreOrder) CheckJFProduct(req *JFPreRequestInfo) (types.XMap, error) {
	db := o.c.GetRegularDB()

	rows, _, _, err := db.Query(sql.CheckJFProduct, map[string]interface{}{
		"channel_no": req.ChannelNO,
		"line_id":    req.LineID,
	})
	if err != nil || rows.Len() != 1 {
		return nil, fmt.Errorf("检查积分产品配置数据异常，count:%d,err:%v", rows.Len(), err)
	}

	JFChannels, _, _, err := db.Query(sql.GetJFUpChannel, map[string]interface{}{})
	if err != nil || JFChannels.Len() != 1 {
		return nil, fmt.Errorf("查询积分上游异常,count:%d,err:%v", JFChannels.Len(), err)
	}

	pro := rows.Get(0)
	pro.Merge(JFChannels.Get(0))
	return pro, nil
}

//LockJFProduct 锁积分产品
func (o *DBJFPreOrder) LockJFProduct(dbtrans db.IDBTrans, product types.XMap, req *JFPreRequestInfo) (types.XMap, error) {
	pros, _, _, err := dbtrans.Query(sql.LockJFProduct, map[string]interface{}{
		"product_id": product.GetString("product_id"),
	})
	if err != nil || pros.IsEmpty() {
		return nil, fmt.Errorf("锁积分产品异常，count:%d,err:%v", pros.Len(), err)
	}

	//检查预下单
	rows, _, _, err := dbtrans.Query(sql.CheckJFPreOrder, map[string]interface{}{
		"channel_no":     req.ChannelNO,
		"pre_request_no": req.PreRequestNO,
	})
	if err != nil {
		return nil, err
	}
	if !rows.IsEmpty() {
		return rows.Get(0), nil
	}

	return nil, nil
}

//CreateJFPreOrder 创建积分预下单
func (o *DBJFPreOrder) CreateJFPreOrder(dbtrans db.IDBTrans, product types.XMap, req *JFPreRequestInfo) (types.XMap, error) {

	// 2.创建预下单订单
	_, row, sqlStr, args, err := utils.Insert(dbtrans, sql.GetNewJFPreOrderID, sql.CreateJFPreOrder, map[string]interface{}{
		"channel_no":      req.ChannelNO,
		"pre_request_no":  req.PreRequestNO,
		"line_id":         req.LineID,
		"down_shelf_id":   product.GetString("shelf_id"),
		"down_product_id": product.GetString("product_id"),
		"point_type":      req.PointType,
		"num":             req.Num,
		"user_no":         req.UserNo,
		"over_time":       req.OverTime,
	})
	if err != nil || row != 1 {
		return nil, fmt.Errorf("创建积分预下单订单信息异常，err:%v,sql:%s,args:%v", err, sqlStr, args)
	}

	pres, _, _, err := dbtrans.Query(sql.CheckJFPreOrder, map[string]interface{}{
		"channel_no":     req.ChannelNO,
		"pre_request_no": req.PreRequestNO,
	})
	if err != nil || pres.IsEmpty() {
		return nil, fmt.Errorf("查询预下单信息异常，count:%d,err:%v", pres.Len(), err)
	}

	return pres.Get(0), nil
}

//CancelPreOrderQuery 预下单信息查询
func (o *DBJFPreOrder) CancelPreOrderQuery(channelNo string, preRequestNo string) (types.XMap, error) {
	db := o.c.GetRegularDB()
	pres, _, _, err := db.Query(sql.CancelPreOrderQuery, map[string]interface{}{
		"channel_no":     channelNo,
		"pre_request_no": preRequestNo,
	})
	if err != nil {
		return nil, err
	}
	if pres.IsEmpty() {
		return nil, fmt.Errorf("待激活预下单不存在或已激活")
	}
	preOrder := pres.Get(0)

	JFChannels, _, _, err := db.Query(sql.GetJFUpChannel, map[string]interface{}{})
	if err != nil || JFChannels.Len() != 1 {
		return nil, fmt.Errorf("查询积分上游异常,count:%d,err:%v", JFChannels.Len(), err)
	}
	preOrder.Merge(JFChannels.Get(0))

	return preOrder, nil
}

//CancelJFPreOrder 取消预下单积分
func (o *DBJFPreOrder) CancelJFPreOrder(dbtrans db.IDBTrans, channelNo string, preRequestNo string) error {
	count, _, _, err := dbtrans.Execute(sql.CancelJFPreOrder, map[string]interface{}{
		"channel_no":     channelNo,
		"pre_request_no": preRequestNo,
	})
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("预下单不存在或已激活")
	}

	return nil
}

//CancelOrderCheck 取消预下单订单检查
func (o *DBJFPreOrder) CancelOrderCheck(dbtrans db.IDBTrans, preInfo types.XMap) error {
	orders, _, _, err := dbtrans.Query(sql.CancelOrderCheck, map[string]interface{}{
		"create_time":  preInfo.GetString("create_time"),
		"pre_order_id": preInfo.GetString("pre_order_id"),
	})
	if err != nil {
		return err
	}
	if !orders.IsEmpty() {
		return fmt.Errorf("预下单已下单，不能取消")
	}

	return nil
}
