package main

import (
	"fmt"
	"gvb_server/models"

	"github.com/fatih/structs"
)

//
type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`     // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法"` // 图片
	IsShow *bool   `json:"isShow" binding:"required" msg:"请选择是否展示" structs:"is_show"`    // 是否展示
}

func main()  {
	ul := AdvertRequest{
		Title:"xxx",
		Href:"xxx",
		Images:"xxx",
		IsShow:true,
	}
	m3 := structs.Map(&ul)
	fmt.Println(m3)
}