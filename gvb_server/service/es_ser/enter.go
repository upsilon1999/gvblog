package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//获取es分页列表数据
func CommList(key string, page int, limit int)(list []models.ArticleModel,count int,err error){
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
	  boolSearch.Must(
		//查询title值为key的数据，由于我们的title设置为text类型，所以支持模糊查询
		elastic.NewMatchQuery("title", key),
	  )
	}
	if limit == 0 {
	  limit = 10
	}
	if from == 0 {
	  from = 1
	}
  
	//注意这里面的FetchSourceContext,我们在里面写了过滤文章内容字段的逻辑
	// res, err := global.ESClient.
    // Search(models.ArticleModel{}.Index()).
    // Query(boolSearch).
    // FetchSourceContext(elastic.NewFetchSourceContext(true).Exclude("content")).
    // From((from - 1) * limit).
    // Size(limit).
    // Do(context.Background())

	res, err := global.ESClient.
    Search(models.ArticleModel{}.Index()).
    Query(boolSearch).
    From((from - 1) * limit).
    Size(limit).
    Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
		return nil,0,err
	}

	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []models.ArticleModel{}
	for _,hit := range res.Hits.Hits{
		var model models.ArticleModel
		data,err := hit.Source.MarshalJSON()
		if err!=nil{
			logrus.Error(err.Error())
			continue
		}

		err = json.Unmarshal(data,&model)
		if err!=nil{
			logrus.Error(err)
		}
		model.ID = hit.Id 
		demoList = append(demoList, model)
	}
	fmt.Println(demoList,count)
	return demoList,count,err
}

//根据id获取es详情
func CommDetail(id string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Get().
	  Index(models.ArticleModel{}.Index()).
	  Id(id).
	  Do(context.Background())
	if err != nil {
	  return
	}
	err = json.Unmarshal(res.Source, &model)
	if err != nil {
	  return
	}
	model.ID = res.Id
	return
  }

//根据keyword，即文章标题获取文章详情
func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
	  Index(models.ArticleModel{}.Index()).
	  Query(elastic.NewTermQuery("keyword", key)).
	  Size(1).
	  Do(context.Background())
	if err != nil {
	  return
	}
	if res.Hits.TotalHits.Value == 0 {
	  return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]
  
	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
	  return
	}
	model.ID = hit.Id
	return
}
  
  