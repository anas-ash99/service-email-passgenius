package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"service-email-passgenius/internal/app"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	emailService := app.EmailService{}
	emailHandler := app.NewEmailHandler(emailService)
	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World!"))

	})
	r.Mount("/", emailHandler.RegisterRoutes())
	// Start the HTTP server
	fmt.Println("Starting server on port 8080...")
	http.ListenAndServe(":8080", r)
}
