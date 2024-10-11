package models

import "time"

type Email struct {
	From    string    `json:"from"`
	To      string    `json:"to"`
	Subject string    `json:"subject"`
	Body    string    `json:"body"`
	IsHTML  bool      `json:"is_html"`
	SentAt  time.Time `json:"sent_at"`
}

func NewEmail() *Email {
	return &Email{
		From:    "",
		To:      "",
		Subject: "",
		Body:    "",
		IsHTML:  false,
		SentAt:  time.Now(),
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
