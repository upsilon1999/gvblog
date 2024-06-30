package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/res"
	"io/fs"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

const (
	//限制图片大小为2M
	limitSize = 2
)


//功能：上传单个图片
//返回值: 返回图片url
//这个是个废案，主要用来展示单个文件上传的方法
func (ImagesApi) OneImageUpload(c *gin.Context){
	// 使用gin封装的上传文件的方法，image是post传递文件对应的字段名
	//只能上传单个文件，如果一次性上传多个文件将只读取第一个
	fileHeader,err := c.FormFile("image")
	if err!=nil{
		res.FailWithMessage(err.Error(),c)
		return
	}

	/*
		fileHeader对象上的属性和方法
		Header 请求头
		Filename 文件名
		Size 文件大小
		Open() 打开文件
	*/
	fmt.Println(fileHeader.Filename)
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)

	// fileHeader.
}

//上传多个文件，返回url列表
func (ImagesApi) ImageUploadView(c *gin.Context){
	// 使用gin封装的上传文件的方法，支持上传多个文件
	form,err := c.MultipartForm()
	if err!=nil{
		res.FailWithMessage(err.Error(),c)
		return
	}

	//form实际上是个文件列表
	//form上有Value和File
	//images是post传递文件对应的字段名
	fmt.Println(form)
	fileList,ok := form.File["images"]

	if !ok {
		res.FailWithMessage("不存在的文件",c)
		return
	}


	// 判断路径是否存在
	//如果uoloads/file路径不存在
	// 不存在就创建
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}

	var resList []FileUploadResponse

	//遍历拿到的图片列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		filePath := path.Join(basePath,file.Filename)
		// 判断大小
		size := float64(file.Size) / float64(1024*1024)

		if size >= float64(limitSize) {
			resList = append(resList, FileUploadResponse{
			  FileName:  file.Filename,
			  IsSuccess: false,
			  Msg:       fmt.Sprintf("图片大小超过设定大小，当前大小为:%.2fMB, 设定大小为：%dMB ", size, limitSize),
			})
			continue
		  }

		// SaveUploadedFile(要写入的文件,要写入的文件路径)
		err =c.SaveUploadedFile(file,filePath)
		//写入失败
		if err!=nil{
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		//写入成功
		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}

	res.OkWithData(resList, c)
}