package routers

import (
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

//初始化路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	// 路由分组
	apiRouterGroup := router.Group("api")

	routerGroupApp := RouterGroup{apiRouterGroup}
	// 路由分层
	// 系统配置api
	routerGroupApp.SettingsRouter()
	routerGroupApp.ImagesRouter()
	routerGroupApp.AdvertRouter()
	return router
}