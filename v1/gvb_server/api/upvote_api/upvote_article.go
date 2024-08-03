package upvote_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
)

func (UpvoteApi) UpvoteArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	// 对长度校验
	// 查es
	redis_ser.Upvote(cr.ID)
	res.OkWithMessage("文章点赞成功", c)
}