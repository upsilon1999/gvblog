package redis_ser

import (
	"gvb_server/core"
	"gvb_server/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

const commentUpvotePrefix = "commentUpvote"

// commentUpvote评论某一篇文章
func CommentUpvote(id string) error {
	hasKey :=global.Redis.HExists(core.RedisCtx,commentUpvotePrefix, id).Val()
	num := 0
	var err error
	if hasKey {
		// fmt.Println("是否进入")
		num, err = global.Redis.HGet(core.RedisCtx,commentUpvotePrefix, id).Int()
		if err!=nil{
			logrus.Errorf("获取id错误,错误为%v",err)
			return err
		}
		num++
	}else{
		num = 1
	}
	// fmt.Printf("num值为%#v\n",num)
	err = global.Redis.HSet(core.RedisCtx,commentUpvotePrefix, id, num).Err()
	if err!=nil{
		logrus.Errorf("设置id下评论点赞数出错,错误为%v",err)
		return err
	}
	return nil
}

// GetUpvote 获取某一篇文章下的评论数
func GetCommentUpvote(id string) int {
	num, err := global.Redis.HGet(core.RedisCtx,commentUpvotePrefix, id).Int()
	if err!=nil{
		logrus.Errorf("获取评论点赞数出错,错误为%v",err)
		return 0
	}
	return num
}

// GetUpvoteInfo 取出评论点赞数据
func GetCommentUpvoteInfo() map[string]int {
	var commentUpvoteInfo = map[string]int{}
	maps := global.Redis.HGetAll(core.RedisCtx,commentUpvotePrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		commentUpvoteInfo[id] = num
	}
	return commentUpvoteInfo
}

//清除评论点赞数据
func CommentUpvoteClear() {
	global.Redis.Del(core.RedisCtx,commentUpvotePrefix)
}
