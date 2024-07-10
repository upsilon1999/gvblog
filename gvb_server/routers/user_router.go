package routers

import "gvb_server/api"

//获取siteInfo配置信息
func (router RouterGroup) UserRouter() {
	UserApi:= api.ApiGroupApp.UserApi
	user := router.Group("user")
	{
	   //邮箱或用户名登录
	   user.POST("emailLogin", UserApi.EmailLoginView)
	   //获取用户列表
	   user.GET("list",UserApi.UserListView)
	}
   
  }