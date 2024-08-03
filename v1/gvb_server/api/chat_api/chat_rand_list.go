package chat_api

import (
	"fmt"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/sirupsen/logrus"
)

func (ChatApi) ChatRandListView(c *gin.Context){
	var cr models.PageInfo
	// fmt.Printf("接收到的前端值为%#v\n",cr)
	err := c.ShouldBindQuery(&cr)
	if err !=nil{
		res.FailWithCode(res.ArgumentError,c)
		return
	}

	cr.Sort = "created_at desc"
	// fmt.Printf("接收到的前端值为%#v\n",cr)
	//过滤是否返回群聊消息
	list,count,err :=common.ComList(models.ChatModel{IsGroup: true},common.Option{
		PageInfo: cr,
	})

	fmt.Printf("数据为%#v\n",list)
	if err!=nil{
		logrus.Errorf("分页查询出错，错误为%#v\n",err)
		res.FailWithMessage("分页查询出错",c)
		return
	}

	data:=filter.Omit("list",list)

	//解决Omit后的数组空值问题
	//因为这个会给数组空值为{},对前端很不友好
	_list,_ := data.(filter.Filter)
	if string(_list.MustMarshalJSON())=="{}"{
		list = make([]models.ChatModel,0)
		res.OkWithList(list,count,c)
		return
	}

	res.OkWithList(_list,count,c)
}