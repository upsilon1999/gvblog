package settings_api

import (
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type SettingsApi struct{

}

//视图函数
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.Ok(map[string]string{},"xxx",c)
	// res.FailWithCode(res.SettingsError,c)
}