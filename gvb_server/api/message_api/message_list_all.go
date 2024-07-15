package message_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

//管理员查看所有的聊天记录
func (MessageApi)MessageListAllView(c *gin.Context){
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr);err !=nil{
		res.FailWithCode(res.ArgumentError,c)
		return
	}

	list,count,_ := common.ComList(models.MessageModel{},common.Option{
		PageInfo:cr,
		//开启打印debug日志
		Debug: true,
	})
	res.OkWithList(list,count,c)
}