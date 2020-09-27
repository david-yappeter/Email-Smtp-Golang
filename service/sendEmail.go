package service

import (
	"bytes"
	"fmt"
	"myapp/template"
	"net/smtp"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_EMAIL = "senderemail@gmail.com" //Email MUST BE CONFIGURE LESS SECURE APP ACCESS TO TRUE
const CONFIG_PASSWORD = "password"

//SendMail Send Mail
func SendMail(to []string, cc []string, subject, name string) error {

	auth := smtp.PlainAuth("", CONFIG_EMAIL, CONFIG_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	var buffered bytes.Buffer

	template.UserList(name, &buffered)

	byteString := buffered.String()

	// bufferByted := buffered.Bytes()

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectSend := "Subject: " + subject + "\n"
	msg := []byte(subjectSend + mime + byteString)
	// msg := bufferByted

	err := smtp.SendMail(smtpAddr, auth, CONFIG_EMAIL, append(to, cc...), msg)
	if err != nil {
		return err
	}

	return nil
}
