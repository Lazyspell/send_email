package email

import (
	"fmt"
	"net/mail"
	"net/smtp"
	"os"
)

type InputExternalEmail struct {
	EmailName    string `json:"email_name" bson:"email_name"`
	EmailAddress string `json:"email_address" bson:"email_address"`
	EmailSubject string `json:"email_subject" bson:"email_subject"`
	EmailMessage string `json:"email_message" bson:"email_message"`
}

func ValidMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return false
	}
	return true
}

func SendEmail(input InputExternalEmail) (string, error) {
	host := "smtp.mail.yahoo.com"
	port := "587"

	toList := []string{os.Getenv("RECEIVING_EMAIL_ADDRESS")}

	emailAddress := input.EmailAddress
	if !ValidMailAddress(emailAddress) {
		return "Failed", nil
	}

	emailSubject := "Subject: " + input.EmailSubject + "\n"
	emailMessage := "\n" + input.EmailName + "\n" + emailAddress + "\n" + input.EmailMessage
	fromEmail := os.Getenv("YAHOO_EMAIL_ADDRESS")
	passwordEmail := os.Getenv("YAHOO_PASSWORD")

	message := []byte(emailSubject + emailMessage)

	auth := smtp.PlainAuth("", fromEmail, passwordEmail, host)
	err := smtp.SendMail(host+":"+port, auth, fromEmail, toList, message)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
		return "Failed", err
	}

	return "Success", err
}
