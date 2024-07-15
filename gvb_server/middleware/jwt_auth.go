package middleware

import (
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

//判断是否携带token的中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		//从redis中判断是否是注销的token
		 // 判断是否在redis中
		 ok,err := redis_ser.CheckLogout(token)
		 if err != nil {
				 global.Log.Error("读取redis失败",err)
				 res.FailWithMessage("读取redis失败", c)
				 c.Abort()
				 return
		 }
		 if ok{
			 res.FailWithMessage("token已注销", c)
			 c.Abort()
			 return
		 }
		// 登录的用户
		c.Set("claims", claims)
	}
}