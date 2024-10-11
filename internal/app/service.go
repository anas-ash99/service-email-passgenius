package app

import (
	"fmt"
	"net/smtp"
	"os"
	"service-email-passgenius/api/models"
)

type EmailService struct {
}

func (s *EmailService) sendEmail(email *models.Email) error {
	var err error
	err = validateBody(email)
	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	username := "anas.ash099@gmail.com"
	password := os.Getenv("PASSGENIUS-GMAIL-APP-PASSWORD")
	fmt.Println(password)

	to := []string{email.To}
	subject := "Subject: " + email.Subject + "\n"
	// Compose the email
	message := []byte(subject + "\n" + email.Body)
	// Set up authentication
	auth := smtp.PlainAuth("", username, password, smtpHost)

	// Send the email
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, username, to, message)

	return err
}

func validateBody(email *models.Email) error {
	if email.To == "" {
		return fmt.Errorf("to address must not be empty")
	}
	if email.Subject == "" {
		return fmt.Errorf("subject must not be empty")
	}
	if email.Body == "" {
		return fmt.Errorf("body must not be empty")
	}
	return nil
}
