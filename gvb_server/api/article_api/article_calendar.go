package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

/*
	用于构造前端所需数据
	dateList:[
		{date:"2023-12-14",count:1},
		{date:"2024-05-01",count:5}
	]
*/
type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}
  
//用于解析获取es的结果
type BucketsType struct {
	Buckets []struct {
	  KeyAsString string `json:"key_as_string"`
	  Key         int64  `json:"key"`
	  DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
 }

//构造出一种格式{时间字符串:count}
//这样可以配合时间遍历获得单天的count
var DateCount = map[string]int{}

func (ArticleApi) ArticleCalendarView(c *gin.Context) {

	// 时间聚合
	agg := elastic.NewDateHistogramAggregation().Field("createdAt").CalendarInterval("day")
  
	// 时间段搜索
	// 从今天开始，到去年的今天
	now := time.Now()
	//获取一年前的时间
	aYearAgo := now.AddDate(-1, 0, 0)
	// fmt.Println("一年前时间",aYearAgo)
  
	format := "2006-01-02 15:04:05"
	// lt 小于  gt 大于
	//时间使用format的原因是因为我们存入的created_at也使用了格式化，为了方便比较，应该格式化为同一格式
	query := elastic.NewRangeQuery("createdAt").
	  Gte(aYearAgo.Format(format)).
	  Lte(now.Format(format))
  
	result, err := global.ESClient.
	  Search(models.ArticleModel{}.Index()).
	  Query(query).
	  Aggregation("calendar", agg).
	  Size(0).
	  Do(context.Background())
	if err != nil {
	  global.Log.Error(err)
	  res.FailWithMessage("查询失败", c)
	  return
	}
	// fmt.Printf("查询结果为%#v",string(result.Aggregations["calendar"]))

	//将获得的结果映射到json中
	var data BucketsType
  	err = json.Unmarshal(result.Aggregations["calendar"], &data)
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("结果解析错误", c)
		return
	}
	//将获得的结果，例如
	/*
		data.Buckets= [
				{
					"key_as_string": "2024-07-23 15:21:00",
					"key": 1721748060000,
					"doc_count": 1
				},
				{
					"key_as_string": "2024-07-23 15:22:00",
					"key": 1721748120000,
					"doc_count": 2
				}
				]
		}
	*/
  	var resList = make([]CalendarResponse, 0)
  	for _, bucket := range data.Buckets {
		//按照我们之前的format格式化时间
		Time, _ := time.Parse(format, bucket.KeyAsString)
		//构造 {时间字符串:count}
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount
  	}

	//now.Sub(aYearAgo).Hours()获得两个时间差之间的总小时数
	//除以24获得天数，由于格式是int64，所以我们转换一下
	days := int(now.Sub(aYearAgo).Hours() / 24)

	//按照天数遍历
	for i := 0; i <= days; i++ {

		//格式化每一天，使用AddDate(0, 0, i)获得递增时间序列
		day := aYearAgo.AddDate(0, 0, i).Format("2006-01-02")

		//由于{时间字符串:count}格式，我们可以通过时间获取数据
		count, _ := DateCount[day]
		resList = append(resList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OkWithData(resList, c)
	
}