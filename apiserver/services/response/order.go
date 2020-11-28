package response

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// OrderNotify 退款查询接入层对象
type OrderNotify struct {
	container component.IContainer
	// q         query.IQuery
}

// NewOrderNotify 构建对象
func NewOrderNotify(container component.IContainer) (u *OrderNotify) {
	return &OrderNotify{
		container: container,
		// q:         query.NewQuery(container),
	}
}

//Handle 接入处理
func (u *OrderNotify) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------订单通知接收--------------------")

	body, err := ctx.Request.GetBody()
	if err != nil {
		return err
	}
	ctx.Log.Infof("订单通知结果：%-v", body)

	return "success"
}
