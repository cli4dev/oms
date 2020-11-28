// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

func init() {
	app.IsDebug = true

	redisAddrs := []string{"192.168.0.111:6379", "192.168.0.112:6379", "192.168.0.113:6379", "192.168.0.114:6379", "192.168.0.115:6379", "192.168.0.116:6379"}
	redisQueueConf := conf.QueueConf(conf.NewRedisQueueConf(redisAddrs, 1).WithTimeout(10, 10, 10, 100))

	app.Conf.API.SetMain(conf.NewAPIServerConf(":8090"))

	app.Conf.Plat.SetQueue(redisQueueConf)

	// app.Conf.API.SetAuthes(conf.NewAuthes().WithServiceAuth(
	// 	conf.NewServiceAuth("/single/apiserver/md5/auth@authserver.sas_debug",
	// 		"/order/query",
	// 		"/order/request",
	// 		"/refund/query",
	// 		"/refund/general",
	// 		"/refund/mandatory",
	// 	).WithUIDAlias("channel_no")))

}
