package user_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

func (UserApi) LogoutView(c *gin.Context) {
	//1.只有登录了之后才能登出
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	//2.我们要通过这个来拿到token过期时间
	//这里返回的是具体时间点，而redis要设置的是time.Duration,也就是时间段
	//所以我们要计算当前到截止时间的耗时
	fmt.Println(claims.ExpiresAt)

	//获取要注销的token
	token := c.Request.Header.Get("token")
	err := service.ServiceApp.UserService.Logout(claims, token)

	if err != nil {
		global.Log.Error("写入redis失败",err)
		res.FailWithMessage("注销失败", c)
		return
	}
	res.OkWithMessage("注销成功", c)

}