package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) UpvoteRouter() {
	upvoteApi:= api.ApiGroupApp.UpvoteApi
	upvote := router.Group("upvote")
	{
	   //添加广告
	   upvote.POST("upvoteById",middleware.JwtAuth(), upvoteApi.UpvoteArticleView)
	}
   
  }