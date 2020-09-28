package service

import (
	"fmt"
	"net/smtp"
	"os"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

var CONFIG_EMAIL = os.Getenv("EMAIL_DUMMY") //Email MUST BE CONFIGURE LESS SECURE APP ACCESS TO TRUE
var CONFIG_PASSWORD = os.Getenv("EMAIL_DUMMY_PASS")

//SendMail Send Mail
func SendMail(to []string, subject string, htmlBody string) error {

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectSend := "Subject: " + subject + "\n"
	msg := []byte(subjectSend + mime + htmlBody)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, to, msg)
	if err != nil {
		return err
	}

	return nil
}
