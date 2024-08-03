package chat_api

import (
	"encoding/json"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
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



// 发送的消息类型
const (
	//进入聊天室消息
	InRoomMsg ctype.MsgType = 1
	//文字消息
	TextMsg ctype.MsgType = 2
	//图片消息
	ImageMsg ctype.MsgType = 3
	//语音消息
	VoiceMsg ctype.MsgType = 4
	//视频消息
	VideoMsg ctype.MsgType=5
	//系统消息
	SystemMsg ctype.MsgType = 6
	
	//离开聊天室消息
	OutRoomMsg ctype.MsgType = 7
)



type GroupRandRequest struct {
	Content string  `json:"content"`  // 聊天的内容
	MsgType ctype.MsgType `json:"msgType"` // 聊天类型
}
type GroupRnadResponse struct {
	NickName string    `json:"nickName"` // 前端自己生成
	Avatar   string    `json:"avatar"`    // 头像
	MsgType  ctype.MsgType   `json:"msgType"`  // 聊天类型
	Content  string    `json:"content"`   // 聊天的内容
	OnlineCount int `json:"onlineCount"` //聊天室在线人数
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
	avatar := fmt.Sprintf("uploads/chat_random_avatar/%s.png", nickNameFirst)

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
			SendGroupMsg(conn,GroupRnadResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				MsgType:  OutRoomMsg,
				Content: fmt.Sprintf("%s 离开聊天室", chatUser.NickName),
				Date:    time.Now(),
				//每发一条消息都获取在线人数
				//离开聊天室应该减少1
				OnlineCount: len(ConnGroupMap)-1,
			})
			break
		}
		// 进行参数绑定
		var request GroupRandRequest
		err = json.Unmarshal(p, &request)
		if err != nil {
			logrus.Errorf("参数绑定出错,错误为%v\n", err)
			// 参数绑定失败
			SendMsg(addr, GroupRnadResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				MsgType:  SystemMsg,
				Content:  "参数绑定失败",
				Date:    time.Now(),
				//每发一条消息都获取在线人数
				OnlineCount: len(ConnGroupMap),
			  })
			// 参数绑定失败
			continue
		}
		// 判断类型
		switch request.MsgType {
		case TextMsg:
			if strings.TrimSpace(request.Content) == "" {
				SendMsg(addr, GroupRnadResponse{
					NickName: chatUser.NickName,
					Avatar:   chatUser.Avatar,
					MsgType:  SystemMsg,
					Content:  "消息不能为空",
					Date:    time.Now(),
					//每发一条消息都获取在线人数
					OnlineCount: len(ConnGroupMap),
				})
				continue
			}
			SendGroupMsg(conn,GroupRnadResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content:  request.Content,
				MsgType:  TextMsg,
				Date:     time.Now(),
				//每发一条消息都获取在线人数
				OnlineCount: len(ConnGroupMap),
			})
		case InRoomMsg:
			SendGroupMsg(conn,GroupRnadResponse{
				NickName: chatUser.NickName,
				Avatar:   chatUser.Avatar,
				Content: fmt.Sprintf("%s 进入聊天室", chatUser.NickName),
				Date:    time.Now(),
				//每发一条消息都获取在线人数
				OnlineCount: len(ConnGroupMap),
			})
		default:
			SendMsg(addr, GroupRnadResponse{
			  NickName: chatUser.NickName,
			  Avatar:   chatUser.Avatar,
			  MsgType:  SystemMsg,
			  Content:  "消息类型错误",
			  Date:    time.Now(),
			  //每发一条消息都获取在线人数
			  OnlineCount: len(ConnGroupMap),
			})
		}

	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(conn *websocket.Conn, response GroupRnadResponse) {
	byteData, _ := json.Marshal(response)
	_addr := conn.RemoteAddr().String()
	ip, addr := getIPAndAddr(_addr)
  
	global.DB.Create(&models.ChatModel{
	  NickName: response.NickName,
	  Avatar:   response.Avatar,
	  Content:  response.Content,
	  IP:       ip,
	  Addr:     addr,
	  IsGroup:  true,
	  MsgType:  response.MsgType,
	})
	for _, chatUser := range ConnGroupMap {
	  chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	}
  }
  
  // SendMsg 给某个用户发消息
  func SendMsg(_addr string, response GroupRnadResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[_addr]
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
	  NickName: response.NickName,
	  Avatar:   response.Avatar,
	  Content:  response.Content,
	  IP:       ip,
	  Addr:     addr,
	  IsGroup:  false,
	  MsgType:  response.MsgType,
	})
	chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
  }
  
  func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	addr = "内网"
	return addrList[0], addr
  }