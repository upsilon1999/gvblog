package rand_test

import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)

//由于缓存的存在，直接运行将不生效
//可以来到此目录下执行 go test -v -count=1 TestRandom
func TestRandom(t *testing.T) {
	//go 1.22设立随机源
	//设立随机种子，这里采用秒级时间戳
	src := rand.NewSource(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	//注册种子生成实例
    r := rand.New(src)
	//获取[0,100]的随机整数
    number := r.Intn(100)
    fmt.Println(number)
}