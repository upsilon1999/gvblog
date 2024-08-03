package core

import (
	"context"
	"gvb_server/global"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var RedisCtx = context.Background()

func ConnectRedis() *redis.Client{
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client{
	
	redisConf :=global.Config.Redis 
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),//连接地址
		Password: redisConf.Password, // 没有密码，默认值
		DB:       db,  // 默认DB 0
		PoolSize: redisConf.PoolSize,//连接池大小
	})

	_,cancel := context.WithTimeout(RedisCtx,500*time.Millisecond)
	defer cancel()

	_, err := rdb.Ping(RedisCtx).Result()
	if err != nil {
		logrus.Errorf("redis连接失败%s",redisConf.Addr())
		return nil
	}
	return rdb
}