module gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api

go 1.12

require (
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/beanpay v0.0.0-00010101000000-000000000000
	github.com/micro-plat/hydra v0.12.2
	github.com/micro-plat/lib4go v0.3.0
	github.com/micro-plat/sso v0.0.0-20190902014109-5276afc9c026
	github.com/tealeg/xlsx v1.0.5
	gitlab.100bm.cn/micro-plat/dds/dds v0.0.0-00010101000000-000000000000
	gitlab.100bm.cn/micro-plat/oms/oms v0.0.0-00010101000000-000000000000
)

replace github.com/micro-plat/lib4go => ../../../../../github.com/micro-plat/lib4go

replace gitlab.100bm.cn/micro-plat/oms/apiserver => ../../../../../gitlab.100bm.cn/micro-plat/oms/apiserver

replace gitlab.100bm.cn/micro-plat/lcs/lcs => ../../../../../gitlab.100bm.cn/micro-plat/lcs/lcs

replace github.com/micro-plat/qtask => ../../../../../github.com/micro-plat/qtask

replace github.com/micro-plat/hydra => ../../../../../github.com/micro-plat/hydra

replace gitlab.100bm.cn/micro-plat/oms/oms => ../../../../../gitlab.100bm.cn/micro-plat/oms/oms

replace gitlab.100bm.cn/micro-plat/dds/dds => ../../../../../gitlab.100bm.cn/micro-plat/dds/dds

replace github.com/micro-plat/beanpay => ../../../../../github.com/micro-plat/beanpay
