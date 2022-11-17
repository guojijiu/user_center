package mail

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"user_center/config"
)

// SendGoMail /*
func SendGoMail(mailAddress []string, subject string, body string) error {
	m := gomail.NewMessage()
	mailConf := config.MailConfig
	mailPort, _ := strconv.Atoi(mailConf.MailPort)
	// 这种方式可以添加别名，即 nickname， 也可以直接用<code>m.SetHeader("From", MAIL_USER)</code>
	nickname := mailConf.MailFromName
	m.SetHeader("From", nickname+"<"+mailConf.MailUser+">")
	// 发送给多个用户
	m.SetHeader("To", mailAddress...)
	// 设置邮件主题
	m.SetHeader("Subject", subject)
	// 设置邮件正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(mailConf.MailHost, mailPort, mailConf.MailUser, mailConf.MailPwd)
	// 发送邮件
	err := d.DialAndSend(m)
	return err
}
