package article_api

import (
	"context"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title" structs:"title"`     // 文章标题
	Abstract string   `json:"abstract" structs:"abstract"`  // 文章简介
	Content  string   `json:"content" structs:"content"`   // 文章内容
	Category string   `json:"category" structs:"category"`  // 文章分类
	Source   string   `json:"source" structs:"source"`    // 文章来源
	Link     string   `json:"link" structs:"link"`      // 原文链接
	BannerID uint     `json:"bannerId" structs:"banner_id"` // 文章封面id
	Tags     []string `json:"tags" structs:"tags"`      // 文章标签
	ID       string   `json:"id" structs:"id"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}
	fmt.Printf("得到的json数据为%#v\n",cr)
	var bannerUrl string
	if cr.BannerID != 0 {
		err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
		if err != nil {
			res.FailWithMessage("banner不存在", c)
			return
		}
	}

	article := models.ArticleModel{
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:     cr.Title,
		Keyword:   cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		BannerID:  cr.BannerID,
		BannerUrl: bannerUrl,
		Tags:      cr.Tags,
	}
    fmt.Printf("结构体数据为%#v\n",article)


	//更新前应该检测文章是否存在
	// if article.ISExistDataById() {
	// 	global.Log.Error(err)
	// 	res.FailWithMessage("文章不存在", c)
	// 	return
	// }


	//这里要转map的主要原因，我们json是驼峰格式，而es中是蛇形命名
	maps := structs.Map(&article)
	// fmt.Printf("获得的map值为%#v",maps)
	//要根据不同类型移除空值，因为structs会把没有值的键都采用零值
	var DataMap = map[string]any{}
	// 去掉空值
	for key, v := range maps {
		switch val := v.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array:
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		DataMap[key] = v
	}
	fmt.Printf("去掉空值的map值为%#v",DataMap)

	

	_, err = global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Doc(DataMap).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		res.FailWithMessage("更新失败", c)
		return
	}
	res.OkWithMessage("更新成功", c)
}