package article_api

import (
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type HighListRequest struct{
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
	Tag string `json:"tags" form:"tag"`
}


func (ArticleApi) ArticleHighListView(c *gin.Context) {
	var cr HighListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	list,count,err := es_ser.CommHighLightList(es_ser.Option{
		Page: cr.Page,
		Limit: cr.Limit,
		Key: cr.Key,
		Sort: cr.Sort,
		Fields: []string{"title","abstract","content"},
		Tag: cr.Tag,
	})
	if err != nil{
		global.Log.Error(err)
		res.FailWithMessage("查询失败",c)
	}
	res.OkWithList(filter.Omit("list", list),int64(count),c)
}