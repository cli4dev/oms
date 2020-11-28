package main

import (
	"github.com/micro-plat/hydra/hydra"
)

type mgrweb struct {
	*hydra.MicroApp
}

func main() {
	app := &mgrweb{
		hydra.NewApp(
			hydra.WithPlatName("oms_yxtx"),
			hydra.WithSystemName("oms-web"),
			hydra.WithServerTypes("web"),
		),
	}

	app.install()
	app.Start()
}

func (s *mgrweb) install() {
	s.IsDebug = false
	s.Conf.WEB.SetMainConf(`{"address":":8088"}`)
	s.Conf.WEB.SetSubConf("static", `{
			"dir":"./static",
			"rewriters":["*"],
			"exts":[".ttf",".woff",".woff2"]			
	}`)
}
