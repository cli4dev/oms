package main

import "github.com/micro-plat/hydra/hydra"

type sdkserver struct {
	*hydra.MicroApp
}

func main() {
	app := &sdkserver{
		hydra.NewApp(
			hydra.WithPlatName("oms"),
			hydra.WithSystemName("sdkserver"),
			hydra.WithServerTypes("api")),
	}
	app.install()
	app.init()
	app.Start()
}
