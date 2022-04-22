package emailStr

import (
	"errors"

	"gopkg.in/gomail.v2"
)

//发邮件
func SendMail(emailBody, stmpKey, formatAddress, emailHeader, emailSecondHeader string, toAddress []string) error {
	if emailBody == "" || stmpKey == "" || formatAddress == "" || toAddress == nil {
		return errors.New("parameters are missing")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(formatAddress, emailHeader))
	m.SetHeader("To", toAddress...)
	m.SetHeader("Subject", emailSecondHeader)
	m.SetBody("text/html", emailBody)
	d := gomail.NewDialer("smtp.163.com", 465, formatAddress, stmpKey)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
