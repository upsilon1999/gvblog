package elasticSearch_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//创建记录
func TestDocCreate(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	//1.我们这里没有传入的值将采用类型零值
	//2.此处的id不是es的id，而是我们这条记录的自定义id
	//3.es要设置id需要采用其他方式
	var data = DemoModel{
		ID: "0",
		Title:"first",
		UserID:18,
	}
	//创建索引不指定id
	indexResponse, err := client.Index().
    Index(data.Index()).
    BodyJson(data).Do(context.Background())

	if err != nil {
		logrus.Error(err.Error())
	}
	//正常的业务中会使用这个id，因为他是文档唯一标识
	data.ID = indexResponse.Id
	logrus.Infof("记录创建成功,id为%s", data.ID)
}

//FindList列表查询
// FindList 列表查询
func TestQueryFindList(t *testing.T)  {
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}


	//构造假的初始值
	//页数
	var page = 1
	//
	key := "first"
	limit := 0


	var demoList []DemoModel

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
  
	res, err := client.
	  Search(DemoModel{}.Index()).
	  Query(boolSearch).
	  From((from - 1) * limit).
	  Size(limit).
	  Do(context.Background())
	if err != nil {
	  logrus.Error(err.Error())
	}

	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
	  var demo DemoModel
	  data, err := hit.Source.MarshalJSON()
	  if err != nil {
		logrus.Error(err.Error())
		continue
	  }
	  err = json.Unmarshal(data, &demo)
	  if err != nil {
		logrus.Error(err)
		continue
	  }
	  //将查到数据的ES的id赋值给结构体的id字段
	  demo.ID = hit.Id
	  demoList = append(demoList, demo)
	}
	logrus.Infof("查询结果为%v,条数为%d", demoList,count)
  }
  
  //这个查询没有生效原因不明
  func TestSourceFindList(t *testing.T)  {
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}


	//构造假的初始值
	//页数
	var page = 1
	//
	key := "first"
	limit := 0


	//结构体切片实例化
	demoList := []DemoModel{}

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
  
	res, err := client.
    Search(DemoModel{}.Index()).
    Query(boolSearch).
    Source(`{"_source": ["title"]}`).
    From((from - 1) * limit).
    Size(limit).
    Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
		logrus.Error(err.Error())
		continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
		logrus.Error(err)
		continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	logrus.Infof("查询结果为%v,条数为%d", demoList,count)
  }

  //根据id更新数据
  func TestUpdateById(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	//正常使用中应该由外部传递
	var id = "T3UcyZABZFHQi8GVwzA9"
	//用于更新的结构体实例
	data := DemoModel{
		Title: "newTime",
	}

	_, err = client.
    Update().
    Index(DemoModel{}.Index()).
    Id(id).
    Doc(map[string]string{
      "title": data.Title,
    }).
    Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
	}
	logrus.Info("更新demo成功")
  }

  func TestMoreDel(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	var idList = []string{"TXX9yJABZFHQi8GVxTCm","TnUSyZABZFHQi8GV9jCT"}

	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}

	res, err := bulkService.Do(context.Background())
  	if err !=nil{
		logrus.Fatalf("删除失败 %s", err.Error())
	}
	logrus.Infof("删除成功，结果为%v", res)

  }