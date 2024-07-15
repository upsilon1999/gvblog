package tag_api

//标签管理相关API
type TagApi struct{
	
}

type TagRequest struct{
	Title string `josn:"title" binding:"required" msg:"请输入标题" structs:"title"`//标题
}