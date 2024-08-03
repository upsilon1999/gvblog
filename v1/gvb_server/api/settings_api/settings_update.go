package settings_api

import (
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

//单一纯享版
//系统信息更新
func (SettingsApi) SettingsInfoBaseUpdateView(c *gin.Context){
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
	err = core.SetYaml()
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage(err.Error(),c)
	}

	res.OkWith(c)
}

//综合api版

// SettingsInfoUpdateView 修改某一项的配置信息
func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	switch cr.Name {
	case "site":
	  var info config.SiteInfo
	  err = c.ShouldBindJSON(&info)
	  if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	  }
	  global.Config.SiteInfo = info
  
	case "email":
	  var info config.Email
	  err = c.ShouldBindJSON(&info)
	  if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	  }
	  global.Config.Email = info
	case "qq":
	  var info config.QQ
	  err = c.ShouldBindJSON(&info)
	  if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	  }
	  global.Config.QQ = info
	case "qiniu":
	  var info config.QiNiu
	  err = c.ShouldBindJSON(&info)
	  if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	  }
	  global.Config.QiNiu = info
	case "jwt":
	  var info config.Jwt
	  err = c.ShouldBindJSON(&info)
	  if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	  }
	  global.Config.Jwt = info
	default:
	  res.FailWithMessage("没有对应的配置信息", c)
	  return
	}
  
	err = core.SetYaml()
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage(err.Error(),c)
	}

	res.OkWith(c)
  }
  