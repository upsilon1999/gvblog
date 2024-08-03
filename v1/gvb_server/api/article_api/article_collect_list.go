package article_api

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
)

type CollResponse struct {
	models.ArticleModel `json:"articleModel"`
	//收藏时间，这里有一点要明确
	//在表中的createdAt是记录创建时间,但是回给前端时为了和文章的createdAt做区分，所以改名为收藏时间
	CollectTime string `json:"collectTime"`
}

func (ArticleApi) ArticleCollListView(c *gin.Context) {

	var cr models.PageInfo

	c.ShouldBindQuery(&cr)

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var articleIDList []interface{}

	list, count, err := common.ComList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})
	// fmt.Printf("list的值为%v\n",list)
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("获取列表失败",c)
	}

	var collMap = map[string]string{}

	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		//从list中获取创建时间并格式化，未来将他赋予收藏时间
		collMap[model.ArticleID] = model.CreatedAt.Format("2006-01-02 15:04:05")
	}

	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)

	var collList = make([]CollResponse, 0)

	// 传id列表，查es
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(result.Hits.TotalHits.Value, articleIDList)

	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		collList = append(collList, CollResponse{
			ArticleModel: article,
			CollectTime:    collMap[hit.Id],
		})
	}
	res.OkWithList(filter.Omit("list",collList), count, c)
}