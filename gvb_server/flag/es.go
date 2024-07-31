package flag

import (
	"fmt"
	"gvb_server/models"
)

func EsCreateIndex() {
	err := models.ArticleModel{}.CreateIndex()
	if err!=nil{
		fmt.Println("创建文章索引失败")
	}
	//由于文章和全文两个是相关的所以操作不分开
	err = models.FullTextModel{}.CreateIndex()
	if err!=nil{
		fmt.Println("创建全文索引失败")
	}
}