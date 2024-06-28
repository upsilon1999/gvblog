package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

/*
go 1.16后废弃ioutil包，各种写法更新

ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.ReadDir -> os.ReadDir
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile

*/

//1.指定配置文件路径
//之所以写成全局变量是为了修改配置文件时也能直接使用
const ConfigFile = "settings.yaml"

// InitConf 读取yaml文件的配置
func InitConf() {
	//ps:因为日志的初始化要先读配置文件，所以这个方法内无法使用global.Log

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


//修改配置文件
func SetYmal() error {
	//读取修改后配置
	/*
		我们在执行修改的时候改变了global.Config全局对象
		这里所做的操作就是读取被修改后的global.Config全局对象
		然后将修改写入配置文件
	*/
	byteData,err := yaml.Marshal(global.Config)
	if err!=nil{
		// 虽然我们抛出了err,但这里也要打印,有助于错误定位
		global.Log.Error(err)
		return err
	}

	//写入配置
	//os.WriteFile(写入路路径，写入内容,权限标识)
	err = os.WriteFile(ConfigFile,byteData,fs.ModePerm)
	if err!=nil{
		global.Log.Error(err)
		return err
	}

	//写入成功的操作
	global.Log.Info("配置文件修改成功")
	return nil
}



