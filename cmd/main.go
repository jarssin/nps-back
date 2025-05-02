package main

import (
	"log"
	"net/http"
	"nps-back/functions"
)

func main() {
	http.HandleFunc("/survey", functions.CreateSurvey)

	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
