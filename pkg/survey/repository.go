package survey

import (
	"context"
	"nps-back/internal/infra/database"
)

type SurveyRepository struct {
	db *database.MongoDB
}

type SurveyRepositoryI interface {
	CreateSurvey(survey DTO) error
}

func NewSurveyRepository(db *database.MongoDB) SurveyRepositoryI {
	return &SurveyRepository{db: db}
}

func (r *SurveyRepository) CreateSurvey(survey DTO) error {
	_, err := r.db.Collection("surveys").InsertOne(context.Background(), survey)
	return err
}
