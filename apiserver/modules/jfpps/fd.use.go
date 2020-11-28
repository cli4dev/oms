package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/const/enum"
)

// JFFd 积分记账
type JFFd struct {
	c  component.IContainer
	db *DBJFFd
}

// NewJFFd 积分记账处理
func NewJFFd(c component.IContainer) *JFFd {
	return &JFFd{
		c:  c,
		db: NewDBJFFd(c),
	}
}

//UseFdRequest 使用记账请求
func (o *JFFd) UseFdRequest(req *JFFDUseRequest) (types.XMap, error) {
	//1.检查请求是否已存在(唯一索引控制并发)
	record, err := o.db.CheckFdRequestRecord(req.ChannelNO, req.RequestNo, req.ChannelPayNo, enum.JFFDType.OrderUse)
	if err != nil || record != nil {
		return record, err
	}

	//2.检查支付记录对应订单与积分
	orders, err := o.db.CheckUseFdOrders(req)
	if err != nil {
		return nil, err
	}

	//3.创建记账请求记录与记账记录
	return o.db.CreateUseFDRecord(req, orders)
}

//VoidFdRequest 作废记账请求
func (o *JFFd) VoidFdRequest(req *JFFDVoidRequest) (types.XMap, error) {
	//1.检查请求记录是否已存在(唯一索引控制并发)
	record, err := o.db.CheckFdRequestRecord(req.ChannelNO, req.RequestNo, req.GainNo, enum.JFFDType.OrderUse)
	if err != nil || record != nil {
		return record, err
	}

	//2.检查获取记录与剩余积分
	gainInfo, err := o.db.JFVoidCheck(o.c.GetRegularDB(), req)
	if err != nil {
		return nil, err
	}

	//3.检查积分获取信息与作废信息是否一致
	err = o.db.CompareVoidAndGain(req, gainInfo)
	if err != nil {
		return nil, err
	}

	//4.创建记账请求记录与记账记录
	return o.db.CreateJFVoidFDRecord(req, gainInfo)
}
