// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

func (s *sdkserver) install() {
	s.IsDebug = true

	redisAddrs := []string{"192.168.0.111:6379", "192.168.0.112:6379", "192.168.0.113:6379", "192.168.0.114:6379", "192.168.0.115:6379", "192.168.0.116:6379"}
	redisQueueConf := conf.QueueConf(conf.NewRedisQueueConf(redisAddrs, 1).WithTimeout(10, 10, 10, 100))

	s.Conf.API.SetMain(conf.NewAPIServerConf(":1111"))
	s.Conf.Plat.SetDB(conf.NewOracleConf("oms", "123456", "orcl136"))
	s.Conf.Plat.SetQueue(redisQueueConf)

}
