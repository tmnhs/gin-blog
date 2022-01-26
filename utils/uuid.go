package utils

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)
//工具，生成唯一字符串（激活码）
func CreateUUID() string{
	id:=uuid.NewV4()
	return id.String()
}
//生成四位数验证码
func CreateVcode() string {
	codeHub:="1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	rand.Seed(time.Now().Unix())
	var code strings.Builder
	for i:=0;i<4;i++{
		index:=rand.Intn(len(codeHub))
		fmt.Fprintf(&code,"%c",codeHub[index])
	}
	fmt.Println(code.String())
	return code.String()
}

