package repository

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(ctx context.Context, uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)

	ctxConnect, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxConnect, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	if err := client.Ping(ctxConnect, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	log.Println("MongoDB connected successfully.")
	return client
}
