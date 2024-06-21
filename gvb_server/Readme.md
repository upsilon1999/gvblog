# 后台开发文档

## 文件目录结构

```sh
├─api   接口目录
├─config  存放记录配置的结构体目录
├─core  初始化操作
├─docs  swag生成的api文档目录
├─flag  命令行相关的初始化
├─global 全局变量的包
├─middleware 中间件
├─models 表结构
├─routers  gin路由的目录
├─service  项目与服务有关的目录
├─testdata 测试文件的目录
├─utils 常用的一些工具合集
├─main.go 入口文件
└─settings.yaml 配置文件
```

**拓展**windows 下查看文件树
WIndow 平台要想打印目录树，可以用 cmd 工具或者 power shell 的 tree 命令实现

tree 命令格式和参数：

```sh
TREE [drive:][path] [/F] [/A]
```

/F 显示每个文件夹中文件的名称。（带扩展名）
/A 使用 ASCII 字符，而不使用扩展字符。(如果要显示中文，例如 tree /f /A >tree.txt)
比如：

> tree /f >tree.txt

导出当前目录的文件夹/文件的目录树到 tree.txt 文件中。

```sh
Tip：
要是目录很深文件很多生成的树大了去!，比如我在D盘根目录使用49MB的tree.txt；
没找到其它参数可以只打印一级或者二级目录类似的参数
```
