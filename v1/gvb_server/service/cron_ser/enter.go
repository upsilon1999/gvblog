package cron_ser

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func CronInit() {
	timezone, err := time.LoadLocation("Asia/Beijing")

	if err!=nil{
		logrus.Error(err.Error())
		return
	}

	//第一个参数支持秒级，第二个参数设定时区
	Cron := cron.New(cron.WithSeconds(),cron.WithLocation(timezone))
	//在每日的0点0分0秒同步文章数据到es
	Cron.AddFunc("0 0 0 * * *",SyncArticleData)
	//在每日的0点0分0秒同步评论点赞数据到mysql
	Cron.AddFunc("0 0 0 * * *",SyncCommentData)
	Cron.Start()

	//为什么不用阻塞
	/*
		demo中使用阻塞的原因:主进程很快就走完了，不阻塞协程不会执行

		真实项目中
		我们在main函数中启动定时任务，
		由于项目启动后会一直开着主进程，所以不需要阻塞
	*/
}