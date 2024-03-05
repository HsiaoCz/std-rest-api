package main

import (
	"log"
	"net/smtp"
)

func SendMailSimple() {
	auth := smtp.PlainAuth(
		"",
		// from the email address
		"984274788@qq.com",
		"123456",
		"smtp.gmail.com",
	)
	message := []byte("hello this is the message!")
	if err := smtp.SendMail("smtp.gmail.com:587", auth, "984274788@qq.com", []string{"1299720482@qq.com"}, message); err != nil {
		log.Fatal(err)
	}
}
