package testdata

import (
	"fmt"
	"gvb_server/utils"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Printf(utils.HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	var str = utils.HashPwd("1234")
	fmt.Println(utils.CheckPwd(str,"1234"))
}