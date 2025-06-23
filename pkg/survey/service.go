package survey

import (
	"fmt"
	"log"
)

type SurveyServiceI interface {
	CreateSurvey(surveyData any) error
}

type SurveyService struct {
	npsService  SurveyServiceI
	csatService SurveyServiceI
}

func NewSurveyService(npsService SurveyServiceI, csatService SurveyServiceI) *SurveyService {
	return &SurveyService{npsService, csatService}
}

func (s *SurveyService) CreateSurvey(surveyType string, payload any) error {
	switch surveyType {
	case "nps":
		err := s.npsService.CreateSurvey(payload)
		if err != nil {
			log.Printf("error creating NPS survey: %v", err)
			return fmt.Errorf("error creating survey: %v", err)
		}
		return nil
	case "csat":
		err := s.csatService.CreateSurvey(payload)
		if err != nil {
			log.Printf("error creating CSAT survey: %v", err)
			return fmt.Errorf("error creating survey: %v", err)
		}
		return nil
	default:
		return fmt.Errorf("invalid survey type")
	}
}
