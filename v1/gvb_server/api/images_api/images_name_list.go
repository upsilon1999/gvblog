package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct{
	ID  uint `json:"id"`
	Path      string          `json:"path"`                        // 图片路径
	Name      string          `json:"name"`         // 图片名称
}

// ImageListView 图片名称列表
// @Tags 图片管理
// @Summary 图片名称列表
// @Description 图片名称列表
// @Router /api/images/listNames [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {
	//获取简易的图片列表信息，只包含id、路径、图片名
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id","path","name").Scan(&imageList)
	res.OkWithData(imageList, c)
}