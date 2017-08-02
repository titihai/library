package util

import "github.com/go-gomail/gomail"

var EmailFrom string
var EmailServer string
var EmailPort int
var EmailAccount string
var EmailPassword string

func SendEmail(to string, content string, subject string) error {
	var m = gomail.NewMessage()
	m.SetHeader("From", EmailFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)

	var d = gomail.NewDialer(EmailServer, EmailPort, EmailAccount, EmailPassword)

	var err = d.DialAndSend(m)
	if err != nil {
		//util.Logger.Debug(err)
	}
	return err
}
