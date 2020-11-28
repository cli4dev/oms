package jfpps

import (
	"fmt"
	"strings"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
)

//JFPreOrder 积分预下单
type JFPreOrder struct {
	c  component.IContainer
	db *DBJFPreOrder
}

//NewJFPreOrder 初始化积分预下单
func NewJFPreOrder(c component.IContainer) *JFPreOrder {
	return &JFPreOrder{
		c:  c,
		db: NewDBJFPreOrder(c),
	}
}

//Request 预下单请求
func (o *JFPreOrder) Request(req *JFPreRequestInfo) (types.XMap, error) {
	//1.检查预下单是否存在
	preInfo, err := o.db.CheckJFPreOrder(req)
	if err != nil || preInfo != nil {
		return preInfo, err
	}

	//2.检查产品配置
	product, err := o.db.CheckJFProduct(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(product)
	//3.锁积分产品,检查预下单
	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return nil, err
	}
	preInfo, err = o.db.LockJFProduct(dbTrans, product, req)
	if err != nil || preInfo != nil {
		dbTrans.Rollback()
		return preInfo, err
	}

	//4.创建预下单
	preInfo, err = o.db.CreateJFPreOrder(dbTrans, product, req)
	if err != nil {
		dbTrans.Rollback()
		return nil, err
	}

	//5.请求积分系统发放待激活积分
	postMap := map[string]interface{}{
		"up_channel_no":      product.GetString("up_channel_no"),
		"request_channel_no": product.GetString("ext_channel_no"),
		"pre_order_id":       preInfo.GetInt64("pre_order_id"),
		"point_type":         req.PointType,
		"point_num":          req.Num,
		"over_time":          req.OverTime,
		"ext_activity_no":    req.ActivityNo,
		"ext_user_no":        req.UserNo,
		"ext_memo":           req.Memo,
	}
	fmt.Println(postMap)
	status, resultStr, param, err := o.c.Request("/preorder/request@jfpps.oms_debug", "GET", nil, postMap, true)
	if err != nil || status != 200 || !strings.Contains(resultStr, "success") {
		dbTrans.Rollback()
		return nil, fmt.Errorf("积分预下单失败,status:%d,err:%+v,param:%+v,resultStr:%s", status, err, param, resultStr)
	}

	dbTrans.Commit()
	return preInfo, nil
}

//Cancel 待激活积分取消
func (o *JFPreOrder) Cancel(channelNo string, preRequestNo string) error {
	//1.获取预下单信息
	preInfo, err := o.db.CancelPreOrderQuery(channelNo, preRequestNo)
	if err != nil {
		return err
	}

	dbTrans, err := o.c.GetRegularDB().Begin()
	if err != nil {
		return err
	}
	//2.积分预下单设置为已取消
	err = o.db.CancelJFPreOrder(dbTrans, channelNo, preRequestNo)
	if err != nil {
		dbTrans.Rollback()
		return nil
	}

	//3.检查是否存在订单
	err = o.db.CancelOrderCheck(dbTrans, preInfo)
	if err != nil {
		dbTrans.Rollback()
		return nil
	}

	//4.请求积分系统取消积分
	postMap := map[string]interface{}{
		"up_channel_no":      preInfo.GetString("up_channel_no"),
		"request_channel_no": preInfo.GetString("ext_channel_no"),
		"ext_user_no":        preInfo.GetString("user_no"),
		"pre_order_id":       preInfo.GetInt64("pre_order_id"),
	}
	fmt.Println(postMap)
	status, resultStr, param, err := o.c.Request("/preorder/cancel@jfpps.oms_debug", "GET", nil, postMap, true)
	if err != nil || status != 200 || !strings.Contains(resultStr, "success") {
		dbTrans.Rollback()
		return fmt.Errorf("积分预下单撤销失败,status:%d,err:%+v,param:%+v,resultStr:%s", status, err, param, resultStr)
	}

	dbTrans.Commit()
	return nil
}
