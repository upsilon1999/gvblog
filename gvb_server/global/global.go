package global

import (
	"gvb_server/config"

	"gorm.io/gorm"
)

var (
	// 配置文件全局变量
	Config *config.Config
	//数据库全局变量
	DB     *gorm.DB
)
