// +build prod
// +build oracle

package main

import (
	"github.com/micro-plat/hydra/conf"
	_ "github.com/zkfy/go-oci8"
)

func init() {
	app.Conf.Plat.SetDB(conf.NewOracleConfForProd("#db_connString"))
}
