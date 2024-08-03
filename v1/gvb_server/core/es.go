package core

import (
	"gvb_server/global"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func EsConnect() {
	var err error
	//是否开启集群监听，单机模式需要关闭
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		//读取封装好的Es配置
		elastic.SetURL(global.Config.Es.URL()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.Es.Username, global.Config.Es.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	//连接成功后将es实例赋值给全局的变量
	global.ESClient = c
}