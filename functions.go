package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"

	"github.com/Jardessomonster/nps-back/internal/infra/database"
	"github.com/Jardessomonster/nps-back/internal/infra/middlewares"
	"github.com/Jardessomonster/nps-back/pkg/survey"
)

var (
	surveyService survey.SurveyServiceI
)

func init() {
	functions.HTTP("CreateSurvey", middlewares.CorsMiddleware(CreateSurvey))

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("MongoDB: ", os.Getenv("MONGODB_URL"))
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
		http.Error(w, fmt.Sprintf("Invalid body: %v", err), http.StatusBadRequest)
		return
	}

	if err := payload.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
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
