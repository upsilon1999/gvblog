package article_api

import (
	"context"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type IDListRequest struct {
	IDList []string `json:"idList"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	/*
		如果文章被删除了，用户收藏的这篇文章怎么办？
		有两种方案可以选择
		1.顺带把与这个文章关联的收藏也删除掉
		2.用户收藏表，新增一个字段，表示文章是否删除，
		
		第二个方法存在新问题:
		用户可以删除这个收藏记录，但是找不到文章去修改收藏数
	*/
	bulkService := global.ESClient.Bulk().Index(models.ArticleModel{}.Index()).Refresh("true")
	for _, id := range cr.IDList {
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)

		//同时删除全文搜索
		go es_ser.DeleteFullTextByArticleID(id)
	}
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
}