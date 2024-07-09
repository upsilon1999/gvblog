package testdata

import (
	"fmt"
	"gvb_server/utils/jwt"
	"testing"
)
func TestJwt(t *testing.T) {
	token,err:=jwt.GenTokenforTest(jwt.JwtPayLoad{
		UserID: 1,
		Role: 1,
		Username: "upsilon",
		NickName: "lmryBC01",
	})

	fmt.Printf(token,err)
}