package routers

import (
	"gvb_server/api"
)

//获取siteInfo配置信息
func (router RouterGroup) SettingsRouter() {
  settingsApi := api.ApiGroupApp.SettingsApi
  //单一纯享版
  router.GET("puresettings", settingsApi.SettingsInfoBaseView) //查询系统信息
  router.PUT("puresettings", settingsApi.SettingsInfoBaseUpdateView) //修改系统信息

  //综合api版
  router.GET("settings/:name", settingsApi.SettingsInfoView) //查询系统信息
  router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView) //修改系统信息
}