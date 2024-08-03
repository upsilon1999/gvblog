# 本项目简介

本项目采用的 es 是 7.x 版本，所以使用的库是

```sh
go get -u github.com/olivere/elastic/v7
```

版本 2 将升级为 es8，同时采用官方库

# 测试目录

## 连接 es

```go
package elasticSearch_test

import (
	"testing"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func TestConnect(t *testing.T) {
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
