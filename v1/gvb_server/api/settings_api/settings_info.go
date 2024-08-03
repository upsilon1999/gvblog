package settings_api

import (
	"gvb_server/global"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

/*
	单一纯享版
	查询 系统信息就用SettingsInfoBaseView
	查询 email配置就用 EmailBaseView
	查询 jwt配置就用 JwtBaseView
	...

	配套修改的纯享版
	修改 系统信息就用SettingsInfoBaseUpdateView
	修改 email配置就用 EmailBaseUpdateView
	修改 jwt配置就用 JwtBaseUpdateView
	...

	优势：可读性强
	缺点：需要写的接口太多,扩展性差

	如果多一项配置，需要多一项查询修改，以及多两个路由(一查一改)
*/
//系统信息查询
func (SettingsApi) SettingsInfoBaseView(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo,c)
	// res.FailWithCode(res.SettingsError,c)
}

/*
	综合API版，根据前端传递的名称来返回查询项
	配套修改API SettingsInfoUpdateView

	优点:易拓展
	缺点: 接口的入参和出参不统一
*/ 
// SettingsInfoView 显示某一项的配置信息
func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
	}
}