package log_stash

import (
	"gvb_server/global"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwts.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}

	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   "内网",
		userId: userID,
	}
}

//入库的方法
func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

//提示信息的方法
func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warn(content string) {
	l.send(WarnLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

//发送日志信息的方法
func (l Log) send(level Level, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}