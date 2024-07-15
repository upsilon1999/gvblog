package config

import "fmt"

//6.0版本前的redis没有用户名
type Redis struct {
	Ip       string `json:"ip" yaml:"ip"` //ip
	Port     int    `json:"port" yaml:"port"` //端口
	Password string `json:"password" yaml:"password"` //密码
	PoolSize int    `json:"poolSize" yaml:"pool_size"` //连接池大小
}

func (r Redis) Addr()string{
	return fmt.Sprintf("%s:%d",r.Ip,r.Port)
}