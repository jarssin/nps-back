package functions

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/joho/godotenv"

	"github.com/jarssin/nps-back/internal/infra/database"
	"github.com/jarssin/nps-back/internal/infra/middlewares"
	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
	"github.com/jarssin/nps-back/pkg/mrrobot"
	"github.com/jarssin/nps-back/pkg/ordersync"
	"github.com/jarssin/nps-back/pkg/person"
	"github.com/jarssin/nps-back/pkg/survey"
	"github.com/jarssin/nps-back/pkg/survey/csat"
	"github.com/jarssin/nps-back/pkg/survey/nps"
)

var (
	npsService         survey.SurveyServiceI
	csatService        survey.SurveyServiceI
	personService      person.ServiceI
	cardapioWebService ordersync.ServiceI
)

func init() {
	functions.HTTP("CreateSurvey", middlewares.CorsMiddleware(CreateSurvey))

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	log.Println("MongoDB: ", os.Getenv("MONGODB_URL"))
	log.Println("MongoDB Database: ", os.Getenv("MONGODB_DATABASE"))
	database, err := database.Connect()

	if err != nil {
		log.Printf("failed to create MongoDB client: %e", err)
	}

	htppClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		},
	}

	npsRepository := nps.NewSurveyRepository(database)
	npsService = nps.NewSurveyService(npsRepository)

	csatRepository := csat.NewSurveyRepository(database)
	csatService = csat.NewSurveyService(csatRepository)

	mrRobotClient := mrrobot.NewMrRobotClient(htppClient, os.Getenv("MRROBOT_BASE_URL"))
	personService = person.NewService(mrRobotClient)

	cardapiowebClient := cardapioweb.NewClient(htppClient, os.Getenv("CARDAPIOWEB_BASE_URL"), os.Getenv("CARDAPIOWEB_TOKEN"))
	cardapioWebService = ordersync.NewService(cardapiowebClient, database)
}

func CreateSurvey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var npsPayload nps.ToCreateDTO
	var csatPayload csat.ToCreateDTO

	params := r.URL.Query()
	surveyType := params.Get("type")
	log.Printf("Received survey type: %s\n", surveyType)

	surveyService := survey.NewSurveyService(npsService, csatService)

	switch surveyType {
	case "nps", "":
		if err := json.NewDecoder(r.Body).Decode(&npsPayload); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
			return
		}

		if err := surveyService.CreateSurvey(surveyType, npsPayload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "csat":
		if err := json.NewDecoder(r.Body).Decode(&csatPayload); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
			return
		}

		if err := surveyService.CreateSurvey(surveyType, csatPayload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	log.Printf("Survey of type %s created successfully\n", surveyType)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Survey created successfully",
	})
}

func SendSurvey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	people, err := person.ToDTOFromCSV(csvReader)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing CSV: %v", err), http.StatusBadRequest)
		return
	}

	chunkSize := 100
	for i := 0; i < len(people); i += chunkSize {
		end := min(i+chunkSize, len(people))
		go personService.SendSurvey(people[i:end])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

func CardapioWebSubscriber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()
	err := cardapioWebService.SendSurveyToLastOrders(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending surveys: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "CardapioWeb orders pulled successfully",
	})
}
