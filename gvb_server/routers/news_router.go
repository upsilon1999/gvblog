package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) NewsRouter() {
	newsApi:= api.ApiGroupApp.NewsApi
	news := router.Group("news")
	{
	   //获取广告列表
	   news.GET("list",middleware.JwtAuth(),newsApi.NewListGetView)
	}
   
  }