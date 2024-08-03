package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) MsgRouter() {
	msgApi:= api.ApiGroupApp.MessageApi
	msg := router.Group("msg")
	{
	   //发送消息
	   msg.POST("create",middleware.JwtAuth(), msgApi.MessageCreateView)
	   //管理员查看所有消息记录
	   msg.GET("allList",middleware.JwtAuth(),msgApi.MessageListAllView)
	   //用户查看自己的消息记录
	   msg.GET("list",middleware.JwtAuth(),msgApi.MessageListView)
	   //查看消息详情
	   msg.GET("record",middleware.JwtAuth(),msgApi.MessageRecordView)
	   
	}
   
  }