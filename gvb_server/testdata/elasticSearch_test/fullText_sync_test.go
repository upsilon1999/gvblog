package elasticSearch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/models"
	"gvb_server/service/es_ser"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestSyncFullText(t *testing.T) {
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	boolSearch := elastic.NewMatchAllQuery()
	res, _ := client.
	  Search(models.ArticleModel{}.Index()).
	  Query(boolSearch).
	  Size(1000).
	  Do(context.Background())
	
	for _, hit := range res.Hits.Hits {
	  var article models.ArticleModel
	  _ = json.Unmarshal(hit.Source, &article)
	
	  indexList := es_ser.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)
	
	  bulk :=client.Bulk()
	  for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	  }
	  result, err := bulk.Do(context.Background())
	  if err != nil {
		logrus.Error(err)
		continue
	  }
	  fmt.Println(article.Title, "添加成功", "共", len(result.Succeeded()), " 条！")
	}
}