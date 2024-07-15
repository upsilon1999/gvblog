package tag_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description 标签列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/tag/list [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context){
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr);err !=nil{
		res.FailWithCode(res.ArgumentError,c)
		return
	}

	list,count,_ := common.ComList(models.TagModel{},common.Option{
		PageInfo:cr,
		//开启打印debug日志
		Debug: true,
	})
	res.OkWithList(list,count,c)
}