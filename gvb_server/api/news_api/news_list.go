package news_api

import (
	"encoding/json"
	"fmt"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/requests"
	"time"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

type NewsData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

type NewsResponse struct {
	Code int       `json:"code"`
	Data []redis_ser.NewsData `json:"data"`
	Msg  string    `json:"msg"`
}

//get请求对应结构体
type QueryForGet struct{
	ID string `json:"id"`
	Lang string `json:"lang"`
	Size int    `json:"size"`
}
const newAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

func (NewsApi) NewListPostView(c *gin.Context) {
	var cr params
	var headers header
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("json解析出错", c)
		return
	}
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Size == 0 {
		cr.Size = 1
	}
	httpResponse, err := requests.Post(newAPI, cr, structs.Map(headers), timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	var response NewsResponse
	err = json.Unmarshal(httpResponse, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	res.OkWithData(response.Data, c)
	// return
}

func (NewsApi) NewListGetView(c *gin.Context){
	var cr QueryForGet

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMessage("json解析出错", c)
		return
	}

	if cr.Size == 0 {
		cr.Size = 1
	}

	//如果已经有新闻了就直接走缓存
	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	newsData, _ := redis_ser.GetNews(key)
	if len(newsData) != 0 {
	  res.OkWithData(newsData, c)
	  return
	}


	//没有新闻缓存再调用接口
	var headers map[string]interface{}
	httpResponse, err := requests.Get(newAPI, requests.QueryForGet(cr), headers, timeout)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	var response NewsResponse
	err = json.Unmarshal(httpResponse, &response)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMessage(response.Msg, c)
		return
	}
	res.OkWithData(response.Data, c)
	
	redis_ser.SetNews(key, response.Data)
}