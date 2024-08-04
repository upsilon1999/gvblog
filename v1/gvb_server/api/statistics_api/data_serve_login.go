package statistics_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"time"

	"github.com/gin-gonic/gin"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList  []string `json:"dateList"`
	LoginData []int    `json:"loginData"`
	SignData  []int    `json:"signData"`
}

func (StatisticsApi) SevenLogin(c *gin.Context) {
	var loginDateCount, signDateCount []DateCount

	// 按照天数，统计每天的登录人数，时间是七日内,
	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)
	//按照天数，统计每天的用户数，时间是七日内,
	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= created_at").
		Select("date_format(created_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)
	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int
	now := time.Now()
	for _, i2 := range loginDateCount {
		loginDateCountMap[i2.Date] = i2.Count
	}
	for _, i2 := range signDateCount {
		signDateCountMap[i2.Date] = i2.Count
	}

	// 构造时间列表
	var dateList []string
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02")
		loginCount := loginDateCountMap[day]
		signCount := signDateCountMap[day]
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}

	res.OkWithData(DateCountResponse{
		DateList:  dateList,
		LoginData: loginCountList,
		SignData:  signCountList,
	}, c)

}