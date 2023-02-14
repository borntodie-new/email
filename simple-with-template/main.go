package main

import (
	"bytes"
	"github.com/borntodie-new/email"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

func main() {
	from := email.Email
	tos := []string{"ccailatiao@gmail.com"} // 可以多个，一个切片
	pwd := email.Pwd                        // 不是登陆密码，类似app key
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", "发送邮件测试")

	data := struct {
		Username string
		Code     int
	}{
		Username: "Jason",
		Code:     123456,
	}

	t, err := template.ParseFiles("./template/email.html")
	if err != nil {
		log.Printf("parse file to template failed, and error is %s\n", err.Error())
		return
	}
	buf := bytes.NewBufferString("")
	if err := t.Execute(buf, data); err != nil {
		log.Printf("execute template faile, and error is %s\n", err.Error())
		return
	}
	m.SetBody("text/html", buf.String())
	d := gomail.NewDialer("smtp.qq.com", 465, from, pwd)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("send email to %s error, and error is %s\n", tos[0], err.Error())
		return
	}
	log.Println("send email success!")
}
