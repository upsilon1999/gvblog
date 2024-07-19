package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

// 连接es
// func Connect() (*elastic.Client,error) {
// 	//连接es
// 	sniffOpt := elastic.SetSniff(false)
// 	host := "http://127.0.0.1:9200"
// 	client, err := elastic.NewClient(
// 		elastic.SetURL(host),
// 		sniffOpt,
// 		elastic.SetBasicAuth("", ""),
// 	)
// 	if err !=nil{
// 		logrus.Fatalf("es连接失败 %s", err.Error())
// 		return nil,err
// 	}
// 	return client, nil
// }

//测试创建索引
func TestCreatIndex(t *testing.T) {
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	//结构体实例
	var demo = new(DemoModel)

	createIndex, err := client.
    CreateIndex(demo.Index()).
    BodyString(demo.Mapping()).
    Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
	}
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
	}
	logrus.Infof("索引 %s 创建成功", demo.Index())
}

//测试索引是否存在
func TestIndexExists(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	var demo = new(DemoModel)
	//测试索引是否存在
	exists, err := client.
	  IndexExists(demo.Index()).
	  Do(context.Background())
	if err != nil {
	  logrus.Error(err.Error())
	}
	if exists {
		logrus.Infof("索引已存在")
	}else{
		logrus.Infof("索引不存在")
	}
}

//删除索引
func TestRemoveIndex(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	var demo = new(DemoModel)
	// 删除索引
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
	  logrus.Error("删除索引失败")
	  logrus.Error(err.Error())
	}
	if !indexDelete.Acknowledged {
	  logrus.Error("删除索引失败")
	}
	logrus.Info("索引删除成功")
}