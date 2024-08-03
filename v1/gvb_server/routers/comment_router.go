package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) CommentRouter() {
	commentApi:= api.ApiGroupApp.CommentApi
	comment := router.Group("comment")
	{
	   //发布评论
	   comment.POST("create",middleware.JwtAuth(), commentApi.CommentCreateView)
	   //获取评论列表
	   comment.GET("list",middleware.JwtAuth(), commentApi.CommentListView)
	   //获取评论列表
	   comment.GET("upvote/:id",middleware.JwtAuth(), commentApi.CommentUpvoteView)
	   //删除评论
	   comment.DELETE("remove/:id",middleware.JwtAuth(), commentApi.CommentRemoveView)
	   
	}
   
  }