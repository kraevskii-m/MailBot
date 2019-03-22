package main

import (
	"log"
	"net/smtp"
)

func mailSender(to string, body string, subject string) {
	from := "fatherofbots@yandex.ru"
	pass := "lermonter07"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject:" + subject + "\n" +
		body

	err := smtp.SendMail("smtp.yandex.ru:465",
		smtp.PlainAuth("", from, pass, "smtp.yandex.ru"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("Message sent!")
}
