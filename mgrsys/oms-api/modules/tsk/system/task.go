package system

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/const/sql"
)

//QueryTask 查询任务表
type QueryTask struct {
	CreateTime string `json:"create_time" form:"create_time" m2s:"create_time"` //CreateTime 创建时间
	Name       string `json:"name" form:"name" m2s:"name"`                      //Name 名称
	QueueName  string `json:"queue_name" form:"queue_name" m2s:"queue_name"`    //QueueName 消息队列
	Status     string `json:"status" form:"status" m2s:"status"`                //Status 状态

	Pi string `json:"pi" form:"pi" m2s:"pi" valid:"required"`
	Ps string `json:"ps" form:"ps" m2s:"ps" valid:"required"`
}

//IDbTask  任务表接口
type IDbTask interface {
	//Get 单条查询
	Get(taskId string) (db.QueryRow, error)
	//Query 列表查询
	Query(input *QueryTask) (data db.QueryRows, count int, err error)
}

//DbTask 任务表对象
type DbTask struct {
	c component.IContainer
}

//NewDbTask 创建任务表对象
func NewDbTask(c component.IContainer) *DbTask {
	return &DbTask{
		c: c,
	}
}

//Get 查询单条数据任务表
func (d *DbTask) Get(taskId string) (db.QueryRow, error) {

	db := d.c.GetRegularDB()
	data, q, a, err := db.Query(sql.GetTskSystemTask, map[string]interface{}{
		"task_id": taskId,
	})
	if err != nil {
		return nil, fmt.Errorf("获取任务表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data.Get(0), nil
}

//Query 获取任务表列表
func (d *DbTask) Query(input *QueryTask) (data db.QueryRows, count int, err error) {

	db := d.c.GetRegularDB()
	c, q, a, err := db.Scalar(sql.QueryTskSystemTaskCount, map[string]interface{}{
		"create_time": input.CreateTime,
		"name":        input.Name,
		"queue_name":  input.QueueName,
		"status":      input.Status,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取任务表列表条数发生错误(err:%v),sql:%s,输入参数:%v", err, q, a)
	}

	data, q, a, err = db.Query(sql.QueryTskSystemTask, map[string]interface{}{
		"create_time": input.CreateTime,
		"name":        input.Name,
		"queue_name":  input.QueueName,
		"status":      input.Status,
		"pi":          input.Pi,
		"ps":          input.Ps,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("获取任务表数据表发生错误(err:%v),sql:%s,输入参数:%v,", err, q, a)
	}

	return data, types.GetInt(c, 0), nil
}
