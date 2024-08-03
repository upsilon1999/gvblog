package randomavatar_test

import (
	"fmt"
	"image/png"
	"os"
	"path"
	"testing"
	"unicode/utf8"

	"github.com/DanPlayer/randomname"
	"github.com/disintegration/letteravatar"
	"github.com/golang/freetype"
)

//随机生成用户名加头像的组合，生成的头像在/uploads/chat_avatar目录下
//等到未来实装群聊功能就不用使用这种方式了

func TestCreateAvatar(t *testing.T){
	//调用在uploads/chat_avatar目录下生成文字头像
	//实际上我们也可以自己上传很多图片来代替，这里主要测试功能
	GenerateNameAvatar()
}

func GenerateNameAvatar() {
	dir := "../../uploads/chat_random_avatar"
	for _, s := range randomname.AdjectiveSlice {
	  DrawImage(string([]rune(s)[0]), dir)
	}
	for _, s := range randomname.PersonSlice {
	  DrawImage(string([]rune(s)[0]), dir)
	}
  }
  
func DrawImage(name string, dir string) {
	fontFile, err := os.ReadFile("../../uploads/fontTtf/HYShangWeiShouShuW.ttf")
	if err != nil {
		fmt.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontFile)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	options := &letteravatar.Options{
	  Font: font,
	}
	// 绘制文字
	firstLetter, _ := utf8.DecodeRuneInString(name)
	img, err := letteravatar.Draw(140, firstLetter, options)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	// 保存
	filePath := path.Join(dir, name+".png")
	file, err := os.Create(filePath)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	err = png.Encode(file, img)
	if err != nil {
	  fmt.Println(err)
	  return
	}
  }