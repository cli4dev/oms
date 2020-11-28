package utils

import (
	"github.com/micro-plat/hydra/component"
)

// CallBackFuncs 调用回调函数
func CallBackFuncs(container component.IContainer, f ...func(c interface{}) error) error {
	for _, v := range f {
		if v == nil {
			continue
		}
		if err := v(container); err != nil {
			return err
		}
	}
	return nil
}
