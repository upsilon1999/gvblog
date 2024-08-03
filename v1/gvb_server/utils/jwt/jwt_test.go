package jwts

import (
	"fmt"
	"testing"
)
func TestJwt(t *testing.T) {
	//由于token需要配置文件，所以我们要提前加载
	//如果不加载会报错invalid memory address or nil pointer dereference
	// core.InitConf()
	// global.Log = core.InitLogger()
	token,err:=GenTokenforTest(JwtPayLoad{
		UserID: 1,
		Role: 1,
		Username: "upsilon",
		NickName: "lmryBC01",
	})

	fmt.Printf(token,err)

	
}