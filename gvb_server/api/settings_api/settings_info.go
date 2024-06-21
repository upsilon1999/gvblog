package settings_api

import "github.com/gin-gonic/gin"

type SettingsApi struct{

}

//视图函数
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "xxx",
	})
}