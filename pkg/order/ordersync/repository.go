package ordersync

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderSyncRepository struct {
	db *database.MongoDB
}

type OrderSyncRepositoryI interface {
	FindOrderSyncLastDate(ctx context.Context) (string, error)
	UpsertOrderSyncLastDate(ctx context.Context, lastSyncDate string) error
}

func NewOrderSyncRepository(db *database.MongoDB) OrderSyncRepositoryI {
	return &OrderSyncRepository{db: db}
}

func (r *OrderSyncRepository) FindOrderSyncLastDate(ctx context.Context) (string, error) {
	var result struct {
		LastSyncDate string `bson:"last_sync_date"`
	}

	err := r.db.Collection("order_sync").FindOne(ctx, nil).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.LastSyncDate, nil
}

func (r *OrderSyncRepository) UpsertOrderSyncLastDate(
	ctx context.Context,
	lastSyncDate string,
) error {

	filter := bson.M{
		"_id": "order_sync_state",
	}

	update := bson.M{
		"$set": bson.M{
			"last_sync_date": lastSyncDate,
		},
	}

	opts := options.Update().SetUpsert(true)

	_, err := r.db.
		Collection("order_sync").
		UpdateOne(ctx, filter, update, opts)

	return err
}
