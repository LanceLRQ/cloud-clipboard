package conf

import "github.com/gookit/config/v2"

// ServerConfig HTTP服务端配置信息
var ServerConfig = ServerConfigStruct{}

// SecurityConfig 安全配置信息
var SecurityConfig = SecurityConfigStruct{}

func LoadServerConfig() error {
	err := config.BindStruct("server", &ServerConfig)
	if err != nil {
		return err
	}
	err = config.BindStruct("security", &SecurityConfig)
	if err != nil {
		return err
	}
	return nil
}

var ServerDebugging = false
