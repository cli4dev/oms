// +build prod

package main

import "github.com/micro-plat/hydra/conf"

//bindConf 绑定启动配置， 启动时检查注册中心配置是否存在，不存在则引导用户输入配置参数并自动创建到注册中心
func init() {

	//-------------------api配置信息--------------------
	app.Conf.API.SetMain(conf.NewAPIServerConf(":9091"))
	app.Conf.API.SetSubConf("app", `
	{	
		"secret":"B128F779D5741E701923346F7FA9F95C",
		"sso_api_host":"http://api.sso.100bm.cn",
		"ident":"oms_yxtx",
		"type":"md5",
		"plat_name":"sas",
		"server_name":"cipherserver"
	}`)
	app.Conf.API.SetHeaders(conf.NewHeader().WithCrossDomain().WithAllowHeaders("X-Requested-With", "Content-Type", "__jwt__", "X-Request-Id"))
	app.Conf.API.SetAuthes(
		conf.NewAuthes().WithJWT(
			conf.NewJWT("__jwt__", "HS512", "66c76999f5529f7601b8dab046214ad1", 36000, "/sso/login/verify", "/image/upload", "/dds/dictionary/get").WithHeaderStore()))

	//-------------------plat配置信息--------------------
	app.Conf.Plat.SetDB(conf.NewMysqlConfForProd("#db_connString").WithConnect(20, 10, 600))
}
