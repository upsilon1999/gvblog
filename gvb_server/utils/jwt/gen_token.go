package jwt

import (
	"fmt"
	"gvb_server/global"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	fmt.Println(global.Config.Jwt)
	MySecret := []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
	  user,
	  jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires))), // 默认2小时过期
		Issuer:    global.Config.Jwt.Issuer,                                                     // 签发人
	  },
	}
  
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
  }
  

  //配合测试使用
  func GenTokenforTest(user JwtPayLoad) (string, error) {
	MySecret := []byte("upsilon")
	claim := CustomClaims{
	  user,
	  jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(48))), // 默认2小时过期
		Issuer:    "zgry",                                                     // 签发人
	  },
	}
  
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
  }