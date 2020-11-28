package main

import "github.com/micro-plat/hydra/hydra"

var app = hydra.NewApp(
	hydra.WithPlatName("oms"),
	hydra.WithSystemName("flowserver"),
	hydra.WithServerTypes("mqc-cron-rpc"))

func main() {
	app.Start()
}
