package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/chat_api"
	"gvb_server/api/comment_api"
	"gvb_server/api/images_api"
	"gvb_server/api/log_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/news_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/upvote_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi images_api.ImagesApi
	AdvertApi advert_api.AdvertApi
	MenuApi menu_api.MenuApi
	UserApi user_api.UserApi
	TagApi tag_api.TagApi
	MessageApi message_api.MessageApi
	ArticleApi article_api.ArticleApi
	UpvoteApi upvote_api.UpvoteApi
	CommentApi comment_api.CommentApi
	NewsApi news_api.NewsApi
	ChatApi chat_api.ChatApi
	LogApi log_api.LogApi
}

var ApiGroupApp = new(ApiGroup)