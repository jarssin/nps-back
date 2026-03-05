package ordersync

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
)

type ServiceI interface {
	SendSurveyToLastOrders(ctx context.Context) error
}

type ClientI interface {
	PullOrders(ctx context.Context, lastCallDate string) ([]cardapioweb.OrderDTO, error)
	OrderDetails(ctx context.Context, orderID int64) (cardapioweb.OrderDetailsDTO, error)
}
