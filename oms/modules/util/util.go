package util

import (
	"fmt"
	"gitlab.100bm.cn/micro-plat/oms/oms/modules/const/confs"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/db"
)

//SDKDB sdk数据库操作对象
type SDKDB struct {
	DB        db.IDBExecuter
	Container context.IContainer
	IsTrans   bool
	DBTrans   db.IDBTrans
}

//NewSDKDB 初始化对象
func NewSDKDB(c interface{}, data ...interface{}) (*SDKDB, error) {
	switch v := c.(type) {
	case *context.Context:
		i := v.GetContainer()
		db, err := i.GetDB(confs.DBName)
		return &SDKDB{DB: db, Container: i, IsTrans: false}, err
	case component.IContainer:
		db, err := v.GetDB(confs.DBName)
		return &SDKDB{DB: db, Container: v, IsTrans: false}, err
	case db.IDB:
		i, err := getContainer(data)
		return &SDKDB{DB: v, Container: i, IsTrans: false}, err
	case db.IDBTrans:
		i, err := getContainer(data)
		return &SDKDB{DB: v, Container: i, DBTrans: v.(db.IDBTrans), IsTrans: true}, err
	default:
		return nil, fmt.Errorf("不支持的参数类型")
	}
}

//TransBegin 开启事物
func (d *SDKDB) TransBegin() error {
	t, err := d.DB.(db.IDB).Begin()
	if err != nil {
		return err
	}
	d.DBTrans = t

	return nil
}

//TransClose 关闭对象
func (d *SDKDB) TransClose(err error) {
	if err != nil {
		d.DBTrans.Rollback()
		return
	}

	d.DBTrans.Commit()
	return
}

func getContainer(data ...interface{}) (context.IContainer, error) {
	if len(data) <= 0 {
		return nil, fmt.Errorf("入参是数据库链接时,不定参数不能为空")
	}

	switch v := data[0].(type) {
	case *context.Context:
		i := v.GetContainer()
		return i, nil
	case component.IContainer:
		return v, nil
	default:
		return nil, fmt.Errorf("不定参数入参是不支持的参数类型")
	}
}
