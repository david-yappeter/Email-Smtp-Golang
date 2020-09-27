package main

import (
	"log"
	"myapp/service"
)

func main() {

	to := []string{"receiver@gmail.com"}
	cc := []string{"cc@gmail.com"}
	subject := "testing subject"
	message := "Testing 1"

	err := service.SendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}
