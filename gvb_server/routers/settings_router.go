package routers

import (
	"gvb_server/api"
)

//获取siteInfo配置信息
func (router RouterGroup) SettingsRouter() {
  settingsApi := api.ApiGroupApp.SettingsApi
  router.GET("settings", settingsApi.SettingsInfoView) //查询系统信息
  router.PUT("settings", settingsApi.SettingsInfoUpdateView) //修改系统信息

}