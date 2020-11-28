package notify

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

// IDBNotify 退款通知DB接口
type IDBNotify interface {
	CheckByDB(notifyID string) error
	BeginNotifyByDB(notifyID string) error
	UpdateNotifyToSuccessByDB(input map[string]interface{}) error
	UpdateNotifyToFailByDB(input map[string]interface{}) error
	GetNotifyInfoByDB(notifyID string) (types.IXMap, error)
	BuildNotifyParamByDB(data types.IXMap) (map[string]interface{}, error)
	RemoteNotify(param map[string]interface{}, url string) (string, error)
}

// DBNotify 退款通知对象
type DBNotify struct {
	c component.IContainer
}

//NewDBNotify 创建退款通知信息对象
func NewDBNotify(c component.IContainer) *DBNotify {
	return &DBNotify{
		c: c,
	}
}

// CheckByDB 检查通知状态
func (n *DBNotify) CheckByDB(notifyID string) error {
	dbt := n.c.GetRegularDB()
	// 获取等待通知记录
	notify, q, a, err := dbt.Query(sql.SqlGetWaitNotifyInfo, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil {
		return fmt.Errorf("获取等待通知记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notify.Len())
	}
	if notify.IsEmpty() {
		return context.NewErrorf(errorcode.NOTIFY_STATUS_ERROR.Code, "退款通知状态异常")
	}
	return nil
}

// BeginNotifyByDB 开启通知
func (n *DBNotify) BeginNotifyByDB(notifyID string) error {
	dbt := n.c.GetRegularDB()
	// 开启通知
	count, q, a, err := dbt.Execute(sql.SqlUpdateNotifyToProcess, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || count <= 0 {
		return fmt.Errorf("开启通知错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	return nil
}

// UpdateNotifyToSuccessByDB 更新通知为成功
func (n *DBNotify) UpdateNotifyToSuccessByDB(input map[string]interface{}) error {
	dbTrans, err := n.c.GetRegularDB().Begin()
	if err != nil {
		return fmt.Errorf("事务开启失败:err:%+v", err)
	}
	// 更新退款表告知状态为成功
	count, q, a, err := dbTrans.Execute(sql.SqlUpdateRefundNotifyToSuccess, input)
	if err != nil {
		dbTrans.Rollback()
		return fmt.Errorf("更新退款表告知状态成功错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	// 更新通知表通知状态为成功
	count, q, a, err = dbTrans.Execute(sql.SqlUpdateNotifyToSuccess, input)
	if err != nil || count <= 0 {
		dbTrans.Rollback()
		return fmt.Errorf("更新通知表通知状态错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	dbTrans.Commit()
	return nil
}

// UpdateNotifyToFailByDB 通知失败
func (n *DBNotify) UpdateNotifyToFailByDB(input map[string]interface{}) error {
	dbt := n.c.GetRegularDB()
	// 通知失败
	count, q, a, err := dbt.Execute(sql.SqlUpdateFailNotify, input)
	if err != nil || count <= 0 {
		return fmt.Errorf("更新通知结果错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, count)
	}
	return nil
}

// GetNotifyInfoByDB 查询通知信息
func (n *DBNotify) GetNotifyInfoByDB(notifyID string) (types.IXMap, error) {
	dbt := n.c.GetRegularDB()
	// 获取通知详情
	notify, q, a, err := dbt.Query(sql.SqlGetRequestNotifyInfo, map[string]interface{}{
		"notify_id": notifyID,
	})
	if err != nil || notify.IsEmpty() {
		return nil, fmt.Errorf("获取等待通知记录错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, notify.Len())
	}

	// 获取退款信息
	refunds, q, a, err := dbt.Query(sql.SqlGetRefundInfo, notify.Get(0))
	if err != nil || refunds.IsEmpty() {
		return nil, fmt.Errorf("获取退款信息错误(err:%v),sql:%s,input:%+v,count:%d", err, q, a, refunds.Len())
	}

	result := refunds.Get(0)
	result.SetValue("notify_id", notifyID)
	result.SetValue("notify_url", notify.Get(0).GetString("notify_url"))
	return result, nil
}

// BuildNotifyParamByDB 构建通知参数
func (n *DBNotify) BuildNotifyParamByDB(notifyInfo types.IXMap) (map[string]interface{}, error) {
	status, code, msg := errorcode.SetFlowRecordStatus(notifyInfo.GetInt("status"), errorcode.RequestFlowType.Refund, notifyInfo.GetString("fail_code"), notifyInfo.GetString("fail_msg"))

	res := map[string]interface{}{
		"channel_no":  notifyInfo.GetString("channel_no"),
		"request_no":  notifyInfo.GetString("request_no"),
		"refund_id":   notifyInfo.GetInt64("refund_id"),
		"refund_no":   notifyInfo.GetString("refund_no"),
		"refund_num":  notifyInfo.GetInt("refund_num"),
		"status":      status,
		"failed_code": code,
		"failed_msg":  msg,
	}

	if notifyInfo.GetString("account_no") != "" {
		res["account_no"] = notifyInfo.GetString("account_no")
	}
	if notifyInfo.GetInt("point_num") > 0 && notifyInfo.GetInt("status") == enum.RefundResultStatus.Success {
		pointList := types.NewXMaps()
		if notifyInfo.GetInt("activity_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": notifyInfo.GetInt("activity_send_num"),
				"point_type":  enum.PointType.Activity,
			})
		}
		if notifyInfo.GetInt("buy_send_num") > 0 {
			pointList.Append(map[string]interface{}{
				"point_count": notifyInfo.GetInt("buy_send_num"),
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
		return content, fmt.Errorf("调用退款通知接口错误,err:%v,status:%v,content:%v", err, status, content)
	}
	return content, nil
}
