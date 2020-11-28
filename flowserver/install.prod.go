// +build prod

package main

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	fasconf "gitlab.100bm.cn/micro-plat/fas/fas/modules/const/conf"
	"gitlab.100bm.cn/micro-plat/fas/fas/sdk"
)

func init() {

	redisQueueConf := conf.QueueConf(conf.NewRedisCacheConfForProd(1, "#redis_address").WithTimeout(10, 10, 10, 100))

	app.Ready(func() error {
		redisQueues := conf.NewQueues()
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:pay", app.PlatName), "/order/pay/down")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:bind", app.PlatName), "/order/bind")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:delivery", app.PlatName), "/order/delivery/start")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v.order.delivery_finish", app.PlatName), "/order/delivery/finish")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:up_pay", app.PlatName), "/order/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:notify", app.PlatName), "/order/notify")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:overtime", app.PlatName), "order/overtime")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:delivery_unknown", app.PlatName), "order/overtime/delivery")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return", app.PlatName), "/refund/product/return")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return_complete", app.PlatName), "/refund/product/complete")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:pay", app.PlatName), "/refund/pay/down")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:up_pay", app.PlatName), "/refund/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:order_fail", app.PlatName), "/refund/pay/orderfail")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:notify", app.PlatName), "/refund/notify")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:overtime", app.PlatName), "/refund/overtime")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return_unknown", app.PlatName), "/refund/overtime/unknown")

		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:order:uppay", app.PlatName), "/point/get/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:return:finish", app.PlatName), "/point/get/return/finish")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:refund:uppay", app.PlatName), "/point/get/refund/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:use:refund:fd", app.PlatName), "/point/use/refund/fd")
		app.Conf.MQC.SetQueues(redisQueues)
		return nil
	})

	app.Conf.RPC.SetMain(conf.NewRPCServerConf("#rpc_port"))
	app.Conf.MQC.SetMain(conf.NewMQCServerConf().WithEnable())
	app.Conf.MQC.SetServer(&redisQueueConf)

	app.Conf.CRON.SetMain(conf.NewCronServerConf().WithEnable())
	app.Conf.Plat.SetQueue(redisQueueConf)
	app.Conf.Plat.SetVarConf("fd", "fd", `
	{"server_name":"apiserver",
	"plat_name":"fas",
    "order_source": "#order_source"}
	`)
	sdk.SetFdConf(&fasconf.InnerConf{
		FdRequestURL: "@apiserver.fas_debug",
	})
}
