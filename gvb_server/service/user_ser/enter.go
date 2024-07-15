package user_ser

import (
	"gvb_server/service/redis_ser"
	jwts "gvb_server/utils/jwt"
	"time"
)

type UserService struct {
}

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	//Sub方法返回的就是Duration类型
	//时间点A.Time.Sub(时间点B) 得到两个时间点的时间差
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}