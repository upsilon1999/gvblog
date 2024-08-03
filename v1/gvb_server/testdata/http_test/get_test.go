package http_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {

	url := "https://api.codelife.cc/api/top/list"

	req, _ := http.NewRequest("GET", url, nil)
	// 增加请求参数
    params := req.URL.Query()
    params.Add("id", "mproPpoq6O")
	params.Add("lang","cn")
    req.URL.RawQuery = params.Encode()

	response, err := http.DefaultClient.Do(req)
	if err != nil{
		fmt.Printf("错误为%#v\n",err)
	}
	// 客户端发起的请求必须在结束的时候关闭 response body
    defer response.Body.Close()
    body, err := io.ReadAll(response.Body)
	if err != nil{
		fmt.Printf("错误为%#v\n",err)
	}
	fmt.Printf("请求结果为%#v\n",string(body))
}