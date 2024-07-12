package middleware

import (
	"gvb_server/models/res"
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
		// 登录的用户
		c.Set("claims", claims)
	}
}