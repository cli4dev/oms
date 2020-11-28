package response

import (
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
)

// RefundNotify 退款查询接入层对象
type RefundNotify struct {
	container component.IContainer
	// q         query.IQuery
}

// NewRefundNotify 构建对象
func NewRefundNotify(container component.IContainer) (u *RefundNotify) {
	return &RefundNotify{
		container: container,
		// q:         query.NewQuery(container),
	}
}

//Handle 接入处理
func (u *RefundNotify) Handle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("---------------退款通知接收--------------------")

	body, err := ctx.Request.GetBody()
	if err != nil {
		return err
	}
	ctx.Log.Infof("订单通知结果：%-v", body)

	return "success"
}
