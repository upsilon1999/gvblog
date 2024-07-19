package flag

import (
	"fmt"
	"gvb_server/models"
)

func EsCreateIndex() {
	err := models.ArticleModel{}.CreateIndex()
	if err!=nil{
		fmt.Println("创建es索引失败")
	}
}