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

type Option struct{
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
	Fields []string 
	Tag string `form:"tag"`
} 
func (op Option)GetFrom()int{
	if op.Limit == 0 {
		op.Limit = 10
	}
	if op.Page == 0 {
		op.Page = 1
	}
	return (op.Page - 1)*op.Limit
}



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
  


//获取分页并高亮标题
func CommHighLightList(option Option)(list []models.ArticleModel,count int,err error){
	boolSearch := elastic.NewBoolQuery()

	if option.Key != "" {
	  boolSearch.Must(
		//构造多字段查询
		elastic.NewMultiMatchQuery(option.Key, option.Fields...),
	  )
	}

	//标签搜索
	if option.Tag != "" {
		boolSearch.Must(
			//构造多字段查询
		  elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	  }
	
	if option.Limit == 0 {
		option.Limit = 10
	}

	res, err := global.ESClient.
    Search(models.ArticleModel{}.Index()).
    Query(boolSearch).
	Highlight(elastic.NewHighlight().Field("title")).
    From(option.GetFrom()).
    Size(option.Limit).
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
			continue
		}

		
		if title, ok := hit.Highlight["title"];ok {
			model.Title = title[0]
		}

		model.ID = hit.Id 
		demoList = append(demoList, model)
	}
	fmt.Println(demoList,count)
	return demoList,count,err
}