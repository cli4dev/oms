package queue

import "fmt"

//QueueShelTag 队列对应的伙计tag
var QueueShelTag = &struct {
	JFTag string
}{"jfpps"}

// Queue 消息队列
type Queue struct {
	Name  string
	Route string
}

// GetName 获取消息队列
func (q *Queue) GetName(platName string) string {
	if platName == "" {
		return q.Name
	}
	return fmt.Sprintf("%s:%s", platName, q.Name)
}

// GetRoute 获取路由
func (q *Queue) GetRoute() string {
	return q.Route
}

//GetOrderUpPayQueue 订单上游支付队列获取
func GetOrderUpPayQueue(shelfTag string) *Queue {
	switch shelfTag {
	case QueueShelTag.JFTag:
		return JFGetUpPay
	default:
		return OrderUpPay
	}
}

//GetFinishReturnQueue 退货完成处理队列
func GetFinishReturnQueue(shelfTag string) *Queue {
	switch shelfTag {
	case QueueShelTag.JFTag:
		return JFGetFinishUpReturn
	default:
		return FinishUpReturn
	}
}

//GetUpRefundQueue 获取上游退款队列
func GetUpRefundQueue(shelfTag string) *Queue {
	switch shelfTag {
	case QueueShelTag.JFTag:
		return JFGetUpRefund
	default:
		return UpRefund
	}
}
