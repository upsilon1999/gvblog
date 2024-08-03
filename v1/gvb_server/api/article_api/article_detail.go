package article_api

import (
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	"gvb_server/service/redis_ser"

	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

//通过es的id来获取数据
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}

	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
	  res.FailWithMessage(err.Error(), c)
	  return
	}

	//每次查看文章详情就增加一次文章的浏览量
	redis_ser.Look(model.ID)
	res.OkWithData(model, c)
}

// 根据keyword，即文章标题获取文章详情
type ArticleDetailRequest struct {
	Title string `json:"title" form:"title"`
}
  
func (ArticleApi) ArticleDetailByTitleView(c *gin.Context) {
	var cr ArticleDetailRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}

	

	model, err := es_ser.CommDetailByKeyword(cr.Title)
	if err != nil {
	  res.FailWithMessage(err.Error(), c)
	  return
	}

	//每次查看文章详情就增加一次文章的浏览量
	redis_ser.Look(model.ID)

	res.OkWithData(model, c)
}