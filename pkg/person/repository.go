package person

import (
	"context"

	"github.com/jarssin/nps-back/internal/infra/database"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *database.MongoDB
}

type RepositoryI interface {
	Find(ctx context.Context, phone string) (DTO, error)
	BulkCreate(ctx context.Context, people []DTO) error
	Create(ctx context.Context, person DTO) error
}

func NewRepository(db *database.MongoDB) RepositoryI {
	return &Repository{db: db}
}

func (r *Repository) Find(ctx context.Context, phone string) (DTO, error) {
	cursor := r.db.Collection("people").FindOne(ctx, map[string]any{"phone": phone})
	var person DTO
	err := cursor.Decode(&person)
	return person, err
}

func (r *Repository) BulkCreate(ctx context.Context, people []DTO) error {
	var documents []any
	for _, person := range people {
		documents = append(documents, person)
	}

	opts := options.InsertMany().SetOrdered(false)
	_, err := r.db.Collection("people").InsertMany(ctx, documents, opts)
	return err
}

func (r *Repository) Create(ctx context.Context, person DTO) error {
	_, err := r.db.Collection("people").InsertOne(ctx, person)
	return err
}
