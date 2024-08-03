package es_ser

import (
	"context"
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
)

type SearchData struct {
	Key   string `json:"key"` //文章关联id
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

//构造全文搜索
func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n")
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title))
	for _, s := range dataList {
		// #{1,6}
		// 判断一下是否是代码块
		if strings.HasPrefix(s, "```") {
			isCode = !isCode
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))
			//if strings.TrimSpace(body) != "" {
			bodyList = append(bodyList, getBody(body))
			//}
			body = ""
			continue
		}
		body += s
	}
	bodyList = append(bodyList, getBody(body))
	ln := len(headList)
	for i := 0; i < ln; i++ {
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key: id,
		})
	}
	b, _ := json.Marshal(searchDataList)
	fmt.Println(string(b))
	return searchDataList
}

func getHeader(head string) string {
	head = strings.ReplaceAll(head, "#", "")
	head = strings.ReplaceAll(head, " ", "")
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}



//通过全文搜索到es
//第一版设定这个会抛出错误，然后文章添加时会阻塞
//第二版设定，这个会提示错误，但是不抛出错误，不干扰文章的创建，在文章添加时使用协程
func AysncFullText(id, title, content string){
	  //创建全文搜索内容
	  indexList := GetSearchIndexDataByContent(id, title, content)
	
	  bulk :=global.ESClient.Bulk()
	  for _, indexData := range indexList {
		req := elastic.NewBulkIndexRequest().Index(models.FullTextModel{}.Index()).Doc(indexData)
		bulk.Add(req)
	  }
	  result, err := bulk.Do(context.Background())
	  if err != nil {
		logrus.Error(err)
		// return err
	  }
	  logrus.Infof( "%s添加成功, 共%d条", title, len(result.Succeeded()))
	//   return nil
}

//删除全文搜索
// DeleteFullTextByArticleID 删除全文搜索数据
func DeleteFullTextByArticleID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	res, _ := global.ESClient.
	  DeleteByQuery().
	  Index(models.FullTextModel{}.Index()).
	  Query(boolSearch).
	  Do(context.Background())
	//如果id所代表的数据不存在也不会报错，而是删除0条
	logrus.Infof("成功删除 %d 条记录", res.Deleted)
  }