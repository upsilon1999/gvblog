package utils

import "strings"

// 手机号脱敏
//就是将手机号加入星号返回给前端显示
func DesensitizationTel(tel string)string  {
	//这里还可以加入区号的操作，主要看数据库里如何存储
	/*
		脱敏指的是我们把数据从数据库拿出来后，
		对于敏感数据要进行混淆处理再返回给前端
	*/


	//例如正常手机号 18825540000
	//脱敏后为 188 **** 0000
	
	//先判断手机号位数，如果不是11位就返回空
	if len(tel) != 11{
		return ""
	}

	// 然后对手机号进行操作
	return tel[:3]+"****"+tel[7:]
}

func DesensitizationEmail(email string)  string{
	//根据个人需求，我们的目标保留首字母和邮箱后缀
	//12457@qq.com== 1*****@qq.com
	elist := strings.Split(email, "@")

	//如果有一部分不存在
	if len(elist) != 2{
		return ""
	}

	return elist[0][:1]+"****@"+elist[1]

}