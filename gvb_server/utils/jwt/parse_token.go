package jwts

import (
	"errors"
	"fmt"
	"gvb_server/global"

	"github.com/dgrijalva/jwt-go/v4"
)

// ParseToken 解析 token
  func ParseToken(tokenStr string) (*CustomClaims, error) {
	//传入token需要的密钥，要与生成token使用的一样
	MySecret := []byte(global.Config.Jwt.Secret)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	  return MySecret, nil
	})
	if err != nil {
	  global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
	  return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
	  return claims, nil
	}
	return nil, errors.New("invalid token")
  }

  func ParseTokenForTest(tokenStr string) (*CustomClaims, error) {
	MySecret := []byte("upsilon")
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	  return MySecret, nil
	})
	if err != nil {
	  global.Log.Error(fmt.Sprintf("token parse err: %s", err.Error()))
	  return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
	  return claims, nil
	}
	return nil, errors.New("invalid token")
  }