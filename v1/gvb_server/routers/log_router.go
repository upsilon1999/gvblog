package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) LogRouter() {
	logApi:= api.ApiGroupApp.LogApi
	log := router.Group("log")
	{
	   //获取日志列表列表
	   log.GET("list",middleware.JwtAuth(),logApi.LogListView)
	   //日志删除
	   log.DELETE("remove",middleware.JwtAuth(),logApi.LogRemoveListView)

	}
   
  }