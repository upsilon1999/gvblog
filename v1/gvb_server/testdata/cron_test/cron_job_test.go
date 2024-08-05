package cron_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct {
}

func (Job) Run() {
	fmt.Println("使用鸭子类型实现接口",time.Now())
}
func TestCronJob(t *testing.T) {
	//创建cron实例，默认只支持到分钟
	//使用cron.WithSeconds()使得支持秒级
	Cron := cron.New(cron.WithSeconds())

	//创建携程，第一个参数是定时表达式
	//第二个参数job,根据源码我们用一个结构体来实现这个接口
	Cron.AddJob("* * * * * *", Job{})

	//启动该实例中所有的定时任务
	Cron.Start()

	//阻塞主进程使得协程得以被执行，否则主进程结束，协程就被杀死了
	// select{}
}