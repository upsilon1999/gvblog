package comment_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `form:"articleId"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	rootCommentList := FindArticleCommentList(cr.ArticleID)
	res.OkWithData(filter.Select("c", rootCommentList), c)
}

func FindArticleCommentList(articleID string) (RootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&RootCommentList, "article_id = ? and parent_comment_id is null", articleID)

	// 获取评论点赞数
	upvoteInfo := redis_ser.GetCommentUpvoteInfo()

	// 遍历根评论，递归查根评论下的所有子评论
	for _, model := range RootCommentList {
		var subCommentList, newSubCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		//遍历子评论列表，并同步对应的点赞数
		for _, commentModel := range subCommentList {
		  upvote := upvoteInfo[fmt.Sprintf("%d", commentModel.ID)]
		  commentModel.UpvoteCount = commentModel.UpvoteCount + upvote
		  newSubCommentList = append(newSubCommentList, commentModel)
		}

		//父评论列表也要同步点赞数
		modelDigg := upvoteInfo[fmt.Sprintf("%d", model.ID)]
		model.UpvoteCount = model.UpvoteCount + modelDigg
		model.SubComments = newSubCommentList
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList)
	}
}