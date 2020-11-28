// +build !prod
// +build !oracle

package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra/conf"
)

func init() {
	app.Conf.Plat.SetDB(conf.NewMysqlConf("oms_t", "123456", "192.168.0.36", "oms_t"))
}
