# 文件上传

## 单文件上传

单文件上传使用 gin.Context 的 FormFile()方法，该方法的值为 POST 请求中文件上传字段的名称：
例如我们在 post 请求中用`myFile`字段上传文件

```go
func (ImagesApi) OneFileUpload(c *gin.Context){
	//此处的"myFile"是post请求中上传文件对应的字段
	fileHeader,err := c.FormFile("myFile")
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

**存储文件**
调用 gin.Context 的 SaveUploadedFile()方法可以将文件保存到某个目录下：

```go
 dst := "./uploads/" + file.Filename
 // SaveUploadedFile(要写入的文件,要写入的文件路径)
 c.SaveUploadedFile(file,"./uploadFile")
```

**设置缓冲区大小**
Go 默认文件上传缓冲区为 32M，当有大量文件上传时，服务器内存的压力会很大，因此可以通过 MaxMultipartMemory 属性来设置缓冲区大小：

```go
engine := gin.Default()
//8M
engine.MaxMultipartMemory = 8 << 20
```

上传文件时，不限制文件大小可以会导致服务存储空间暴涨，因为有必须限制上传文件大小：

```go
fileMaxSize := 4 << 20 //4M
 if int(file.Size) > fileMaxSize {
   c.String(http.StatusBadRequest, "文件不允许大小于4M")
   return
 }
```

**限制文件类型**
对文件类型，也可以进行限制：

```go
reader, err := file.Open()
 if err != nil {
   fmt.Println(err)
   return
 }
 b := make([]byte, 512)
 reader.Read(b)
 ​
 contentType := http.DetectContentType(b)
 if contentType != "image/jpeg" && contentType != "image/png" {
   c.String(http.StatusOK, "文件格式错误")
   return
 }
```

在上面的代码中，我们读取文件的前 512 个字节，再调用 http.DetectContentType()便可以获取文件的 MIME 值。
**完整案例**

```go
package main
 ​
 import (
   "fmt"
   "log"
   "net/http"
 ​
   "github.com/gin-gonic/gin"
 )
 ​
 func main() {
   engine := gin.Default()
   //8M
   engine.MaxMultipartMemory = 8 << 20
   engine.POST("/upload", func(c *gin.Context) {
     file, err := c.FormFile("myFile")
     if err != nil {
       log.Println(err)
       c.String(http.StatusBadRequest, "文件上传失败")
       return
     }
     fileMaxSize := 4 << 20 //4M
     if int(file.Size) > fileMaxSize {
       c.String(http.StatusBadRequest, "文件不允许大小于32KB")
       return
     }
 ​
     reader, err := file.Open()
     if err != nil {
       fmt.Println(err)
       return
     }
     b := make([]byte, 512)
     reader.Read(b)
 ​
     contentType := http.DetectContentType(b)
     if contentType != "image/jpeg" && contentType != "image/png" {
       c.String(http.StatusOK, "文件格式错误")
       return
     }
 ​
     dst := "./uploads/" + file.Filename
     c.SaveUploadedFile(file, dst)
     c.String(http.StatusOK, fmt.Sprintf("'%s' 上传成功!", file.Filename))
   })
   engine.Run()
 }
```

## 多文件上传

如果要上传多个文件，多次调用 gin.Context 的 FormFile()方法也是可以的，但更好的方式是使用 gin.Context 的 MultipartForm()方法：

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
	fileList,ok := form.File["myFile"]

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

我们可以通过遍历来操作每一个文件对象

```go
    //遍历拿到的文件列表
	// file实际上就是fileHeader类型的实例
	for _,file := range fileList {
		fmt.Println(file.Filename)
		fmt.Println(file.Header)
		fmt.Println(file.Size)
	}
```

**案例演示**

```go
 package main
 ​
 import (
   "fmt"
   "log"
   "net/http"
 ​
   "github.com/gin-gonic/gin"
 )
 ​
 func main() {
   engine := gin.Default()
   engine.POST("/uploadMul", func(c *gin.Context) {
     form, err := c.MultipartForm()
     if err != nil {
       log.Println(err)
       c.String(http.StatusBadRequest, "文件上传失败")
       return
     }
     files := form.File["upload"]
     for _, file := range files {
       fmt.Println(file.Filename)
     }
     c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
   })
   engine.Run()
 }
```
