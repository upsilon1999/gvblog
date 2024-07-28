package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

const lookPrefix = "look"

// Look 浏览某一篇文章
func Look(id string) error {
	hasKey :=global.Redis.HExists(core.RedisCtx,lookPrefix, id).Val()
	num := 0
	var err error
	if hasKey {
		num, err = global.Redis.HGet(core.RedisCtx,lookPrefix, id).Int()
		if err!=nil{
			logrus.Errorf("获取id错误,错误为%v",err)
			return err
		}
		num++
	}else{
		num = 1
	}
	// fmt.Printf("num值为%#v\n",num)
	err = global.Redis.HSet(core.RedisCtx,lookPrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("设置id下浏览数出错,错误为%v",err)
		return err
	}
	return nil
}



// GetLook 获取某一篇文章下的浏览数
func GetLook(id string) int {
	num, _ := global.Redis.HGet(core.RedisCtx,lookPrefix, id).Int()
	return num
}

// GetLookInfo 取出浏览量数据
func GetLookInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(core.RedisCtx,lookPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

//清空redis缓存
func LookClear() {
	global.Redis.Del(core.RedisCtx,lookPrefix)
}