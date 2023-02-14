package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	from := "发件人邮箱号"
	tos := []string{"收件人邮箱号"} // 可以多个，一个切片
	pwd := "发件人邮箱密码" // 不是登陆密码，类似app key
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", "发送邮件测试")
	m.SetBody("text/html", fmt.Sprintf("<h1>hi: %s</h1>", tos[0]))
	d := gomail.NewDialer("smtp.qq.com", 465, from, pwd)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send email to %s error, and error is %s\n", tos[0], err.Error())
		return
	}
	log.Println("send email success!")
}
