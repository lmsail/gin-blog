package Task

import (
	"gin-lmsail/app/Helpers"
	"gopkg.in/gomail.v2"
	"strconv"
)

// 注册发邮件
func SendRegisterMail(mailTo []string, subject, body string) {
	mailConfig := Helpers.ConfigMultiple("mail")
	mailConn := map[string]string{
		"user": mailConfig["MAIL_USER"],
		"pass": mailConfig["MAIL_PASSWORD"],
		"host": mailConfig["MAIL_HOST"],
		"port": mailConfig["MAIL_PORT"],
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", "XD Game"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                            //发送给多个用户
	m.SetHeader("Subject", subject)                         //设置邮件主题
	m.SetBody("text/html", body)                            //设置邮件正文
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	if err != nil {
		Helpers.Fatal("邮件发送失败", err)
	}
}
