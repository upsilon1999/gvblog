package routers

import "gvb_server/api"

//获取siteInfo配置信息
func (router RouterGroup) MenuRouter() {
	menuApi:= api.ApiGroupApp.MenuApi
	menu := router.Group("menu")
	{
	   //添加菜单
	   menu.POST("create", menuApi.MenuCreateView)
	}
   
  }