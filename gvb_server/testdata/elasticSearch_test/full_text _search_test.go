package elasticSearch_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestFullText(t *testing.T) {
	//连接es
	// client, err := Connect()
	// if err != nil {
	// 	logrus.Fatalf("es连接失败 %s", err.Error())
	// }

	//要解析的数据
	data:="## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录：\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
	//1.对数据按行进行切割
	dataList := strings.Split(data,"\n")
	var isCode bool = false
	//2.然后对每行进行判断，如果是#号开头就是标题，反之是正文
	// var headList,bodyList []string
	for _,str := range dataList{
		/*
			实际上这样会错误捕获到两种情况
			情形1:#123中间无空格 
			对应解决方案使用正则，例如 #{1,6} *

			情形2：
			```
			# 在代码块中的注释
			```
			对应解决方案，判断是否是代码块，如果是就不纳入

			但是捕获到他们对我们的目的影响不大
			因为我们就是要构造全文搜索的标题
		*/
		if strings.HasPrefix(str, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(str,"#")&&!isCode{
			fmt.Println(str)
		}
	}
}