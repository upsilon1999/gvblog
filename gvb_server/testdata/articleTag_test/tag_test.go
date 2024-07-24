package articletag_test

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/models"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//返回给前端的
//这是要用结构体映射json，最终json是给前端的所以驼峰
type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"articleIdList"`
}
  
//用于接收聚类分析结果的
//这里用于将es分析的结果json转为结构体，所以json映射得和es分析结果类似
//es分析结果是蛇形命名，所以采用蛇形命名映射
type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
	  Key      string `json:"key"`
	  DocCount int    `json:"doc_count"`
	  Articles struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
		  Key      string `json:"key"`
		  DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	  } `json:"articles"`
	} `json:"buckets"`
}

//连接es
func Connect() (*elastic.Client, error) {
	//连接es
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
		return nil, err
	}
	return client, nil
}

func TestTag(t *testing.T) {
	// 我们最终要实现的文章标签形式为
	//[{tag:"python",articleCount:2,articleList:[]}]
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	ctx := context.Background()

	//查询总数
	result, err := client.
    Search(models.ArticleModel{}.Index()).
    Aggregation("mytags", elastic.NewValueCountAggregation().Field("tags")).
    Size(0).
    Do(ctx)

	if err != nil{
		logrus.Error(err)
		return 
	}
	cTag, _ := result.Aggregations.Cardinality("mytags")
	count := int64(*cTag.Value)
	fmt.Printf("总数为%d\n",count)

	//按照tags标签聚类
	agg := elastic.NewTermsAggregation().Field("tags")
	//添加一个子聚类，在之前聚类的基础上，根据keyword再次分组
	//得到keyword，也就是标题
	// agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	//得到id
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("_id"))
	//对数据进行分页，例如查询第1页前2条
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(1).Size(2))

	query := elastic.NewBoolQuery()

	result, err = client.
	  Search("article_index").
	  Query(query).
	  Aggregation("tags", agg).
	  Size(0).
	  Do(ctx)

	if err != nil{
		logrus.Error(err)
		return 
	}

	fmt.Printf("查询结果为%#v\n",string(result.Aggregations["tags"]))
	//根据查询的到聚类结果来构造json数据
	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
	//将聚类结果进行解析和映射
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	for _, bucket := range tagType.Buckets {

		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}

		tagList = append(tagList, TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
	}
	fmt.Printf("获得的标签列表%v\n",tagList)
}