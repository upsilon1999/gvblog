package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)



func TestRedisString(t *testing.T) {
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

	/*
	//Set设置
	//语法 redis实例.Set(ctx,键名,键值,过期时间段)
	//Set方法的最后一个参数表示过期时间，0表示永不过期
	err = rdb.Set(ctx, "key1", "value1", 0).Err()
	if err != nil {
		panic(err)
	}

	//key2将会在两分钟后过期失效
	err = rdb.Set(ctx, "key2", "value2", time.Minute * 2).Err()
	if err != nil {
		panic(err)
	}
	*/


	/*
	//SetNX设置
	//设置键的同时设置过期时间
	//语法 redis实例.SetEX(ctx,键名,键值,过期时间段)
	err = rdb.SetNX(ctx, "key3", "value", time.Hour * 2).Err()
	if err != nil {
		panic(err)
	}
	*/

	/*
	//Get获取
	//语法 redis实例.Get(ctx,键名)
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key1: %v\n", val)
	
	val2, err := rdb.Get(ctx, "key-not-exist").Result()
	if err == redis.Nil {
		fmt.Println("key不存在")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("值为: %v\n", val2)
	}
	*/

	//GetRange获取
	//语法 redis实例.GetRange(ctx,键名,开始下标,结束下标)
	val, err := rdb.GetRange(ctx, "key1",1,3).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("key1: %v\n", val)
}