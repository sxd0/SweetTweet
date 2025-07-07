package main

import (
	"context"
	"log"
	"os"

	"github.com/sxd0/SweetTweet/internal/user/repository"
)

func main() {
	ctx := context.Background()
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	_ = repository.NewMongoClient(ctx, mongoURI)
}
