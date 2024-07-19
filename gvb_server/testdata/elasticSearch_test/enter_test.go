package elasticSearch_test

import (
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