// +build !prod

package main

import (
	"fmt"

	"github.com/micro-plat/hydra/conf"
	fasconf "gitlab.100bm.cn/micro-plat/fas/fas/modules/const/conf"
	"gitlab.100bm.cn/micro-plat/fas/fas/sdk"
)

func init() {
	app.IsDebug = true

	redisAddrs := []string{"192.168.0.111:6379", "192.168.0.112:6379", "192.168.0.113:6379", "192.168.0.114:6379", "192.168.0.115:6379", "192.168.0.116:6379"}
	redisQueueConf := conf.QueueConf(conf.NewRedisQueueConf(redisAddrs, 1).WithTimeout(10, 10, 10, 100))

	app.Ready(func() error {
		redisQueues := conf.NewQueues()
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:pay", app.PlatName), "/order/pay/down")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:bind", app.PlatName), "/order/bind")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:delivery", app.PlatName), "/order/delivery/start")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:delivery:finish", app.PlatName), "/order/delivery/finish")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:uppay", app.PlatName), "/order/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:notify", app.PlatName), "/order/notify")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:overtime", app.PlatName), "/order/overtime")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:order:delivery:unknown", app.PlatName), "/order/overtime/delivery")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return", app.PlatName), "/refund/product/return")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return:finish", app.PlatName), "/refund/product/finish")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:pay", app.PlatName), "/refund/pay/down")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:uppay", app.PlatName), "/refund/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:order:fail", app.PlatName), "/refund/pay/orderfail")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:notify", app.PlatName), "/refund/notify")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:overtime", app.PlatName), "/refund/overtime")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:refund:return:unknown", app.PlatName), "/refund/overtime/unknown")

		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:order:uppay", app.PlatName), "/point/get/pay/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:return:finish", app.PlatName), "/point/get/return/finish")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:get:refund:uppay", app.PlatName), "/point/get/refund/up")
		redisQueues = redisQueues.Append(fmt.Sprintf("%v:point:use:refund:fd", app.PlatName), "/point/use/refund/fd")

		app.Conf.MQC.SetQueues(redisQueues)
		return nil
	})
	app.Conf.RPC.SetMain(conf.NewRPCServerConf("8083"))

	app.Conf.MQC.SetMain(conf.NewMQCServerConf().WithEnable())
	app.Conf.MQC.SetServer(redisQueueConf)

	app.Conf.Plat.SetQueue(redisQueueConf)
	app.Conf.Plat.SetVarConf("fd", "fd", `
	{"server_name":"apiserver",
	"plat_name":"fas_debug",
    "order_source": 211}
	`)
	sdk.SetFdConf(&fasconf.InnerConf{
		FdRequestURL: "@apiserver.fas_debug",
	})
}
