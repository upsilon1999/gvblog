package article_api

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"articleIdList"`
}

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

func (ArticleApi) ArticleTagListView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	//获取总数，由于上下文的存在，所以求取总数在这里做，不受后续聚类或分页影响
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Aggregation("tags", elastic.NewValueCountAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	cTag, _ := result.Aggregations.Cardinality("tags")
	count := int64(*cTag.Value)

	//按照tags标签聚类
	agg := elastic.NewTermsAggregation().Field("tags")

	//添加一个子聚类，在之前聚类的基础上，根据keyword再次分组
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))
	// agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))
	//分页操作
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit))
	query := elastic.NewBoolQuery()

	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	var tagType TagsType
	var tagList = make([]TagsResponse, 0)
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

	res.OkWithList(tagList, count, c)
}