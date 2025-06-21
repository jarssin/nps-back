package nps

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SurveyService struct {
	surveyRepository SurveyRepositoryI
}

func NewSurveyService(surveyRepository SurveyRepositoryI) *SurveyService {
	return &SurveyService{surveyRepository: surveyRepository}
}

func (s *SurveyService) CreateSurvey(surveyData DTO) error {
	survey := DTO{
		Id:                 primitive.NewObjectID(),
		Comment:            surveyData.Comment,
		Score:              surveyData.Score,
		VisitorId:          surveyData.VisitorId,
		Phone:              surveyData.Phone,
		JourneyEvaluations: surveyData.JourneyEvaluations,
		CreatedAt:          time.Now(),
	}

	return s.surveyRepository.CreateSurvey(survey)
}
