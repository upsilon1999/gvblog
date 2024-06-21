package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 执行读取配置文件的操作
	core.InitConf()
	//查看配置文件全局变量
	// fmt.Println(global.Config)


	//连接数据库，注册数据库实例
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
