package main

import (
	"log"
	"net/http"

	functions "github.com/jarssin/nps-back"
	"github.com/jarssin/nps-back/internal/infra/middlewares"
)

func main() {
	http.HandleFunc("/create-survey", middlewares.CorsMiddleware(functions.CreateSurvey))
	http.HandleFunc("/send-survey", middlewares.CorsMiddleware(functions.SendSurvey))

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
