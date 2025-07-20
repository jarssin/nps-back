package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
}

func Connect() (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

func (p *MongoDB) Close() error {
	return p.client.Disconnect(context.Background())
}
