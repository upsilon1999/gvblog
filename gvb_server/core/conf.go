package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// InitConf 读取yaml文件的配置
func InitConf() {
	//1.读取配置文件
	const ConfigFile = "settings.yaml"

	//关联到我们的配置文件结构体
	c := &config.Config{}

	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")

	// fmt.Println(c)
	//全局变量，就是将读取的到配置文件存储为全局变量
	global.Config = c
}

//ps:因为日志的初始化要先读配置文件，所以这里无法使用global.Log