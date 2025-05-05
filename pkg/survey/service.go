package survey

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SurveyService struct {
	surveyRepository SurveyRepositoryI
}

type SurveyServiceI interface {
	CreateSurvey(survey DTO) error
}

func NewSurveyService(surveyRepository SurveyRepositoryI) SurveyServiceI {
	return &SurveyService{surveyRepository: surveyRepository}
}

func (s *SurveyService) CreateSurvey(surveyData DTO) error {
	survey := DTO{
		Id:        primitive.NewObjectID(),
		Comment:   surveyData.Comment,
		Score:     surveyData.Score,
		VisitorId: surveyData.VisitorId,
		CreatedAt: time.Now(),
	}

	return s.surveyRepository.CreateSurvey(survey)
}
