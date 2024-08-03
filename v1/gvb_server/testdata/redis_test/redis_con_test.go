package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)



func TestRedis(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("连接redis出错,错误信息：%v", err)
		return
	}
	fmt.Println("成功连接redis", pong)

}