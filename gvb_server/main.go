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

	//初始化日志
	global.Log = core.InitLogger()
	//测试全局日志
	// global.Log.Warnln("警告")
	// global.Log.Error("错误")
	// global.Log.Info("信息")

	//连接数据库，注册数据库实例
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
