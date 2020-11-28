package utils

import (
	"encoding/json"
	"fmt"

	"github.com/micro-plat/lib4go/types"
)

// GetExtendParams 处理拓展信息
func GetExtendParams(source string, input map[string]interface{}) (string, error) {
	var info types.XMap
	if source != "" && source != "{}" {
		e, err := types.NewXMapByJSON(source)
		if err != nil {
			return "", fmt.Errorf("source:%v,json转map失败", source)
		}
		info = e
	}
	var bytes []byte
	var err error
	if info.IsEmpty() {
		bytes, err = json.Marshal(input)
		return string(bytes), err
	}
	info.MergeMap(input)
	bytes, err = json.Marshal(info)
	return string(bytes), nil
}

// GetExtendParamsByString 处理拓展信息
func GetExtendParamsByString(source string, input string) (string, error) {
	info := types.NewXMap()
	if source != "" && source != "{}" {
		e, err := types.NewXMapByJSON(source)
		if err != nil {
			return "", fmt.Errorf("source:%v,json转map失败", source)
		}
		info = e
	}
	if input != "" && input != "{}" {
		ext, err := types.NewXMapByJSON(input)
		if err != nil {
			return "", fmt.Errorf("input:%v,json转map失败", input)
		}
		info.MergeMap(ext)
	}
	if !info.IsEmpty() {
		bytes, err := json.Marshal(info)
		if err != nil {
			return "", fmt.Errorf("map转json失败")
		}
		return string(bytes), nil
	}
	return "", nil
}
