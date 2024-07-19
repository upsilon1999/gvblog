package flag

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

type Option struct {
  DB bool
  /*
    预计形式 -u admin 就是admin用户，user就是普通用户
  */
  User string 
  /*
    -es create 创建索引
    -es delete 删除索引
  */
  Es string
}

// Parse 解析命令行参数
func Parse() Option {
  db := sys_flag.Bool("db", false, "初始化数据库")
  user:= sys_flag.String("u","","创建用户")
  es := sys_flag.String("es","","es操作")
  // 解析命令行参数写入注册的flag里
  sys_flag.Parse()
  return Option{
    DB: *db,
    User: *user,
    Es:*es,
  }
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
  maps := structs.Map(&option)
  for _, v := range maps {
    switch val := v.(type) {
    case string:
      if val != "" {
        f = true
      }
    case bool:
      if val == true {
        f = true
      }
    }
  }
  return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
  if option.DB {
    Makemigrations()
    return
  }

  if option.User =="admin"||option.User =="user" {
    CreateUser(option.User)
    return
  }

  //在这里解析值，例如 -es create
  //实际上键是es 值是es后面跟的内容
  if option.Es == "create"{
    EsCreateIndex()
    return
  }

  //不要忘记前面的return
  //不符合预期走这里
  sys_flag.Usage()
}
