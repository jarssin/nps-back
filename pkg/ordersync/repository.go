package ordersync

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/database"
	"github.com/jarssin/nps-back/internal/infra/providers/cardapioweb"
)

type OrderSyncRepository struct {
	db *database.MongoDB
}

type OrderSyncRepositoryI interface {
	CreateOrder(ctx context.Context, order cardapioweb.OrderDetailsDTO) error
}

func NewOrderSyncRepository(db *database.MongoDB) OrderSyncRepositoryI {
	return &OrderSyncRepository{db: db}
}

func (r *OrderSyncRepository) CreateOrder(ctx context.Context, order cardapioweb.OrderDetailsDTO) error {
	_, err := r.db.Collection("orders").InsertOne(ctx, order)
	return err
}
