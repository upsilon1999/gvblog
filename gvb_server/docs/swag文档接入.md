## 下载swag

```sh
# 方法1
go install github.com/swaggo/swag/cmd/swag@v1.8.12


# 方法2
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

// 生成api文档
# 他会自动生成再项目的docs目录下
swag init

浏览器访问
http://ip:host/swagger/index.html
```

代码部分

```go
import (
  _ "gvb_server/docs"  // swag init生成后的docs路径
)
// @title API文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:9000
// @BasePath /
func main() {}
```

路由

```go
import (
  "github.com/gin-gonic/gin"
  swaggerFiles "github.com/swaggo/files"
  gs "github.com/swaggo/gin-swagger"
)
func Routers() *gin.Engine {
  Router := gin.Default()
  PublicGroup := Router.Group("")
  PublicGroup.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}
```

## 使用

注释

```go
// @Tags 标签
// @Summary 标题
// @Description 描述，可以有多个
// @Param limit query string false "表示单个参数"
// @Param data body request.Request    true  "表示多个参数"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} gin.H{"msg": "响应"}
func (a *Api) UsersApi(c *gin.Context) {}
```

## 广告增删改查的接口文档

```go
// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}


// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{}


// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Param token header string  true  "token"
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/adverts/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}


// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo    false  "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}

```

## 注意事项

在执行`swag init`时提示

```sh
swag : 无法将“swag”项识别为 cmdlet、函数、脚本文件或可运行程序的名称。
```

这可能有两个原因

```sh
1.GoPath没有写入环境变量
2.我们的项目GoPath配置错误
```

1.先用`go env`查看gopath目录，然后在该目录的bin目录下寻找是否有`swag.exe`,如果有则可能是我们没有将gopath目录写入环境变量，写入即可。

2.如果gopath目录下没有`swag.exe`则需要在该gopath目录下执行安装swag包的命令，然后再检查是否写入了环境变量

```sh
# 查看swag版本
swag -v
```

3.当我们处理完后用cmd发现可以识别swag，但是用vscode终端还是报错，这时候可以有以下两个操作

```sh
1.查看vscode执行策略
get-ExecutionPolicy
如果不是RemoteSigned，需要执行
set-ExecutionPolicy RemoteSigned


2.如果执行策略没有问题，那么就是vscode没有读取系统变量，将所有vscode关闭再次打开，或者重新开机
```

## gin使用swag可能遇到的问题

### 1. 生成的 swagger.json 为空

https://github.com/swaggo/swag

https://github.com/swaggo/gin-swagger

这两个项目可以根据注释来生成 `swagger` 文档。具体如何生成的，看一下项目文档就能明白。

执行 `swag init` 时会生成一个 `doc` 目录里面包含 `doc.go、swagger.json、swagger.yaml`

#### 如果你的项目目录是这样的

```go
├── cmd
│   ├── main.go
│   └── server
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── internal
    ├── server
    │   ├── grpc.go
    │   ├── http.go
```

如果你的项目目录大概是这个样子的。`main.go` 没有在项目根目录下，`internal/server/http.go` 是你的"注释"存放的位置。然后你在 `cmd` 目录下执行 `swag init --output ../docs` 是没有办法生成你想要的内容的。

因为执行 `swag init` 时，会在当前目录下搜索注释内容并生成 `swagger.json` 文档。

所以：你可以在 `cmd` 和 `internal` 的上层目录执行：

```sh
swag init -g ./cmd/main.go
```

默认会在`cmd` 同级目录生成出来 `docs` 文档目录

#### 看一下 swag init 的参数

```sh
swag init -h
NAME:
   swag init - Create docs.go

USAGE:
   swag init [command options] [arguments...]

OPTIONS:
   --generalInfo value, -g value       API通用信息所在的go源文件路径，如果是相对路径则基于API解析目录 (默认: "main.go")
   
   --dir value, -d value               API解析目录 (默认: "./")
   
   --propertyStrategy value, -p value  结构体字段命名规则，三种：snakecase,camelcase,pascalcase (默认: "camelcase")
   
   --output value, -o value            文件(swagger.json, swagger.yaml and doc.go)输出目录 (默认: "./docs")
   --parseVendor                       是否解析vendor目录里的go源文件，默认不
   --parseDependency                   是否解析依赖目录中的go源文件，默认不
   --markdownFiles value, --md value   指定API的描述信息所使用的markdown文件所在的目录
   --generatedTime                     是否输出时间到输出文件docs.go的顶部，默认是
```

也就是说，如果你没有通过 -g 指定 main.go 的位置，则默认查找当前目录。如果你没有通过 -o 指定 docs 的生成位置，则默认是当前目录。