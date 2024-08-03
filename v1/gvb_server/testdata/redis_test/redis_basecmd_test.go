package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/redis/go-redis/v9"
)




func TestBaseCmd(t *testing.T) {
	//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
	//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
	var ctx = context.Background()
	//连接redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",//连接地址
		Password: "123456789", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})


	//*表示获取所有的key
	//得到的时一个切片
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)


	/*
	//`Type()`方法用户获取一个key对应值的类型
	//语法 redis实例.Type(ctx,键名)
	//返回值 对应键值的类型
	vType, err := rdb.Type(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vType) //string
	*/
	
	/*
	//Del()删除缓存项，
	//语法 redis实例.Del(ctx,键名1,键名2)
	//返回值 删除缓存的项数,错误
	n, err := rdb.Del(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("成功删除了 %v 个\n", n)
	*/


	/*
	//Exists():检测缓存项是否存在
	//语法 redis实例.Exists(ctx,键名1,键名2,...).Result()
	n, err := rdb.Exists(ctx, "age").Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
	*/



	/*
	//设置有效时间段
	//在某个时间段后过期，例如 2分钟后过期
	//语法 redis实例.Expire(ctx,键名1,有效期).Result()
	res, err := rdb.Expire(ctx, "name", time.Minute * 2).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
	
	//设置有效时间点
	//在某个时间点过期，例如 现在过期
	//语法 redis实例.Expire(ctx,键名1,过期时间点).Result()
	res, err = rdb.ExpireAt(ctx, "age", time.Now()).Result()
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("设置成功")
	} else {
		fmt.Println("设置失败")
	}
	*/


	/*
	//获取剩余有效期,单位:秒(s)
	//语法 redis实例.TTL(ctx,键名)
	ttl, err := rdb.TTL(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(ttl)

	//获取剩余有效期,单位:毫秒(ms)
	//语法 redis实例.PTTL(ctx,键名)
	pttl, err := rdb.PTTL(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pttl)
	*/


	/*
	//查看当前数据库key的数量
	//语法: redis实例.DBSize(ctx)
	num, err := rdb.DBSize(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("数据库有 %v 个缓存项\n", num)
	*/

	/*
	 //清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
	 //语法 redis实例.FlushDB(ctx)
	 res, err := rdb.FlushDB(ctx).Result()
	 if err != nil {
		 panic(err)
	 }
	 fmt.Println(res)//OK
	*/


	/*
	//清空该连接上所有数据库
	// 语法 redis实例.FlushAll(ctx)
	res, err := rdb.FlushAll(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)//OK
	*/
}