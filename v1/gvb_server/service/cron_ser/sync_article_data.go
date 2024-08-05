package cron_ser

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"

	"github.com/olivere/elastic/v7"
)

//同步文章数据到es
func SyncArticleData() {
	//1.查询es中的全部数据,为后面的数据更新做准备
	result,err := global.ESClient.
		Search(models.ArticleModel{}.
		Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())

	if err != nil {
		global.Log.Error(err)
		return
	}


	//2.拿到redis中的缓存数据，例如点赞数、浏览数、评论数
	upvoteInfo := redis_ser.GetUpvoteInfo()
	lookInfo := redis_ser.GetLookInfo()
	commentInfo := redis_ser.GetCommentInfo()

	//3.遍历es数据
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source,&article)
		if err!=nil{
			global.Log.Error()
			continue
		}

		/*
			这里有一个巧妙的设置，upvoteInfo是map
			如果拿不到就是零值，那么就是0
		*/
		upvote := upvoteInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]
		// 3.1.计算新的数据 旧数据加上缓存中的数据
		newUpvote := article.UpvoteCount + upvote
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment

		//3.2 判断新数据是否和缓存中旧数据一样
		if upvote == 0 && look==0&& comment==0{
			global.Log.Infof("%s 无变化",article.Title)
		}
		//3.3 更新es数据
		_,err = global.ESClient.Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"lookCount":newLook,
				"commentCount":newComment,
				"upvoteCount":newUpvote,
			}).Do(context.Background())
		
		if err!=nil{
			global.Log.Error(err)
			continue
		}
		global.Log.Infof("%s 更新成功 点赞数为:%d 评论数为%d 浏览量为%d\n",article.Title,newUpvote,newComment,newLook)
	}

	//4.清除redis中的数据
	redis_ser.CommentClear()
	redis_ser.LookClear()
	redis_ser.UpvoteClear()
}