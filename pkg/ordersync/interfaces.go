package ordersync

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
)

type ServiceI interface {
	SendSurveyToLastOrders(ctx context.Context) error
}

type ClientI interface {
	PullOrders(lastCallDate string) ([]cardapioweb.OrderDTO, error)
	OrderDetails(orderID string) (cardapioweb.OrderDetailsDTO, error)
}
