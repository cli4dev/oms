// +build !oracle

package sql

// GetNewID 获取id
const GetNewID = `REPLACE INTO oms_sequence (sequence_type,create_time) VALUES(1,NOW())`

// GetNewDeliveryID 获取新发货ID
const GetNewDeliveryID = `REPLACE INTO oms_sequence (sequence_type,create_time) VALUES(1,NOW())`
