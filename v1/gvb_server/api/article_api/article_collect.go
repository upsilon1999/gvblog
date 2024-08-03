package article_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

// ArticleCollCreateView 用户收藏文章，或取消收藏
func (ArticleApi) ArticleCollCreateView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
  
	//通过id来获取文章详情
	model, err := es_ser.CommDetail(cr.ID)
	if err != nil {
	  res.FailWithMessage("文章不存在", c)
	  return
	}
  
	var collect models.UserCollectModel
	//查询
	err = global.DB.Take(&collect, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
	  // 没有找到 收藏文章
	  global.DB.Create(&models.UserCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ID,
	  })
	  // 给文章的收藏数 +1
	  num = 1
	}
	// 取消收藏
	// 文章数 -1
	global.DB.Delete(&collect)
  
	// 更新文章收藏数
	//收藏数是针对整篇文章的，和具体用户无关
	err = es_ser.ArticleUpdate(cr.ID, map[string]any{
	  "collectsCount": model.CollectsCount + num,
	})
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("更新收藏数据失败",c)
	}

	if num == 1 {
	  res.OkWithMessage("收藏文章成功", c)
	} else {
	  res.OkWithMessage("取消收藏成功", c)
	}
  }