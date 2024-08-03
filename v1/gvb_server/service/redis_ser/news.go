package redis_ser

import (
	"encoding/json"
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"time"
)

const newsIndex = "news_index"

type NewsData struct {
	Index    string `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

// SetNews 设置某一个数据，重复执行，重复累加
func SetNews(key string, newData []NewsData) error {
	byteData, _ := json.Marshal(newData)
	err := global.Redis.Set(core.RedisCtx,fmt.Sprintf("%s_%s", newsIndex, key), byteData, 10*time.Second).Err()
	return err
}

func GetNews(key string) (newData []NewsData, err error) {
	res := global.Redis.Get(core.RedisCtx,fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newData)
	return
}