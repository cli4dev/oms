module gitlab.100bm.cn/micro-plat/oms/omscli

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/micro-plat/beanpay v0.1.0
	github.com/micro-plat/hydra v0.12.2
	github.com/micro-plat/lib4go v0.2.1
	github.com/micro-plat/qtask v0.0.0-20190708085554-931550afc18b
	github.com/urfave/cli v1.22.1
	github.com/zkfy/go-oci8 v0.0.0-20180327092318-ad9f59dedff0
	gitlab.100bm.cn/micro-plat/vds/vdscli v0.0.0-20191119033033-05530456318a
)

replace github.com/micro-plat/qtask => ../../../../github.com/micro-plat/qtask

replace github.com/micro-plat/beanpay => ../../../../github.com/micro-plat/beanpay

replace gitlab.100bm.cn/micro-plat/vds/vdscli => ../../../../gitlab.100bm.cn/micro-plat/vds/vdscli
