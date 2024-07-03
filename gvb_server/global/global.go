package global

import (
	"gvb_server/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// 配置文件全局变量
	Config *config.Config
	//数据库全局变量
	DB     *gorm.DB
	//日志全局变量
	Log *logrus.Logger
	//全局mysql日志
	MysqlLog logger.Interface
)

