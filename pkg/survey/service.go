package survey

import (
	"errors"
	"fmt"

	"github.com/jarssin/nps-back/pkg/survey/csat"
	"github.com/jarssin/nps-back/pkg/survey/nps"
)

type SurveyServiceI[T any] interface {
	CreateSurvey(surveyData T) error
}

type SurveyService struct {
	npsService  SurveyServiceI[nps.DTO]
	csatService SurveyServiceI[csat.DTO]
}

func NewSurveyService(npsService SurveyServiceI[nps.DTO], csatService SurveyServiceI[csat.DTO]) *SurveyService {
	return &SurveyService{npsService, csatService}
}

func (s *SurveyService) CreateSurvey(surveyType string, payload any) (nps.DTO, csat.DTO, error) {
	switch surveyType {
	case "nps":
		npsResult, err := s.createNpsSurvey(payload)
		if err != nil {
			return nps.DTO{}, csat.DTO{}, fmt.Errorf("error creating survey: %v", err)
		}
		return npsResult, csat.DTO{}, nil
	case "csat":
		csatResult, err := s.createCsatSurvey(payload)
		if err != nil {
			return nps.DTO{}, csat.DTO{}, fmt.Errorf("error creating survey: %v", err)
		}
		return nps.DTO{}, csatResult, nil
	default:
		return nps.DTO{}, csat.DTO{}, fmt.Errorf("invalid survey type")
	}
}

func (s *SurveyService) createNpsSurvey(payload any) (nps.DTO, error) {
	npsPayload, ok := payload.(nps.ToCreateDTO)
	if !ok {
		return nps.DTO{}, fmt.Errorf("invalid NPS payload type")
	}

	npsDTO := npsPayload.ToDTO()

	if err := npsDTO.Validate(); err != nil {
		return nps.DTO{}, fmt.Errorf("validation error: %v", err)
	}

	if err := s.npsService.CreateSurvey(npsDTO); err != nil {
		return nps.DTO{}, errors.New("internal error creating NPS survey")
	}

	return npsDTO, nil
}

func (s *SurveyService) createCsatSurvey(payload any) (csat.DTO, error) {
	csatPayload, ok := payload.(csat.ToCreateDTO)

	if !ok {
		return csat.DTO{}, errors.New("invalid CSAT payload type")
	}

	csatDTO := csatPayload.ToDTO()

	if err := csatDTO.Validate(); err != nil {
		fmt.Println("Validation error:", err)
		return csat.DTO{}, fmt.Errorf("validation error: %v", err)
	}

	if err := s.csatService.CreateSurvey(csatDTO); err != nil {
		return csat.DTO{}, errors.New("internal error creating CSAT survey")
	}

	return csatDTO, nil
}
