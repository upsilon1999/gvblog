package global

import (
	"gvb_server/config"

	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
