package main

import (
	"github.com/borntodie-new/email"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	from := email.Email
	tos := []string{"ccailatiao@gmail.com"} // 可以多个，一个切片
	pwd := email.Pwd                        // 不是登陆密码，类似app key
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", "发送邮件测试")

	file, err := os.Open("template/email.html")
	if err != nil {
		log.Printf("open file failed, and error is %s\n", err.Error())
		return
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("read file context failed, and error is %s\n", err.Error())
		return
	}
	m.SetBody("text/html", string(body))
	d := gomail.NewDialer("smtp.qq.com", 465, from, pwd)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send email to %s error, and error is %s\n", tos[0], err.Error())
		return
	}
	log.Println("send email success!")
}
