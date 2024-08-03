package advert_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Param token header string  true  "token"
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/advert/update/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")


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

	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil&&err==gorm.ErrRecordNotFound {
	  res.FailWithMessage("该广告不存在", c)
	  return
	}else if err !=nil{
		global.Log.Error(err)
		res.FailWithMessage("查询出错",c)
		return
	}

	maps := structs.Map(&cr)
	fmt.Println(maps)
	//通过map来修改数据
	//注意map的键对应的是数据库字段名
	err =global.DB.Model(&advert).Updates(maps).Error
	if err!=nil{
		global.Log.Error(err)
		res.FailWithMessage("修改广告失败",c)
		return
	}

	res.OkWithMessage("修改广告成功",c)
}