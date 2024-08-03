package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	//限制图片大小为2M
	limitSize = 2
)


var (
	// WhiteImageList 图片上传的白名单
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
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

		//判断图片是否位于白名单列表中
		//根据截取后缀来判断,
		//1.先将文件名按点分割，得到列表
		imageNameList := strings.Split(file.Filename,".")
		//2.获取最后的一个值，即后缀，并且转为小写，因为文件后缀不区分大小写
		suffix := strings.ToLower(imageNameList[len(imageNameList)-1])
		//3.判断后缀是否位于白名单内，
		isInWhite := utils.InList(suffix,WhiteImageList)
		//4.根据是否在白名单列表中进行处理
		if !isInWhite {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:"非法文件",
			  })
			  continue
		}


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

		//获取到图片内容
		fileObj,err := file.Open()
		if err!=nil{
			global.Log.Error(err)
		}
		byteData,err :=io.ReadAll(fileObj)
		if err!=nil{
			global.Log.Error("读取出错",err)
		}
		//得到md5值
		imageHash := utils.Md5(byteData)
		fmt.Println(imageHash)

		//根据imageHash去数据库查询图片是否存在
		var bannerModel models.BannerModel
		/*
			查询
			情况1: 查询出错  res.Error!=nil&&res.Error!=record
			情况2:图片已存在 res.RowsAffected >0 记录数大于0
			情况3:图片不存在且不报错
		*/
		res := global.DB.Take(&bannerModel,"hash=?",imageHash)
		//查询出错
		if res.Error!=nil&&res.Error!=gorm.ErrRecordNotFound{
			global.Log.Error("查询出错",err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       res.Error.Error(),
			})
			continue
		}
		//图片已存在
		if res.RowsAffected>0{
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       "图片已存在",
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

		//写入文件成功后，将图片内容写入数据库
		global.DB.Create(&models.BannerModel{
			Path:filePath,
			Hash:imageHash,
			Name:file.Filename,
		})
	}

	res.OkWithData(resList, c)
}