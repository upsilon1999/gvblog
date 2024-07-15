package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) TagRouter() {
	tagApi:= api.ApiGroupApp.TagApi
	tag := router.Group("tag")
	{
	   //添加广告
	   tag.POST("create",middleware.JwtAuth(), tagApi.TagCreateView)
	   //获取广告列表
	   tag.GET("list",middleware.JwtAuth(),tagApi.TagListView)
	   //修改广告
	   tag.PUT("update/:id",middleware.JwtAuth(),tagApi.TagUpdateView)
	   //删除广告
	   tag.DELETE("delete",middleware.JwtAuth(),tagApi.TagRemoveView)
	}
   
  }