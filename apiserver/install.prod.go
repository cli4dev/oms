// +build prod

package main

import "github.com/micro-plat/hydra/conf"

func init() {
	redisQueueConf := conf.QueueConf(conf.NewRedisCacheConfForProd(1, "#redis_address").WithTimeout(10, 10, 10, 100))

	app.Conf.API.SetMain(conf.NewAPIServerConf("#api_port"))

	app.Conf.Plat.SetQueue(redisQueueConf)

	app.Conf.API.SetAuthes(conf.NewAuthes().WithServiceAuth(
		conf.NewServiceAuth("/single/apiserver/md5/auth@authserver.sas",
			"/order/query",
			"/order/request",
			"/refund/query",
			"/refund/general",
			"/refund/mandatory",
		).WithUIDAlias("channel_no")))

}
