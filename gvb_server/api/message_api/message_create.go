package message_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id,可以直接从token拿
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)


	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error("json解析错误",err)
	  res.FailWithError(err, &cr, c)
	  return
	}
	var senUser, recvUser models.UserModel
  
	count := global.DB.Take(&senUser, claims.UserID).RowsAffected
	if count == 0 {
		global.Log.Error("发送人不存在",err)
	  res.FailWithMessage("发送人不存在", c)
	  return
	}
	count1 := global.DB.Take(&recvUser, cr.RevUserID).RowsAffected
	if count1==0 {
		global.Log.Error("接收人不存在",err)
	  res.FailWithMessage("接收人不存在", c)
	  return
	}
  
	err = global.DB.Create(&models.MessageModel{
	  SendUserID:       claims.UserID,
	  SendUserNickName: senUser.NickName,
	  SendUserAvatar:   senUser.Avatar,
	  RevUserID:        cr.RevUserID,
	  RevUserNickName:  recvUser.NickName,
	  RevUserAvatar:    recvUser.Avatar,
	  IsRead:           false,
	  Content:          cr.Content,
	}).Error
	if err != nil {
	  global.Log.Error("消息发送失败",err)
	  res.FailWithMessage("消息发送失败", c)
	  return
	}
	res.OkWithMessage("消息发送成功", c)
  }