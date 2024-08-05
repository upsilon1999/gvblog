package cron_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func Func1(){
	fmt.Println("func1",time.Now())
}
func TestCronTest(t *testing.T) {
	//创建cron实例，默认只支持到分钟
	//使用cron.WithSeconds()使得支持秒级
	Cron := cron.New(cron.WithSeconds())

	//创建携程，第一个参数是定时表达式
	//第二个参数是要执行的任务
	Cron.AddFunc("* * * * * *",Func1)

	//启动该实例中所有的定时任务
	Cron.Start()


	//阻塞主进程使得协程得以被执行，否则主进程结束，协程就被杀死了
	// select{}
}