package flag

import sys_flag "flag"

type Option struct {
  DB bool
  /*
    预计形式 -u admin 就是admin用户，user就是普通用户
  */
  User string 
}

// Parse 解析命令行参数
func Parse() Option {
  db := sys_flag.Bool("db", false, "初始化数据库")
  user:= sys_flag.String("u","","创建用户")
  // 解析命令行参数写入注册的flag里
  sys_flag.Parse()
  return Option{
    DB: *db,
    User: *user,
  }
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
  //在创建数据库时停下
  if option.DB {
    return true
  }
  //在命令行创建用户时停下
  if option.User =="admin"||option.User =="user"{
    return true
  }
  return false
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

  //不符合预期走这里
  sys_flag.Usage()
}
