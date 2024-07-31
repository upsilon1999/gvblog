package article_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/es_ser"
	jwts "gvb_server/utils/jwt"
	"strings"
	"time"

	"math/rand"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

func (ArticleApi) ArticleCreateView(c *gin.Context){
	//登录后从token拿去用户数据
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName
	//对文章内容content的处理
	/*
		要进行xxs过滤:
		1.把前端传递来的markdown形式content转成html
		2.进行xxs过滤后再转回markdown，方便前端回显
	*/
	 // 处理content，把markdown转为html
	 unsafe := blackfriday.MarkdownCommon([]byte(cr.Content))
	
	 doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	 //fmt.Println(doc.Text())
	 // 是不是有script标签
	 nodes := doc.Find("script").Nodes
	 //如果有script标签就过滤掉然后再将过滤后结果转回markdown
	 if len(nodes) > 0 {
	   // 有script标签
	   doc.Find("script").Remove()
	   converter := md.NewConverter("", true, nil)
	   html, _ := doc.Html()
	   markdown, _ := converter.ConvertString(html)
	   cr.Content = markdown
	 }




	// 对文章简介的处理
	// 如果文章简介没传，就默认选文章内容的前30个字符
	if cr.Abstract == "" {
		// 汉字的截取不一样
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
		  cr.Abstract = string(abs[:100])
		} else {
		  cr.Abstract = string(abs)
		}
	  }

	// 不传banner_id,后台就随机去选择一张
	//当我们不传时将采用零值
	if cr.BannerID == 0 {
		var bannerIDList []uint
		global.DB.Model(models.BannerModel{}).Select("id").Scan(&bannerIDList)
		if len(bannerIDList) == 0 {
			res.FailWithMessage("没有banner数据", c)
			return
		}

		//go 1.22设立随机源
		//然后调用随机源上的方法
		r:=rand.New(rand.NewSource(time.Now().UnixNano()))
		// fmt.Println("随机数生成",r.Intn(20))
		cr.BannerID = bannerIDList[r.Intn(len(bannerIDList))]
	}
	// 查banner_id下的banner_url
	var bannerUrl string
	err = global.DB.Model(models.BannerModel{}).Where("id = ?", cr.BannerID).Select("path").Scan(&bannerUrl).Error
	if err != nil {
		res.FailWithMessage("banner不存在", c)
		return
	}

	// 查用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}


	//发布文章
	//获取当前时间并格式化
	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        cr.Title,
		Keyword: cr.Title,//这一步是必要的，否则es中keyword就无值
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		BannerID:     cr.BannerID,
		BannerUrl:    bannerUrl,
		Tags:         cr.Tags,
	}

	//在创建es文章数据前，对keyword进行判断
	if article.ISExistTitle(){
		res.FailWithMessage("文章已存在",c)
		return
	}

	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}


	//在创建文章时使用协程创建全文搜索
	go es_ser.AysncFullText(article.ID,article.Title,article.Content)
	res.OkWithMessage("文章发布成功", c)
}