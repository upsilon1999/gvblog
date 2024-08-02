package routers

import (
	"gvb_server/api"
)

//获取siteInfo配置信息
func (router RouterGroup) ChatRouter() {
	chatApi:= api.ApiGroupApp.ChatApi
	chat := router.Group("chat")
	{
	   //测试webscoket
	   chat.GET("baseConnect", chatApi.ChatBaseView)
	   //群聊接口，头像昵称随机生成
	   chat.GET("randGroup",chatApi.ChatGroupRandView)
	   //获取聊天记录
	   chat.GET("groupRecords",chatApi.ChatRandListView)
	}
   
  }