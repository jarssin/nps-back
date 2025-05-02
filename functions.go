package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nps-back/internal/infra/database"
	"nps-back/pkg/survey"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	surveyService survey.SurveyServiceI
)

func init() {
	err := godotenv.Load()
	fmt.Println("MongoDB: ", os.Getenv("MONGODB_URL"))
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	connection, err := database.Connect()

	if err != nil {
		fmt.Printf("failed to create MongoDB client: %e", err)
	}

	surveyRepository := survey.NewSurveyRepository(connection)
	surveyService = survey.NewSurveyService(surveyRepository)
}

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload survey.DTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	if err := surveyService.CreateSurvey(payload); err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Survey created successfully",
	})
}
