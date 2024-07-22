package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) AdvertRouter() {
	advertApi:= api.ApiGroupApp.AdvertApi
	advert := router.Group("advert")
	{
	   //添加广告
	   advert.POST("create",middleware.JwtAuth(), advertApi.AdvertCreateView)
	   //获取广告列表
	   advert.GET("list",middleware.JwtAuth(),advertApi.AdvertListView)
	   //修改广告
	   advert.PUT("update/:id",middleware.JwtAuth(),advertApi.AdvertUpdateView)
	   //删除广告
	   advert.DELETE("delete",middleware.JwtAuth(),advertApi.AdvertRemoveView)
	}
   
  }