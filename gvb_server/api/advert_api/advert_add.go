package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

//增加广告
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err!= nil{
		//这个封装逻辑有一个问题，例如我们前端传入了herf或is_show
		//也就是值传错了，我们暴露给前端的报错信息不是字段错误，而是msg信息
		//其实不利于错误定位
		global.Log.Error(err)
		res.FailWithError(err,&cr,c)
		return
	}

	// 重复的判断
	var advert models.AdvertModel
	count := global.DB.Take(&advert, "title = ?", cr.Title).RowsAffected
	if count != 0 {
	  res.FailWithMessage("该广告已存在", c)
	  return
	}

	//添加广告入数据库
	err =global.DB.Create(&models.AdvertModel{
		Title: cr.Title,
		Href: cr.Href,
		Images: cr.Images,
		IsShow: *cr.IsShow,
	}).Error
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败",c)
		return
	}

	res.OkWithMessage("添加广告成功",c)
}