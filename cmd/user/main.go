package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sxd0/SweetTweet/internal/user/cache"
	"github.com/sxd0/SweetTweet/internal/user/handler"
	"github.com/sxd0/SweetTweet/internal/user/repository"
	"github.com/sxd0/SweetTweet/internal/user/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Mongo error:", err)
	}

	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	redisCache := cache.NewCache(redisAddr)

	repo := repository.NewProfileRepository(mongoClient)
	svc := service.NewProfileService(repo, redisCache)
	h := handler.NewProfileHandler(svc)

	r := gin.Default()
	h.RegisterRoutes(r)

	log.Println("User service running on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Server error:", err)
	}
}
