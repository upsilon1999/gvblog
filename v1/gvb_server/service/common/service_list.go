package common

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"

	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string //模糊匹配的字段
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

	DB = DB.Where(model)

	/*
	这里做模糊匹配查询，根据传入的模糊匹配列表来查
	注意 select * from table where id = 1 and ip like "%0.1%" or addr like "%网%"

	一般查询与模糊匹配之间用and，模糊匹配与模糊匹配之间用or
	*/
	for idx,col := range option.Likes{
		if idx == 0 {
			DB.Where(fmt.Sprintf("%s like ?",col),fmt.Sprintf("%%%s%%",option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?",col),fmt.Sprintf("%%%s%%",option.Key))
	}

	/*
	//由于Select("id")的影响，query变成了只有id一列，我们有两种解决方案
	//1.将Select("id")去掉，相当于select *
	//2.再次给query赋值，相当于复位
	count = query.Select("id").Find(&list).RowsAffected
	query = DB.Where(model)
	*/
	count = DB.Find(&list).RowsAffected
	//设置默认值
	//因为新版的gorm不传默认为0

	query := DB.Where(model)
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