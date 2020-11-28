package jfpps

import (
	"fmt"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"gitlab.100bm.cn/micro-plat/oms/apiserver/modules/jfpps"
)

// FDHandler 使用记账结构体
type FDHandler struct {
	c  component.IContainer
	jf *jfpps.JFFd
}

// NewFDHandler 构建Handler
func NewFDHandler(c component.IContainer) *FDHandler {
	return &FDHandler{
		c:  c,
		jf: jfpps.NewJFFd(c),
	}
}

//UseHandle 使用积分记账请求
func (o *FDHandler) UseHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------使用积分记账请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &jfpps.JFFDUseRequest{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求记账")
	fmt.Println("input", input)
	res, err := o.jf.UseFdRequest(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return res
}

//VoidHandle 作废积分记账请求
func (o *FDHandler) VoidHandle(ctx *context.Context) (r interface{}) {
	ctx.Log.Info("-------------作废积分记账请求----------------")

	ctx.Log.Info("1.参数检验")
	input := &jfpps.JFFDVoidRequest{}
	if err := ctx.Request.Bind(input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.请求记账")
	res, err := o.jf.VoidFdRequest(input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.返回数据")
	return res
}
