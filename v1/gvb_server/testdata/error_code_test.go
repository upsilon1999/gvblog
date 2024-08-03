package testdata

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

//文件路径
const filePath = "gvb_server\\models\\res\\err_code.json"

type ErrMap map[string]string
func TestErrorCode(t *testing.T)  {
	byteData,err := os.ReadFile(filePath)
	if(err!=nil){
		logrus.Error(err)
		return 
	}

	var errMap = ErrMap{}
	// 解析json,并赋予实例对象
	err = json.Unmarshal(byteData,&errMap)
	if(err!=nil){
		logrus.Error(err)
		return 
	}
	fmt.Println(errMap)

}