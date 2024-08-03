package config

import "fmt"

type Es struct {
	Host     string `yaml:"host"`     //服务地址
	Port     int    `yaml:"port"`     //端口
	Username string `yaml:"username"` //数据库用户名
	Password string `yaml:"password"` //数据库密码
}

// Es连接配置，
func (es *Es) URL() string {
	return fmt.Sprintf("%s:%d",es.Host,es.Port)
}