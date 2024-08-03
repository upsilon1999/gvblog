package htmlandmdtest

import (
	"fmt"
	"strings"
	"testing"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
)

//这个文件用于测试markdown与html互转

func TestMdToHtml(t *testing.T){
	//将md转html
	unsafe := blackfriday.MarkdownCommon([]byte("### 你好\n ```python\nprint('你好')\n```\n - 123 \n \n<script>alert(123)</script>\n\n ![图片](http://xxx.com)"))
	fmt.Println(string(unsafe))

	//将html转为可读文本格式
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	//移除script标签，过滤xxs
	doc.Find("script").Remove()
	fmt.Println(doc.Text())


	//html转md
	converter := md.NewConverter("", true, nil)
	html, _ := doc.Html()
	markdown, err := converter.ConvertString(html)
	fmt.Println(markdown, err)
}