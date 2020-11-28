package main

import (
	"gitlab.100bm.cn/micro-plat/oms/oms/sdk"

	"github.com/micro-plat/hydra/component"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func (a *sdkserver) init() {
	a.Initializing(func(c component.IContainer) error {
		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}
		//检查消息队列配置
		if _, err := c.GetQueue(); err != nil {
			return err
		}

		return nil
	})

	//服务注册
	a.Micro("/sdk", sdk.NewTestHandle) //下单和查询订单
}
