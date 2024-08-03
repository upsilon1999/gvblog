package elasticSearch_test

import (
	"context"
	"fmt"
	"gvb_server/models"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestHighLight(t *testing.T) {
	//连接es
	client, err := Connect()
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	/*
		NewMultiMatchQuery(要搜索的内容,被搜索字段1,被搜索字段2,...)
		elastic.NewHighlight().Field(要高亮的字段名)
	*/
	//单字段高亮测试
	// result, err := client.
	// 	Search(models.ArticleModel{}.Index()).
	// 	Query(elastic.NewMultiMatchQuery("node", "title", "abstract", "content")).
	// 	Highlight(elastic.NewHighlight().Field("title")).
	// 	Size(100).
	// 	Do(context.Background())
	// if err != nil {
	// 	logrus.Error(err)
	// 	return
	// }


	//多字段高亮测试
	var title = elastic.NewHighlighterField("title")
	var abstract = elastic.NewHighlighterField("abstract")
	var content = elastic.NewHighlighterField("content")
	result, err := client.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMultiMatchQuery("nodejs", "title", "abstract", "content")).
		Highlight(elastic.NewHighlight().Fields(title,abstract,content)).
		Size(100).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}


	for _, hit := range result.Hits.Hits {
		fmt.Println(string(hit.Source))
		fmt.Println(hit.Highlight)
	}
}