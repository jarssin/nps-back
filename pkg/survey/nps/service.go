package nps

import (
	"fmt"
	"log"
)

type SurveyService struct {
	surveyRepository SurveyRepositoryI
}

func NewSurveyService(surveyRepository SurveyRepositoryI) *SurveyService {
	return &SurveyService{surveyRepository: surveyRepository}
}

func (s *SurveyService) CreateSurvey(npsPayload any) error {
	npsToCreate, ok := npsPayload.(ToCreateDTO)
	if !ok {
		log.Println("Invalid NPS payload type")
		return fmt.Errorf("invalid NPS payload type")
	}

	npsDTO := npsToCreate.ToDTO()
	if err := npsDTO.Validate(); err != nil {
		log.Println("Validation error:", err)
		return fmt.Errorf("validation error: %v", err)
	}

	return s.surveyRepository.CreateSurvey(npsDTO)
}
