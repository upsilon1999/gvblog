package tag_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// TagCreateView 添加标签
// @Tags 标签管理
// @Summary 创建标签
// @Description 创建标签
// @Param data body TagRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/tag/create [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagCreateView(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err!= nil{
		global.Log.Error(err)
		res.FailWithError(err,&cr,c)
		return
	}

	// 重复的判断
	var tag models.TagModel
	count := global.DB.Take(&tag, "title = ?", cr.Title).RowsAffected
	if count != 0 {
	  res.FailWithMessage("该标签已存在", c)
	  return
	}

	//添加标签入数据库
	err =global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败",c)
		return
	}

	res.OkWithMessage("添加标签成功",c)
}