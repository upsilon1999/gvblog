package routers

import (
	"gvb_server/api"
)

//获取siteInfo配置信息
func (router RouterGroup) ChatRouter() {
	chatApi:= api.ApiGroupApp.ChatApi
	chat := router.Group("chat")
	{
	   //添加广告
	   chat.GET("connect", chatApi.ChatGroupView)
	}
   
  }