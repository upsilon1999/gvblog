package models

// TagModel 标签表
type TagModel struct {
	MODEL
	Title string `gorm:"size:16" json:"title"` // 标签的名称
	// Articles []ArticleModel `grom:"many2many:artice_tag_models" json:"-"` //关联的文章列表
}
