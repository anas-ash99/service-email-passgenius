package app

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"service-email-passgenius/api/models"
)

type EmailHandler struct {
	service EmailService
}

func NewEmailHandler(service EmailService) *EmailHandler {

	return &EmailHandler{service: service}
}

func (h *EmailHandler) RegisterRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.sendEmail)

	return r
}

func (h *EmailHandler) sendEmail(w http.ResponseWriter, r *http.Request) {
	email := models.NewEmail()
	var err error
	err = json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		log.Printf("Error decoding json body: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.sendEmail(email)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(email)
}
