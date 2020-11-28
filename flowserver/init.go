package main

import (
	"gitlab.100bm.cn/micro-plat/oms/flowserver/services/invoice"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/services/jfpps"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/services/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/services/red"

	"gitlab.100bm.cn/micro-plat/oms/flowserver/services/refund"
	"gitlab.100bm.cn/micro-plat/vds/vds/jcsdk"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/qtask/qtask"
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

		//每隔10秒将未完成的任务放入队列，删除3天前的任务,app:*hydrapp.MicroApp
		qtask.Bind(app, 10)
		//初始化  注入基础发货系统服务

		if err := jcsdk.Bind(app, c, app.PlatName, "", "", 10); err != nil {
			return err
		}

		return nil
	})

	// 服务注册
	// 订单
	app.Flow("/order/pay", order.NewPayHandler)           //上、下游支付流程
	app.Flow("/order/bind", order.NewBindHandler)         //绑定流程
	app.Flow("/order/delivery", order.NewDeliveryHandler) //发货流程
	app.Flow("/order/notify", order.NewNotifyHandler)     //通知流程
	app.Flow("/order/overtime", order.NewOvertimeHandler) //超时流程

	// 退货
	app.Flow("/refund/product", refund.NewReturnHandler) //退货
	// 退款
	app.Flow("/refund/pay", refund.NewRefundHandler) //退款
	// 退款通知
	app.Flow("/refund/notify", refund.NewNotifyHandler) //退款通知
	// 退款超时
	app.Flow("/refund/overtime", refund.NewOverTimeHandler) //退款超时处理

	app.Flow("/point/get/pay/up", jfpps.NewPayHandler)         //积分领取记账
	app.Flow("/point/use/refund/fd", jfpps.NewRefundHandler)   //积分使用退款记账创建
	app.Flow("/point/get/return", jfpps.NewJFGetReturnHandler) //积分发放退货完成
	app.Flow("/point/get/refund", jfpps.NewJFGetRefundHandler) //积分发放退款

	// 开票流程
	app.Flow("/invoice/start", invoice.NewStartHandler)   //开始开票
	app.Flow("/invoice/finish", invoice.NewFinishHandler) //完成开票
	app.Flow("/invoice/notify", invoice.NewNotifyHandler) //开票通知
	app.Flow("/invoice/pay", invoice.NewPayHandler)       //开票记账通知

	// 开票冲红流程
	app.Flow("/red/start", red.NewStartHandler)       //开始冲红
	app.Flow("/red/finish", red.NewFinishHandler)     //完成冲红
	app.Flow("/red/notify", red.NewNotifyHandler)     //冲红通知
	app.Flow("/invoice/refund", red.NewRefundHandler) //开票记账通知

}
