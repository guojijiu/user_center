package config

import "os"

type mailConfig struct {
	MailHost     string
	MailPort     string
	MailUser     string
	MailPwd      string
	MailFromName string
}

var MailConfig = mailConfig{
	MailHost:     os.Getenv("MAIL_HOST"),
	MailPort:     os.Getenv("MAIL_PORT"),
	MailUser:     os.Getenv("MAIL_USER"),
	MailPwd:      os.Getenv("MAIL_PWS"),
	MailFromName: os.Getenv("MAIL_FROM_NAME"),
}
