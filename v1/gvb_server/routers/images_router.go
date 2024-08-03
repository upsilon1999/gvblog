package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	images := router.Group("images")
	{
		//上传单个图片 由于多图上传包含单图的功能，所以暂时废案
		images.POST("single",middleware.JwtAuth(), imagesApi.OneImageUpload) 
		images.POST("uploads",middleware.JwtAuth(), imagesApi.ImageUploadView) //上传多个图片
		images.GET("list",middleware.JwtAuth(), imagesApi.ImageListView) 
		images.DELETE("delete",middleware.JwtAuth(), imagesApi.ImageRemoveView)//批量删除图片
		images.PUT("updateName",middleware.JwtAuth(),imagesApi.ImageUpdateView)//修改图片名称
		images.GET("listNames",middleware.JwtAuth(), imagesApi.ImageNameListView) //获取简易图片列表
		
	}
}