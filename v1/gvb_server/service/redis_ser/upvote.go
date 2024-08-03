package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

const upvotePrefix = "upvote"

// Upvote 点赞某一篇文章
func Upvote(id string) error {
	hasKey :=global.Redis.HExists(core.RedisCtx,upvotePrefix, id).Val()
	num := 0
	var err error
	if hasKey {
		// fmt.Println("是否进入")
		num, err = global.Redis.HGet(core.RedisCtx,upvotePrefix, id).Int()
		if err!=nil{
			logrus.Errorf("获取id错误,错误为%v",err)
			return err
		}
		num++
	}else{
		num = 1
	}
	// fmt.Printf("num值为%#v\n",num)
	err = global.Redis.HSet(core.RedisCtx,upvotePrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("设置id下点赞数出错,错误为%v",err)
		return err
	}
	return nil
}

// GetUpvote 获取某一篇文章下的点赞数
func GetUpvote(id string) int {
	num, err := global.Redis.HGet(core.RedisCtx,upvotePrefix, id).Int()
	if err!=nil{
		logrus.Errorf("获取点赞数出错,错误为%v",err)
		return 0
	}
	return num
}

// GetUpvoteInfo 取出点赞数据
func GetUpvoteInfo() map[string]int {
	var UpvoteInfo = map[string]int{}
	maps := global.Redis.HGetAll(core.RedisCtx,upvotePrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		UpvoteInfo[id] = num
	}
	return UpvoteInfo
}

//清除点赞数据
func UpvoteClear() {
	global.Redis.Del(core.RedisCtx,upvotePrefix)
}
