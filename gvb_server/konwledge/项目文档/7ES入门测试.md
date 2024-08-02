本项目第一版本使用的es包

```sh
github.com/olivere/elastic/v7
```

## 连接es

```go
package elasticSearch_test

import (
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestConnet(t *testing.T) {
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	logrus.Fatalf("es连接成功 %s", c)
}
```

实际应用中可以在连接成功后抛出es实例，具体类型可以可以直接查看源码

## 索引操作

对标mysql的数据库操作

### 创建索引

为了实现和mysql类似的性质，我们需要创建关键的mapping

```go
package elasticSearch_test

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

//这里面的字段要根据mapping设计，达到映射效果
type DemoModel struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
func (DemoModel)Index() string{
	return "demo_index"
}

//这个是这个github.com/olivere/elastic/v7包的关键
//也是他与官网elastic的区别，他主要构筑了类似mysql的结构
/*
	我们解释一下下面模板字符串的设计
	settings(配置项)-->index(索引)-->max_result_window(能查到记录的最大数量)
	10万条足够我们用了

	mapping下面的properties是字段集合，例如
	title 字段名
	type 对应字段类型 
	{
		text，可以进行模糊匹配
		integer,可以比大小
		date 时间类型
	}
	 
	null_value 空值
	format  将值格式化成什么样存储在es中

	我们拿这个要存储的mapping和我们的结构体或者map相映射
	就和操作mysql差不多
*/
func (DemoModel)Mapping()string{
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}


// IndexExists 索引是否存在
func (demo DemoModel) IndexExists(client *elastic.Client) bool {
	exists, err := client.
	  IndexExists(demo.Index()).
	  Do(context.Background())
	if err != nil {
	  logrus.Error(err.Error())
	  return exists
	}
	return exists
}
```

### mapping解读

mapping是一切映射的关键，我们来解读一下

```json
{
  //配置项
  "settings": {
      //索引配置
    "index":{
        //能查到的最大索引条数，中小项目100000够用了
      "max_result_window": "100000"
    }
  }, 
    //映射，主要跟我们的结构体匹配，就和mysql的表结构一样
  "mappings": {
      //字段集合
    "properties": {
        //字段名
      "title": { 
          //字段类型，不一样的字段类型有着不同的功能
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
          //设置空值
        "null_value": "null",
          //格式化值的形式
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
```

这里说一下常见的mapping内字段类型

```json
"mappings": {
	"properties": {
				"id": {  //整形字段, 允许精确匹配
					"type": "integer",
				},
				"name": {
					"type":            "text",  //字符串类型且进行分词, 允许模糊匹配
					"analyzer":        "ik_smart", //设置分词工具
					"search_analyzer": "ik_smart",
					"fields": {    //当需要对模糊匹配的字符串也允许进行精确匹配时假如此配置
						"keyword": {
							"type":         "keyword",
							"ignore_above": 256,
						},
					},
				},
				"date_field": {  //时间类型, 允许精确匹配
					"type": "date",
				},
				"keyword_field": { //字符串类型, 允许精确匹配
					"type": "keyword",
				},
				"nested_field": { //嵌套类型
					"type": "nested",
					"properties": {
						"id": {
							"type": "integer",
						},
						"start_time": { //长整型, 允许精确匹配
							"type": "long",
						},
						"end_time": {
							"type": "long",
						},
					},
				},
			},
		},
}
```

### 映射结构体

```go
//这里面的字段要根据mapping设计，达到映射效果
type DemoModel struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}

//注意索引名只能采取蛇形命名法
func (DemoModel)Index() string{
	return "demo_index"
}
```

这个和数据库的model很相似

### 完整创建索引的测试

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type DemoModel struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
func (DemoModel)Index() string{
	return "demo_index"
}

func (DemoModel)Mapping()string{
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
// IndexExists 索引是否存在
func (demo DemoModel) IndexExists(client *elastic.Client) bool {
	exists, err := client.
	  IndexExists(demo.Index()).
	  Do(context.Background())
	if err != nil {
	  logrus.Error(err.Error())
	  return exists
	}
	return exists
}

//创建索引
func TestCreatIndex(t *testing.T) {

	//连接es
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)

	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	//实例化结构体
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
```

### 查询索引是否存在

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type DemoModel struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
func (DemoModel)Index() string{
	return "demo_index"
}

func (DemoModel)Mapping()string{
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
// 连接es
func Connect() (*elastic.Client,error) {
	//连接es
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
		return nil,err
	}
	return client, nil
}



//测试索引是否存在
func TestIndexExists(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	var demo = new(DemoModel)
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
```

### 删除索引

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type DemoModel struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
func (DemoModel)Index() string{
	return "demo_index"
}

func (DemoModel)Mapping()string{
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}
// 连接es
func Connect() (*elastic.Client,error) {
	//连接es
	sniffOpt := elastic.SetSniff(false)
	host := "http://127.0.0.1:9200"
	client, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
		return nil,err
	}
	return client, nil
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
```

## 文档操作

文档的增删改查操作，就相当于mysql中的记录(注意es自V7开始删除了type，也就没有了对标mysql表的概念)

### 创建记录

没有赋值，则全部采用结构体的零值

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestDocCreate(t *testing.T){
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	var data = new(DemoModel)
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
```

赋值之后，没有赋值到的将采用零值，注意我们插入的是记录数据，不要与ES的id混淆，Es的id设置需要其他API

```go
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
```

### Query方式列表查询

```go
package elasticSearch_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)
// FindList 列表查询
func TestFindList(t *testing.T)  {
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
```

### Source方式列表查询

这个查询没有生效原因不明

```go
package elasticSearch_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)
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
```

### 更新记录

```go
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
```

### 批量删除记录

根据id列表批量删除记录

```go
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
```



## 关于测试的小结

测试文件编译时会遍历包下的所有test文件，但凡有一个文件有问题也会提示编译不通过，例如

```go
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

	var demo = new(DemoModel)
    //这里的client不存在
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
```

虽然我们执行的是对`TestCreatIndex`的测试,但是由于`TestIndexExists`有错误，也会导致编译出错，而且最大的问题是，go-test不会帮我们定位到错误。

### 全局方法或变量

test执行的时候是扫描本包的所有`_test`文件，所以写在非test文件中的全局变量和全局方法都是识别不到的，

```sh
【诡异的点】
我们可以用"包名.方法名"或"包名.变量名"的方式引用其他包的方法和变量

但是没有用同样的方法引入本包的，所以`本包下的test全局变量/方法`无法与`本包的全局变量/方法`共通
```

#### 同包非test

`enter.go`

```go
package elasticSearch_test

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

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
```

`doc_test.go`

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestDocCreate(t *testing.T){
    //无法识别的Connect方法
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	var data = new(DemoModel)
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
```

#### 同包同test

`enter_test.go`

```go
package elasticSearch_test

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

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
```

`doc_test.go`

```go
package elasticSearch_test

import (
	"context"
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestDocCreate(t *testing.T){
    //可识别到Connect方法
	client,err := Connect()
	if err !=nil{
		logrus.Fatalf("es连接失败 %s", err.Error())
	}

	var data = new(DemoModel)
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
```

