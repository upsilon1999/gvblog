package routers

import "gvb_server/api"

//获取siteInfo配置信息
func (router RouterGroup) AdvertRouter() {
	advertApi:= api.ApiGroupApp.AdvertApi
	advert := router.Group("advert")
	{
	   //添加广告
	   advert.POST("create", advertApi.AdvertCreateView)
	   //获取广告列表
	   advert.GET("list",advertApi.AdvertListView)
	   //修改广告
	   advert.PUT("update/:id",advertApi.AdvertUpdateView)

	}
   
  }