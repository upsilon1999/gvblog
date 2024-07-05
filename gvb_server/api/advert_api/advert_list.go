package advert_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"strings"

	"github.com/gin-gonic/gin"
)

//获取广告列表
func (AdvertApi) AdvertListView(c *gin.Context){
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr);err !=nil{
		res.FailWithCode(res.ArgumentError,c)
		return
	}

	 // 判断 Referer 是否包含admin，如果是，就全部返回，不是，就返回is_show=true
	 referer := c.GetHeader("Referer")
	 isShow := true
	 //参数1字符串是否包含子串参数2
	 if strings.Contains(referer, "admin") {
	   // admin来的
	   isShow = false
	 }

	list,count,_ := common.ComList(models.AdvertModel{IsShow: isShow},common.Option{
		PageInfo:cr,
		//开启打印debug日志
		Debug: true,
	})
	res.OkWithList(list,count,c)
}