package chat_api

import (
	"encoding/json"
	"fmt"
	"gvb_server/models/res"
	"net/http"
	"strings"
	"time"

	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

/*
	chat_group.throw 是由前端传递头像和昵称的废弃版本，保留作为参照组

	这个文件是由后端随机生成头像和昵称

*/

//存储所有的群聊信息
var ConnGroupMap = map[string]ChatUser{}

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}


type MsgType int

// 发送的消息类型
const (
	//文字消息
	TextMsg MsgType = 1
	//图片消息
	ImageMsg MsgType = 2
	//系统消息
	SystemMsg MsgType = 3
	//进入聊天室消息
	InRoomMsg MsgType = 4
	//离开聊天室消息
	OutRoomMsg MsgType = 5
)



type GroupRandRequest struct {
	Content string  `json:"content"`  // 聊天的内容
	MsgType MsgType `json:"msg_type"` // 聊天类型
}
type GroupRnadResponse struct {
	NickName string    `json:"nick_name"` // 前端自己生成
	Avatar   string    `json:"avatar"`    // 头像
	MsgType  MsgType   `json:"msg_type"`  // 聊天类型
	Content  string    `json:"content"`   // 聊天的内容
	Date     time.Time `json:"date"`      // 消息的时间
}

func (ChatApi) ChatGroupRandView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	nickNameFirst := string([]rune(nickName)[0])
	avatar := fmt.Sprintf("uploads/chat_avatar/%s.png", nickNameFirst)

	chatUser := ChatUser{
		Conn:     conn,
		NickName: nickName,
		Avatar:   avatar,
	}
	ConnGroupMap[addr] = chatUser
	// 需要去生成昵称，根据昵称首字关联头像地址
	// 昵称关联 addr

	logrus.Infof("%s 连接成功", addr)
	for {
		// 消息类型，消息，错误
		_, p, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(GroupRnadResponse{
				Content: fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				Date:    time.Now(),
			})
			break
		}
		// 进行参数绑定
		var request GroupRandRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			// 参数绑定失败
			continue
		}
		// 判断类型
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				continue
			}
			SendGroupMsg(GroupRnadResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content:  request.Content,
				MsgType:  TextMsg,
				Date:     time.Now(),
			})
		case InRoomMsg:
			SendGroupMsg(GroupRnadResponse{
				Content: fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:    time.Now(),
			})
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(response GroupRnadResponse) {
	byteData, _ := json.Marshal(response)
	for _, chatUser := range ConnGroupMap {
		chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
}