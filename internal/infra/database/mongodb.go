package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func Connect(ctx context.Context) (*MongoDB, error) {
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MONGODB_URL")),
	)

	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client: client,
	}, nil
}

func (p *MongoDB) Collection(name string) *mongo.Collection {
	return p.client.Database(os.Getenv("MONGODB_DATABASE")).Collection(name)
}

func (p *MongoDB) Close(ctx context.Context) error {
	return p.client.Disconnect(ctx)
}
