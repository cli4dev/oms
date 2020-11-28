package order

import (
	"encoding/json"
	"fmt"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/enum"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/errorcode"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/sql"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/types"
)

// IDBNotify 数据层接口
type IDBNotify interface {
	RemoteNotify(param map[string]interface{}, url string) (string, error)
	CheckOrderAndNotifyByDB(notifyID int64) (types.XMap, bool, error)
	StartByDB(param map[string]interface{}) error
	SuccessByDB(param map[string]interface{}) error
	FailedByDB(param map[string]interface{}) error
	BuildNotifyParamByDB(data types.XMap) (map[string]interface{}, error)
}

// DBNotify 通知数据层
type DBNotify struct {
	c component.IContainer
}

// NewDBNotify 构建DBNotify
func NewDBNotify(c component.IContainer) *DBNotify {
	return &DBNotify{c: c}
}

// CheckOrderAndNotifyByDB 检查订单和通知
func (n *DBNotify) CheckOrderAndNotifyByDB(notifyID int64) (types.XMap, bool, error) {

	db := n.c.GetRegularDB()
	// 检查通知
	notifys, sqlStr, args, err := db.Query(sql.CheckNotify, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("检查通知发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if notifys.IsEmpty() {
		return nil, true, fmt.Errorf("通知不存在或已完成通知,notify_id:%d", notifyID)
	}

	// 检查订单
	orderID := notifys.Get(0).GetInt64("order_id")

	orders, sqlStr, args, err := db.Query(sql.ChechOrderForNotify, map[string]interface{}{
		"order_id": orderID,
	})
	if err != nil {
		return nil, false, fmt.Errorf("通知检查订单发生异常,err:%v,sql:%v,args:%v", err, sqlStr, args)
	}
	if orders.IsEmpty() {
		return nil, true, fmt.Errorf("订单不存在或已完成通知,order_id:%d", orderID)
	}
	order := orders.Get(0)

	if order.GetInt("status") == 90 && order.GetInt("point_num") != order.GetInt("buy_send_num")+order.GetInt("activity_send_num") {
		return nil, false, fmt.Errorf("订单积分记账未创建")
	}

	order.Merge(notifys.Get(0))
	return order, false, nil
}

// StartByDB 开始通知
func (n *DBNotify) StartByDB(param map[string]interface{}) error {
	db := n.c.GetRegularDB()
	row, sqlStr, args, err := db.Execute(sql.StartNotify, param)
	if err != nil || row != 1 {
		return fmt.Errorf("修改通知状态为正在通知发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// SuccessByDB 成功通知
func (n *DBNotify) SuccessByDB(param map[string]interface{}) error {
	trans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}

	row, sqlStr, args, err := trans.Execute(sql.SuccessOrderNotify, param)
	if err != nil || row != 1 {
		trans.Rollback()
		return fmt.Errorf("修改订单通知状态为通知成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	row, sqlStr, args, err = trans.Execute(sql.SuccessNotify, param)
	if err != nil || row != 1 {
		trans.Rollback()
		return fmt.Errorf("修改通知状态为成功发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}

	trans.Commit()
	return nil
}

// FailedByDB 通知失败
func (n *DBNotify) FailedByDB(param map[string]interface{}) error {
	db := n.c.GetRegularDB()
	row, sqlStr, args, err := db.Execute(sql.FailedNotify, param)
	if err != nil || row != 1 {
		return fmt.Errorf("修改通知状态发生异常,cnt:%d,err:%v,sql:%v,args:%v", row, err, sqlStr, args)
	}
	return nil
}

// BuildNotifyParamByDB 构建通知参数
func (n *DBNotify) BuildNotifyParamByDB(data types.XMap) (map[string]interface{}, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(data.GetInt("status"), errorcode.RequestFlowType.Order, data.GetString("fail_code"), data.GetString("fail_msg"))
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
	if data.GetInt("point_num") > 0 && data.GetInt("status") == enum.OrderResultStatus.Failed {
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

//RemoteNotify 远程通知
func (n *DBNotify) RemoteNotify(param map[string]interface{}, url string) (string, error) {
	content, status, err := utils.HTTPRequest(param, "post", url)
	if err != nil || status != context.ERR_OK || content != "success" {
		return "", fmt.Errorf("调用通知接口错误,err:%v,status:%v,content:%v,param:%+v", err, status, content, param)
	}
	return content, nil
}
