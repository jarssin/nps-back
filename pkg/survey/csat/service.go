package csat

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

func (s *SurveyService) CreateSurvey(csatPayload any) error {
	csatToCreate, ok := csatPayload.(ToCreateDTO)
	if !ok {
		log.Println("Invalid CSAT payload type")
		return fmt.Errorf("invalid CSAT payload type")
	}

	csatDTO := csatToCreate.ToDTO()
	if err := csatDTO.Validate(); err != nil {
		log.Println("Validation error:", err)
		return fmt.Errorf("validation error: %v", err)
	}

	return s.surveyRepository.CreateSurvey(csatDTO)
}
