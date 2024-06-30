package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	//上传单个图片 由于多图上传包含单图的功能，所以暂时废案
	router.POST("imageAlone", imagesApi.OneImageUpload) 
	router.POST("images", imagesApi.ImageUploadView) //上传多个图片
}