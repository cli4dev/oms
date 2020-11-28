package main

import (
	"os"
	"strings"

	beanpay "github.com/micro-plat/beanpay/beanpay/const/sql/creator"
	"github.com/micro-plat/hydra/conf"
	_ "github.com/micro-plat/hydra/hydra"
	ldb "github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/logger"
	"github.com/micro-plat/lib4go/types"
	qtask "github.com/micro-plat/qtask/modules/const/sql/creator"

	vds "gitlab.100bm.cn/micro-plat/vds/vdscli/sql"
	"gitlab.100bm.cn/micro-plat/oms/omscli/sql"

	_ "github.com/urfave/cli"
)

func main() {

	//处理输入参数
	defer logger.Close()
	logger := logger.New("oms")
	if len(os.Args) < 1 {
		logger.Error("命令错误，请使用 ./oms [mysql;oracle]#[maxOpen]#[maxIdle]#[数据库连接串] 的形式启动程序")
		return
	}
	argsParam := strings.Split(os.Args[1], "#")
	if len(argsParam) != 4 {
		logger.Error("命令错误, 数据库链接串不正确")
		return
	}

	dbConf := conf.DBConf{
		Provider:   argsParam[0],
		ConnString: argsParam[3],
		LifeTime:   600,
	}

	switch argsParam[0] {
	case "oracle":
		dbConf.MaxOpen = types.GetInt(argsParam[1], 200)
		dbConf.MaxIdle = types.GetInt(argsParam[2], 100)
	case "mysql":
		dbConf.MaxOpen = types.GetInt(argsParam[1], 20)
		dbConf.MaxIdle = types.GetInt(argsParam[2], 10)
	default:
		logger.Error("命令错误, 数据库类型错误,只支持oracle/mysql")
		return
	}

	//构建数据库对象
	xdb, err := ldb.NewDB(dbConf.Provider, dbConf.ConnString, dbConf.MaxOpen, dbConf.MaxIdle, dbConf.LifeTime)
	if err != nil {
		logger.Error("构建数据链接对象失败,err:%+v", err)
		return
	}
	// 创建qtask数据库
	logger.Info("开始创建qtask数据库")
	if err = qtask.CreateDB(xdb); err != nil {
		logger.Error(err)
		return
	}
	logger.Info("qtask数据库创建完成")
	logger.Info("开始创建beanpay数据库")
	// 创建beanpay数据库
	if err := beanpay.CreateDB(xdb); err != nil {
		logger.Error(err)
		return
	}
	logger.Info("beanpay数据库创建完成")
	logger.Info("开始创建vds数据库")
	// 创建vds数据库
	err = vds.CreateDB(xdb)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("vds数据库创建完成")

	// logger.Info("开始创建lcs数据库")
	// // 创建lcs数据库
	// err = vds.CreateDB(xdb)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Info("lcs数据库创建完成")

	// logger.Info("开始创建dds数据库")
	// // 创建dds数据库
	// err = vds.CreateDB(xdb)
	// if err != nil {
	// 	logger.Error(err)
	// 	return
	// }
	// logger.Info("dds数据库创建完成")

	logger.Info("开始创建oms数据库")
	//创建oms数据库
	err = sql.CreateDB(xdb)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("oms数据库创建完成")
	logger.Info("数据表创建成功")
}
