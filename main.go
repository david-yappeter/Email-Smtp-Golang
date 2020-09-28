package main

import (
	"bytes"
	"log"
	"myapp/service"
	"myapp/template"
	"os"
)

func main() {

	to := []string{os.Getenv("RECEIVER_EMAIL")}
	// cc := []string{"cc@gmail.com"}
	subject := "testing subject"

	var buffer bytes.Buffer

	template.TestTemplate("dummy message", &buffer)

	message := buffer.String()

	err := service.SendMail(to, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
