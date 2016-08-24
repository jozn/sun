package main

import (
    "log"
    "net/smtp"
)

func main() {
    // Set up authentication information.
    auth := smtp.PlainAuth(
        "",
        "7atashfeshan@gmail.com",
        "kxwz115hi",
        "smtp.gmail.com",
    )
    // Connect to the server, authenticate, set the sender and recipient,
    // and send the email all in one step.
    err := smtp.SendMail(
        "smtp.gmail.com:587",
        auth,
        "7atashfeshan@gmail.com",
        []string{"webspycat@gmail.com"},
        []byte("This is the email body."),
    )
    if err != nil {
        log.Fatal(err)
    }
}