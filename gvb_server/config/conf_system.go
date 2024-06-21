package config

import "fmt"

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

//拼接程序运行的IP与端口，供监听使用
func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
