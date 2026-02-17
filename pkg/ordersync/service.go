package ordersync

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/database"
)

type Service struct {
	client ClientI
	db     *database.MongoDB
}

func NewService(client ClientI, db *database.MongoDB) ServiceI {
	return &Service{client, db}
}

func (s *Service) SendSurveyToLastOrders(ctx context.Context) error {
	return nil
}
