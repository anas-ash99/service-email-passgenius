package app

import (
	"fmt"
	"service-email-passgenius/api/models"
)

type EmailService struct {
}

func (s *EmailService) sendEmail(email *models.Email) error {

	if email.From == "" {
		return fmt.Errorf("from address must not be empty")
	}
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
