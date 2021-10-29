package model

import (
	"demo1/MyBlog/utils"
	"demo1/MyBlog/utils/errmsg"
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmail(to string,title string,text string) int{
	m := gomail.NewMessage()

	m.SetHeader("From", utils.FromEmail)
	m.SetHeader("To", to)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", title)
	m.SetBody("text/html", text)
	//m.Attach("/home/Alex/lolcat.jpg")
	serverPort,_:= strconv.Atoi(utils.ServerPort)
	d := gomail.NewDialer(utils.ServerHost, serverPort , "1685290935@qq.com", utils.FromPassword)

	if err := d.DialAndSend(m); err != nil {
		return errmsg.ERROR_EMAIL_SEND
	}
	fmt.Println("发送结束")
	return errmsg.SUCCESS
}



