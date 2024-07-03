package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	images := router.Group("images")
	{
		//上传单个图片 由于多图上传包含单图的功能，所以暂时废案
		images.POST("single", imagesApi.OneImageUpload) 
		images.POST("uploads", imagesApi.ImageUploadView) //上传多个图片
		images.GET("list", imagesApi.ImageListView) 
		images.DELETE("images", imagesApi.ImageRemoveView)//批量删除图片
		images.PUT("updateName",imagesApi.ImageUpdateView)
	}
}