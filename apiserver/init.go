package main

import (
	"github.com/micro-plat/hydra/component"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/services/invoice"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/services/jfpps"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/services/order"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/services/refund"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/services/response"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	app.Initializing(func(c component.IContainer) error {
		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}
		//检查消息队列配置
		if _, err := c.GetQueue(); err != nil {
			return err
		}

		return nil
	})

	//服务注册
	app.Micro("/order", order.NewOrderHandler) //下单和查询订单

	// 退款申请
	app.Micro("/refund", refund.NewRefundHandle) //退款接口

	app.Micro("/invoice", invoice.NewInvoiceHandler) //开票接口
	// 通知接收（测试）
	app.Micro("/refund/response", response.NewRefundNotify)
	app.Micro("/order/response", response.NewOrderNotify)

	app.Micro("/jf/preorder", jfpps.NewPreOrderHandler)       //待激活积分下单
	app.Micro("/jf/order/request", jfpps.NewJFOrderHandler)   //积分发放下单
	app.Micro("/jf/refund/request", jfpps.NewJFRefundHandler) //积分发放退款
	app.Micro("/jf/fd", jfpps.NewFDHandler)                   //积分记账请求
}
