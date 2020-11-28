package task

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/lib4go/db"
	"github.com/micro-plat/lib4go/types"
	"github.com/micro-plat/qtask/qtask"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/const/queue"
	"gitlab.100bm.cn/micro-plat/oms/flowserver/modules/utils"
)

//TaskType 任务处理分类
var TaskType = &struct {
	DownPayTask           string
	OrderOverTimeTask     string
	BindTask              string
	DeliveryTask          string
	UpPaymentTask         string
	NotifyTask            string
	ReturnTask            string
	RefundOverTimeTask    string
	DownRefundPaymentTask string
	UpRefundPaymentTask   string
	RefundNotifyTask      string
	OrderFailRefundTask   string
	DeliveryUnknownTask   string
	ReturnUnknownTask     string
	PointUseRefundFDTask  string
	InvoiceTask           string
	InvoiceNotifyTask     string
	InvoiceRedNotifyTask  string
	InvoicePayTask        string
	InvoiceRefundTask     string
}{"down_pay", "order_over_time", "bind", "delivery", "up_pay", "notify", "return", "refund_over_time", "down_refund_pay", "up_refund_pay",
	"refud_notify", "order_fail_refund", "delivery_unknown", "refund_unknown", "point_use_refund_fd", "invoice", "invoice_notify", "invoice_red_notify", "invoice_pay", "invoice_refund"}

// TaskRemark 任务说明
var TaskRemark = &struct {
	DownPayTask           string
	OrderOverTimeTask     string
	BindTask              string
	DeliveryTask          string
	UpPaymentTask         string
	NotifyTask            string
	ReturnTask            string
	RefundOverTimeTask    string
	RefundNotifyTask      string
	RefundTask            string
	DeliveryOvertimeTask  string
	UpRefundPaymentTask   string
	DownRefundPaymentTask string
	RefundUnknownTask     string
	PointUseRefundFDTask  string
	InvoiceTask           string
	InvoiceNotifyTask     string
	InvoiceRedNotifyTask  string
	InvoicePayTask        string
	InvoiceRefundTask     string
}{"下游支付任务", "订单超时处理任务", "绑定任务", "发货流程任务", "上游支付任务", "通知任务", "退货任务", "退款超时处理任务", "退款通知任务", "订单退款任务", "发货超时任务", "上游退款任务",
	"下游退款任务", "退款未知任务", "积分使用退款记账任务", "开票流程任务", "开票通知任务", "开票冲红通知任务", "开票记账任务", "冲红退款任务"}

// TaskIntervalTime 任务间隔时间
const TaskIntervalTime = 300

//QTask 任务处理对象
type QTask struct {
	c component.IContainer
}

//NewQTask 初始化任务创建对象
func NewQTask(c component.IContainer) *QTask {
	return &QTask{
		c: c,
	}
}

//IQTask 任务处理接口
type IQTask interface {
	CreateTask(trans db.IDBTrans, tasks []string, info types.XMap) (funcs []func(c interface{}) error, err error)
	CreateBatchTasks(trans db.IDBTrans, tasks []string, list types.XMaps) (funcs []func(c interface{}) error, err error)
	AppendTasks(soureTasks []string, newTasks ...string) []string
}

//CreateTask 创建任务
func (q *QTask) CreateTask(trans db.IDBTrans, tasks []string, info types.XMap) (funcs []func(c interface{}) error, err error) {
	if tasks == nil {
		return
	}

	for _, v := range tasks {
		switch v {
		case TaskType.DownPayTask: //订单下游支付任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.DownPayTask,
				map[string]interface{}{"order_id": info.GetInt64("order_id")},
				TaskIntervalTime,
				queue.OrderDownPay.GetName(q.c.GetPlatName()),
				qtask.WithDeadline(info.GetInt("flow_overtime")),
			)
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.OrderOverTimeTask: //订单超时处理任务创建
			fmt.Println("order_overtime", info)
			fmt.Println("order_overtime", info.GetInt("order_overtime"))
			_, err = qtask.Delay(trans,
				TaskRemark.OrderOverTimeTask,
				map[string]interface{}{"order_id": info.GetInt64("order_id")},
				info.GetInt("order_overtime")+1,
				TaskIntervalTime,
				queue.OrderOvertime.GetName(q.c.GetPlatName()),
			)
			if err != nil {
				return nil, err
			}

		case TaskType.BindTask: //订单绑定任务创建
			_, f, err := qtask.Create(trans, TaskRemark.BindTask, map[string]interface{}{"order_id": info.GetInt64("order_id")},
				30, queue.OrderBind.GetName(q.c.GetPlatName()), qtask.WithDeadline(info.GetInt("flow_overtime")))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.DeliveryTask: //订单发货任务创建
			_, f, err := qtask.Create(trans, TaskRemark.DeliveryTask,
				map[string]interface{}{"delivery_id": info.GetInt64("delivery_id"), "order_id": info.GetInt64("order_id")},
				120, queue.OrderStartDelivery.GetName(q.c.GetPlatName()),
				qtask.WithDeadline(utils.GetDeliveryOvertime(info.GetInt("delivery_overtime"), info.GetInt("flow_overtime"))))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.UpPaymentTask: //上游支付任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.UpPaymentTask,
				map[string]interface{}{"delivery_id": info.GetInt64("delivery_id"), "order_id": info.GetInt64("order_id")},
				TaskIntervalTime,
				queue.GetOrderUpPayQueue(info.GetString("pre_tag")).GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.NotifyTask: //通知任务创建
			_, f, err := qtask.Create(trans, TaskRemark.NotifyTask,
				map[string]interface{}{"notify_id": info.GetInt64("notify_id"), "order_id": info.GetInt64("order_id")},
				TaskIntervalTime,
				queue.OrderNotify.GetName(q.c.GetPlatName()),
				qtask.WithDeadline(info.GetInt("flow_overtime")))

			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.ReturnTask: //退货任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.ReturnTask,
				map[string]interface{}{"return_id": info.GetString("return_id")},
				TaskIntervalTime,
				queue.StartUpReturn.GetName(q.c.GetPlatName()),
				qtask.WithDeadline(info.GetInt("return_overtime")))

			if err != nil {
				return nil, fmt.Errorf("创建退货任务失败,err:%+v", err)
			}
			funcs = append(funcs, f)

		case TaskType.RefundOverTimeTask: //退款超时处理任务创建
			_, err = qtask.Delay(trans,
				TaskRemark.RefundOverTimeTask,
				map[string]interface{}{"refund_id": info.GetString("refund_id")},
				info.GetInt("refund_flow_overtime")+1,
				TaskIntervalTime,
				queue.RefundOvertime.GetName(q.c.GetPlatName()),
			)
			if err != nil {
				return nil, fmt.Errorf("创建退款超时处理任务失败,err:%+v", err)
			}
		case TaskType.DownRefundPaymentTask: //下游退款任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.DownRefundPaymentTask,
				map[string]interface{}{"refund_id": info.GetString("refund_id")},
				TaskIntervalTime,
				queue.DownRefund.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		case TaskType.UpRefundPaymentTask: //上游退款任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.UpRefundPaymentTask,
				map[string]interface{}{"return_id": info.GetString("return_id")},
				TaskIntervalTime,
				queue.GetUpRefundQueue(info.GetString("pre_tag")).GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.RefundNotifyTask: //退款通知任务创建
			fmt.Printf("\ninfo:%+v", info)
			_, f, err := qtask.Create(trans,
				TaskRemark.RefundNotifyTask,
				map[string]interface{}{"notify_id": info.GetString("notify_id")},
				TaskIntervalTime,
				queue.RefundNotify.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.OrderFailRefundTask: //订单失败退款任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.RefundTask,
				map[string]interface{}{"order_id": info.GetInt64("order_id")},
				TaskIntervalTime,
				queue.OrderFailedRefund.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.DeliveryUnknownTask: //订单超时发货结果未知的处理任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.DeliveryOvertimeTask,
				map[string]interface{}{"order_id": info.GetInt64("order_id"), "delivery_id": info.GetInt64("delivery_id")},
				TaskIntervalTime,
				queue.OrderDeliveryOvertime.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.ReturnUnknownTask: //退款超时退货结果未知的处理任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.RefundUnknownTask,
				map[string]interface{}{
					"return_id": info.GetString("return_id"),
					"refund_id": info.GetInt64("refund_id"),
					"order_id":  info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.ReturnOvertime.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.PointUseRefundFDTask: //积分使用退款记账任务创建
			_, f, err := qtask.Create(trans,
				TaskRemark.PointUseRefundFDTask,
				map[string]interface{}{"order_id": info.GetString("order_id"),
					"refund_id": info.GetString("refund_id"),
				},
				TaskIntervalTime,
				queue.PointUseRefundFD.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)

		case TaskType.InvoiceTask: //开始开票任务
			fmt.Printf("info:%+v", info)
			_, f, err := qtask.Create(trans,
				TaskRemark.InvoiceTask,
				map[string]interface{}{
					"invoice_id": info.GetInt64("invoice_id"),
					"order_id":   info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.InvoiceStart.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		case TaskType.InvoiceNotifyTask: //开票通知任务
			_, f, err := qtask.Create(trans,
				TaskRemark.InvoiceNotifyTask,
				map[string]interface{}{
					"invoice_id": info.GetString("invoice_id"),
					"order_id":   info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.InvoiceNotify.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		case TaskType.InvoiceRedNotifyTask: //开票冲红通知任务
			_, f, err := qtask.Create(trans,
				TaskRemark.InvoiceRedNotifyTask,
				map[string]interface{}{
					"invoice_id": info.GetString("invoice_id"),
					"order_id":   info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.InvoiceRedNotify.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		case TaskType.InvoicePayTask: //开票记账任务
			_, f, err := qtask.Create(trans,
				TaskRemark.InvoicePayTask,
				map[string]interface{}{
					"invoice_id": info.GetString("invoice_id"),
					"order_id":   info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.InvoicePay.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		case TaskType.InvoiceRefundTask: //冲红退款任务
			_, f, err := qtask.Create(trans,
				TaskRemark.InvoiceRefundTask,
				map[string]interface{}{
					"invoice_id": info.GetString("invoice_id"),
					"order_id":   info.GetInt64("order_id"),
				},
				TaskIntervalTime,
				queue.InvoiceRedFefund.GetName(q.c.GetPlatName()))
			if err != nil {
				return nil, err
			}
			funcs = append(funcs, f)
		default:
			return nil, fmt.Errorf("请求创建的任务类型不存在")
		}
	}

	return
}

//CreateBatchTasks 创建一批次任务
func (q *QTask) CreateBatchTasks(trans db.IDBTrans, tasks []string, list types.XMaps) (funcs []func(c interface{}) error, err error) {
	if tasks == nil {
		return
	}

	for _, v := range list {
		fs, err := q.CreateTask(trans, tasks, v)
		if err != nil {
			return nil, err
		}
		funcs = append(funcs, fs...)
	}

	return funcs, nil
}

//AppendTasks 合并任务
func (q *QTask) AppendTasks(soureTasks []string, newTasks ...string) []string {
	if newTasks == nil {
		return soureTasks
	}

	for _, v := range newTasks {
		if v != "" {
			soureTasks = append(soureTasks, v)
		}
	}

	return soureTasks
}
