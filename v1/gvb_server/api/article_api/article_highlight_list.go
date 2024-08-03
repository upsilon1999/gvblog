package article_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
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
	Tag string `json:"tags" form:"tags"`
}

//搜索列表
/*
	查询多个字段，并高亮对应字段
*/
func (ArticleApi) ArticleHighListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	list,count,err := es_ser.CommHighLightList(cr.Key,cr.Page,cr.Limit)
	if err != nil{
		global.Log.Error(err)
		res.FailWithMessage("查询失败",c)
	}
	res.OkWithList(filter.Omit("list", list),int64(count),c)
}


//搜索列表
/*
	查询多个字段，仅高亮标题
*/
func (ArticleApi) ArticleHighTitleView(c *gin.Context) {
	var cr HighListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	fmt.Printf("获取到的值为%#v\n",cr)
	list,count,err := es_ser.CommHighTitileList(es_ser.Option{
		Page: cr.Page,
		Limit: cr.Limit,
		Key: cr.Key,
		Sort: cr.Sort,
		Fields: []string{"title","abstract"},
		Tag: cr.Tag,
	})
	if err != nil{
		global.Log.Error(err)
		res.FailWithMessage("查询失败",c)
	}
	res.OkWithList(filter.Omit("list", list),int64(count),c)
}