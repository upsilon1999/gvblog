package message_api

import "time"

type MessageApi struct {
}

type MessageRequest struct {
	// SendUserID uint   `json:"sendUserId"` // 发送人id,可以直接从token获取
	RevUserID  uint   `json:"revUserId" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`    // 消息内容
}

type Message struct {
	SendUserID       uint      `json:"sendUserId"` // 发送人id
	SendUserNickName string    `json:"sendUserNickName"`
	SendUserAvatar   string    `json:"sendUserAvatar"`
	RevUserID        uint      `json:"revUserId"` // 接收人id
	RevUserNickName  string    `json:"revUserNickName"`
	RevUserAvatar    string    `json:"revUserAvatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"createdAt"`    // 最新的消息时间
	MessageCount     int       `json:"messageCount"` // 消息条数
}