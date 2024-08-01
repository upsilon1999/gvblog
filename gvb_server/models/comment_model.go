package models

// CommentModel 评论表
type CommentModel struct {
	//json:",select(c)" 这种tag字段为空的方式会直接把该结构体展开，当作匿名结构体处理
	//此处select(c) 是针对comment的缩写
	MODEL              `json:",select(c)"`
	SubComments        []CommentModel `gorm:"foreignkey:ParentCommentID" json:"subComments,select(c)"` // 子评论列表
	ParentCommentModel *CommentModel  `gorm:"foreignkey:ParentCommentID" json:"commentModel"`          // 父级评论
	ParentCommentID    *uint          `json:"parentCommentId,select(c)"`                              // 父评论id
	Content            string         `gorm:"size:256" json:"content,select(c)"`                        // 评论内容
	UpvoteCount          int            `gorm:"size:8;default:0;" json:"upvoteCount,select(c)"`            // 点赞数
	CommentCount       int            `gorm:"size:8;default:0;" json:"commentCount,select(c)"`         // 子评论数
	ArticleID          string         `gorm:"size:32" json:"articleId,select(c)"`                      // 文章id
	User               UserModel      `json:"user,select(c)"`                                           // 关联的用户
	UserID             uint           `json:"userId,select(c)"`                                        // 评论的用户
}
