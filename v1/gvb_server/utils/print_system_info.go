package utils

import "gvb_server/global"

func PrintSysInfo() {
	ip := global.Config.System.Host
	port := global.Config.System.Port

	if ip == "0.0.0.0"{
		ipList := GetIpList()
		for _, v := range ipList {
			global.Log.Infof("gvb_server 运行在:http://%s:%d/api", v,port)
			global.Log.Infof("api文档运行在http://%s:%d/swagger/index.html#",v,port)
		}
	}else{
		global.Log.Infof("gvb_server 运行在:http://%s:%d/api", ip,port)
		global.Log.Infof("api文档运行在http://%s:%d/swagger/index.html#",ip,port)
	}
	
}