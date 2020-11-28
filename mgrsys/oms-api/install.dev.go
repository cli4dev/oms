// +build !prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {
	app.IsDebug = true

	//-------------------api配置信息--------------------
	app.Conf.API.SetMain(conf.NewAPIServerConf(":9092"))
	app.Conf.API.SetSubConf("app", `
	{	
		"secret":"311124b57e468ff88e4f1c8743354314",
		"sso_api_host":"http://api.sso.18jiayou1.com:6689",
		"ident":"oms_t",
		"plat_name":"sas_debug",
		"type":"md5",
		"server_name":"cipherserver"
	}`)
	app.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain())
	app.Conf.API.SetAuthes(
		conf.NewAuthes().WithJWT(
			conf.NewJWT("__jwt__", "HS512", "800369fda1f8e8cbe9225a1421452148", 36000, "/sso/login/verify", "/image/upload", "/dds/dictionary/get", "/oms/report/upchannel/export", "/oms/report/downchannel/export", "/oms/report/profit/export").WithEnable().WithHeaderStore()))

	//-------------------plat配置信息--------------------

	app.Conf.Plat.SetDB(conf.NewOracleConf("oms_t", "123456", "orcl136"))
}
