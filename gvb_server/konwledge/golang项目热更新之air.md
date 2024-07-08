## 安装air

```go
# 1.16版本及之后
go install github.com/cosmtrek/air@latest
```

## 使用Air

默认安装完后会到gopath目录下，来到我们的项目中

```sh
# 首先在当前目录下查找 `.air.conf`配置文件，如果找不到就使用默认的
air -c .air.conf
```

推荐的使用方式

```sh
# 1. 在当前目录创建一个新的配置文件.air.conf
touch .air.conf

# 2. 复制 `air.conf.example` 中的内容到这个文件，然后根据你的需要去修改它

# 3. 使用你的配置运行 air, 如果文件名是 `.air.conf`，只需要执行 `air`。
air
```

#### air_example.conf示例

```sh
# [Air](https://github.com/cosmtrek/air) TOML 格式的配置文件

# 工作目录
# 使用 . 或绝对路径，请注意 `tmp_dir` 目录必须在 `root` 目录下
root = "."
tmp_dir = "tmp"

[build]
# 只需要写你平常编译使用的shell命令。你也可以使用 `make`
cmd = "go build -o ./tmp/main ."
# 由`cmd`命令得到的二进制文件名
bin = "tmp/main"
# 自定义的二进制，可以添加额外的编译标识例如添加 GIN_MODE=release
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# 监听以下文件扩展名的文件.
include_ext = ["go", "tpl", "tmpl", "html"]
# 忽略这些文件扩展名或目录
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# 监听以下指定目录的文件
include_dir = []
# 排除以下文件
exclude_file = []
# 如果文件更改过于频繁，则没有必要在每次更改时都触发构建。可以设置触发构建的延迟时间
delay = 1000 # ms
# 发生构建错误时，停止运行旧的二进制文件。
stop_on_error = true
# air的日志文件名，该日志文件放置在你的`tmp_dir`中
log = "air_errors.log"

[log]
# 显示日志时间
time = true

[color]
# 自定义每个部分显示的颜色。如果找不到颜色，使用原始的应用程序日志。
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# 退出时删除tmp目录
clean_on_exit = true
```

然后就完成了 使用以下命令启动

```sh
air  #启动项目  取代之前的 go run main.go
```

## 关于配置文件

也可以使用

```sh
air init
```

在项目根目录下生成一个toml配置文件，然后进行修改
