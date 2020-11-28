package jfpps

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//JFExtParams 积分拓展参数
type JFExtParams struct {
	PreOrderID          int64
	OverTime            string
	UserNo              string
	PointType           int
	SubstituteChannelNo string
	Memo                string
}

// DBJFOrder 订单结构体
type DBJFOrder struct {
	*order.DBOrder
	ps *JFExtParams
	c  component.IContainer
}

// NewDBJFOrder 构建DBOrder结构体
func NewDBJFOrder(c component.IContainer, ps *JFExtParams) *DBJFOrder {
	db := &DBJFOrder{
		ps: ps,
		c:  c,
	}
	db.DBOrder = order.NewDBOrder(c)
	return db
}

//QueryPorductByDB 覆盖产品检查
func (o *DBJFOrder) QueryPorductByDB(info *order.RequestInfo) ([]string, types.XMap, error) {
	_, data, err := o.DBOrder.QueryPorductByDB(info)

	return []string{task.TaskType.BindTask, task.TaskType.OrderOverTimeTask}, data, err
}

//CreateOrderAndNotifyByDB 覆盖创建订单
func (o *DBJFOrder) CreateOrderAndNotifyByDB(dbTrans db.IDBTrans, param map[string]interface{}) (types.XMap, error) {
	//1.构建订单拓展信息
	ext := types.NewXMapByMap(map[string]interface{}{
		"pre_order_id": o.ps.PreOrderID,
		"over_time":    o.ps.OverTime,
		"user_no":      o.ps.UserNo,
		"point_type":   o.ps.PointType,
		"memo":         o.ps.Memo,
	})

	bt, err := json.Marshal(ext)
	if err != nil {
		return nil, err
	}
	param["extend_info"] = string(bt)

	// 2.创建订单
	orderID, row, sqlStr, args, err := utils.Insert(dbTrans, sql.GetNewOrderID, sql.JFOrderCreate, param)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("创建订单信息发送异常，err:%v,sql:%s,args:%v", err, sqlStr, args)
	}

	// 3.创建通知
	param["order_id"] = orderID
	if param["notify_url"] != "" {
		row, sqlStr, args, err := dbTrans.Execute(sql.NotifyCreate, param)
		if err != nil || row != 1 {
			return nil, fmt.Errorf("创建订单通知发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
		}
	}

	// 4.查询订单信息
	datas, sqlStr, args, err := dbTrans.Query(sql.CheckOrder, param)
	if err != nil || datas.IsEmpty() {
		return nil, fmt.Errorf("订单查询信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}

	data := datas.Get(0)
	data.MergeMap(param)
	return data, nil
}

//------------------------------------------------新增函数----------------------------------------------------------------

//QueryJFPreOrderInfo 查询积分预下单详情
func (o *DBJFOrder) QueryJFPreOrderInfo(req *JFRequestInfo) (types.XMap, error) {
	db := o.c.GetRegularDB()

	rows, _, _, err := db.Query(sql.QueryJFPreOrderInfo, map[string]interface{}{
		"channel_no":     req.ChannelNO,
		"pre_request_no": req.PreRequestNO,
	})
	if err != nil || rows.IsEmpty() {
		return nil, fmt.Errorf("查询积分预下单异常:count:%d,err:%v", rows.Len(), err)
	}

	return rows.Get(0), nil
}
