package advert_api

// 广告管理相关API
type AdvertApi struct{
	
}

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"跳转链接非法" structs:"href"`   // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` // 图片
	//使用指针形式的原因:可以传入零值
	IsShow *bool   `json:"isShow" binding:"required" msg:"请选择是否展示" structs:"is_show"`  // 是否展示
}