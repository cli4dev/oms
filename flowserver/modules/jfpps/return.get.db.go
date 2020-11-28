package jfpps

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/order"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/refund/refund"
)

//DBJFReturn 退货
type DBJFReturn struct {
	*refund.DBReturn
	c component.IContainer
}

//NewDBJFReturn 构建DBJFReturn
func NewDBJFReturn(c component.IContainer) *DBJFReturn {
	db := &DBJFReturn{c: c}
	db.DBReturn = refund.NewDBReturn(c)
	return db
}

//DoForAllSuccessByDB 重写退款完成成功函数
func (o *DBJFReturn) DoForAllSuccessByDB(dbTrans db.IDBTrans, data types.IXMap, returnSuccessAmountInfo types.XMap, input *order.DeliveryResult) ([]string, error) {
	tasks, err := o.DBReturn.DoForAllSuccessByDB(dbTrans, data, returnSuccessAmountInfo, input)
	if err != nil {
		return nil, err
	}
	tasks = tasks[1:]
	fmt.Println("tasks", tasks)

	return tasks, nil
}
