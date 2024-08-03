package common

import (
	"gvb_server/global"
	"gvb_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

//列表查询
func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}

	query := DB.Where(model)

	/*
	//由于Select("id")的影响，query变成了只有id一列，我们有两种解决方案
	//1.将Select("id")去掉，相当于select *
	//2.再次给query赋值，相当于复位
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model)
	*/
	count = query.Find(&list).RowsAffected
	//设置默认值
	//因为新版的gorm不传默认为0
	if option.Page == 0{
		option.Page =1
	}
	if option.Limit ==0{
		option.Limit = 10
	}
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}