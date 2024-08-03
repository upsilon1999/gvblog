package redis_es_test

import (
	"context"
	"encoding/json"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestSyncRedis(t *testing.T) {
	// 读取配置文件
	core.InitConf()
	// 初始化日志
	global.Log = core.InitLogger()

	global.Redis = core.ConnectRedis()
	core.EsConnect()

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	//从redis获取点赞数据
	upvoteInfo := redis_ser.GetUpvoteInfo()

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		//对于我们来说这里还有一个问题，
		//es中存的是upvote_count，而我们json映射是upvoteCount
		//所以读取时需要用map接收
		err = json.Unmarshal(hit.Source, &article)
		if err!=nil{
			logrus.Error(err)
			continue
		}

		//获取每个id的对应点赞数
		upvote := upvoteInfo[hit.Id]

		newUpvote := article.UpvoteCount + upvote
		if article.UpvoteCount == newUpvote {
			logrus.Info(article.Title, "点赞数无变化")
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"upvote_count": newUpvote,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Info(article.Title, "点赞数据同步成功， 点赞数", newUpvote)
	}
	redis_ser.UpvoteClear()
}