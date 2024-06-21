package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`      //服务地址
	Port     int    `yaml:"port"`      //端口
	DB       string `yaml:"db"`        //数据库名
	Username string `yaml:"username"`  //数据库用户名
	Password string `yaml:"password"`  //数据库密码
	Config   string `yaml:"config"`      //高级配置,例如charset
	LogLevel string `yaml:"log_level"` // 日志等级，debug就是输出全部sql，dev, release
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
