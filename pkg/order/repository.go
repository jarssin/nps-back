package order

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/database"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryI interface {
	BulkCreate(ctx context.Context, orders []DTO) error
}

type Repository struct {
	db *database.MongoDB
}

func NewRepository(db *database.MongoDB) RepositoryI {
	return &Repository{db: db}
}

func (r *Repository) BulkCreate(ctx context.Context, orders []DTO) error {
	var documents []any
	for _, order := range orders {
		documents = append(documents, order)
	}

	opts := options.InsertMany().SetOrdered(false)
	_, err := r.db.Collection("orders").InsertMany(ctx, documents, opts)
	return err
}
