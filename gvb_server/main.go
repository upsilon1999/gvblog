package main

import (
	"gvb_server/core"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
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
	// fmt.Println(global.DB)


	//命令行参数绑定
	// go run main.go -db
	//如果在者停止了web服务，后面有关web的就不该执行
	option := flag.Parse()
	if flag.IsWebStop(option){
		flag.SwitchOption(option)
		return
	}

	



	//注册路由
	router := routers.InitRouter()
	// 根据system配置来设定监听目标
	addr:=global.Config.System.Addr()
	global.Log.Info("gvb_server正在监听:%s",addr)
	err :=router.Run(addr)
	if(err!=nil){
		global.Log.Fatalf(err.Error())
	} 
}
