package es_ser

import (
	"context"
	"encoding/json"
	"errors"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/service/redis_ser"
	"strings"

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
	//获取文章点赞数据
	upvoteInfo := redis_ser.GetUpvoteInfo()
	//获取文章浏览数
	lookInfo := redis_ser.GetLookInfo()
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
		model.ID = hit.Id 
		//同步每一条的点赞数据
		upvote := upvoteInfo[hit.Id]
		look := lookInfo[hit.Id]
		model.UpvoteCount += upvote
		model.LookCount+=look
		demoList = append(demoList, model)
	}
	// fmt.Println(demoList,count)
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
	  logrus.Error(err)
	  return
	}
	model.ID= res.Id
	//同步浏览量
	model.LookCount += redis_ser.GetLook(res.Id)
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
	  logrus.Error(err)
	  return
	}
	model.ID = hit.Id
	model.LookCount += redis_ser.GetLook(hit.Id)
	return
}
  


//获取分页并高亮多个字段
func CommHighLightList(key string, page int, limit int)(list []models.ArticleModel,count int,err error){
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
	  boolSearch.Must(
		//构造多字段查询
		elastic.NewMultiMatchQuery(key, "title", "abstract"),
	  )
	}
	if limit == 0 {
	  limit = 10
	}
	if from == 0 {
	  from = 1
	}


	var title = elastic.NewHighlighterField("title")
	var abstract = elastic.NewHighlighterField("abstract")
	res, err := global.ESClient.
    Search(models.ArticleModel{}.Index()).
    Query(boolSearch).
	Highlight(elastic.NewHighlight().Fields(title,abstract)).
    From((from - 1) * limit).
    Size(limit).
    Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
		return nil,0,err
	}

	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []models.ArticleModel{}
	//获取文章点赞数据
	upvoteInfo := redis_ser.GetUpvoteInfo()
	//获取文章浏览数
	lookInfo := redis_ser.GetLookInfo()
	for _,hit := range res.Hits.Hits{
		var model models.ArticleModel
		data,err := hit.Source.MarshalJSON()
		if err!=nil{
			logrus.Error(err.Error())
			continue
		}
		
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		//要高亮哪些字段就在这里添加
		//只有在这里添加的才会返回到前端
		if title, ok := hit.Highlight["title"];ok {
			model.Title = title[0]
		}
		if abstract, ok := hit.Highlight["abstract"];ok {
			model.Abstract = abstract[0]
		}


		model.ID = hit.Id 
		//同步每一条的点赞数据
		upvote := upvoteInfo[hit.Id]
		//同步浏览数
		look := lookInfo[hit.Id]
		model.UpvoteCount += upvote
		model.LookCount+=look
		demoList = append(demoList, model)
	}
	// fmt.Println(demoList,count)
	return demoList,count,err
}

//分页搜索，但仅高亮标题
func CommHighTitileList(option Option)(list []models.ArticleModel,count int,err error){
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
	
	//排序相关操作
	//该结构体的来源是Sort需要的参数类型
	type SortField struct{
		//按照哪个字段排序
		Field string
		//排序方式
		Ascending bool
	}
	//构造默认值
	sortField := SortField{
		Field: "created_at",
		//true是升序，即从小到大 
		//false是降序，即从大到小
		Ascending: false, 
	}

	/*
		当前端传递了排序时，由于sort的格式为

		字段名:排序方式

		例如:created_at:desc
	*/
	if option.Sort != "" {
		_list := strings.Split(option.Sort, ":")
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
		  sortField.Field = _list[0]
		  //desc降序
		  if _list[1] == "desc" {
			sortField.Ascending = false
		  }
		  //asc升序
		  if _list[1] == "asc" {
			sortField.Ascending = true
		  }
		}
	}

	// fmt.Printf("接收到的数据为%#v\n",option)

	// Sort(sortField.Field, sortField.Ascending).
	if option.Limit == 0{
		option.Limit=10
	}

	//Highlight加入高亮搜索
	//Sort加入排序搜索
	res, err := global.ESClient.
    Search(models.ArticleModel{}.Index()).
    Query(boolSearch).
	Highlight(elastic.NewHighlight().Field("title")).
	Sort(sortField.Field, sortField.Ascending).
    From(option.GetFrom()).
    Size(option.Limit).
    Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
		return nil,0,err
	}

	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []models.ArticleModel{}

	//获取文章点赞数据
	upvoteInfo := redis_ser.GetUpvoteInfo()

	//获取文章浏览数
	lookInfo := redis_ser.GetLookInfo()
	for _,hit := range res.Hits.Hits{
		var model models.ArticleModel
		data,err := hit.Source.MarshalJSON()
		if err!=nil{
			logrus.Error(err.Error())
			continue
		}
		
		err = json.Unmarshal(data, &model)
		if err != nil {
			logrus.Error(err)
			continue
		}
		// fmt.Printf("每条数据为%#v\n",model)
		if title, ok := hit.Highlight["title"];ok {
			model.Title = title[0]
		}

		model.ID = hit.Id 

		//同步每一条的点赞数据
		upvote := upvoteInfo[hit.Id]
		look := lookInfo[hit.Id]
		model.UpvoteCount += upvote
		model.LookCount+=look

		demoList = append(demoList, model)
	}
	// fmt.Println(demoList,count)
	return demoList,count,err
}

//更新记录
func ArticleUpdate(id string,data map[string]any)error{
	_,err := global.ESClient.
	Update().
	Index(models.ArticleModel{}.Index()).
	Id(id).
	Doc(data).
	Do(context.Background())
	return err
}