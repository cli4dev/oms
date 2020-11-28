package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/hydra"
)

var app = hydra.NewApp(
	hydra.WithPlatName("oms_yxtx"),
	hydra.WithSystemName("oms-api"),
	hydra.WithServerTypes("api"))

func main() {
	app.Start()
}
