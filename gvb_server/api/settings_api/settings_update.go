package settings_api

import (
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

//系统信息更新
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context){
	var cr config.SiteInfo

	//解析json数据
	err:=c.ShouldBindJSON(&cr)
	if(err!=nil){
		//参数错误
		res.FailWithCode(res.ArgumentError,c)
		return
	}

	//测试
	// fmt.Println("before",global.Config)
	global.Config.SiteInfo = cr
	// fmt.Println("after",global.Config)

	//修改配置文件
	err = core.SetYmal()
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage(err.Error(),c)
	}

	res.OkWith(c)
}