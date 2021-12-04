package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)
var Code = make(map[string]int) //key为Email,value为随机码

func SendEmail(email string) int {
	//smtp.PlainAuth()
	// 参数1：Usually identity should be the empty string, to act as username
	// 参数2：username
	//参数3：password
	//参数4：host
	auth := smtp.PlainAuth("", "18510122672@163.com", "此处填写邮箱密码", "smtp.163.com")
	to := []string{"1354428522@qq.com"}
	//发送随机数为验证码
	// Seed uses the provided seed value to initialize the default Source to a
	// deterministic state. If Seed is not called, the generator behaves as
	// if seeded by Seed(1). Seed values that have the same remainder when
	// divided by 2^31-1 generate the same pseudo-random sequence.
	// Seed, unlike the Rand.Seed method, is safe for concurrent use.
	rand.Seed(time.Now().Unix())
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	num := rand.Intn(10000)
	//发送内容使用base64 编码，单行不超过80字节，需要插入\r\n进行换行
	//The msg headers should usually include
	// fields such as "From", "To", "Subject", and "Cc".  Sending "Bcc"
	// messages is accomplished by including an email address in the to
	// parameter but not including it in the msg headers.
	str := fmt.Sprintf("From:18510122672@163.com\r\nTo:1354428522@qq.com\r\nSubject:verifycode\r\n\r\n,%d,\r\n.", num)//邮件格式
	msg := []byte(str)
	err := smtp.SendMail("smtp.163.com:25", auth, "18510122672@163.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

