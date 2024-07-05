package routers

import "gvb_server/api"

//获取siteInfo配置信息
func (router RouterGroup) AdvertRouter() {
	advertApi:= api.ApiGroupApp.AdvertApi
	settings := router.Group("advert")
	{
	   //添加广告
	   settings.POST("create", advertApi.AdvertCreateView)
	}
   
  }