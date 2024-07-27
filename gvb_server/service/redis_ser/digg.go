package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

const diggPrefix = "digg"

// Digg 点赞某一篇文章
func Digg(id string) error {
	num, err := global.Redis.HGet(core.RedisCtx,diggPrefix, id).Int()
	if err!=nil{
		logrus.Errorf("获取id错误,错误为%v",err)
		return err
	}
	num++
	err = global.Redis.HSet(core.RedisCtx,diggPrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("获取id下点赞数出错,错误为%v",err)
		return err
	}
	return nil
}

// GetDigg 获取某一篇文章下的点赞数
func GetDigg(id string) int {
	num, err := global.Redis.HGet(core.RedisCtx,diggPrefix, id).Int()
	if err!=nil{
		logrus.Errorf("获取点赞数出错,错误为%v",err)
		return 0
	}
	return num
}

// GetDiggInfo 取出点赞数据
func GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(core.RedisCtx,diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

//清除点赞数据
func DiggClear() {
	global.Redis.Del(core.RedisCtx,diggPrefix)
}
