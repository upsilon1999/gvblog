package statistics_api

import (
	"context"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataSumResponse struct {
	//用户数
	UserCount      int `json:"userCount"`
	//文章总数
	ArticleCount   int `json:"articleCount"`
	//消息总数
	MessageCount   int `json:"messageCount"`
	//群聊总数
	ChatGroupCount int `json:"chatGroupCount"`
	//今日登录用户数
	NowLoginCount  int `json:"nowLoginCount"`
	//今日注册用户数
	NowSignCount   int `json:"nowSignCount"`
}

func (StatisticsApi) DataSumView(c *gin.Context) {

	var userCount, articleCount, messageCount, ChatGroupCount int
	var nowLoginCount, nowSignCount int

	//查询文章总数
	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	articleCount = int(result.Hits.TotalHits.Value) //搜索到结果总条数


	global.DB.Model(models.UserModel{}).Select("count(id)").Scan(&userCount)
	global.DB.Model(models.MessageModel{}).Select("count(id)").Scan(&messageCount)
	global.DB.Model(models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&ChatGroupCount)
	global.DB.Model(models.LoginDataModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowLoginCount)
	global.DB.Model(models.UserModel{}).Where("to_days(created_at)=to_days(now())").
		Select("count(id)").Scan(&nowSignCount)

	res.OkWithData(DataSumResponse{
		UserCount:      userCount,
		ArticleCount:   articleCount,
		MessageCount:   messageCount,
		ChatGroupCount: ChatGroupCount,
		NowLoginCount:  nowLoginCount,
		NowSignCount:   nowSignCount,
	}, c)
}