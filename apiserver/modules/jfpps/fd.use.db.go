package jfpps

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

// DBJFFd 积分记账
type DBJFFd struct {
	c component.IContainer
}

// NewDBJFFd 积分记账处理
func NewDBJFFd(c component.IContainer) *DBJFFd {
	return &DBJFFd{c: c}
}

//CheckFdRequestRecord 检查请求是否已存在
func (o *DBJFFd) CheckFdRequestRecord(channelNO string, requestNo string, relationNo string, requestType int) (types.XMap, error) {
	db := o.c.GetRegularDB()

	rows, _, _, err := db.Query(sql.CheckFdRequestRecord, map[string]interface{}{
		"down_channel_no": channelNO,
		"request_type":    requestType,
		"request_no":      requestNo,
	})
	if err != nil {
		return nil, err
	}
	if !rows.IsEmpty() {
		return rows.Get(0), nil
	}

	count, _, _, err := db.Scalar(sql.CheckRelationNo, map[string]interface{}{
		"down_channel_no": channelNO,
		"request_type":    requestType,
		"relation_no":     relationNo,
	})
	if err != nil || types.GetInt(count) > 0 {
		return nil, fmt.Errorf("检查支付订单号异常,count:%v,err:%v", count, err)
	}
	return nil, nil
}

//CheckUseFdOrders 检查支付记录对应订单与积分
func (o *DBJFFd) CheckUseFdOrders(req *JFFDUseRequest) (types.XMaps, error) {
	db := o.c.GetRegularDB()

	orders, sqlStr, args, err := db.Query(sql.CheckUseFdOrders, map[string]interface{}{
		"down_channel_no": req.ChannelNO,
		"channel_pay_no":  req.ChannelPayNo,
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(sqlStr, args)
	fmt.Println("orders", orders)

	reqBuyNum := req.BuyPointNum
	reqActNum := req.ActivityPointNum
	for k, v := range orders {

		fmt.Println("v", v)
		data := types.NewXMap()
		orderPoint := v.GetInt("point_num")

		fmt.Println(orderPoint)
		orderPoint, reqBuyNum, data = PointCalc(orderPoint, reqBuyNum, enum.JFPointType.BuyPoint, data)
		if orderPoint <= 0 {
			continue
		}
		fmt.Println("1", orderPoint, reqBuyNum, data)
		orderPoint, reqActNum, data = PointCalc(orderPoint, reqActNum, enum.JFPointType.ActivityPoint, data)
		if orderPoint > 0 {
			return nil, fmt.Errorf("积分数与订单积分数不一致")
		}

		orders[k].Merge(data)
	}

	if reqBuyNum != 0 || reqActNum != 0 {
		return nil, fmt.Errorf("积分数与订单积分数不一致")
	}

	fmt.Println("comporder", orders)
	return orders, nil
}

//CreateUseFDRecord 创建使用记账记录
func (o *DBJFFd) CreateUseFDRecord(req *JFFDUseRequest, orders types.XMaps) (types.XMap, error) {
	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	//查询积分上游
	JFChannels, _, _, err := dbTrans.Query(sql.GetJFUpChannel, map[string]interface{}{})
	if err != nil || JFChannels.Len() != 1 {
		dbTrans.Rollback()
		return nil, fmt.Errorf("查询积分上游异常,count:%d,err:%v", JFChannels.Len(), err)
	}

	//添加请求记录
	requestID, count, _, _, err := utils.Insert(dbTrans, sql.GetNewFDRequestID, sql.CreateJFFDRequestRecord, map[string]interface{}{
		"down_channel_no":   req.ChannelNO,
		"request_no":        req.RequestNo,
		"request_type":      enum.JFFDType.OrderUse,
		"relation_no":       req.ChannelPayNo,
		"buy_send_num":      req.BuyPointNum,
		"activity_send_num": req.ActivityPointNum,
	})
	if err != nil || count != 1 {
		dbTrans.Rollback()
		return nil, fmt.Errorf("创建积分记账请求记录异常,count:%d,err:%v", count, err)
	}

	//添加订单积分记账记录
	for _, v := range orders {

		_, count, _, _, err := utils.Insert(dbTrans, sql.GetNewFDOrderID, sql.CreateJFFDOrderRecord, map[string]interface{}{
			"down_channel_no":   req.ChannelNO,
			"fd_channel_no":     JFChannels.Get(0).GetString("up_channel_no"),
			"request_type":      enum.JFFDType.OrderUse,
			"order_id":          v.GetString("order_id"),
			"request_id":        requestID,
			"point_num":         v.GetInt("point_num"),
			"buy_send_num":      v.GetInt(enum.JFPointType.BuyPoint),
			"activity_send_num": v.GetInt(enum.JFPointType.ActivityPoint),
		})
		if err != nil || count != 1 {
			dbTrans.Rollback()
			return nil, fmt.Errorf("创建积分订单记账记录异常,count:%d,err:%v", count, err)
		}
	}

	//查询添加的请求记录
	rows, _, _, err := dbTrans.Query(sql.CheckFdRequestRecord, map[string]interface{}{
		"down_channel_no": req.ChannelNO,
		"request_no":      req.RequestNo,
		"request_type":    enum.JFFDType.OrderUse,
	})
	if err != nil || rows.IsEmpty() {
		dbTrans.Rollback()
		return nil, fmt.Errorf("查询添加的记账请求记录异常,count:%d,err:%v", rows.Len(), err)
	}

	dbTrans.Commit()
	return rows.Get(0), nil
}

//JFVoidCheck 积分过期作废检查
func (o *DBJFFd) JFVoidCheck(dbTrans db.IDBExecuter, req *JFFDVoidRequest) (types.XMap, error) {
	gains, _, _, err := dbTrans.Query(sql.GetGainPointsInfo, map[string]interface{}{
		"gain_no": req.GainNo,
	})
	if err != nil {
		return nil, err
	}
	if gains.IsEmpty() {
		return nil, fmt.Errorf("积分获取记录不存在")
	}
	gainInfo := gains.Get(0)
	extend := types.NewXMap()
	err = json.Unmarshal([]byte(gainInfo.GetString("extend_info")), &extend)
	if err != nil {
		return nil, err
	}
	gainInfo.Merge(extend)

	voids, _, _, err := dbTrans.Query(sql.GetHasVoidpointInfo, map[string]interface{}{
		"down_channel_no": req.ChannelNO,
		"request_type":    enum.JFFDType.Void,
		"gain_no":         req.GainNo,
	})
	if err != nil {
		return nil, err
	}
	if voids.IsEmpty() {
		return gainInfo, nil
	}
	if gainInfo.GetInt("num") < voids.Get(0).GetInt("buy_send_num")+voids.Get(0).GetInt("activity_send_num") {
		return nil, fmt.Errorf("请求作废积分超过发放积分")
	}

	return gainInfo, nil
}

//CreateJFVoidFDRecord 添加积分作废记录
func (o *DBJFFd) CreateJFVoidFDRecord(req *JFFDVoidRequest, gainInfo types.XMap) (types.XMap, error) {
	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}

	//1.查询积分上游
	JFChannels, _, _, err := dbTrans.Query(sql.GetJFUpChannel, map[string]interface{}{})
	if err != nil || JFChannels.Len() != 1 {
		dbTrans.Rollback()
		return nil, fmt.Errorf("查询积分上游异常,count:%d,err:%v", JFChannels.Len(), err)
	}

	//2.添加作废请求记录
	requestID, count, _, _, err := utils.Insert(dbTrans, sql.GetNewFDRequestID, sql.CreateJFFDRequestRecord, map[string]interface{}{
		"down_channel_no":   req.ChannelNO,
		"request_no":        req.RequestNo,
		"request_type":      enum.JFFDType.Void,
		"relation_no":       req.GainNo,
		"buy_send_num":      req.BuyPointNum,
		"activity_send_num": req.ActivityPointNum,
	})
	if err != nil || count != 1 {
		dbTrans.Rollback()
		return nil, fmt.Errorf("创建积分作废记账请求记录异常,count:%d,err:%v", count, err)
	}

	//3.检查作废积分是否超过取出积分
	_, err = o.JFVoidCheck(dbTrans, req)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	//4.添加作废记账记录
	_, count, _, _, err = utils.Insert(dbTrans, sql.GetNewFDOrderID, sql.CreateJFFDOrderRecord, map[string]interface{}{
		"down_channel_no":   req.ChannelNO,
		"fd_channel_no":     JFChannels.Get(0).GetString("up_channel_no"),
		"request_type":      enum.JFFDType.Void,
		"order_id":          gainInfo.GetString("order_id"),
		"request_id":        requestID,
		"point_num":         gainInfo.GetInt("num"),
		"buy_send_num":      req.BuyPointNum,
		"activity_send_num": req.ActivityPointNum,
	})
	if err != nil || count != 1 {
		dbTrans.Rollback()
		return nil, fmt.Errorf("创建积分订单记账记录异常,count:%d,err:%v", count, err)
	}

	//5.查询添加的请求记录
	rows, _, _, err := dbTrans.Query(sql.CheckFdRequestRecord, map[string]interface{}{
		"down_channel_no": req.ChannelNO,
		"request_type":    enum.JFFDType.Void,
		"request_no":      req.RequestNo,
	})
	if err != nil || rows.IsEmpty() {
		dbTrans.Rollback()
		return nil, fmt.Errorf("查询添加的记账请求记录异常,count:%d,err:%v", rows.Len(), err)
	}

	dbTrans.Commit()
	return rows.Get(0), nil
}

//CompareVoidAndGain 比较作废与获取记录信息
func (o *DBJFFd) CompareVoidAndGain(req *JFFDVoidRequest, gainInfo types.XMap) error {
	if req.ActivityPointNum > 0 && gainInfo.GetInt("point_type") != 1 {
		return fmt.Errorf("作废积分类型不匹配")
	}
	if req.BuyPointNum > 0 && gainInfo.GetInt("point_type") != 2 {
		return fmt.Errorf("作废积分类型不匹配")
	}

	if req.ActivityPointNum <= 0 && req.BuyPointNum <= 0 {
		return fmt.Errorf("积分数据不正确")
	}

	return nil
}

//PointCalc 积分计算
func PointCalc(orderRestPoints, fdPoints int, pointType string, data types.XMap) (int, int, types.XMap) {
	if orderRestPoints <= 0 || fdPoints <= 0 {
		return orderRestPoints, fdPoints, data
	}

	if orderRestPoints > fdPoints {
		data.SetValue(pointType, fdPoints)
		return (orderRestPoints - fdPoints), 0, data
	}

	data.SetValue(pointType, orderRestPoints)
	return 0, (fdPoints - orderRestPoints), data
}
