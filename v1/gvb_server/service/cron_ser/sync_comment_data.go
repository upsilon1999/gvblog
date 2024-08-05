package cron_ser

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"

	"gorm.io/gorm"
)

//同步评论点赞数
func SyncCommentData() {
	//1.获取redis中的数据
	commentUpvoteInfo := redis_ser.GetCommentUpvoteInfo()

	//2.遍历评论点赞信息,更新mysql数据
	for key,count := range commentUpvoteInfo{
		var comment models.CommentModel
		err := global.DB.Take(&comment,key).Error
		if err !=nil{
			global.Log.Error(err)
			continue
		}

		err = global.DB.Model(&comment).
			Update("upvote_count",gorm.Expr("upvote_count + ?",count)).Error
		
		if err !=nil{
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功，评论点赞数为%d\n",comment.Content,comment.UpvoteCount)
	}

	//3.清除评论点赞数据
	redis_ser.CommentUpvoteClear()
}