package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/micro-plat/lib4go/net/http"
)

// HTTPRequest http请求
func HTTPRequest(info map[string]interface{}, method string, url string) (content string, status int, err error) {
	bt, err := json.Marshal(info)
	if err != nil {
		return "", 0, fmt.Errorf("map转json失败err:%v", err)
	}
	connTimeout := http.WithConnTimeout(30 * time.Second)
	reqRimeout := http.WithRequestTimeout(30 * time.Second)
	client, err := http.NewHTTPClient(connTimeout, reqRimeout)
	if err != nil {
		return "", 0, fmt.Errorf("构建http对象失败err:%v", err)
	}
	content, status, err = client.Request(method, url, string(bt), "utf-8", map[string][]string{"Content-Type": {"application/json"}})
	return
}
