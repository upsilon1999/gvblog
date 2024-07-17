package elasticSearch_test

import (
	"context"

	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

//这里面的字段要根据mapping设计，达到映射效果
type DemoIndex struct{
	ID        string `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"userId"`
	CreatedAt string `json:"createdAt"`
}
func (DemoIndex)Index() string{
	return "demoIndex"
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
func (DemoIndex)Mapping()string{
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
func (demo DemoIndex) IndexExists(client *elastic.Client) bool {
	exists, err := client.
	  IndexExists(demo.Index()).
	  Do(context.Background())
	if err != nil {
	  logrus.Error(err.Error())
	  return exists
	}
	return exists
}