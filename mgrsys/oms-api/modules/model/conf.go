package model

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/sso/sdk/sso"
)

//Conf 应用程序配置
type Conf struct {
	Secret     string `json:"secret" valid:"ascii,required"`
	SsoApiHost string `json:"sso_api_host" valid:"ascii,required"`
	Ident      string `json:"ident"`
	PlatName   string `json:"plat_name"`
	Type       string `json:"type"`
	Group      string `json:"group"`
	ServerName string `json:"server_name"`
}

//Valid 验证配置参数是否合法
func (c Conf) Valid() error {
	if b, err := govalidator.ValidateStruct(&c); !b {
		return fmt.Errorf("app 配置文件有误:%v", err)
	}
	return nil
}

//SaveConf 保存当前应用程序配置
func SaveConf(c component.IContainer, m *Conf) {
	c.Set("__AppConf__", m)
}

//GetConf 获取当前应用程序配置
func GetConf(c component.IContainer) *Conf {
	return c.Get("__AppConf__").(*Conf)
}

//SaveSSOClient  保存sso client
func SaveSSOClient(c component.IContainer, m *sso.Client) {
	c.Set("__SsoClient__", m)
}

//GetSSOClient  获取sso client
func GetSSOClient(c component.IContainer) *sso.Client {
	return c.Get("__SsoClient__").(*sso.Client)
}
