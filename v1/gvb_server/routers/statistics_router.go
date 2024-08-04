package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

//获取siteInfo配置信息
func (router RouterGroup) StatisticsRouter() {
	statisticsApi:= api.ApiGroupApp.StatisticsApi
	statistics := router.Group("statistics")
	{
	   //获取用户登录数据
	   statistics.GET("dataLogin",middleware.JwtAuth(),statisticsApi.SevenLogin)
	   //获取总数统计
	   statistics.GET("dataSum",middleware.JwtAuth(),statisticsApi.DataSumView)

	}
   
  }