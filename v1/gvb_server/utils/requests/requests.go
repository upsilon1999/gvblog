package requests

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type QueryForGet struct{
	ID string `json:"id"`
	Lang string `json:"lang"`
	Size int    `json:"size"`
}


//封装转发用的post请求
func Post(url string, data any, headers map[string]interface{}, timeout time.Duration) (body []byte, err error) {
	reqParam, _ := json.Marshal(data)
	reqBody := strings.NewReader(string(reqParam))
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) {
		case string:
		httpReq.Header.Add(key, v)
		case int:
		httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}
	client := http.Client{
		Timeout: timeout,
	}
	httpResp, err := client.Do(httpReq)
	if err != nil {
		return
	}
	// 客户端发起的请求必须在结束的时候关闭 response body
	// 这步是必要的，防止以后的内存泄漏，切记
    defer httpResp.Body.Close()
    body, err = io.ReadAll(httpResp.Body)
	return body, err
}


func Get(url string,urlQuery QueryForGet,headers map[string]interface{}, timeout time.Duration)(body []byte,err error){

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	
	for key, val := range headers {
		switch v := val.(type) {
		case string:
		httpReq.Header.Add(key, v)
		case int:
		httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}

	var query = httpReq.URL.Query()
	query.Add("id",urlQuery.ID)
	query.Add("lang",urlQuery.Lang)
	query.Add("size",strconv.Itoa(urlQuery.Size))
	// 增加请求参数
    httpReq.URL.RawQuery = query.Encode()


	client := http.Client{
		Timeout: timeout,
	}
	httpResp, err := client.Do(httpReq)

	if err != nil {
		return
	}
	// 客户端发起的请求必须在结束的时候关闭 response body
    defer httpResp.Body.Close()
    body, err = io.ReadAll(httpResp.Body)
	return body, err
}