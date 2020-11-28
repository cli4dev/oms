package order

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

// IDBOrder 数据层接口
type IDBOrder interface {
	CreateOrderAndNotifyByDB(dbTrans db.IDBTrans, param map[string]interface{}) (types.XMap, error)
	CheckOrderByDB(channelNO, requestNO string) (types.XMap, error)
	QueryPorductByDB(info *RequestInfo) ([]string, types.XMap, error)
	QueryByDB(info *QueryInfo) (types.XMap, error)
	ProcessNotifyStatus(info types.XMap) error
	LockProductByDB(trans db.IDBTrans, param map[string]interface{}) (types.XMap, error)
	CheckChannelAccount(channelNo string, accountNo string) error
	BuildResult(data types.IXMap) types.XMap
	BuildQueryResult(data types.IXMap) (types.IXMap, error)
}

// DBOrder 订单结构体
type DBOrder struct {
	c component.IContainer
}

// NewDBOrder 构建DBOrder结构体
func NewDBOrder(c component.IContainer) *DBOrder {
	return &DBOrder{
		c: c,
	}
}

// CreateOrderAndNotifyByDB 创建订单和通知
func (d *DBOrder) CreateOrderAndNotifyByDB(dbTrans db.IDBTrans, param map[string]interface{}) (types.XMap, error) {

	// 1.创建订单
	orderID, row, sqlStr, args, err := utils.Insert(dbTrans, sql.GetNewOrderID, sql.OrderCreate, param)
	if err != nil || row != 1 {
		return nil, fmt.Errorf("创建订单信息发送异常，err:%v,sql:%s,args:%v", err, sqlStr, args)
	}

	// 2.创建通知
	param["order_id"] = orderID
	if param["notify_url"] != "" {
		row, sqlStr, args, err := dbTrans.Execute(sql.NotifyCreate, param)
		if err != nil || row != 1 {
			return nil, fmt.Errorf("创建订单通知发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
		}
	}

	// 3.查询订单信息
	datas, sqlStr, args, err := dbTrans.Query(sql.CheckOrder, param)
	if err != nil || datas.IsEmpty() {
		return nil, fmt.Errorf("订单查询信息发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}

	data := datas.Get(0)
	data.MergeMap(param)
	return data, nil
}

// CheckOrderByDB 订单检查是否存在
func (d *DBOrder) CheckOrderByDB(channelNO, requestNO string) (types.XMap, error) {
	db := d.c.GetRegularDB()
	datas, sqlStr, args, err := db.Query(sql.CheckOrder, map[string]interface{}{
		"request_no": requestNO,
		"channel_no": channelNO,
	})
	if err != nil {
		return nil, fmt.Errorf("订单检查是否存在发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil
	}
	return datas.Get(0), nil
}

// QueryPorductByDB 查询下游商品
func (d *DBOrder) QueryPorductByDB(info *RequestInfo) ([]string, types.XMap, error) {
	if info.PointNum < 0 || (info.PointNum > 0 && info.PayOrderNo == "") {
		return nil, nil, fmt.Errorf("积分信息错误")
	}

	info.ProvinceNo = types.DecodeString(info.ProvinceNo, "", "*")
	info.CityNO = types.DecodeString(info.CityNO, "", "*")
	param, err := types.Struct2Map(info)
	if err != nil {
		return nil, nil, err
	}

	db := d.c.GetRegularDB()
	datas, sqlStr, args, err := db.Query(sql.QueryProduct, param)
	if err != nil {
		return nil, nil, fmt.Errorf("查询下游商品是否存在发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil, context.NewError(errorcode.DOWN_CHANNEL_NO_EXIST_PRODUCT.Code, fmt.Sprintf("下游渠道(%s)不存在面值(%v)的商品err:%v,sql:%v,args:%v,", param["channel_no"], param["face"], err, sqlStr, args))
	}

	// 检查折扣差异
	data := datas.Get(0)

	diff := data.GetFloat64("diff")
	if diff > 0.1 || diff < -0.1 {
		return nil, nil, context.NewErrorf(errorcode.DISCOUNT_DIFF_ERROR.Code, fmt.Sprintf("折扣差异为%f,销售金额%f,下游支付金额%f,总面值%f",
			diff, data.GetFloat64("sell_amount"), param["amount"], data.GetFloat64("total_face")))
	}

	if types.GetFloat64(param["num"]) > 1 && data.GetInt("can_split_order") == 0 && data.GetFloat64("face") != data.GetFloat64("split_order_face") {
		return nil, nil, context.NewError(errorcode.DOWN_CHANNEL_NO_EXIST_PRODUCT.Code, fmt.Sprintf("下游渠道(%s)不存在面值(%v)的商品", param["channel_no"], param["face"]))
	}
	data.MergeMap(param)

	return []string{task.TaskType.DownPayTask, task.TaskType.OrderOverTimeTask}, data, nil
}

// QueryByDB 订单查询
func (d *DBOrder) QueryByDB(info *QueryInfo) (types.XMap, error) {
	dbt := d.c.GetRegularDB()
	param, err := types.Struct2Map(info)
	if err != nil {
		return nil, err
	}

	// 1.查询订单信息
	datas, sqlStr, args, err := dbt.Query(sql.QueryOrderInfo, param)
	if err != nil {
		return nil, fmt.Errorf("查询订单信息发生异常,cnt:%d,err:%v,sql:%v,args:%v", datas.Len(), err, sqlStr, args)
	}

	if datas.IsEmpty() {
		return nil, context.NewErrorf(context.ERR_NO_CONTENT, "订单不存在,channel_no:%v,request_no:%v", info.ChannelNO, info.RequestNO)
	}
	return datas.Get(0), nil
}

// ProcessNotifyStatus 处理通知状态
func (d *DBOrder) ProcessNotifyStatus(info types.XMap) error {

	dbt := d.c.GetRegularDB()

	if info.GetInt("notify_status") != enum.NotifyStatus.Processing {
		return nil
	}

	row, sqlStr, args, err := dbt.Execute(sql.SuccessOrderNotify, info)
	if err != nil || row != 1 {
		return fmt.Errorf("修改订单通知状态为查询成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// LockProductByDB 锁产品
func (d *DBOrder) LockProductByDB(trans db.IDBTrans, param map[string]interface{}) (types.XMap, error) {

	// 1.锁产品
	products, sqlStr, args, err := trans.Query(sql.LockProduct, param)
	if err != nil || products.IsEmpty() {
		return nil, fmt.Errorf("锁产品发生异常,cnt:%d,err:%v,sql:%v,args:%v", products.Len(), err, sqlStr, args)
	}

	// 2.检查订单是否存在，存在返回，不存在继续往下走
	datas, sqlStr, args, err := trans.Query(sql.CheckOrder, param)
	if err != nil {
		return nil, fmt.Errorf("订单检查是否存在发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if datas.IsEmpty() {
		return nil, nil
	}
	return datas.Get(0), nil
}

//CheckChannelAccount 检查下游渠道收款信息
func (d *DBOrder) CheckChannelAccount(channelNo string, accountNo string) error {
	//查询渠道支付账户信息
	db := d.c.GetRegularDB()

	accounts, _, _, err := db.Query(sql.GetChannelAccountInfo, map[string]interface{}{
		"down_channel_no": channelNo,
		"down_account_no": accountNo,
	})
	if err != nil || accounts.IsEmpty() {
		return fmt.Errorf("查询渠道账户信息异常，count:%d,err:%v", accounts.Len(), err)
	}

	channelAccountCount := accounts.Get(0).GetInt("channel_account_count")
	findCount := accounts.Get(0).GetInt("find_count")

	if channelAccountCount > 0 && findCount == 0 {
		return fmt.Errorf("渠道收款账户错误")
	}

	if channelAccountCount == 0 && accountNo != "" {
		return fmt.Errorf("渠道不存在收款账户配置")
	}

	return nil
}

// BuildResult 构建返回参数
func (d *DBOrder) BuildResult(data types.IXMap) types.XMap {
	status, code, msg := errorcode.SetFlowRecordStatus(data.GetInt("order_status"), errorcode.RequestFlowType.Order, data.GetString("fail_code"), data.GetString("fail_msg"))
	res := map[string]interface{}{
		"order_id":    data.GetInt64("order_id"),
		"request_no":  data.GetString("request_no"),
		"channel_no":  data.GetString("channel_no"),
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	}

	if data.GetString("account_no") != "" {
		res["account_no"] = data.GetString("account_no")
	}
	return res
}

// BuildQueryResult 构建查询返回参数
func (d *DBOrder) BuildQueryResult(data types.IXMap) (types.IXMap, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(data.GetInt("order_status"), errorcode.RequestFlowType.Order, data.GetString("fail_code"), data.GetString("fail_msg"))
	res := types.NewXMapByMap(map[string]interface{}{
		"order_id":    data.GetInt64("order_id"),
		"request_no":  data.GetString("request_no"),
		"channel_no":  data.GetString("channel_no"),
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	})
	if data.GetString("account_no") != "" {
		res.SetValue("account_no", data.GetString("account_no"))
	}
	if data.GetInt("point_num") > 0 && data.GetInt("order_status") == enum.OrderStatus.Failed {
		pointList := types.NewXMaps()
		if data.GetInt("activity_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": data.GetInt("activity_send_num"),
				"point_type":  enum.PointType.Activity,
			})
		}
		if data.GetInt("buy_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": data.GetInt("buy_send_num"),
				"point_type":  enum.PointType.Buy,
			})
		}
		if pointList.IsEmpty() {
			return nil, fmt.Errorf("积分数据异常")
		}
		bt, err := json.Marshal(pointList)
		if err != nil {
			return nil, err
		}

		res["ext_params"] = string(bt)
	}
	return res, nil
}
