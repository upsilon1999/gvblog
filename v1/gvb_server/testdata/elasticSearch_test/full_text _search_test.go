package elasticSearch_test

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

func getHead(str string) string{
	//将#号替换为空格
	head := strings.ReplaceAll(str,"#","")
	//将所有空格移除
	head = strings.ReplaceAll(head," ","")
	//构造出#标题的形式,适用于hash路由
	return head
}

//构造hash路由
func getSlug(slug string) string {
	return "#" + slug
}

func getBody(str string) string{
	unsafe := blackfriday.MarkdownCommon([]byte(str))
	doc,_ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	fmt.Println(doc.Text())
	return doc.Text()
}

func TestFullText(t *testing.T) {
	

	//要解析的数据
	data:="## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录：\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
	//1.对数据按行进行切割
	dataList := strings.Split(data,"\n")
	var isCode bool = false
	//2.然后对每行进行判断，如果是#号开头就是标题，反之是正文
	var headList,bodyList []string
	var body string
	//加一个文章标题，应该由前端传入
	headList = append(headList, "文章标题")

	for _,str := range dataList{
		
		//移除代码块的影响
		if strings.HasPrefix(str, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(str,"#")&&!isCode{
			headList = append(headList, getHead(str))
			// if strings.TrimSpace(body) != "" {
				bodyList = append(bodyList, getBody(body))
			// }
			body = ""
			continue
		}
		body += str
	}
	bodyList = append(bodyList, getBody(body))
	
	var searchDataList []SearchData
	//文章的id，未来也由前端传入
	id := "/article/asdxxccvv"
	ln := len(headList)
	for i := 0; i < ln; i++ {
	  searchDataList = append(searchDataList, SearchData{
		Title: headList[i],
		Body:  bodyList[i],
		Slug:  id + getSlug(headList[i]),
	  })
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	fmt.Println(searchDataList)
}