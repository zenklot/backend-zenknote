package helper

import (
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(email string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", Config("CONFIG_AUTH_EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	port, _ := strconv.Atoi(Config("CONFIG_SMTP_PORT"))
	d := gomail.NewDialer(Config("CONFIG_SMTP_HOST"), port, Config("CONFIG_AUTH_EMAIL"), Config("CONFIG_AUTH_PASSWORD"))

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
