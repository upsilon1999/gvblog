package comment_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

func (CommentApi) CommentUpvoteView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("评论不存在", c)
		return
	}

	//记得将id转为string类型
	redis_ser.CommentUpvote(fmt.Sprintf("%d", cr.ID))

	res.OkWithMessage("评论点赞成功", c)

}