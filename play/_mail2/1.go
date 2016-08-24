package main

import (
	"log"
	"net/smtp"
)

func main() {
	send("hello there")
}

func send(body string) {
	from := "7atashfeshan@gmail.com"
	pass := "kxwz115hi"
	to := "webspycat@gmail.com"

	msg := "From: " + "webs@gdmail.com" + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent, visit http://foobarbazz.mailinator.com")
}
