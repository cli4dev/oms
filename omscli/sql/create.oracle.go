// +build oracle

package sql

import (
	"github.com/micro-plat/lib4go/db"
)

func CreateDB(xdb db.IDB) error {
	return db.CreateDB(xdb, "src/gitlab.100bm.cn/micro-plat/oms/omscli/sql/oracle")
}
