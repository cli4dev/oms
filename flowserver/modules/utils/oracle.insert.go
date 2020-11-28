// +build oracle

package utils

import (
	"fmt"

	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
)

// Insert 创建
func Insert(dbTrans db.IDBExecuter, getNewIDSQL string, createSQL string, param map[string]interface{}) (int64, int64, string, []interface{}, error) {
	// 1.获取订单号
	var id int64
	if getNewIDSQL != "" {
		newId, sqlStr, args, err := dbTrans.Scalar(getNewIDSQL, map[string]interface{}{})
		if err != nil || types.GetInt64(newId) == 0 {
			return 0, 0, sqlStr, args, err
		}
		id = types.GetInt64(newId)
	}

	param["id"] = id
	row, sqlStr, args, err := dbTrans.Execute(createSQL, param)
	if err != nil || row != 1 {
		return 0, 0, sqlStr, args, fmt.Errorf("执行sql异常count:%d,err:%v", row, err)
	}
	return id, row, sqlStr, args, err
}
