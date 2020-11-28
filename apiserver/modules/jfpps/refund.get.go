package jfpps

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/refund"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/task"
)

// JFRefund JFRefund积分退款结构体
type JFRefund struct {
	*refund.Refund
	c component.IContainer
}

// NewJFRefund 积分退款
func NewJFRefund(c component.IContainer) *JFRefund {
	oo := &JFRefund{
		c: c,
	}

	return oo
}

//RefundRequest 积分退款请求
func (o *JFRefund) RefundRequest(input *JFRefundRequestBody) (types.IXMap, error) {
	res, ps := o.BuildJFParams(input)

	o.Refund = refund.NewRefund(o.c, NewDBJFRefund(o.c, ps), task.NewQTask(o.c))

	return o.Refund.GeneralRequest(res)
}

//BuildJFParams 构建积分退款参数
func (o *JFRefund) BuildJFParams(input *JFRefundRequestBody) (*refund.RequestBody, *JFRefundExtParams) {
	res := &refund.RequestBody{
		ChannelNo:       input.ChannelNo,
		RefundNo:        input.RefundNo,
		RequestNo:       input.RequestNo,
		RefundNum:       input.RefundNum,
		NotifyURL:       input.NotifyURL,
		RefundMerAmount: types.GetFloat64(input.RefundNum) * 0.01,
	}
	ps := &JFRefundExtParams{
		JFAmount: input.JFAmount,
	}

	return res, ps
}
