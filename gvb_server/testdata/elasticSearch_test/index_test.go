package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)


func TestCreatIndex(t *testing.T) {
	//连接es
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)

	var demo = DemoIndex

	createIndex, err := client.
    CreateIndex(demo.Index()).
    BodyString(demo.Mapping()).
    Do(context.Background())
  if err != nil {
    logrus.Error("创建索引失败")
    logrus.Error(err.Error())
    return err
  }
  if !createIndex.Acknowledged {
    logrus.Error("创建失败")
    return err
  }
  logrus.Infof("索引 %s 创建成功", demo.Index())
}