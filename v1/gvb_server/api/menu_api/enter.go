package menu_api

import "gvb_server/models/ctype"

type MenuApi struct {
}

type ImageSort struct {
	ImageID uint `json:"imageId"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	TitleEn         string      `json:"titleEn" binding:"required" msg:"请完善菜单英文名称" structs:"title_en"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstractTime" structs:"abstract_time"`                // 切换的时间，单位秒
	BannerTime    int         `json:"bannerTime" structs:"banner_time"`                    // 切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"` // 菜单的序号
	ImageSortList []ImageSort `json:"imageSortList" structs:"-"`                          // 具体图片的顺序
}
