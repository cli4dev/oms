package main

import "github.com/micro-plat/hydra/hydra"

var app = hydra.NewApp(
	hydra.WithPlatName("oms"),
	hydra.WithSystemName("apiserver"),
	hydra.WithServerTypes("api"))

func main() {
	app.Start()
}
