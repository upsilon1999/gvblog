package email_conn

import (
	"testing"

	"gopkg.in/gomail.v2"
)

type Subject string

const (
	Code  Subject = "平台验证码"
	Note  Subject = "操作通知"
	Alorm Subject = "告警通知"
)

type Api struct {
	Subject Subject
}

func (a Api) Send(name, body string) error {
	return send(name, string(a.Subject), body)
}

func NewCode() Api {
	return Api{
		Subject: Code,
	}
}

func NewNote() Api {
	return Api{
		Subject: Note,
	}
}

func NewAlorm() Api {
	return Api{
		Subject: Alorm,
	}
}

// send 发给谁，主题，正文
func send(name, subject, body string) error {
	// e := global.Config.Email

	// return sendMail(
	// 	e.User,
	// 	e.Password,
	// 	e.Host,
	// 	e.Port,
	// 	name,
	// 	e.DefaultFromEmail,
	// 	subject,
	// 	body,
	// )

	return sendMail(
		"1459894855@qq.com",
		"sgwdqybkhlbyhfee",
		"smtp.qq.com",
		25,
		name,
		"xxx",
		subject,
		body,
	)
}

func sendMail(userName,authCode,host string,port int,mailTo,sendName string,subject,body string) error{
	m := gomail.NewMessage()
	m.SetHeader("Form",m.FormatAddress(userName,sendName))//谁发的
	m.SetHeader("To",mailTo)//发送给谁
	m.SetHeader("Subject",subject)
	m.SetBody("text/html",body)
	d:= gomail.NewDialer(host,port,userName,authCode)
	err := d.DialAndSend(m)
	return err
}

func TestSendEmail(t *testing.T){
	NewCode().Send("1519327186@qq.com","验证码是2056")
}