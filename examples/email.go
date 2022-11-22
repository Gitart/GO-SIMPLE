package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "al@meta.ua", "123", "www.meta.ua")


	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to  := []string{"arthur.savch@gmail.com"}
	msg := []byte("OOO")
	
	err := smtp.SendMail("smtp.meta.ua:465", auth, "alerc@meta.ua", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
