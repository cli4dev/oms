package main

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/sso/sdk/sso"
	"gitlab.100bm.cn/micro-plat/dds/dds"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/model"

	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/beanpay/account"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/dds/dictionary"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/lcs/life"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/audit"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/canton"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/down"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/notify"
	oms "gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/order"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/product"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/refund"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/oms/up"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/report"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/tsk/system"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/vds/channel"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/services/vds/order"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {

	//每个请求执行前执行
	app.Handling(func(ctx *context.Context) (rt interface{}) {
		//验证jwt
		jwt, err := ctx.Request.GetJWTConfig()
		if err != nil {
			return err
		}
		for _, u := range jwt.Exclude {
			if u == ctx.Service {
				return nil
			}
		}

		if err := sso.CheckAndSetMember(ctx); err != nil {
			return err
		}
		return nil
	})

	app.Initializing(func(c component.IContainer) error {
		var conf model.Conf
		if err := c.GetAppConf(&conf); err != nil {
			return err
		}
		if err := conf.Valid(); err != nil {
			return err
		}
		model.SaveConf(c, &conf)
		dds.Bind(app, "db")
		//检查db配置是否正确
		if _, err := c.GetDB(); err != nil {
			return err
		}

		if err := sso.Bind(app, conf.SsoApiHost, conf.Ident, conf.Secret); err != nil {
			return err
		}
		app.Micro("/oms/down/channel", down.NewChannelHandler, "*")
		app.Micro("/oms/order/info", oms.NewInfoHandler, "*")
		app.Micro("/vds/order/info", order.NewInfoHandler, "*")
		app.Micro("/vds/order/notify", order.NewNotifyHandler, "*")
		app.Micro("/oms/refund/info", refund.NewInfoHandler, "*")
		app.Micro("/oms/up/channel", up.NewChannelHandler, "*")
		app.Micro("/tsk/system/task", system.NewTaskHandler, "*")
		app.Micro("/oms/canton/info", canton.NewInfoHandler, "*")
		app.Micro("/vds/order/query", order.NewQueryHandler, "*")
		app.Micro("/beanpay/account/info", account.NewInfoHandler, "*")
		app.Micro("/oms/down/product", down.NewProductHandler, "*")
		app.Micro("/oms/notify/info", notify.NewInfoHandler, "*")
		app.Micro("/oms/up/product", up.NewProductHandler, "*")
		app.Micro("/lcs/life/time", life.NewTimeHandler, "*")
		app.Micro("/oms/down/shelf", down.NewShelfHandler, "*")
		app.Micro("/dds/dictionary/info", dictionary.NewInfoHandler, "*")
		app.Micro("/vds/order/exp", order.NewExpHandler, "*")
		app.Micro("/oms/up/shelf", up.NewShelfHandler, "*")
		app.Micro("/oms/refund/up/return", refund.NewUpReturnHandler, "*")
		app.Micro("/oms/audit/info", audit.NewInfoHandler, "*")
		app.Micro("/vds/channel/info", channel.NewInfoHandler, "*")
		app.Micro("/oms/product/line", product.NewLineHandler, "*")
		app.Micro("/vds/channel/error/code", channel.NewErrorCodeHandler, "*")
		app.Micro("/oms/order/delivery", oms.NewDeliveryHandler, "*")
		app.Micro("/beanpay/account/record", account.NewRecordHandler, "*")
		app.Micro("/oms/report/downchannel", report.NewReportDownchannelHandler, "*")
		app.Micro("/oms/report/upchannel", report.NewReportUpchannelHandler, "*")
		app.Micro("/oms/report/profit", report.NewReportProfitHandler, "*")

		return nil
	})
}
