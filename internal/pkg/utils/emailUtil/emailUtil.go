package emailUtil

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(userName, password, host, portStr, sendName, mailTo string, subject, body string) error {
	port, _ := strconv.Atoi(portStr)
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(userName, sendName))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(host, port, userName, password)
	err := d.DialAndSend(m)
	return err
}
