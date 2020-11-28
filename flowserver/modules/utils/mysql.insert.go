// +build !oracle

package utils

import (
	"github.com/micro-plat/lib4go/db"
)

// Insert 创建
func Insert(dbTrans db.IDBExecuter, getNewIDSQL string, createSQL string, param map[string]interface{}) (int64, int64, string, []interface{}, error) {
	return dbTrans.Executes(createSQL, param)
}
