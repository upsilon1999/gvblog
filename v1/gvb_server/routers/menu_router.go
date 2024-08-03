package routers

import "gvb_server/api"

//获取siteInfo配置信息
func (router RouterGroup) MenuRouter() {
	menuApi:= api.ApiGroupApp.MenuApi
	menu := router.Group("menu")
	{
	   //添加菜单
	   menu.POST("create", menuApi.MenuCreateView)
	   //获取菜单列表
	   menu.GET("list", menuApi.MenuListView)
	   //菜单名称列表查询
	   menu.GET("nameList",menuApi.MenuNameList)
	   //更新菜单
	   menu.PUT("update/:id",menuApi.MenuUpdateView)
	   //菜单删除
	   menu.DELETE("remove",menuApi.MenuRemoveView)
	   //查看菜单详情
	   menu.GET("detail/:id",menuApi.MenuDetailView)
	}
   
  }