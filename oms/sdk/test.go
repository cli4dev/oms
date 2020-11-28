package sdk

import (
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/oms/modules/audit"

	"github.com/micro-plat/hydra/component"
)

// TestHandle 普通退款接入层对象
type TestHandle struct {
	c component.IContainer
}

// NewTestHandle 构建对象
func NewTestHandle(container component.IContainer) (u *TestHandle) {
	return &TestHandle{
		c: container,
	}
}

// Test1Handle 发货未知审核为成功测试
func (u *TestHandle) Test1Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.构建参数")
	input := &audit.RequestParams{
		OrderID:    1100003820,
		DeliveryID: 21460,
		AuditID:    3,
		AuditBy:    "111",
		AuditMsg:   "测试发货未知审核为成功",
	}

	ctx.Log.Info("2.执行审核")
	funcs, err := DeliveryUnknownAuditSucc(u.c, input, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("3.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test2Handle 发货未知审核为失败测试
func (u *TestHandle) Test2Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.构建参数")
	input := &audit.RequestParams{
		OrderID:    1100003820,
		DeliveryID: 21460,
		AuditID:    3,
		AuditBy:    "111",
		AuditMsg:   "测试发货未知审核为失败",
	}

	ctx.Log.Info("2.执行审核")
	funcs, err := DeliveryUnknownAuditFail(u.c, input, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("3.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test3Handle 退货未知审核为成功测试
func (u *TestHandle) Test3Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.构建参数")
	input := &audit.RequestParams{
		OrderID:    1100003070,
		RefundID:   1,
		DeliveryID: 1,
		AuditID:    2,
		AuditBy:    "111",
		AuditMsg:   "测试退货未知审核为成功",
	}

	ctx.Log.Info("2.执行审核")
	funcs, err := ReturnUnknownAuditSucc(u.c, input, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("3.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test4Handle 退货未知审核为失败测试
func (u *TestHandle) Test4Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.构建参数")
	input := &audit.RequestParams{
		OrderID:    1100003070,
		RefundID:   1,
		DeliveryID: 1,
		AuditID:    2,
		AuditBy:    "111",
		AuditMsg:   "测试退货未知审核为失败",
	}

	ctx.Log.Info("2.执行审核")
	funcs, err := ReturnUnknownAuditFail(u.c, input, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("3.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test5Handle 审核订单为部分成功
func (u *TestHandle) Test5Handle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("1.执行审核")
	funcs, err := OrderPartialSuccessAudit(u.c, 1100003071, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test6Handle 审核退货为部分成功
func (u *TestHandle) Test6Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := RefundPartialSuccessAudit(u.c, 1100003485, 2, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test7Handle 发货假成功审核
func (u *TestHandle) Test7Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := DeliveryFalseSuccAudit(u.c, 20058, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test8Handle 发货绑定重试测试
func (u *TestHandle) Test8Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := DeliveryBindRetry(u.c, 1100003100, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test9Handle 上游支付重试
func (u *TestHandle) Test9Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := UpPayRetry(u.c, 21314, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test10Handle 订单通知重试
func (u *TestHandle) Test10Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := OrderNotifyRetry(u.c, 1100003080, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test11Handle 退货重试
func (u *TestHandle) Test11Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := RefundReturnRetry(u.c, 1, 1, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test12Handle 上游退款重试
func (u *TestHandle) Test12Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := UpRefundRetry(u.c, 1, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test13Handle 下游退款重试
func (u *TestHandle) Test13Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := DownRefundRetry(u.c, 1, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

//Test14Handle 退款通知重试
func (u *TestHandle) Test14Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("1.执行审核")
	funcs, err := RefundNotifyRetry(u.c, 1100003070, 1, u.c.GetPlatName())
	if err != nil {
		return err
	}

	ctx.Log.Info("2.执行回调")
	for _, v := range funcs {
		err := v(u.c)
		if err != nil {
			return err
		}
	}

	return "success"
}

// //Test15Handle 创建订单测试
// func (u *TestHandle) Test15Handle(ctx *context.Context) (r interface{}) {
// 	info := &order.RequestInfo{
// 		ChannelNO:       "ycdown",
// 		RequestNO:       "tt1110000001",
// 		LineID:          2,
// 		CarrierNO:       "ZSH",
// 		ProvinceNo:      "CQ",
// 		Num:             1,
// 		Face:            200,
// 		Amount:          196,
// 		NotifyURL:       "http://xxxx.com/notify",
// 		RechargeAccount: "1110111110",
// 	}
// 	ctx.Log.Info("1.执行审核")
// 	res, err := CreateOrder(u.c, info, u.c.GetPlatName())
// 	if err != nil {
// 		return err
// 	}

// 	return res
// }

//Test16Handle 订单查询测试
// func (u *TestHandle) Test16Handle(ctx *context.Context) (r interface{}) {
// 	info := &order.QueryInfo{
// 		ChannelNO: "ycdown",
// 		RequestNO: "tt1110000001",
// 	}
// 	ctx.Log.Info("1.执行审核")
// 	res, err := QueryOrder(u.c, info, u.c.GetPlatName())
// 	if err != nil {
// 		return err
// 	}

// 	return res
// }
