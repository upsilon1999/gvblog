package routers

import (
	"gvb_server/api"
)

//获取siteInfo配置信息
func (router RouterGroup) SettingsRouter() {
  settingsApi := api.ApiGroupApp.SettingsApi
  settings := router.Group("settings")
  {
     //单一纯享版
     settings.GET("puresettings", settingsApi.SettingsInfoBaseView) //查询系统信息
     settings.PUT("puresettings", settingsApi.SettingsInfoBaseUpdateView) //修改系统信息

    //综合api版
    settings.GET("/:name", settingsApi.SettingsInfoView) //查询系统信息
    settings.PUT("/:name", settingsApi.SettingsInfoUpdateView) //修改系统信息
  }
 
}