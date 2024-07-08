package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// ImageRemoveView 批量删除图片
// @Tags 图片管理
// @Summary 批量删除图片
// @Description 批量删除图片
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "图片id列表"
// @Router /api/images/delete [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err!=nil{
		res.FailWithCode(res.ArgumentError,c)
		return 
	}
	var imageList []models.BannerModel
	//查看查到的记录数
	count:=global.DB.Find(&imageList,cr.IDList).RowsAffected
	if count == 0{
		res.FailWithMessage("图片不存在",c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除%d张图片",count),c)
}