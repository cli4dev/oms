// +build !oracle

package sql

// GetReturnID 获取id
const GetReturnID = `REPLACE INTO oms_sequence (sequence_type,create_time) VALUES(1,NOW())`
