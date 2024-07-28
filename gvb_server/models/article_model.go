package models

import (
	"context"
	"encoding/json"
	"gvb_server/global"
	"gvb_server/models/ctype"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

//json中的omit和select是来自json-filter包的，警告可忽略
type ArticleModel struct {
	ID        string `json:"id" structs:"id" mapstructure:"id"`                 // es的id
	CreatedAt string `json:"createdAt" structs:"created_at" mapstructure:"created_at"` // 创建时间
	UpdatedAt string `json:"updatedAt" structs:"updated_at" mapstructure:"updated_at"` // 更新时间

	Title    string `json:"title" structs:"title" mapstructure:"title"`                // 文章标题
	Keyword  string `structs:"keyword" mapstructure:"keyword" json:"keyword,omit(list)"` // 关键字
	Abstract string `json:"abstract" mapstructure:"abstract" structs:"abstract"`          // 文章简介
	Content  string `structs:"content" mapstructure:"content" json:"content,omit(list)"` // 文章内容

	LookCount     int `json:"lookCount" structs:"look_count" mapstructure:"look_count"`         // 浏览量
	CommentCount  int `json:"commentCount" structs:"comment_count"  mapstructure:"comment_count"`   // 评论量
	UpvoteCount     int `json:"upvoteCount" structs:"upvote_count" mapstructure:"upvote_count"`         // 点赞量
	CollectsCount int `json:"collectsCount" structs:"collects_count" mapstructure:"collects_count"` // 收藏量

	UserID       uint   `json:"userId" structs:"user_id" mapstructure:"user_id"`               // 用户id
	UserNickName string `json:"userNickName" structs:"user_nick_name" mapstructure:"user_nick_name"` //用户昵称
	UserAvatar   string `json:"userAvatar" structs:"user_avatar" mapstructure:"user_avatar"`       // 用户头像

	Category string `json:"category" structs:"category" mapstructure:"category"`        // 文章分类
	Source   string `json:"source" structs:"source" mapstructure:"source"` // 文章来源
	Link     string `json:"link" structs:"link" mapstructure:"link"`     // 原文链接

	BannerID  uint   `json:"bannerId" structs:"banner_id" mapstructure:"banner_id"`   // 文章封面id
	BannerUrl string `json:"bannerUrl" structs:"banner_url" mapstructure:"banner_url"` // 文章封面

	Tags ctype.Array `json:"tags" structs:"tags" mapstructure:"tags"` // 文章标签
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
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
        "type": "text",
		"analyzer": "ik_max_word",
		"search_analyzer": "ik_smart"
      },
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text",
		"analyzer": "ik_max_word",
		"search_analyzer": "ik_smart"
      },
      "content": { 
        "type": "text",
		"analyzer": "ik_max_word",
		"search_analyzer": "ik_smart"
      },
      "lookCount": {
        "type": "integer"
      },
      "commentCount": {
        "type": "integer"
      },
      "upvoteCount": {
        "type": "integer"
      },
      "collectsCount": {
        "type": "integer"
      },
      "userId": {
        "type": "integer"
      },
      "userNickName": { 
        "type": "keyword"
      },
      "userAvatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "bannerId": {
        "type": "integer"
      },
      "bannerUrl": { 
        "type": "keyword"
      },
      "tags": { 
        "type": "keyword"
      },
      "createdAt":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "updatedAt":{
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
func (a ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(a.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return exists
	}
	return exists
}

// CreateIndex 创建索引
func (a ArticleModel) CreateIndex() error {
	if a.IndexExists() {
		// 有索引
		a.RemoveIndex()
	}
	// 没有索引
	// 创建索引
	createIndex, err := global.ESClient.
		CreateIndex(a.Index()).
		BodyString(a.Mapping()).
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
	logrus.Infof("索引 %s 创建成功", a.Index())
	return nil
}

// RemoveIndex 删除索引
func (a ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := global.ESClient.DeleteIndex(a.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// Create 添加文章的方法
func (a *ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(a.Index()).
		BodyJson(a).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	a.ID = indexResponse.Id
	return nil
}

// ISExistData 是否存在该文章
func (a ArticleModel) ISExistTitle() bool {
	//NewTermQuery(key,value) 精确匹配键值
	res, err := global.ESClient.
		Search(a.Index()).
		Query(elastic.NewTermQuery("keyword", a.Title)).
		Size(1).
		Do(context.Background())

	//这里报错，虽然是操作es出错了，但我们也认为是没存在文章
	//错误根据错误捕捉来人工修改
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	//查询结果
	if res.Hits.TotalHits.Value > 0 {
		return true
	}
	return false
}

// func (a ArticleModel) ISExistDataById() bool {
// 	//NewTermQuery(key,value) 精确匹配键值
// 	_, err := global.ESClient.
// 		Get().
// 		Index(a.Index()).
// 		Id(a.ID).
// 		Do(context.Background())
// 	if err != nil {
// 		return true
// 	}
// 	return false
// }
func (a *ArticleModel) GetDataByID(id string) error {
	res, err := global.ESClient.
		Get().
		Index(a.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return err
	}
	err = json.Unmarshal(res.Source, a)
	return err
}
