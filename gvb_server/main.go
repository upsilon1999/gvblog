package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
)

func main() {
	// 读取配置文件
	core.InitConf()
	//连接数据库
	global.DB = core.InitGorm()
	fmt.Println(global.Config)
	fmt.Println(global.DB)
}
