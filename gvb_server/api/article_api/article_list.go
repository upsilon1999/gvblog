package article_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	list,count,err := es_ser.CommList(cr.Key,cr.Page,cr.Limit)
	if err != nil{
		global.Log.Error(err)
		res.FailWithMessage("查询失败",c)
	}
	res.OkWithList(filter.Omit("list", list),int64(count),c)
}