# 图片上传

## 单图片上传

单图片上传使用 gin.Context 的 FormFile()方法，该方法的值为 POST 请求中文件上传字段的名称：
例如我们在 post 请求中用`images`字段上传文件

```go
func (ImagesApi) OneFileUpload(c *gin.Context){
	//此处的"images"是post请求中上传文件对应的字段
	fileHeader,err := c.FormFile("images")
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
```

我们来看一下 FormFile 的源码

```go
// FormFile returns the first file for the provided form key.
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}
// A FileHeader describes a file part of a multipart request.
type FileHeader struct {
	Filename string
	Header   textproto.MIMEHeader
	Size     int64

	content   []byte
	tmpfile   string
	tmpoff    int64
	tmpshared bool
}

// Open opens and returns the FileHeader's associated File.
func (fh *FileHeader) Open() (File, error) {
	if b := fh.content; b != nil {
		r := io.NewSectionReader(bytes.NewReader(b), 0, int64(len(b)))
		return sectionReadCloser{r, nil}, nil
	}
	if fh.tmpshared {
		f, err := os.Open(fh.tmpfile)
		if err != nil {
			return nil, err
		}
		r := io.NewSectionReader(f, fh.tmpoff, fh.Size)
		return sectionReadCloser{r, f}, nil
	}
	return os.Open(fh.tmpfile)
}
```

由于多图上传的限制和单图上传一样，所以我们的拓展部分放到多图上传

## 多图片上传

如果要上传多个图片，多次调用 gin.Context 的 FormFile()方法也是可以的，但更好的方式是使用 gin.Context 的 MultipartForm()方法：

```go
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
	//image是post传输文件对应的字段
	fileList,ok := form.File["images"]

	if !ok {
		res.FailWithMessage("不存在的文件",c)
		return
	}

	//循环拿到的文件列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		fmt.Println(file.Filename)
		fmt.Println(file.Header)
		fmt.Println(file.Size)
	}
}
```

我们来看一下源码

```go
// MultipartForm is the parsed multipart form, including file uploads.
func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory)
	return c.Request.MultipartForm, err
}
type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
}
```

我们通过遍历来操作每一个文件对象

```go
    //遍历拿到的文件列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		fmt.Println(file.Filename)
		fmt.Println(file.Header)
		fmt.Println(file.Size)
	}
```

### **简易存储图片**

我们存储文件既可以用 go 原生的文件写入方法，也可以使用 gin 提供的 API，例如

```go
//遍历拿到的图片列表
// file实际上就是fileHeader类型的实例
for _,file := range fileList {
    filepath := path.Join("uploads",file.Filename)
    // SaveUploadedFile(要写入的文件,要写入的文件路径)
    err =c.SaveUploadedFile(file,filepath)
    if err!=nil{
        global.Log.Error(err)
        continue
    }
}
```

SaveUploadedFile 就是写入文件的 API，源码为

```go
// SaveUploadedFile uploads the form file to specific dst.
func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
```

### **判断存放目录路径是否存在**

我们已经在配置文件中加入了文件上传目录，如下

```go
global.Config.upload.Path = "uploads/file"
```

现在来判断文件路径是否存在

```go
 basePath := global.Config.Upload.Path
  _, err = os.ReadDir(basePath)
  if err != nil {
    // 递归创建
    err = os.MkdirAll(basePath, fs.ModePerm)
    if err != nil {
      global.Log.Error(err)
    }
  }
```

所谓递归创建，就是如下形式

```go
err = os.MkdirAll("uploads/file/xxx", fs.ModePerm)
```

如果不存在，就继续进行，我们先定义一下文件上传结构体，方便后续接收

```go
type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
}
```

然后进行多图的创建

```go
	//如果路径不存在
	// 不存在就创建
	var resList []FileUploadResponse

	//遍历拿到的图片列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		filePath := path.Join("uploads",file.Filename)
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
```

完整代码

```go
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

type FileUploadResponse struct {
	FileName  string `json:"file_name"`  // 文件名
	IsSuccess bool   `json:"is_success"` // 是否上传成功
	Msg       string `json:"msg"`        // 消息
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
```

目前没有对于文件已存在的处理，如果上传同名文件将覆盖原文件，我们可以在后续加入 md5 校验或者给上传的文件名字加入时间戳。

### 限制文件大小

我们知道文件对象的一个属性是 Size,该属性的单位是字节
直接贴出比较的逻辑

```go
const (
	//限制图片大小为2M
	limitSize = 2
)

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


	//遍历拿到的图片列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		filePath := path.Join("uploads",file.Filename)

		// 判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(limitSize) {
			global.Log.Error("图片大小超过限制")
			continue
		}


		// SaveUploadedFile(要写入的文件,要写入的文件路径)
		err =c.SaveUploadedFile(file,filePath)
		//写入失败
		if err!=nil{
			global.Log.Error(err)
			continue
		}
	}

	res.OkWith(c)
}
```

配合文件是否存在的完整代码

```go
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
	//如果uploads/file路径不存在
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
			//   %.2f 保留两位小数
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
```

### 白名单、黑名单

黑名单:判断文件名后缀，如果与黑名单中的后缀符合，那就拒绝上传

白名单:只能上传在白名单中出现的文件后缀

我们事先封装了一个方法，用于比较值是否在列表中，例如

```go
package utils

// InList 判断key是否存在与列表中
func InList(key string, list []string) bool {
	for _, s := range list {
	  if key == s {
		return true
	  }
	}
	return false
  }
```



我们设置白名单列表,处理逻辑如下

1.设置白名单列表

```go
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
```

2.根据获取到的文件名进行处理

```go
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
```

### 写入数据库

在我们读取完成后，需要将上传图片的信息写入数据库，

对应models

```go
package models

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"os"
)

type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        // 图片路径
	Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
	Name      string          `gorm:"size:38" json:"name"`         // 图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型， 本地还是七牛
}

func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
```

存入数据库

```go
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
```

###  md5校验

**golang的md5校验**

值得完善的一点，关于golang的md5校验

```go
func Md5(src []byte){
    m:= md5.New()
    m.Write(src)
    res:=hex.EncodeToString(m.Sum(nil))
    return res
}
```

获取到图片的内容并进行md5校验

```go
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
```

## 图片列表查询

**简易查询返回结果**

```go
func (ImagesApi) ImageListView(c *gin.Context) {
	var imagesList []models.BannerModel

    //select * from banner_model
	global.DB.Find(&imagesList)

	res.OkWithData(imagesList,c)
}
```

获取列表的长度、分页的逻辑

```go
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err!=nil{
		res.FailWithCode(res.ArgumentError,c)
		return 
	}
	var imagesList []models.BannerModel

	//获取总条数
	//select count(*) from banner_model
	count:=global.DB.Find(&imagesList).RowsAffected

	// cr.Page当前页码 cr.Limit每页取多少条
	offset := (cr.Page-1)*cr.Limit
    if offset<0{
        offset = 0
    }

	//分页操作，一页取一条数据
	//select * from banner_model limit 1 offset 1
	//global.DB.Limit(1).Offset(1).Find(&imagesList)
	global.DB.Limit(cr.Limit).Offset(offset).Find(&imagesList)
    
    res.OkWithData(gin.H{"count":count,"list":imagesList},c)
}
```

单独封装一个分页的方法

```go
type Option struct {
  models.PageInfo
  Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

  DB := global.DB
    //开启mysql日志打印模式
  if option.Debug {
    DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
  }
  if option.Sort == "" {
    option.Sort = "created_at desc" // 默认按照时间往前排
  }

  count = DB.Select("id").Find(&list).RowsAffected
  offset := (option.Page - 1) * option.Limit
  if offset < 0 {
    offset = 0
  }
  err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

  return list, count, err
}
```

使用封装好的函数进行返回

```go
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

func OkWithList(list any, count int64, c *gin.Context) {
	OkWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
	  res.FailWithCode(res.ArgumentError, c)
	  return
	}
  
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
	  PageInfo: cr,
	  Debug:    false,
	})
  
	res.OkWithList(list, count, c)
  
	return
}
```

## 图片删除

使用到了钩子函数

```go
type BannerModel struct {
  MODEL
  Path      string          `json:"path"`                        // 图片路径
  Hash      string          `json:"hash"`                        // 图片的hash值，用于判断重复图片
  Name      string          `gorm:"size:38" json:"name"`         // 图片名称
  ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` // 图片的类型， 本地还是七牛
}

//删除图片信息前一并删除本地图片
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
  if b.ImageType == ctype.Local {
    // 本地图片，删除，还要删除本地的存储
    err = os.Remove(b.Path)
    if err != nil {
      global.Log.Error(err)
      return err
    }
  }
  return nil
}
```

这里的BeforeDelete是gorm钩子函数的命名，具体的可以访问gorm官网文档查询hook

使用到了批量删除

```go
package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type RemoveRequest struct{
	IDList []uint `json:"id_list"`
}


func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr RemoveRequest
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
```

## 图片信息修改

这里以修改图片名为例

```go
package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}
//修改图片信息
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
	  res.FailWithError(err, &cr, c)
	  return
	}
	var imageModel models.BannerModel
	count := global.DB.Take(&imageModel, cr.ID).RowsAffected
	if count == 0 {
	  res.FailWithMessage("文件不存在", c)
	  return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
	  res.FailWithMessage(err.Error(), c)
	  return
	}
	res.OkWithMessage("图片名称修改成功", c)
	return
  
  }
```

gorm版本差异,在老版本中

```go
count := global.DB.Take(&imageModel, cr.ID).RowsAffected
```

可以替换成

```go
err = global.DB.Take(&imageModel, cr.ID).Error()
if err != nil {
	  res.FailWithMessage("文件不存在", c)
	  return
}
```

但是在新版本中如果没有找到,err也会报错，所以采用记录数进行判断

## 图片上传到七牛云

代码展示

```go
import (
  "bytes"
  "context"
  "errors"
  "fmt"
  "github.com/qiniu/go-sdk/v7/auth/qbox"
  "github.com/qiniu/go-sdk/v7/storage"
  "gvb_server/config"
  "gvb_server/global"
  "time"
)

// 获取上传的token
func getToken(q config.QiNiu) string {
  accessKey := q.AccessKey
  secretKey := q.SecretKey
  bucket := q.Bucket
  putPolicy := storage.PutPolicy{
    Scope: bucket,
  }
  mac := qbox.NewMac(accessKey, secretKey)
  upToken := putPolicy.UploadToken(mac)
  return upToken
}

// 获取上传的配置
func getCfg(q config.QiNiu) storage.Config {
  cfg := storage.Config{}
  // 空间对应的机房
  zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
  cfg.Zone = &zone
  // 是否使用https域名
  cfg.UseHTTPS = false
  // 上传是否使用CDN上传加速
  cfg.UseCdnDomains = false
  return cfg

}

// UploadImage 上传图片  文件数组，前缀
func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
  if !global.Config.QiNiu.Enable {
    return "", errors.New("请启用七牛云上传")
  }
  q := global.Config.QiNiu
  if q.AccessKey == "" || q.SecretKey == "" {
    return "", errors.New("请配置accessKey及secretKey")
  }
  if float64(len(data))/1024/1024 > q.Size {
    return "", errors.New("文件超过设定大小")
  }
  upToken := getToken(q)
  cfg := getCfg(q)

  formUploader := storage.NewFormUploader(&cfg)
  ret := storage.PutRet{}
  putExtra := storage.PutExtra{
    Params: map[string]string{},
  }
  dataLen := int64(len(data))

  // 获取当前时间
  now := time.Now().Format("20060102150405")
  key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)

  err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
  if err != nil {
    return "", err
  }
  return fmt.Sprintf("%s%s", q.CDN, ret.Key), nil

}
```

