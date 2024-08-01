package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

const commentPrefix = "comment"

// Comment评论某一篇文章
func Comment(id string) error {
	hasKey :=global.Redis.HExists(core.RedisCtx,commentPrefix, id).Val()
	num := 0
	var err error
	if hasKey {
		// fmt.Println("是否进入")
		num, err = global.Redis.HGet(core.RedisCtx,commentPrefix, id).Int()
		if err!=nil{
			logrus.Errorf("获取id错误,错误为%v",err)
			return err
		}
		num++
	}else{
		num = 1
	}
	// fmt.Printf("num值为%#v\n",num)
	err = global.Redis.HSet(core.RedisCtx,commentPrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("设置id下评论数出错,错误为%v",err)
		return err
	}
	return nil
}
// CommentCount 评论数计算
//这里的调用时该评论数我们默认是大于等于1的
func CommentCount(id string,count int) error {
	hasKey :=global.Redis.HExists(core.RedisCtx,commentPrefix, id).Val()
	num := 0
	var err error
	if hasKey {
		// fmt.Println("是否进入")
		num, err = global.Redis.HGet(core.RedisCtx,commentPrefix, id).Int()
		if err!=nil{
			logrus.Errorf("获取id错误,错误为%v",err)
			return err
		}
	}
	num+=count
	// fmt.Printf("num值为%#v\n",num)
	err = global.Redis.HSet(core.RedisCtx,commentPrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("设置id下评论数出错,错误为%v",err)
		return err
	}
	return nil
}

// GetUpvote 获取某一篇文章下的评论数
func GetComment(id string) int {
	num, err := global.Redis.HGet(core.RedisCtx,commentPrefix, id).Int()
	if err!=nil{
		logrus.Errorf("获取评论数出错,错误为%v",err)
		return 0
	}
	return num
}

// GetUpvoteInfo 取出评论数据
func GetCommentInfo() map[string]int {
	var CommentInfo = map[string]int{}
	maps := global.Redis.HGetAll(core.RedisCtx,commentPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		CommentInfo[id] = num
	}
	return CommentInfo
}

//清除评论数据
func CommentClear() {
	global.Redis.Del(core.RedisCtx,commentPrefix)
}
